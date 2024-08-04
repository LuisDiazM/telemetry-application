import { useAuth0 } from '@auth0/auth0-react';
import { Box, Button, Typography } from '@mui/material';
import React from 'react';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};


const RegistryComponent = React.forwardRef(() => {
  const { loginWithRedirect } = useAuth0();

    return (
      <>
        <Box sx={style}>
          <Typography id='modal-modal-title' variant='h6' component='h2'>
            Registrarse
          </Typography>
          <Typography id='modal-modal-description' sx={{ mt: 2 }}>
            A continuación se registrará en modo trial usando un proveedor de
            identidad
          </Typography>
          <Button color='info' onClick={() => loginWithRedirect()}>
            {' '}
            Continuar
          </Button>
        </Box>
      </>
    );
});

export default RegistryComponent;
