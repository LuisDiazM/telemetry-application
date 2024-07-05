from abc import ABC, abstractmethod
from typing import Dict, List



class IMeasuresDBRepo(ABC):

    @abstractmethod
    def get_measures_by_sensor(self, sensor_id: str) -> List[Dict]:
        pass