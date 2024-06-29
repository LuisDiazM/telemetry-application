type RegistryProps = {
  setRegistry: React.Dispatch<React.SetStateAction<boolean>>;
};
const RegistryComponent: React.FC<RegistryProps> = ({ setRegistry }) => {
  return (
    <>
      <div>Registry component</div>
      <button onClick={() => setRegistry(false)}> Enviar </button>
    </>
  );
};

export default RegistryComponent;
