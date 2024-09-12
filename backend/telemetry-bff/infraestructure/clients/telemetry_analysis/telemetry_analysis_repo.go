package telemetry_analysis

import (
	context "context"
	"flag"

	"github.com/LuisDiazM/backend/telemetry-bff/domain/telemetryAnalysis/entities"
	"github.com/LuisDiazM/backend/telemetry-bff/domain/telemetryAnalysis/repositories"
	"github.com/phuslu/log"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TelemetryAnalysisGRPCClientRepo struct {
	conn        *grpc.ClientConn
	serviceHost string
}

func NewTelemetryAnalysisGRPCClient(settings *TelemetryAnalysisSettings) repositories.ITelemetryAnalysisGRPC {
	serverAddr := flag.String("addrTelemetryManager", settings.TelemetryAnalysisHosts, "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	return &TelemetryAnalysisGRPCClientRepo{conn: conn, serviceHost: *serverAddr}
}

func (repo *TelemetryAnalysisGRPCClientRepo) GetMeasuresBySensor(ctx context.Context, deviceId string) *entities.MeasureData {
	client := NewMeasuresIoTClient(repo.conn)
	req := MeasuresByDeviceRequest{DeviceId: deviceId}
	response, err := client.GetMeasuresBySensor(ctx, &req)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil
	}
	return ParseMeasuresToEntities(response)
}
