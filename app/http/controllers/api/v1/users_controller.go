package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/auth"
	"gohub/pkg/response"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}