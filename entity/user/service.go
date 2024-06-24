package user

import (
	"errors"
	"guestbook/common"
	"guestbook/entity"
)

type Service struct {
	userRepository *Repository
}

func InitService(userRepository *Repository) *Service {
	return &Service{userRepository: userRepository}
}

func (service *Service) SignUp(
	handle string, githubId entity.ID, profileUrl string) (entity.ID, error) {
	userRepository := service.userRepository
	if userRepository.ExistsByGithubId(githubId) {
		return entity.DefaultId(), errors.New("[ERROR] user Already Exist")
	}

	user, err := createFreeUser(handle, githubId, profileUrl)
	if err != nil {
		return entity.DefaultId(), err
	}

	savedUser, err := userRepository.Save(*user)
	if err != nil {
		return entity.DefaultId(), err
	}

	return savedUser.id, err
}

func createFreeUser(handle string, githubId entity.ID, profileUrl string) (*Entity, error) {
	url, err := common.CreateUrl(profileUrl)
	if err != nil {
		return nil, err
	}

	return &Entity{
		handle:     handle,
		nickname:   handle,
		githubId:   githubId,
		profileUrl: url,
		plan:       FreePlan(),
	}, nil
}
