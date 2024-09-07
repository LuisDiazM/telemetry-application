import { useEffect, useState } from "react";
import StatusComponent from "../../components/devices/status/statusComponent";
import Pusher from "pusher-js";
import { config } from "../../config/config";

const RealtimeDevices = () => {
    const [status, setStatus] = useState<boolean>(false)

  useEffect(() => {
    const pusher = new Pusher(config.PUSHER_KEY, {
      cluster: config.PUSHER_CLUSTER,
    });

    const channel = pusher.subscribe('devices');
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    channel.bind('status', (data: any) => {
      const statusWs = data?.status ?? false;
      console.log(data)
      setStatus(statusWs)
    });
    return () => {
      channel.unbind_all();
      channel.unsubscribe();
    };
  }, []);
  
  return (
    <div style={{ display: 'flex' }}>
    <StatusComponent status={status} device='Clear'></StatusComponent>
  </div>
  )
}

export default RealtimeDevices