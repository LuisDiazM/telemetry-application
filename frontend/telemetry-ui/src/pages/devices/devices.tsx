import { MouseEventHandler, useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import MenuAppBar from '../../shared/nav-bar/navbar';
import api from '../../services/apiService';

import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import { List, ListItemButton, ListItemText } from '@mui/material';
import { DeviceData } from './entities/devices';
import { Location } from './entities/devices';
import { useNavigate } from 'react-router-dom';
import Pusher from 'pusher-js';
import { config } from '../../config/config';
import { addStatusDevice } from './entities/WssStatusDevices.store';

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
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const devicesStatusStore = useSelector((store: any) => store.wsStatus);
  const navigate = useNavigate();

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

  const dispatcher = useDispatch();

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

    const pusher = new Pusher(config.PUSHER_KEY, {
      cluster: config.PUSHER_CLUSTER,
    });

    const channel = pusher.subscribe('devices');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    channel.bind('status', (data: any) => {
      dispatcher(addStatusDevice(data?.device_id));
    });

    return () => {
      channel.unbind_all();
      channel.unsubscribe();
    };
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  const checkStatus = (deviceId: string): boolean => {
    const status = devicesStatusStore?.devicesByWs || [];
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    return status.some((el: any) => el == deviceId);
  };

  const navigateToDashboard = (
    deviceId: string
  ): MouseEventHandler<HTMLDivElement> | undefined => {
    navigate(`/dashboard/${deviceId}`);
    return;
  };

  return (
    <>
      <MenuAppBar />
      <div style={{ height: '600px', width: '80%' }}>
        {location && (
          <MapContainer
            zoom={13}
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
              <ListItemButton
                onClick={() => navigateToDashboard(device.device.sensor_id)}
                key={device.device.sensor_id}
                style={
                  checkStatus(device.device.sensor_id) ? blinkingStyle : {}
                }
              >
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
