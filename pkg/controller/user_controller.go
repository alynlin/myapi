/**
 * @Author: litsky
 * @Date:   2025/4/5
 * @Last Modified by:   litsky
 * @Last Modified time: 2025/4/5
 * @License: MIT
 */

package controller

import (
	"context"
	v1 "github.com/alynlin/myapi/pkg/model/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAPIController struct {
	Service v1.UserAPIService
}

func (ctr *UserAPIController) ListUsers(c *gin.Context) {

	ctx := context.Background()

	res, _ := ctr.Service.ListUsers(ctx, 10)

	c.JSON(http.StatusOK, res)

}
