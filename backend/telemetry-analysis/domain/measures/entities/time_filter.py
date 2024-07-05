from pydantic import BaseModel
from datetime import datetime

class TimeFilter(BaseModel):
    start: datetime
    end: datetime