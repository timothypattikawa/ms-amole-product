package exception

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/product-service/internal/dto"
	"net/http"
)

func CostumeErrorAdvice(err error, e echo.Context) {
	if isErrorNotFound(err, e) {
		return
	} else if isErrorBadRequest(err, e) {
		return
	} else {
		isErrorInternalServer(err, e)
	}
}

func isErrorInternalServer(err error, c echo.Context) bool {
	var e *InternalServerError
	ok := errors.As(err, &e)
	if ok {
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Err: e.Error(),
		})
		return true
	}
	return false
}

func isErrorBadRequest(err error, c echo.Context) bool {
	var e *BadRequestError
	ok := errors.As(err, &e)
	if ok {
		c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Err: e.Error(),
		})
		return true
	}
	return false
}

func isErrorNotFound(err error, c echo.Context) bool {
	var e *NotFoundError
	ok := errors.As(err, &e)
	if ok {
		c.JSON(http.StatusNotFound, dto.BaseResponse{
			Err: e.Error(),
		})
		return true
	}
	return false
}
