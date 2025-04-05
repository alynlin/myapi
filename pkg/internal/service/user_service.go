/**
 * @Author: litsky
 * @Date:   2025/4/4
 * @Last Modified by:   litsky
 * @Last Modified time: 2025/4/4
 * @License: MIT
 */

package service

import (
	"context"
	"errors"
	c_err "github.com/alynlin/myapi/pkg/errors"
	"github.com/alynlin/myapi/pkg/internal/repository"
	"github.com/alynlin/myapi/pkg/logging"
	"github.com/alynlin/myapi/pkg/model"
	v1 "github.com/alynlin/myapi/pkg/model/v1"
	"gorm.io/gorm"
	"net/http"
)

type UserService struct {
	repo   repository.UserRepository
	logger logging.Logger
}

func (s *UserService) ListUsers(ctx context.Context, limit int) (model.ImplResponse, error) {
	users, err := s.repo.FindAll(limit)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error(ctx, "user不存在")
		} else {
			s.logger.Error(ctx, "get products error %v", err)
			return model.ImplResponse{}, c_err.New(err, int(c_err.DB_ERROR), "get users error")
		}
	}
	us := []v1.User{}

	for _, user := range users {
		us = append(us, v1.User{
			Id:   int32(user.ID),
			Name: user.Name,
		})
	}

	//todo
	return model.ImplResponse{
		Code: http.StatusOK,
		Body: v1.UsersResponse{Data: us},
	}, nil
}
