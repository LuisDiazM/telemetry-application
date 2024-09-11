import { createSlice } from '@reduxjs/toolkit';

export interface WsState {
  devicesByWs: Array<string>;
}

export const initalWsStates: WsState = {
  devicesByWs: [],
};

export const statusDevicesSlice = createSlice({
  name: 'websocket_status',
  initialState: {
    ...initalWsStates,
  },
  reducers: {
    addStatusDevice: (state, action) => {
      const currentDevices = [...state.devicesByWs];
      const deviceExists = currentDevices.some(
        (el) => el == action.payload
      );
      if (deviceExists) {
        return { ...state };
      } else {
        return { ...state, devicesByWs: [...state.devicesByWs, action.payload] };
      }
    },
    resetWsStates: () => {
      return {
        ...initalWsStates,
      };
    },
  },
});

export const { addStatusDevice, resetWsStates } = statusDevicesSlice.actions;
