import { configureStore } from '@reduxjs/toolkit';
import { statusDevicesSlice } from './pages/devices/entities/WssStatusDevices.store';

export const store = configureStore({
  reducer: {
    wsStatus: statusDevicesSlice.reducer,
  },
});
