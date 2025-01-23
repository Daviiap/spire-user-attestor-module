package usecases

import "user_attestor_module/internal/domain"

type UserInteractor struct {
	// Add any repositories or services here
}

func (uc *UserInteractor) GetUserInfo() (*domain.UserInfo, error) {
	// Example dummy data fetching logic
	return &domain.UserInfo{
		UnixInfo: domain.UnixInfo{
			UID:   "1001",
			User:  "johndoe",
			GID:   "1001",
			Group: "users",
			Supplementary: []domain.Supplementary{
				{GID: "1002", Group: "wheel"},
			},
		},
		BasicAuth: domain.BasicAuth{
			User:     "johndoe",
			Password: "securepassword",
		},
	}, nil
}
