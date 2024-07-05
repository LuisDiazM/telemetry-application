from datetime import datetime, timedelta
from domain.measures.entities.time_filter import TimeFilter


def get_last_week(timestamp: datetime) -> TimeFilter:
    last_week = timedelta(days=7)
    start_date = timestamp - last_week
    return TimeFilter(end=timestamp, start=start_date)
