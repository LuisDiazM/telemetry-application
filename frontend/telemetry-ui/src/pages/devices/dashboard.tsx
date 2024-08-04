import { Divider } from '@mui/material';
import MenuAppBar from '../../shared/nav-bar/navbar';
import Charts from './components/charts';

const DashboardPage = () => {




  const xAxisData: Array<Date> = [ new Date('2023-01-02'), new Date('2023-01-03'), new Date('2023-01-20')]
  const yAxisData: Array<number> = [20, 25, -5]
  const xLabel: string = "Tiempo"
  const yLabel: string = "°C"
  const title: string = "temperatura";

  return (
    <>
      <MenuAppBar></MenuAppBar>
      <div>
        <Charts title={title} xAxisData={xAxisData} xLabel={xLabel} yAxisData={yAxisData} yLabel={yLabel}></Charts>

        <div></div>
        <Divider></Divider>

      
      </div>
    </>
  );
};

export default DashboardPage;
