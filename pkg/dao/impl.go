package dao

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/AmazingTalker/go-cache"
	"github.com/AmazingTalker/go-rpc-kit/daokit"
	"github.com/AmazingTalker/go-rpc-kit/metrickit"
)

const (
	pfxRecord = "members"
)

var (
	met = metrickit.NewWithPkgName(
		metrickit.EnableAutoFillInFuncName(true),
	)
)

type impl struct {
	mysql MySqlMemberDAO
	cache cache.Cache
}

func NewMemberDAO(db *gorm.DB, cacheSrv cache.Service) MemberDAO {
	im := &impl{mysql: NewMySqlMemberDAO(db)}

	im.cache = cacheSrv.Create([]cache.Setting{
		{
			Prefix: pfxRecord,
			CacheAttributes: map[cache.Type]cache.Attribute{
				cache.SharedCacheType: {TTL: time.Minute},
				cache.LocalCacheType:  {TTL: 10 * time.Second},
			},
		},
	})

	/*
		Use cases:
			ctx := context.Background()

			1) Get()

			record := Record{}
			if err := im.cache.Get(ctx, pfxRecord, "key", &record); err != nil {
				return err
			}

			---

			2) GetM()

			records := []*Record{}
			res, err := im.cache.MGet(ctx, pfxRecord, "key1", "key2", "key3")
			if err != nil {
				return err
			}

			for i := 0; i < res.Len(); i++ {
				r := &Record{}
				if err := res.Get(ctx, i, r); err != nil {
					return err // It may be ErrCacheMiss or other errors
				}
				records = append(records, r)
			}

		More examples: https://github.com/AmazingTalker/go-cache
	*/

	return im
}

func (im *impl) CreateMember(ctx context.Context, member *Member, enrich ...daokit.Enrich) (*Member, error) {
	defer met.RecordDuration([]string{"time"}, map[string]string{}).End()

	return im.mysql.CreateMember(ctx, member, enrich...)
}

func (im *impl) UpdateMember(ctx context.Context, id int64, name string, birthday *time.Time) (*Member, error) {
	return im.mysql.UpdateMember(ctx, id, name, birthday)
}

func (im *impl) ListMembers(ctx context.Context) ([]Member, error) {
	return im.mysql.ListMembers(ctx)
}

func (im *impl) DeleteMember(ctx context.Context, id int64) error {
	return im.mysql.DeleteMember(ctx, id)
}
