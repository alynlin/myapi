/*
 * OpenAPI User API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

import (
	"github.com/gin-gonic/gin"
)

type UserAPI struct {
}

// Get /users
func (api *UserAPI) ListUsers(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

