from typing import Optional

from fastapi import APIRouter
from pydantic import BaseModel, ConfigDict
from models.paginated_data_response import PaginatedDataResponse
from psycopg.rows import dict_row
import pypika
from models.pagination import Pagination
from models.player import Player
from pydantic.alias_generators import to_camel
from db import DatabaseDep
from utils.get_total import get_total

players_router = APIRouter("/players")


class _Body(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)
    pagination: Pagination
    search: Optional[str] = None


@players_router.post("", response_model=None)
def users(body: _Body, db: DatabaseDep) -> PaginatedDataResponse[Player]:
    cursor = db.cursor(row_factory=dict_row)
    players = pypika.Table("players")
    base_query = pypika.Query.from_("players").select(
        "id, created_at, updated_at, name"
    )

    if body.search is not None:
        base_query = base_query.where(players.name.regex(body.search))

    total = get_total(cursor, base_query)

    paginated_query = body.pagination.apply(base_query)

    cursor.execute(paginated_query.get_sql())
    results = cursor.fetchall()

    data = [Player(**row) for row in results]

    return PaginatedDataResponse(data=data, total=total)
