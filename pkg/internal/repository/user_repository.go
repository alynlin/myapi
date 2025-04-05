/**
 * @Author: litsky
 * @Date:   2025/4/4
 * @Last Modified by:   litsky
 * @Last Modified time: 2025/4/4
 * @License: MIT
 */

package repository

import (
	"github.com/alynlin/myapi/pkg/internal/repository/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(limit int) ([]model.User, error)
}

type UserManager struct {
	db gorm.DB
}

func (mgr UserManager) FindAll(limit int) ([]model.User, error) {
	//todo
	return nil, nil
}
