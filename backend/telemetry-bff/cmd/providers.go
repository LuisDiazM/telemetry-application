package cmd

import (
	"github.com/LuisDiazM/backend/telemetry-bff/app"
	"github.com/LuisDiazM/backend/telemetry-bff/infraestructure/web"
	"github.com/google/wire"
)

// infraestructure
var WebServerProvider = wire.NewSet(web.NewHTTPServer)
var SettingsProvider = wire.NewSet(app.GetAppSettings)

// Application
var ApplicationProvider = wire.NewSet(app.NewApplication)
