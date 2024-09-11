import './App.css';
import { RouterProvider } from 'react-router-dom';
import { route } from './routes/route';
import { Provider } from 'react-redux';
import { store } from './store';
function App() {
 
  return (
    <>
      <Provider store={store}>
        <RouterProvider router={route} />
      </Provider>
    </>
  );
}

export default App;
