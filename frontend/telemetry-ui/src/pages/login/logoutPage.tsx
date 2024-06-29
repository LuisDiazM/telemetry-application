import { useAuth0 } from '@auth0/auth0-react';

const LogoutPage = () => {
  const { logout } = useAuth0();

  logout({ logoutParams: { returnTo: window.location.origin } });
  localStorage.clear();
  return <h1>Logout sucessful!!</h1>;
};

export default LogoutPage;
