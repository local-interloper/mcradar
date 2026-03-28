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
from models.players_filters import PlayersFilters

players_router = APIRouter(prefix="/players")


class _Body(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)

    pagination: Pagination
    filters: Optional[PlayersFilters] = None


@players_router.post("", response_model=None)
def users(body: _Body, db: DatabaseDep) -> PaginatedDataResponse[Player]:
    cursor = db.cursor(row_factory=dict_row)

    base_query = pypika.Query.from_("players")
    if body.filters is not None:
        base_query = body.filters.apply(base_query)

    total = get_total(cursor, base_query)

    ordered_query = base_query.orderby("updated_at", order=pypika.Order.desc)

    paginated_query = body.pagination.apply(
        ordered_query.select("id", "created_at", "updated_at", "name")
    )

    cursor.execute(paginated_query.get_sql())
    results = cursor.fetchall()

    data = [Player(**row) for row in results]

    return PaginatedDataResponse(data=data, total=total)
