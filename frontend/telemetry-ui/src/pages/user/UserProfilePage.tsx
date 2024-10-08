import { useAuth0 } from '@auth0/auth0-react';
import { useEffect } from 'react';
import MenuAppBar from '../../shared/nav-bar/navbar';

import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import api from '../../services/apiService';

const UserProfilePage = () => {
  const {
    user,
    isAuthenticated,
    isLoading,
    getAccessTokenSilently,
    getIdTokenClaims,
  } = useAuth0();
  useEffect(() => {
    const fetchUser = async () => {
      try {
        const result = await api.get('/user/me');
        const userData = result.data;
        if (userData == null) {
          const claims = await getIdTokenClaims();
          if (claims) {
            api.post('/user', { user_data: claims.__raw });
          }
        }
      } catch (error) {
        console.error(error);
      }
    };

    if (isAuthenticated) {
      getAccessTokenSilently().then((token) => {
        localStorage.setItem('access_token', token);
        fetchUser();
      });
    }
  }, [isAuthenticated, getAccessTokenSilently]);

  if (isLoading) {
    return <div>Loading ...</div>;
  }

  return (
    isAuthenticated && (
      <>
        <MenuAppBar></MenuAppBar>

        <div
          style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
          }}
        >
          <div style={{ marginTop: '5%' }}>
            <Card sx={{ maxWidth: 345 }}>
              <CardMedia
                sx={{ height: 120, width: 80 }}
                image={user?.picture}
                title={user?.family_name}
              />

              <CardContent>
                <Typography gutterBottom variant='h5' component='div'>
                  {user?.name}
                </Typography>
                <Typography variant='body2' color='text.secondary'>
                  {user?.email}
                </Typography>
              </CardContent>
            </Card>
          </div>
        </div>
      </>
    )
  );
};

export default UserProfilePage;
