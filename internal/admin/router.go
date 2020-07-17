package admin

import (
	ada "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/armors/traceability/internal/admin/pages"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (l *Logic) router(){
	l.Gin.Static("/uploads", "./uploads")
	l.Gin.GET("/admin", ada.Content(func(ctx *gin.Context) (panel types.Panel, e error) {
		if config.GetTheme() == "adminlte" {
			return pages.GetDashBoardContent(ctx)
		} else {
			return pages.GetDashBoard2Content(ctx)
		}
	}))
	l.Gin.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/admin")
	})
}