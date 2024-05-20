package usecase

import (
	error2 "github.com/kimoscloud/user-management-service/internal/core/errors"
	"github.com/kimoscloud/user-management-service/internal/core/model/entity"
	"github.com/kimoscloud/user-management-service/internal/core/ports/logging"
	userRepository "github.com/kimoscloud/user-management-service/internal/core/ports/repository"
	"github.com/kimoscloud/value-types/errors"
)

type GetUSerListByIdsUseCase struct {
	userRepo userRepository.Repository
	logger   logging.Logger
}

func NewGetUSerListByIdsUseCase(
	ur userRepository.Repository,
	logger logging.Logger,
) *GetUSerListByIdsUseCase {
	return &GetUSerListByIdsUseCase{userRepo: ur, logger: logger}
}

func (p *GetUSerListByIdsUseCase) Handler(ids []string) (
	[]entity.User, *errors.AppError,
) {
	result, err := p.userRepo.GetUsersByIds(ids)
	if err != nil {
		p.logger.Error("Error getting users by ids", err)
		return nil, errors.NewInternalServerError(
			"Error getting users by ids", err.Error(), error2.ErrorGettingUserListByIds).AppError
	}
	return result, nil
}
