package rpc

import (
	"context"
	"strconv"
	"unsafe"

	codes "github.com/AmazingTalker/at-error-code"
	"github.com/AmazingTalker/at-member-manager/pkg/dao"
	"github.com/AmazingTalker/at-member-manager/pkg/pb"
	"github.com/AmazingTalker/go-rpc-kit/errorkit"
	"github.com/AmazingTalker/go-rpc-kit/logkit"
	"github.com/AmazingTalker/go-rpc-kit/metrickit"
	"github.com/AmazingTalker/go-rpc-kit/validatorkit"
)

var (
	rpcMet = metrickit.NewWithPkgName(
		metrickit.EnableAutoFillInFuncName(true),
	)
)

type AtMemberManagerServerOpt struct {
	Validator validatorkit.Validator
	MemberDao dao.MemberDAO
}

// AtMemberManagerServer 1. Implement a struct as you like.
// Generate everything with an interface named "AtMemberManagerRPC"
type AtMemberManagerServer struct {
	validator validatorkit.Validator
	memberDao dao.MemberDAO
}

func NewAtMemberManagerServer(opt AtMemberManagerServerOpt) AtMemberManagerServer {
	return AtMemberManagerServer{
		validator: opt.Validator,
		memberDao: opt.MemberDao,
	}
}

// Health 2. Complete these methods.
func (serv AtMemberManagerServer) Health(_ context.Context, _ *pb.HealthReq) (*pb.HealthRes, error) {
	return &pb.HealthRes{Ok: true}, nil
}

func (serv AtMemberManagerServer) CreateMember(ctx context.Context, req *pb.CreateMemberReq) (*pb.CreateMemberRes, error) {
	defer rpcMet.RecordDuration([]string{"time"}, map[string]string{}).End()

	if err := serv.validator.Valid(ctx, req); err != nil {
		logkit.ErrorV2(ctx, "CreateMember validate failed", err, nil)
		return nil, err
	}

	m := &dao.Member{
		Name:     req.Name,
		Birthday: req.Birthday,
	}

	createM, err := serv.memberDao.CreateMember(ctx, m)
	if err != nil {
		logkit.ErrorV2(ctx, "dao.CreateMember failed", err, nil)
		return nil, err
	}

	resp := pb.CreateMemberRes{Member: createM.FormatPb()}
	rpcMet.SetGauge([]string{"resp_size"}, float64(unsafe.Sizeof(resp)), map[string]string{})

	return &resp, nil
}

func (serv AtMemberManagerServer) UpdateMember(ctx context.Context, req *pb.UpdateMemberReq) (*pb.UpdateMemberRes, error) {
	defer rpcMet.RecordDuration([]string{"time"}, map[string]string{}).End()

	if err := serv.validator.Valid(ctx, req); err != nil {
		logkit.ErrorV2(ctx, "UpdateMember validate failed", err, nil)
		return nil, err
	}

	ctx = logkit.EnrichPayload(ctx, logkit.Payload{"id": req.ID})

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logkit.ErrorV2(ctx, "UpdateMember parse id to int64 failed", err, nil)
		err := errorkit.NewFromError(codes.ErrUnqualifiedParameters, err, errorkit.WithHttpStatusCode(400))
		return nil, err
	}
	updateM, err := serv.memberDao.UpdateMember(ctx, id, &dao.Member{Name: req.Name, Birthday: req.Birthday})

	if err != nil {
		logkit.ErrorV2(ctx, "dao.UpdateMember failed", err, nil)
		return nil, err
	}

	resp := pb.UpdateMemberRes{Member: updateM.FormatPb()}
	rpcMet.SetGauge([]string{"resp_size"}, float64(unsafe.Sizeof(resp)), map[string]string{})

	return &resp, nil
}

func (serv AtMemberManagerServer) ListMembers(ctx context.Context, req *pb.ListMembersReq) (*pb.ListMembersRes, error) {
	defer rpcMet.RecordDuration([]string{"time"}, map[string]string{}).End()
	defer rpcMet.IncrCount([]string{"Call per day"}, 1, map[string]string{
		"env":     "dev",
		"service": "at-member-manager",
	})

	members, err := serv.memberDao.ListMembers(ctx)
	if err != nil {
		logkit.ErrorV2(ctx, "dao.ListMembers failed", err, nil)
		return nil, err
	}

	result := make([]*pb.Member, len(members))
	for i, m := range members {
		m := m
		result[i] = m.FormatPb()
	}

	resp := pb.ListMembersRes{Members: result}
	rpcMet.SetGauge([]string{"resp_size"}, float64(unsafe.Sizeof(resp)), map[string]string{})

	return &resp, nil
}

func (serv AtMemberManagerServer) DeleteMember(ctx context.Context, req *pb.DeleteMemberReq) (*pb.DeleteMemberRes, error) {
	defer rpcMet.RecordDuration([]string{"time"}, map[string]string{}).End()

	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logkit.ErrorV2(ctx, "DeleteMember parse id to int64 failed", err, nil)
		err := errorkit.NewFromError(codes.ErrUnqualifiedParameters, err, errorkit.WithHttpStatusCode(400))
		return nil, err
	}

	if err := serv.memberDao.DeleteMember(ctx, id); err != nil {
		logkit.ErrorV2(ctx, "dao.DeleteMember failed", err, nil)
		return nil, err
	}

	resp := pb.DeleteMemberRes{}
	rpcMet.SetGauge([]string{"resp_size"}, float64(unsafe.Sizeof(resp)), map[string]string{})

	return &resp, nil
}
