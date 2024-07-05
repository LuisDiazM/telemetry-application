from typing import Dict
from domain.measures.repositories.db_repository import IMeasuresDBRepo


class MeasuresUsecase:

    def __init__(self, measures_db: IMeasuresDBRepo) -> None:
        self.measures_db = measures_db

    def get_measures_by_device(self, sensor_id: str) -> Dict:
        measures = self.measures_db.get_measures_by_sensor(sensor_id)
        if len(measures) == 0:
            return {}

        metadata = measures[0].get("metadata")
        values = []
        for value in measures:
            val = value.get("values",{})
            data = {"timestamp": value.get("timestamp"), "temperature": val.get(
                "temperature"), "humidity": val.get("humidity")}
            values.append(data)
        return {"sensor_metadata": metadata, "values": values}
