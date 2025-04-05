/**
 * @Author: litsky
 * @Date:   2025/4/4
 * @Last Modified by:   litsky
 * @Last Modified time: 2025/4/4
 * @License: MIT
 */

package v1

import (
	"context"
	"github.com/alynlin/myapi/pkg/model"
)

type UserAPIService interface {
	ListUsers(ctx context.Context, limit int) (model.ImplResponse, error)
}
