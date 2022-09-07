package handler

import (
	"net/http"

	"github.com/chocogem/bigquery-golang/pkg/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	userUseCase  usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase:  uc,
	}
}

// @Description  List all user
// @Tags         User
// @Summary      List all user
// @Produce      json
// @Success      200  {object}  UserResponse
// @Router       /user
func (cr *UserHandler) FindAll(c *gin.Context) {
	users, err := cr.userUseCase.FindAll()
	if err != nil {
		c.Error(err)
	} else {
		var response UserResponse
		copier.Copy(&response.Data, &users)

		c.JSON(http.StatusOK, response)
	}

}
