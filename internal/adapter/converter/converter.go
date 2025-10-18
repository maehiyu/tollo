package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/maehiyu/tollo/gen/go/protos/userservice"
	"github.com/maehiyu/tollo/internal/userservice/domain/user"
)

func ToUserInfo(u *user.User) *userservice.UserInfo {
	if u == nil {
		return nil
	}

	userInfo := &userservice.UserInfo {
		Id: u.ID,
		Name: u.Name,
		Email: string(u.Email),
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}

	switch p := u.Profile.(type){
		case *user.ProfessionalProfile:
			userInfo.Profile = &userservice.UserInfo_ProfessionalProfile{
				ProfessionalProfile: &userservice.ProfessionalProfile{
					ProBadgeUrl: p.ProBadgeURL,
					Biography: p.Biography,
				},
			}
		case *user.GeneralProfile:
			userInfo.Profile = &userservice.UserInfo_General{
				General: &userservice.GeneralProfile{
					Points: p.Points,
					Introduction: p.Introduction,
				},
			}
	}
	return userInfo
}
