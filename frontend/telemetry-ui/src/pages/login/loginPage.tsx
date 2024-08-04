import { useAuth0 } from '@auth0/auth0-react';
import { useState } from 'react';
import RegistryComponent from '../../components/registry/registryComponent';
import Button from '@mui/material/Button';
import devices from './../../assets/images/devices.jpeg';
import {  Modal } from '@mui/material';


const LoginPage = () => {
  const { loginWithRedirect } = useAuth0();
  const [registry, setRegistry] = useState<boolean>(false);


  const handleClose = () => setRegistry(false);

  return (
    <>
    <div style={{textAlign:"center"}}>


      <div style={{display:"inline-grid", marginTop:"5%"}}>
        <h1> Aplicación de telemetría </h1>
        <img width={"340px"} src={devices}></img>


        <Button color='primary' onClick={() => loginWithRedirect()}>Ingresar</Button>
        <h5 style={{textAlign:"center", marginTop:"35px"}}>¿Aún no estás registrado?</h5>
        <Button color='secondary' onClick={() => setRegistry(true)}>Registrar</Button>

      </div>
    </div>

    <Modal
        open={registry}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
       <RegistryComponent></RegistryComponent>
      </Modal>
    </>
  );
};

export default LoginPage;
