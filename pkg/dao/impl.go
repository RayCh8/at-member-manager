package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/AmazingTalker/go-cache"
	"github.com/AmazingTalker/go-rpc-kit/daokit"
	"github.com/AmazingTalker/go-rpc-kit/metrickit"
)

const (
	pfxMember = "members"
)

var (
	met = metrickit.NewWithPkgName(
		metrickit.EnableAutoFillInFuncName(true),
	)
)

type impl struct {
	mysql MySqlMemberDAO
}

func NewMemberDAO(db *gorm.DB, cacheSrv cache.Service) MemberDAO {
	im := &impl{mysql: NewMySqlMemberDAO(db)}

	return im
}

func (im *impl) CreateMember(ctx context.Context, member *Member, enrich ...daokit.Enrich) (*Member, error) {
	defer met.RecordDuration([]string{"time"}, map[string]string{}).End()

	return im.mysql.CreateMember(ctx, member, enrich...)
}

func (im *impl) UpdateMember(ctx context.Context, id int64, member *Member) (*Member, error) {
	defer met.RecordDuration([]string{"time"}, map[string]string{}).End()

	return im.mysql.UpdateMember(ctx, id, member)
}

func (im *impl) ListMembers(ctx context.Context) ([]Member, error) {
	defer met.RecordDuration([]string{"time"}, map[string]string{}).End()

	return im.mysql.ListMembers(ctx)
}

func (im *impl) DeleteMember(ctx context.Context, id int64) error {
	defer met.RecordDuration([]string{"time"}, map[string]string{}).End()

	return im.mysql.DeleteMember(ctx, id)
}
