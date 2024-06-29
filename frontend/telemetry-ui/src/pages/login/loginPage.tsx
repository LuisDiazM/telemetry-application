import { useAuth0 } from '@auth0/auth0-react';
import { useState } from 'react';
import RegistryComponent from '../../components/registry/registryComponent';

const LoginPage = () => {
  const { loginWithRedirect } = useAuth0();
  const [registry, setRegistry] = useState<boolean>(false);

  return (
    <>
      <button onClick={() => loginWithRedirect()}>Ingresar</button>
      <h5>¿Aún no estás registrado?</h5>
      <button onClick={() => setRegistry(true)}>Registrar</button>

      {registry && (
        <RegistryComponent setRegistry={setRegistry}></RegistryComponent>
      )}
    </>
  );
};

export default LoginPage;
