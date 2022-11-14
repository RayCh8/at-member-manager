package dao

import (
	"context"
	"time"

	"github.com/AmazingTalker/at-member-manager/pkg/pb"
	"github.com/AmazingTalker/go-rpc-kit/daokit"
)

type MemberDAO interface {
	CreateMember(context.Context, *Member, ...daokit.Enrich) (*Member, error)
	UpdateMember(context.Context, int64, *Member) (*Member, error)
	ListMembers(context.Context) ([]Member, error)
	DeleteMember(context.Context, int64) error
}

type Member struct {
	ID        int64
	Name      string
	Birthday  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
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
