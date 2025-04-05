/**
 * @Author: litsky
 * @Date:   2025/4/4
 * @Last Modified by:   litsky
 * @Last Modified time: 2025/4/4
 * @License: MIT
 */

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
}
