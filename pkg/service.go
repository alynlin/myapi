/**
 * @Author: litsky
 * @Date:   2025/4/5
 * @Last Modified by:   litsky
 * @Last Modified time: 2025/4/5
 * @License: MIT
 */

package pkg

import (
	"context"
	"github.com/alynlin/myapi/pkg/internal/service"
	"github.com/alynlin/myapi/pkg/logging"
	v1 "github.com/alynlin/myapi/pkg/model/v1"
	"log"
)

func UserService(logger logging.Logger) v1.UserAPIService {

	us, err := service.UserServiceBuild{}.Logger(logger).Build(context.Background())

	if err != nil {
		//todo
		log.Fatalln(err)
	}

	return us

}
