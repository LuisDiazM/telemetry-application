import { useEffect, useState } from 'react';
import MenuAppBar from '../../shared/nav-bar/navbar';
import api from '../../services/apiService';

import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import { List, ListItemButton, ListItemText } from '@mui/material';

interface Location {
  lat: number;
  long: number;
}

interface DeviceLocation {
  type: string;
  coordinates: Array<number>;
}

interface DeviceInfo {
  sensor_id: string;
  location: DeviceLocation;
  type_sensor: string;
  units: Array<string>;
}

interface DeviceData {
  last_ping: Date;
  device: DeviceInfo;
}

const blinkingStyle = {
  animation: 'blink 1s infinite',
};

const styleSheet = document.styleSheets[0];
const keyframes = `
    @keyframes blink {
      0% { background-color: #4caf50; }
      50% { background-color: #a5d6a7; }
      100% { background-color: #4caf50; }
    }
  `;

styleSheet.insertRule(keyframes, styleSheet.cssRules.length);

const DevicesPage = () => {
  const [location, setLocation] = useState<Location | undefined>(undefined);
  const [devices, setDevices] = useState<Array<DeviceData>>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const fetchDevices = async (location: Location) => {
    try {
      const result = await api.get(
        `/status?latitude=${location.lat}&longitude=${location.long}`
      );
      setDevices(result.data || []);
    } catch (error) {
      setError('Error fetching devices.');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { latitude, longitude } = position.coords;
          const newLocation = { lat: latitude, long: longitude };
          setLocation(newLocation);
          fetchDevices(newLocation);
        },
        (error) => {
          setError('Error retrieving location.');
          console.error(error);
          setLoading(false);
        }
      );
    } else {
      setError('Geolocation is not supported by this browser.');
      setLoading(false);
    }
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <>
      <MenuAppBar />
      <div style={{ height: '600px', width: '100%' }}>
        {location && (
          <MapContainer
            zoom={20}
            center={[location.lat, location.long]}
            style={{ height: '100%', width: '100%' }}
          >
            <TileLayer
              url='https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png'
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            />
            {devices.map((device) => (
              <Marker
                key={device.device.sensor_id}
                position={[
                  device.device.location.coordinates[0],
                  device.device.location.coordinates[1],
                ]}
              >
                <Popup>{device.device.sensor_id}</Popup>
              </Marker>
            ))}
          </MapContainer>
        )}
      </div>

      <div>
        <h3>Dispositivos cercanos</h3>
        <List>
          {devices.map((device) => {
            return (
              <ListItemButton style={blinkingStyle}>
                <ListItemText primary={device.device.sensor_id} />
                <ListItemText
                  primary={device.device.type_sensor}
                ></ListItemText>
              </ListItemButton>
            );
          })}
        </List>
      </div>
    </>
  );
};

export default DevicesPage;
