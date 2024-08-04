import { createBrowserRouter } from 'react-router-dom';
import LoginPage from '../pages/login/loginPage';
import UserProfilePage from '../pages/user/UserProfilePage';
import LogoutPage from '../pages/login/logoutPage';
import DashboardPage from '../pages/devices/dashboard';

export const route = createBrowserRouter([
  { path: '/', Component: LoginPage },
  { path: '/login', Component: LoginPage },
  { path: '/logout', Component: LogoutPage },
  { path: '/profile', Component: UserProfilePage },
  { path: '/dashboard', Component: DashboardPage },
]);
