package dao

import (
	"context"
	"time"

	"github.com/AmazingTalker/at-member-manager/pkg/pb"
	"github.com/AmazingTalker/go-rpc-kit/daokit"
)

type MemberDAO interface {
	CreateMember(context.Context, *Member, ...daokit.Enrich) (*Member, error)
	UpdateMember(context.Context, int64, string, *time.Time) (*Member, error)
	ListMembers(context.Context) ([]Member, error)
	DeleteMember(context.Context, int64) error
}

type Member struct {
	ID        int64      `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Name      string     `gorm:"type:varchar(255) NOT NULL;" json:"name,omitempty"`
	Birthday  *time.Time `gorm:"type:timestamp NOT NULL" json:"birthday,omitempty"`
	CreatedAt *time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func (m *Member) FormatPb() *pb.Member {
	return &pb.Member{
		ID:        m.ID,
		Name:      m.Name,
		Birthday:  m.Birthday,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
