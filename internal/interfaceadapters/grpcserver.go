package interfaceadapters

import (
	"context"
	"user_attestor_module/internal/usecases"
	pb "user_attestor_module/pkg/protos/user"
)

type UserServer struct {
	getUserInfoUseCase *usecases.UserInteractor
	getTokenUseCase    *usecases.TokenInteractor
	pb.UnimplementedUserServiceServer
}

func NewUserServer(useCase *usecases.UserInteractor) *UserServer {
	return &UserServer{getUserInfoUseCase: useCase}
}

func (s *UserServer) GetUserInfo(ctx context.Context, req *pb.EmptyMessage) (*pb.UserAttestationData, error) {
	attestationToken, err := s.getTokenUseCase.GetAttestationToken()
	if err != nil {
		return nil, err
	}

	userInfo, err := s.getUserInfoUseCase.GetUserInfo()
	if err != nil {
		return nil, err
	}

	supplementaries := make([]*pb.Supplementary, len(userInfo.UnixInfo.Supplementary))
	for i, supp := range userInfo.UnixInfo.Supplementary {
		supplementaries[i] = &pb.Supplementary{
			Gid:   supp.GID,
			Group: supp.Group,
		}
	}

	return &pb.UserAttestationData{
		AttestationToken: attestationToken,
		UserInfo: &pb.UserInfo{
			UnixInfo: &pb.UnixInfo{
				Uid:           userInfo.UnixInfo.UID,
				User:          userInfo.UnixInfo.User,
				Gid:           userInfo.UnixInfo.GID,
				Group:         userInfo.UnixInfo.Group,
				Supplementary: supplementaries,
			},
			BasicAuth: &pb.BasicAuth{
				User:     userInfo.BasicAuth.User,
				Password: userInfo.BasicAuth.Password,
			},
			External:   &pb.External{},
			Biometrics: &pb.Biometrics{},
		},
	}, nil
}
