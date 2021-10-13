package util

import (
	"breeding/internal/log"
	"breeding/internal/model"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MinInt(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func ConvertETH(s string) *float64 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return nil
	}
	res := f / 1e18
	return &res
}

func Response(c *gin.Context, res interface{}, intErr *model.Error, panicErr interface{}) {
	httpCode := http.StatusOK
	if panicErr != nil {
		intErr = &model.Error{
			HTTPCode:        http.StatusInternalServerError,
			Message:         "Panic Error",
			InternalMessage: fmt.Sprintf("panic: %v", panicErr),
		}
		log.Logger.Debug(string(debug.Stack()))
	}

	if intErr != nil {
		httpCode = intErr.HTTPCode
		res = intErr.ErrorResp()
		if intErr.InternalMessage != "" {
			log.Logger.Error(intErr.InternalMessage)
		}
	}

	c.JSON(httpCode, res)
}
