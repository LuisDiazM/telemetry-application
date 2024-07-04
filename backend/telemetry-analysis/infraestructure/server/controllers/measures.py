from infraestructure.server.protos.measures import measures_pb2, measures_pb2_grpc
from google.protobuf.timestamp_pb2 import Timestamp
from datetime import datetime

class MeasuresIoTServicer(measures_pb2_grpc.MeasuresIoTServicer):
    def __init__(self) -> None:
        super().__init__()
    
    def GetMeasuresBySensor(self, request, context):

        sensor_metadata = measures_pb2.SensorMetadata(
            sensor_id="88888888888888",
            type_sensor="DHT22",
            units=["Celsius", "Relative"],
            location=measures_pb2.Geolocation(
                type="Point",
                coordinates=[7.097549, -73.118515]
            )
        )

        # Crear valores de ejemplo
        timestamp = Timestamp()
        timestamp.FromDatetime(datetime.now())

        values1 = measures_pb2.Values(
            temperature=25.8,
            humidity=0.5,
            timestamp=timestamp
        )

        response = measures_pb2.MeasuresByDeviceResponse(
            values=[values1, values1],
            sensor_metadata=sensor_metadata
        
        )

        return response