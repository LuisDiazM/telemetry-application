import { Card, CardContent, CircularProgress, Divider, Typography } from '@mui/material';
import MenuAppBar from '../../shared/nav-bar/navbar';
import Charts from './components/charts';
import { useParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import api from '../../services/apiService';
import { MeasuresData, Value } from './entities/measures';

const DashboardPage = () => {
  const { device_id } = useParams();

  const [measures, setMeasures] = useState<MeasuresData>();

  useEffect(() => {
    const getMeasures = async () => {
      const result = await api.get(`/measures/${device_id}`);
      if (result.data != null) {
        let data = result.data as MeasuresData;
        const values = data.values.sort(compareByTimestamp);
        data = { ...data, values: [...values] };
        setMeasures(data);
      }
    };
    getMeasures();
    return () => {};
  }, []);

  useEffect(() => {}, [setMeasures]);

  const compareByTimestamp = (a: Value, b: Value): number => {
    return new Date(a.timestamp).getTime() - new Date(b.timestamp).getTime();
  };

  return (
    <>
      <MenuAppBar></MenuAppBar>
      <div>
        <Card sx={{ minWidth: 275 }}>
          <CardContent>
            <Typography
              gutterBottom
              sx={{ color: 'text.primary', fontSize: 20 }}
            >
              Información del dispositivo
            </Typography>

            <Typography sx={{ color: 'text.secondary', mb: 1.5 }}>
              {measures?.sensor_metadata.sensor_id}
            </Typography>
            <Typography variant='body2'>
              {measures?.sensor_metadata.type_sensor}
              <br />
              {measures?.sensor_metadata.units}
            </Typography>
          </CardContent>
        </Card>
        {measures?.values ? (
          <Charts
            title={'Temperatura'}
            xAxisData={
              measures?.values.map((el) => new Date(el.timestamp)) || []
            }
            xLabel={'tiempo'}
            yAxisData={measures?.values.map((el) => el.temperature) || []}
            yLabel={'°C'}
          ></Charts>
        ) : (
          <CircularProgress />
        )}
        <Divider></Divider>
        {measures?.values ? (
          <Charts
            title={'Humedad Relativa'}
            xAxisData={
              measures?.values.map((el) => new Date(el.timestamp)) || []
            }
            xLabel={'tiempo'}
            yAxisData={measures?.values.map((el) => el.humidity) || []}
            yLabel={'%'}
          ></Charts>
        ) : (
          <CircularProgress/>
        )}
      </div>
    </>
  );
};

export default DashboardPage;
