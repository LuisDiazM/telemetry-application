export interface Location {
  lat: number;
  long: number;
}

export interface DeviceLocation {
  type: string;
  coordinates: Array<number>;
}

export interface DeviceInfo {
  sensor_id: string;
  location: DeviceLocation;
  type_sensor: string;
  units: Array<string>;
}

export interface DeviceData {
  last_ping: Date;
  device: DeviceInfo;
}
