import { createBrowserRouter } from 'react-router-dom';
import LoginPage from '../pages/login/loginPage';
import UserProfilePage from '../pages/user/UserProfilePage';
import LogoutPage from '../pages/login/logoutPage';
import DashboardPage from '../pages/devices/dashboard';
import DevicesPage from '../pages/devices/devices';

export const route = createBrowserRouter([
  { path: '/', Component: LoginPage },
  { path: '/login', Component: LoginPage },
  { path: '/logout', Component: LogoutPage },
  { path: '/profile', Component: UserProfilePage },
  { path: '/dashboard/:device_id', Component: DashboardPage },
  { path: '/devices', Component: DevicesPage },
]);
