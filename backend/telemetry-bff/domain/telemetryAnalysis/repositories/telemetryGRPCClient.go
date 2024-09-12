package repositories

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/telemetryAnalysis/entities"
)

type ITelemetryAnalysisGRPC interface {
	GetMeasuresBySensor(ctx context.Context, deviceId string) *entities.MeasureData
}
