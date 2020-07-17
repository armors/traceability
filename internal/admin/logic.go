package admin

import (
	"github.com/GoAdminGroup/components/echarts"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/armors/traceability/internal/plugin/model"
	"github.com/gin-gonic/gin"
)

type Logic struct {
	Config config.Config
	Gin *gin.Engine
	Engine *engine.Engine
}

// New init
func New(cfgPath string) (l *Logic) {

	template.AddComp(chartjs.NewChart())
	template.AddComp(echarts.NewChart())

	gin.SetMode(gin.DebugMode)
	r := gin.New()

	e := engine.Default()

	c := config.ReadFromJson(cfgPath)
	c.Animation = config.PageAnimation{
		Type: "fadeInUp",
		Duration: 0.9,
	}
	c.ColorScheme = adminlte.ColorschemeSkinBlack
	l = &Logic{
		Config: c,
		Gin: r,
		Engine: e,
	}
	l.initAdmin()

	l.router()

	return
}

func (l *Logic) initAdmin(){
	err := l.Engine.
		AddConfig(l.Config).
		ResolveMysqlConnection(model.SetConnection).
		AddDisplayFilterXssJsFilter().
		AddGenerators(model.Generators).
		Use(l.Gin)

	if err != nil {
		panic(err)
	}
}

// Close close resources.
func (l *Logic) Close() {
	l.Engine.MysqlConnection().Close()
}