package telemetry_analysis

type TelemetryAnalysisSettings struct {
	TelemetryAnalysisHosts string `envconfig:"TELEMETRY_ANALYSIS_HOST" required:"true"`
}
