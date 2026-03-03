package app

import (
	"family/internal/models"
	"family/internal/modules"
	"family/internal/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Application struct {
	App *gin.Engine
}

func NewApplication() *Application {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/family?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	app := gin.Default()

	modules := []modules.Module{
		modules.NewUserModule(db),
	}

	routes.RegisterRoutes(app, GetModuleRoutes(modules)...)

	return &Application{
		App: app,
	}
}

func (a *Application) RunApp() error {
	return a.App.Run("127.0.0.1:8080")
}

func GetModuleRoutes(modules []modules.Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))

	for i, module := range modules {
		routeList[i] = module.GetRoutes()
	}

	return routeList
}
