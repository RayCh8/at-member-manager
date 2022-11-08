package rpc

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	codes "github.com/AmazingTalker/at-error-code"
	mockDAO "github.com/AmazingTalker/at-member-manager/internal/pkg/dao"
	"github.com/AmazingTalker/at-member-manager/pkg/dao"
	"github.com/AmazingTalker/at-member-manager/pkg/pb"
	"github.com/AmazingTalker/go-rpc-kit/errorkit"
	"github.com/AmazingTalker/go-rpc-kit/logkit"
	"github.com/AmazingTalker/go-rpc-kit/validatorkit"
)

var (
	mockCTX    = context.Background()
	mockID     = "0"
	mockDate   = time.Now().AddDate(-18, -3, -3)
	mockMember = &dao.Member{
		Name:     "Ray",
		Birthday: &mockDate,
	}
)

type ExpAtError struct {
	ExpStatus int64
	ExpCode   codes.ATErrorCode
}

type rpcSuite struct {
	suite.Suite

	// mocks
	mockMember *mockDAO.MemberDAO

	serv AtMemberManagerServer
}

func (s *rpcSuite) SetupSuite() {
	logkit.RegisterAmazingLogger(&logkit.Config{
		Logger:              logkit.LoggerZap,
		Development:         true,
		IntegrationAirbrake: &logkit.IntegrationAirbrake{},
	})
}

func (s *rpcSuite) TearDownSuite() {
	logkit.Flush()
}

func (s *rpcSuite) SetupTest() {
	// setup mock
	s.mockMember = mockDAO.NewMemberDAO(s.T())

	s.serv = NewAtMemberManagerServer(AtMemberManagerServerOpt{
		Validator: validatorkit.NewGoPlaygroundValidator(),
		MemberDao: s.mockMember,
	})
}

func (s *rpcSuite) TearDownTest() {
	s.mockMember.AssertExpectations(s.T())
}

func TestRPCSuite(t *testing.T) {
	suite.Run(t, new(rpcSuite))
}

func (s *rpcSuite) TestHealth() {
	tests := []struct {
		Desc     string
		Req      *pb.HealthReq
		ExpError *ExpAtError
		ExpRes   *pb.HealthRes
	}{
		{
			Desc:     "normal case",
			ExpError: nil,
			ExpRes:   &pb.HealthRes{Ok: true},
		},
	}

	for _, t := range tests {
		resp, err := s.serv.Health(mockCTX, t.Req)

		if err == nil {
			s.Require().Equal(nil, err, t.Desc)
			s.Require().Equal(t.ExpRes, resp, t.Desc)
		} else {
			atErr := errorkit.FormatError(err)
			s.Require().Equal(t.ExpError.ExpCode, atErr.ATErrorCode(), t.Desc)
			s.Require().Equal(int(t.ExpError.ExpStatus), atErr.HttpStatus(), t.Desc)
			s.Require().Equal(t.ExpRes, resp, t.Desc)
		}

		s.TearDownTest()
	}
}

func (s *rpcSuite) TestCreateMember() {
	tests := []struct {
		Desc      string
		SetupTest func(string)
		Req       *pb.CreateMemberReq
		ExpError  error
		ExpResp   *pb.CreateMemberRes
	}{
		{
			Desc: "create failed",
			SetupTest: func(desc string) {
				s.mockMember.On(
					"CreateMember", mock.Anything, mockMember,
				).Return(
					nil, errors.New("XD"),
				).Once()
			},
			Req: &pb.CreateMemberReq{
				Name:     mockMember.Name,
				Birthday: mockMember.Birthday,
			},
			ExpError: errors.New("XD"),
		},
		{
			Desc: "normal case",
			SetupTest: func(desc string) {
				s.mockMember.On(
					"CreateMember", mock.Anything, mockMember,
				).Return(
					mockMember, nil,
				).Once()
			},
			Req: &pb.CreateMemberReq{
				Name:     mockMember.Name,
				Birthday: mockMember.Birthday,
			},
			ExpError: nil,
			ExpResp: &pb.CreateMemberRes{
				Member: mockMember.FormatPb(),
			},
		},
	}

	for _, t := range tests {
		if t.SetupTest != nil {
			t.SetupTest(t.Desc)
		}

		resp, err := s.serv.CreateMember(mockCTX, t.Req)
		s.Require().Equal(t.ExpError, err, t.Desc)

		if err == nil {
			s.Require().Equal(t.ExpResp, resp, t.Desc)
		}

		s.TearDownTest()
	}
}

func (s *rpcSuite) TestUpdateMember() {
	tests := []struct {
		Desc      string
		SetupTest func(string)
		Req       *pb.UpdateMemberReq
		ExpError  error
		ExpResp   *pb.UpdateMemberRes
	}{
		{
			Desc: "update failed",
			SetupTest: func(desc string) {
				s.mockMember.On(
					"UpdateMember", mock.Anything, mockMember,
				).Return(
					nil, errors.New("XD"),
				).Once()
			},
			Req: &pb.UpdateMemberReq{
				ID:       mockID,
				Name:     mockMember.Name,
				Birthday: mockMember.Birthday,
			},
			ExpError: errors.New("XD"),
		},
		{
			Desc: "normal case",
			SetupTest: func(desc string) {
				s.mockMember.On(
					"UpdateMember", mock.Anything, mockMember,
				).Return(
					mockMember, nil,
				).Once()
			},
			Req: &pb.UpdateMemberReq{
				ID:       mockID,
				Name:     mockMember.Name,
				Birthday: mockMember.Birthday,
			},
			ExpError: nil,
			ExpResp: &pb.UpdateMemberRes{
				Member: mockMember.FormatPb(),
			},
		},
	}

	for _, t := range tests {
		if t.SetupTest != nil {
			t.SetupTest(t.Desc)
		}

		resp, err := s.serv.UpdateMember(mockCTX, t.Req)
		s.Require().Equal(t.ExpError, err, t.Desc)

		if err == nil {
			s.Require().Equal(t.ExpResp, resp, t.Desc)
		}

		s.TearDownTest()
	}
}

func (s *rpcSuite) TestListMember() {
	tests := []struct {
		Desc      string
		SetupTest func(string)
		Req       *pb.ListMembersReq
		ExpError  error
		ExpResp   *pb.ListMembersRes
	}{
		{
			Desc: "update failed",
			SetupTest: func(desc string) {
				s.mockMember.On(
					"ListMembers", mock.Anything, mockMember,
				).Return(
					nil, errors.New("XD"),
				).Once()
			},
			Req:      &pb.ListMembersReq{},
			ExpError: errors.New("XD"),
		},
		{
			Desc: "normal case",
			SetupTest: func(desc string) {
				s.mockMember.On(
					"ListMembers", mock.Anything, mockMember,
				).Return(
					[]dao.Member{*mockMember}, nil,
				).Once()
			},
			Req:      &pb.ListMembersReq{},
			ExpError: nil,
			ExpResp: &pb.ListMembersRes{
				Members: []*pb.Member{mockMember.FormatPb()},
			},
		},
	}

	for _, t := range tests {
		if t.SetupTest != nil {
			t.SetupTest(t.Desc)
		}

		resp, err := s.serv.ListMembers(mockCTX, t.Req)
		s.Require().Equal(t.ExpError, err, t.Desc)

		if err == nil {
			s.Require().Equal(t.ExpResp, resp, t.Desc)
		}

		s.TearDownTest()
	}
}

func (s *rpcSuite) TestDeleteMember() {
	tests := []struct {
		Desc      string
		SetupTest func(string)
		Req       *pb.DeleteMemberReq
		ExpError  error
		ExpResp   *pb.DeleteMemberRes
	}{
		{
			Desc: "delete failed",
			SetupTest: func(desc string) {
				s.mockMember.On(
					"DeleteMember", mock.Anything, mockMember,
				).Return(
					errors.New("XD"),
				).Once()
			},
			Req: &pb.DeleteMemberReq{
				ID: mockID,
			},
			ExpError: errors.New("XD"),
		},
		{
			Desc: "normal case",
			SetupTest: func(desc string) {
				s.mockMember.On(
					"DeleteMember", mock.Anything, mockMember,
				).Return(
					mockMember, nil,
				).Once()
			},
			Req: &pb.DeleteMemberReq{
				ID: mockID,
			},
			ExpError: nil,
			ExpResp:  &pb.DeleteMemberRes{},
		},
	}

	for _, t := range tests {
		if t.SetupTest != nil {
			t.SetupTest(t.Desc)
		}

		resp, err := s.serv.DeleteMember(mockCTX, t.Req)
		s.Require().Equal(t.ExpError, err, t.Desc)

		if err == nil {
			s.Require().Equal(t.ExpResp, resp, t.Desc)
		}

		s.TearDownTest()
	}
}
