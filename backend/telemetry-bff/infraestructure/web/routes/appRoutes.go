package routes

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web/middlewares"
)

const (
	pubv1 = "telemetry-bff/public/v1"
	priv1 = "telemetry-bff/api/v1"
)

func SetUpRoutes(app *app.Application) {
	publicRoutesV1 := app.WebServer.Group(pubv1)
	privateRoutesV1 := app.WebServer.Group(priv1)
	privateRoutesV1.Use(middlewares.JWTAuth0(app.Settings.Auth0Domain))

	//publics
	RegisterHealthRouter(publicRoutesV1)

	//privates
	RegisterUsersRouter(privateRoutesV1, app)
	RegisterDevicesRouter(privateRoutesV1, app)
}
