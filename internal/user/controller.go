package user

import (
	"net/http"
	"strconv"
	"tablelink/internal/pkg/common/http/response"

	"github.com/gin-gonic/gin"
)

type HTTPController interface {
	FindAll(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// NewHTTPController ...
func NewHTTPController(userUseCase UseCase) HTTPController {
	return &httpController{
		userUseCase: userUseCase,
	}
}

type httpController struct {
	userUseCase UseCase
}

func (controller *httpController) FindAll(c *gin.Context) {
	users, err := controller.userUseCase.FindAll(c.Request.Context())

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, users)
}

func (controller *httpController) Add(c *gin.Context) {
	var spec User
	err := c.ShouldBindJSON(&spec)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.userUseCase.Add(c.Request.Context(), spec)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusCreated, result)
}

func (controller *httpController) Update(c *gin.Context) {
	var spec User
	var err error
	err = c.ShouldBindJSON(&spec)

	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	spec.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.userUseCase.Update(c.Request.Context(), spec)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}

func (controller *httpController) Delete(c *gin.Context) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	err = controller.userUseCase.Delete(c.Request.Context(), reqID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, nil)
}
