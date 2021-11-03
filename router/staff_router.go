package router

import (
	v1 "github.com/JEDIAC/server/api/v1"
	"github.com/gin-gonic/gin"
)

// InitStaffRouter ...
func InitStaffRouter(r gin.IRouter) {
	apiv1 := r.Group("/v1")
	apiv1.Any("/", v1.Test)
}
