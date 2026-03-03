package modules

import "family/internal/routes"

type Module interface {
	GetRoutes() routes.Route
}
