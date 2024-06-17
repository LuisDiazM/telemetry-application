

import './App.css'

function App() {
  const environment = import.meta.env.VITE_ENVIRONMENT;

  return (
    <>
      <h1>{environment}</h1>
    </>
  );
}

export default App
