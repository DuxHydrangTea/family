package app

import (
	"family/internal/models"
	"family/internal/modules"
	"family/internal/providers"
	"family/internal/routes"
	"os"

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
		panic("Failed to connect database, check username, password, host:port and database exist")
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database")
	}

	app := gin.Default()

	app.Static("/storage", "./storage")

	uploadProvider := providers.NewS3Provider(LoadConfigR2())
	// uploadProvider := providers.NewLocalProvider()

	allModules := []modules.Module{
		modules.NewUserModule(db, uploadProvider),
	}

	routes.RegisterRoutes(app, GetModuleRoutes(allModules)...)

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

func LoadConfigR2() *providers.S3Config {
	return &providers.S3Config{
		AccountId:   os.Getenv("R2_ACCOUNT_ID"),
		AccessKey:   os.Getenv("R2_ACCESS_KEY"),
		SecretKey:   os.Getenv("R2_SECRET_KEY"),
		BucketName:  os.Getenv("R2_BUCKET_NAME"),
		PublicUrl:   os.Getenv("R2_PUBLIC_URL"),
		ApiEndpoint: os.Getenv("R2_API_END_POINT"),
	}
}
