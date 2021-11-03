package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Test ...
func Test(c *gin.Context) {
	c.String(http.StatusOK, "TEST")
}
