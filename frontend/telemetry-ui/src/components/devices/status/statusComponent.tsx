import { Badge, Box } from '@mui/material';
import React from 'react';

type StatusProps = {
  status: boolean;
  device: string;
};
const StatusComponent: React.FC<StatusProps> = ({ status, device }) => {
  const shapeStyles = { bgcolor: 'primary.main', width: 40, height: 40 };
  const shapeCircleStyles = { borderRadius: '50%' };
  const maxLength = 20;
  let deviceName = device;
  if (device.length > maxLength) {
    deviceName = `${device.slice(0, 20)}...`;
  }
  const circle = (
    <Box component='span' sx={{ ...shapeStyles, ...shapeCircleStyles }} />
  );

 

  return (
    <div style={{ marginLeft: '10px' }}>
      <p>{deviceName}</p>
      {status ? (
        <Badge color='secondary' overlap='circular' badgeContent=' '>
          {circle}
        </Badge>
      ) : (
        <Badge>{circle}</Badge>
      )}
    </div>
  );
};

export default StatusComponent;
