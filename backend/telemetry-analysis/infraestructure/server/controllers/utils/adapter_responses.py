from typing import Dict
from infraestructure.server.protos.measures import measures_pb2
from google.protobuf.timestamp_pb2 import Timestamp
from datetime import datetime

def adapter_measures_response(data: Dict) -> measures_pb2.MeasuresByDeviceResponse:
    meta = data.get("sensor_metadata",{})
    meta_loc = meta.get("location",{})
    metadata = measures_pb2.SensorMetadata(sensor_id=meta.get("sensor_id",""),
                                           location=measures_pb2.Geolocation(type=meta_loc.get("type","Pointer"), coordinates=meta_loc.get("coordinates",[])),
                                           type_sensor=meta.get("type_sensor",""),
                                           units=meta.get("units",[]))
    values = data.get("values",[])
    values_response = []
    for val in values:
        time_data = val.get("timestamp", datetime.now())
        timestamp = Timestamp()
        timestamp.FromDatetime(time_data)
        measure = measures_pb2.Values(timestamp=timestamp, humidity=val.get("humidity", 0.0), temperature=val.get("temperature", 0.0))
        values_response.append(measure)

    return measures_pb2.MeasuresByDeviceResponse(sensor_metadata=metadata, values=values_response)