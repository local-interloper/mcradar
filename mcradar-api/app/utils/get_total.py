from psycopg import Cursor
import pypika
import pypika.functions
from pypika.queries import QueryBuilder


def get_total(cursor: Cursor, base_query: QueryBuilder) -> int:
    total_query = base_query.select(pypika.functions.Count("*"))
    cursor.execute(total_query.get_sql())
    total: int = cursor.fetchone()["count"]

    return total
