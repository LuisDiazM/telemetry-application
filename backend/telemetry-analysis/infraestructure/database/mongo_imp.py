from typing import Dict, List
from pymongo import MongoClient, DESCENDING
import os



class MongoImp:    
    client: MongoClient

    def __init__(self):
        url = os.getenv("MONGO_URL")
        self.client = MongoClient(url)

    def find_documents(self, database: str, collection: str, filter: Dict, options:Dict={}) -> List[Dict]:
        try:
            if not isinstance(database, str) or not isinstance(collection, str) or not isinstance(filter, dict):
                raise TypeError("Input parameters must be of type str for database and collection, and type dict for filter.")
            col = self.client.get_database(database).get_collection(collection)
            if bool(options):
                limit = options.get("limit",500)
                sort_field = options.get("sort","timestamp")
                cursor = col.find(filter).limit(limit).sort(sort_field, DESCENDING)
            else:
                cursor = col.find(filter)
            documents = [doc for doc in cursor]
            return documents
        except Exception as e:
            print(f"An error occurred: {e}")
            return []
    