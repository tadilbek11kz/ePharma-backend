package util

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Respond(c *gin.Context, code int, responseStruct interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")

	c.Writer.WriteHeader(code)
	if responseStruct != nil {
		err := json.NewEncoder(c.Writer).Encode(responseStruct)
		if err != nil {
			logrus.WithError(err).Errorf("cannot send response: %+v", responseStruct)
			c.Writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
