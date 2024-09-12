package usecases

import (
	"context"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/telemetryAnalysis/entities"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/telemetryAnalysis/repositories"
)

type MeasuresUsecase struct {
	measuresRepo repositories.ITelemetryAnalysisGRPC
}

func NewMeasuresUsecase(measuresRepo repositories.ITelemetryAnalysisGRPC) *MeasuresUsecase {
	return &MeasuresUsecase{measuresRepo: measuresRepo}
}

func (usecase *MeasuresUsecase) GetLastWeekMeasuresByDevice(ctx context.Context, deviceId string) *entities.MeasureData {
	return usecase.measuresRepo.GetMeasuresBySensor(ctx, deviceId)
}
