import { LineChart } from '@mui/x-charts/LineChart';
import { useEffect, useState } from 'react';


type ChartsProps = {
  xAxisData: Array<Date>;
  yAxisData: Array<number>;
  xLabel: string;
  yLabel: string;
  title: string;
};
const Charths: React.FC<ChartsProps> = ({xAxisData, xLabel, yLabel, title, yAxisData}) => {
  const [dimensions, setDimensions] = useState({ width: 0, height: 0 });

  useEffect(() => {
    const calculateDimensions = () => {
      const width = window.innerWidth * 0.6;
      const height = window.innerHeight * 0.5;
      setDimensions({ width, height });
    };
    calculateDimensions();
    window.addEventListener('resize', calculateDimensions);
    return () => window.removeEventListener('resize', calculateDimensions);
  }, []);

  return (
    <div>
      <LineChart
        xAxis={[
          {
            label: xLabel,
            data: [
             ...xAxisData
            ],
            scaleType:"time"
          },
        ]}
        series={[
          {
            data: [...yAxisData],
            label: yLabel,
          },
        ]}
        width={dimensions.width}
        height={dimensions.height}
        grid={{ vertical: true, horizontal: true }}
        title= {title}
      />
    </div>
  );
};

export default Charths;
