import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.tsx';
import './index.css';
import { Auth0Provider } from '@auth0/auth0-react';
import { config } from './config/config.ts';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Auth0Provider
      domain={config.DOMAIN_AUTH0}
      clientId={config.CLIENT_ID_AUTH0}
      authorizationParams={{
        redirect_uri: config.REDIRECT_URL_AUTH0,
        audience: `https://${config.DOMAIN_AUTH0}/api/v2/`,
        scope: 'openid profile email update:current_user_metadata read:users read:clients read:client_credentials read:current_user',
      }}
    >
      <App />
    </Auth0Provider>
  </React.StrictMode>
);
