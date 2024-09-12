export interface MeasuresData {
    values:          Value[];
    sensor_metadata: SensorMetadata;
}

export interface SensorMetadata {
    units:       string[];
    sensor_id:   string;
    type_sensor: string;
    location:    Location;
}

export interface Location {
    coordinates: number[];
    type:        string;
}

export interface Value {
    temperature: number;
    humidity:    number;
    timestamp:   Date;
}