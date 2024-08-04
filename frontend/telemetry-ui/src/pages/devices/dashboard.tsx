import { useEffect, useState } from 'react';
import StatusComponent from '../../components/devices/status/statusComponent';
import MenuAppBar from '../../shared/nav-bar/navbar';
import Pusher from 'pusher-js';
import { config } from '../../config/config';

const DashboardPage = () => {
  const [status, setStatus] = useState<boolean>(false)

  useEffect(() => {
    const pusher = new Pusher(config.PUSHER_KEY, {
      cluster: config.PUSHER_CLUSTER,
    });

    const channel = pusher.subscribe('devices');
    channel.bind('status', (data: any) => {
      const statusWs = data?.status ?? false;
      setStatus(statusWs)
    });
    return () => {
      channel.unbind_all();
      channel.unsubscribe();
    };
  }, []);

  return (
    <>
      <MenuAppBar></MenuAppBar>
      <div style={{ display: 'flex' }}>
        <StatusComponent status={status} device='Clear'></StatusComponent>
      </div>
    </>
  );
};

export default DashboardPage;
