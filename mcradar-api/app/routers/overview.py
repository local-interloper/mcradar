from fastapi import APIRouter
from psycopg.rows import dict_row
from pydantic import BaseModel, ConfigDict
import pypika.functions
import pypika
from pydantic.alias_generators import to_camel
from db import DatabaseDep

overivew_router = APIRouter(prefix="/overview")


class _Response(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)

    total_servers: int
    total_players: int


@overivew_router.get("")
def overview(db: DatabaseDep):
    cursor = db.cursor(row_factory=dict_row)

    results = {}

    query = pypika.Query.from_("servers").select(
        pypika.functions.Count("ip").as_("total_servers")
    )
    cursor.execute(query.get_sql())
    results = {**results, **cursor.fetchone()}

    query = pypika.Query.from_("players").select(
        pypika.functions.Count("id").as_("total_players")
    )
    cursor.execute(query.get_sql())
    results = {**results, **cursor.fetchone()}

    return _Response(**results)
