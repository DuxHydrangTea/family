package modules

import (
	"family/internal/handlers"
	"family/internal/providers"
	"family/internal/repositories"
	"family/internal/routes"
	"family/internal/services"

	"gorm.io/gorm"
)

type UserModule struct {
	r routes.Route
}

func NewUserModule(
	db *gorm.DB,
) *UserModule {
	localStorage := providers.NewLocalProvider()
	repo := repositories.NewUserRepository(db)
	services := services.NewUserService(repo, localStorage)
	handlers := handlers.NewUserHandler(services)
	routes := routes.NewUserRoute(handlers)

	return &UserModule{
		r: routes,
	}
}

func (m *UserModule) GetRoutes() routes.Route {
	return m.r
}
