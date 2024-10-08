export const config = {
  ENVIRONMENT: import.meta.env.VITE_ENVIRONMENT,
  DOMAIN_AUTH0: import.meta.env.VITE_DOMAIN_AUTH0,
  CLIENT_ID_AUTH0: import.meta.env.VITE_CLIENT_ID_AUTH0,
  REDIRECT_URL_AUTH0: import.meta.env.VITE_REDIRECT_URL,
  PUSHER_KEY: import.meta.env.VITE_REACT_APP_PUSHER_KEY,
  PUSHER_CLUSTER: import.meta.env.VITE_REACT_APP_PUSHER_CLUSTER,
  SERVER_BFF: import.meta.env.VITE_REACT_SERVER_BFF
};
