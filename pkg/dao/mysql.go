package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/AmazingTalker/go-rpc-kit/daokit"
	"github.com/AmazingTalker/go-rpc-kit/logkit"
)

type MySqlMemberDAO struct {
	db *gorm.DB
}

func NewMySqlMemberDAO(db *gorm.DB) MySqlMemberDAO {
	return MySqlMemberDAO{db: db}
}

func (dao MySqlMemberDAO) CreateMember(ctx context.Context, member *Member, enrich ...daokit.Enrich) (*Member, error) {
	defer met.RecordDuration([]string{"mysql", "time"}, map[string]string{}).End()

	db, _ := daokit.UseTxOrDB(dao.db, enrich...)

	err := db.Create(member).Error

	if err != nil {
		return nil, err
	}
	return member, nil
}

func (dao MySqlMemberDAO) UpdateMember(ctx context.Context, id int64, m *Member) (*Member, error) {
	defer met.RecordDuration([]string{"mysql", "time"}, map[string]string{}).End()

	if m == nil {
		return nil, errors.New("member pointer is nil.")
	}
	member := &Member{}

	err := dao.db.First(member, "id = ?", id).Updates(Member{Name: m.Name, Birthday: m.Birthday}).Error

	if err != nil {
		logkit.Debug(ctx, "update member failed", logkit.Payload{"id": id, "err": err})
		return nil, err
	}

	return member, nil
}

func (dao MySqlMemberDAO) ListMembers(ctx context.Context) ([]Member, error) {
	defer met.RecordDuration([]string{"mysql", "time"}, map[string]string{}).End()

	query := dao.db

	list := []Member{}
	if err := query.Find(&list).Error; err != nil {
		logkit.Debug(ctx, "list member failed", logkit.Payload{"err": err})
		return nil, err
	}

	return list, nil
}

func (dao MySqlMemberDAO) DeleteMember(ctx context.Context, id int64) error {
	defer met.RecordDuration([]string{"mysql", "time"}, map[string]string{}).End()

	if err := dao.db.Delete(&Member{}, id).Error; err != nil {
		logkit.Debug(ctx, "Delete member failed", logkit.Payload{"err": err})
		return err
	}

	return nil
}
