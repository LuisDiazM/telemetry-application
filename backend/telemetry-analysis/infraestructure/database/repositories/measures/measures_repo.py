from datetime import datetime, timezone
from typing import Dict, List
from infraestructure.database.mongo_imp import MongoImp
from domain.measures.utils.time_operations import get_last_week
from domain.measures.repositories.db_repository import IMeasuresDBRepo

DATABASE = "devices-data"
COLLECTION = "measures"
LIMIT = 500
class MeasuresDBRepository(IMeasuresDBRepo):
    def __init__(self, mongo: MongoImp) -> None:
        self.mongo = mongo

    def get_measures_by_sensor(self, sensor_id: str) -> List[Dict]:
        timestamp = datetime.now(tz=timezone.utc)
        time_filter = get_last_week(timestamp)

        filter_query = {
            'metadata.sensor_id': sensor_id,
            'timestamp': {
                '$lte': time_filter.end,
                '$gte': time_filter.start
            }
        }
        options = {"limit":LIMIT, "sort":"timestamp"}
        return self.mongo.find_documents(database=DATABASE, collection=COLLECTION, filter=filter_query, options=options)
