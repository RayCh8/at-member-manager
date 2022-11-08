package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/AmazingTalker/go-cache"
	"github.com/AmazingTalker/go-rpc-kit/cachekit"
	"github.com/AmazingTalker/go-rpc-kit/dockerkit"
	"github.com/AmazingTalker/go-rpc-kit/logkit"
	"github.com/AmazingTalker/go-rpc-kit/migrationkit"
	"github.com/AmazingTalker/go-rpc-kit/mysqlkit"
)

const (
	migrationDir = "../../database/migrations"
	sqlURLFormat = "root:root@tcp(localhost:%s)/mysql?charset=utf8mb4&parseTime=True"
	rdsURLFormat = ":%s"
)

var (
	mockCTX     = context.Background()
	mockID      = int64(0)
	mockTimeNow time.Time
	mockLoc     *time.Location
)

func init() {
	mockLoc, _ = time.LoadLocation("")
	mockTimeNow = time.Unix(1629446406, 0).In(mockLoc)
}

type daoSuite struct {
	suite.Suite

	ring  *redis.Ring
	db    *gorm.DB
	cache cache.Service
	im    *impl

	redisPort string
	mysqlPort string
}

func (s *daoSuite) migrationDir() string {
	if dockerkit.RunCITest() {
		return os.Getenv("MIGRATION_DIR")
	}

	return migrationDir
}

func (s *daoSuite) mysqlURL() string {
	if dockerkit.RunCITest() {
		return os.Getenv("MYSQL_DSN")
	}

	return fmt.Sprintf(sqlURLFormat, s.mysqlPort)
}

func (s *daoSuite) redisAddrs() map[string]string {
	if dockerkit.RunCITest() {
		addrs := strings.Split(os.Getenv("REDIS_ADDRS"), ",")

		m := map[string]string{}
		for _, addr := range addrs {
			strs := strings.SplitN(addr, ":", 2)
			m[strs[0]] = strs[1]
		}

		return m
	}

	return map[string]string{"server1": fmt.Sprintf(rdsURLFormat, s.redisPort)}
}

func (s *daoSuite) SetupSuite() {
	// setup logger
	logkit.RegisterAmazingLogger(&logkit.Config{
		Logger:              logkit.LoggerZap,
		Development:         true,
		IntegrationAirbrake: &logkit.IntegrationAirbrake{},
	})

	// run dockerkit when dealing with go test locally
	if dockerkit.RunLocalTest() {
		ports, err := dockerkit.RunExtDockers(mockCTX, []dockerkit.Image{
			dockerkit.ImageMySQL,
			dockerkit.ImageRedis,
		})
		s.Require().NoError(err)
		s.mysqlPort = ports[0]
		s.redisPort = ports[1]
	}

	// setup mysql
	logkit.Info(mockCTX, "init mysql", logkit.Payload{"dir": s.migrationDir(), "mysqlURL": s.mysqlURL()})
	migration := migrationkit.NewGooseMigrationKit(migrationkit.GooseMysqlDriver, migrationkit.GooseMigrationOpt{
		Dir:      s.migrationDir(),
		DBString: s.mysqlURL(),
	})
	s.Require().NoError(migration.Up())
	migration.Close()

	db, err := mysqlkit.NewMySqlConn(mysqlkit.MySqlConnOpt{
		Config: &mysqlkit.MysqlConnConf{
			DSN: s.mysqlURL(),
		},
	})
	s.Require().NoError(err)
	s.db = db

	// setup redis
	logkit.Info(mockCTX, "init redis", logkit.Payload{"redisAddrs": s.redisAddrs()})
	s.ring = redis.NewRing(&redis.RingOptions{
		Addrs: s.redisAddrs(),
	})
}

func (s *daoSuite) TearDownSuite() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()
	s.ring.Close()

	if dockerkit.RunLocalTest() {
		dockerkit.PurgeExtDockers(mockCTX, []dockerkit.Image{
			dockerkit.ImageRedis,
			dockerkit.ImageMySQL,
		})
	}

	logkit.Flush()
}

func (s *daoSuite) SetupTest() {
	cache.ClearPrefix()
	s.cache = cachekit.NewCache(
		cachekit.NewSharedCache(s.ring),
		cachekit.NewLocalCache(1024),
	)

	s.im = NewMemberDAO(s.db, s.cache).(*impl)
}

func (s *daoSuite) TearDownTest() {
	cache.ClearPrefix()

	// clean all in redis
	s.Require().NoError(s.ring.ForEachShard(mockCTX, func(ctx context.Context, client *redis.Client) error {
		return client.FlushDB(ctx).Err()
	}))

	// clean all in mysql
	s.Require().NoError(s.db.Where("1 = 1").Delete(&Member{}).Error)
}

func TestDAOSuite(t *testing.T) {
	suite.Run(t, new(daoSuite))
}

func (s *daoSuite) TestCreateMember() {
	tests := []struct {
		Desc      string
		Member    *Member
		CheckFunc func(string)
	}{
		{
			Desc: "normal case",
			Member: &Member{
				Name:      "ray",
				Birthday:  &mockTimeNow,
				CreatedAt: &mockTimeNow,
				UpdatedAt: &mockTimeNow,
			},
			CheckFunc: func(desc string) {
				members := []Member{}
				s.Require().NoError(s.db.Find(&members).Error, desc)
				s.Require().Equal(1, len(members), desc)

				member := members[0]
				s.Require().Equal(mockTimeNow, *member.CreatedAt, desc)
				s.Require().Equal(int64(0), member.ID, desc)
				s.Require().Equal("ray", member.Name, desc)
			},
		},
	}

	for _, t := range tests {
		s.SetupTest()

		_, err := s.im.CreateMember(mockCTX, t.Member)
		s.Require().NoError(err, t.Desc)

		if t.CheckFunc != nil {
			t.CheckFunc(t.Desc)
		}

		s.TearDownTest()
	}
}

func (s *daoSuite) TestUpdateMember() {
	tests := []struct {
		Desc      string
		SetupTest func(string)
		ID        int64
		ExpErr    error
		ExpMember *Member
		CheckFunc func(string)
	}{
		{
			Desc:   "not existed",
			ID:     3,
			ExpErr: fmt.Errorf("member not found"),
		},
		{
			Desc: "normal case",
			SetupTest: func(desc string) {
				ms := []Member{
					{CreatedAt: &mockTimeNow, UpdatedAt: &mockTimeNow, Name: "AT", Birthday: &mockTimeNow},
				}
				s.Require().NoError(s.db.Create(&ms).Error, desc)
			},
			ID:     mockID,
			ExpErr: nil,
			ExpMember: &Member{
				ID:        mockID,
				Name:      "ray",
				Birthday:  &mockTimeNow,
				CreatedAt: &mockTimeNow,
				UpdatedAt: &mockTimeNow,
			},
			CheckFunc: func(desc string) {
				// check cache
				b, err := s.ring.Get(mockCTX, fmt.Sprintf("ca:members:%d", mockID)).Bytes()
				s.Require().NoError(err, desc)

				m := Member{}
				s.Require().NoError(json.Unmarshal(b, &m), desc)
				s.Require().Equal(Member{
					ID:        mockID,
					Name:      "ray",
					Birthday:  &mockTimeNow,
					CreatedAt: &mockTimeNow,
					UpdatedAt: &mockTimeNow,
				}, m, desc)
			},
		},
	}

	for _, t := range tests {
		s.SetupTest()

		_, err := s.im.UpdateMember(mockCTX, t.ExpMember.ID, t.ExpMember.Name, t.ExpMember.Birthday)
		s.Require().NoError(err, t.Desc)

		if t.CheckFunc != nil {
			t.CheckFunc(t.Desc)
		}

		s.TearDownTest()
	}
}

func (s *daoSuite) TestListMembers() {
	tests := []struct {
		Desc       string
		SetupTest  func(string)
		ExpErr     error
		ExpMembers []Member
		CheckFunc  func(string)
	}{
		{
			Desc:       "no records",
			ExpErr:     nil,
			ExpMembers: []Member{},
		},
		{
			Desc: "normal case",
			SetupTest: func(desc string) {
				ms := []Member{
					{CreatedAt: &mockTimeNow, UpdatedAt: &mockTimeNow, Name: "AT", Birthday: &mockTimeNow},
				}
				s.Require().NoError(s.db.Create(&ms).Error, desc)
			},
			ExpErr: nil,
			ExpMembers: []Member{
				{
					ID:        mockID,
					Name:      "AT",
					Birthday:  &mockTimeNow,
					CreatedAt: &mockTimeNow,
					UpdatedAt: &mockTimeNow,
				},
			},
			CheckFunc: func(desc string) {
				// check cache
				b, err := s.ring.Get(mockCTX, "ca:members:0-10").Bytes()
				s.Require().NoError(err, desc)

				rs := []Member{}
				s.Require().NoError(json.Unmarshal(b, &rs), desc)
				s.Require().Equal([]Member{{
					ID:        mockID,
					Name:      "ray",
					Birthday:  &mockTimeNow,
					CreatedAt: &mockTimeNow,
					UpdatedAt: &mockTimeNow,
				}}, rs, desc)
			},
		},
	}

	for _, t := range tests {
		s.SetupTest()

		if t.SetupTest != nil {
			t.SetupTest(t.Desc)
		}

		members, err := s.im.ListMembers(mockCTX)
		s.Require().Equal(t.ExpErr, err, t.Desc)
		if err == nil {
			s.Require().Equal(t.ExpMembers, members, t.Desc)
		}

		if t.CheckFunc != nil {
			t.CheckFunc(t.Desc)
		}

		s.TearDownTest()
	}
}

func (s *daoSuite) TestDeleteMember() {
	tests := []struct {
		Desc      string
		SetupTest func(string)
		ID        int64
		ExpErr    error
		CheckFunc func(string)
	}{
		{
			Desc:   "not existed",
			ID:     3,
			ExpErr: fmt.Errorf("record not found"),
		},
		{
			Desc: "normal case",
			SetupTest: func(desc string) {
				ms := []Member{
					{CreatedAt: &mockTimeNow, UpdatedAt: &mockTimeNow, Name: "AT", Birthday: &mockTimeNow},
				}
				s.Require().NoError(s.db.Create(&ms).Error, desc)
			},
			ID:     mockID,
			ExpErr: nil,
			CheckFunc: func(desc string) {
				// check cache
				b, err := s.ring.Get(mockCTX, fmt.Sprintf("ca:members:%d", mockID)).Bytes()
				s.Require().NoError(err, desc)

				r := Member{}
				s.Require().NoError(json.Unmarshal(b, &r), desc)
			},
		},
	}

	for _, t := range tests {
		s.SetupTest()

		err := s.im.DeleteMember(mockCTX, mockID)
		s.Require().NoError(err, t.Desc)

		if t.CheckFunc != nil {
			t.CheckFunc(t.Desc)
		}

		s.TearDownTest()
	}
}
