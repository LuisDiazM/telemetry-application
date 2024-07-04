from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MeasuresByDeviceRequest(_message.Message):
    __slots__ = ("device_id",)
    DEVICE_ID_FIELD_NUMBER: _ClassVar[int]
    device_id: str
    def __init__(self, device_id: _Optional[str] = ...) -> None: ...

class Geolocation(_message.Message):
    __slots__ = ("type", "coordinates")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    COORDINATES_FIELD_NUMBER: _ClassVar[int]
    type: str
    coordinates: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, type: _Optional[str] = ..., coordinates: _Optional[_Iterable[float]] = ...) -> None: ...

class SensorMetadata(_message.Message):
    __slots__ = ("sensor_id", "type_sensor", "units", "location")
    SENSOR_ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_SENSOR_FIELD_NUMBER: _ClassVar[int]
    UNITS_FIELD_NUMBER: _ClassVar[int]
    LOCATION_FIELD_NUMBER: _ClassVar[int]
    sensor_id: str
    type_sensor: str
    units: _containers.RepeatedScalarFieldContainer[str]
    location: Geolocation
    def __init__(self, sensor_id: _Optional[str] = ..., type_sensor: _Optional[str] = ..., units: _Optional[_Iterable[str]] = ..., location: _Optional[_Union[Geolocation, _Mapping]] = ...) -> None: ...

class Values(_message.Message):
    __slots__ = ("temperature", "humidity", "timestamp")
    TEMPERATURE_FIELD_NUMBER: _ClassVar[int]
    HUMIDITY_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    temperature: float
    humidity: float
    timestamp: _timestamp_pb2.Timestamp
    def __init__(self, temperature: _Optional[float] = ..., humidity: _Optional[float] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class MeasuresByDeviceResponse(_message.Message):
    __slots__ = ("values", "sensor_metadata")
    VALUES_FIELD_NUMBER: _ClassVar[int]
    SENSOR_METADATA_FIELD_NUMBER: _ClassVar[int]
    values: _containers.RepeatedCompositeFieldContainer[Values]
    sensor_metadata: SensorMetadata
    def __init__(self, values: _Optional[_Iterable[_Union[Values, _Mapping]]] = ..., sensor_metadata: _Optional[_Union[SensorMetadata, _Mapping]] = ...) -> None: ...
