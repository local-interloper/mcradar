from fastapi import APIRouter
from psycopg.rows import dict_row
from pydantic import BaseModel, ConfigDict
from models import Server
from typing import Optional
import pypika
from pydantic.alias_generators import to_camel
from models.pagination import Pagination
from models.servers_filters import ServersFilters
from models.paginated_data_response import PaginatedDataResponse
from utils.get_total import get_total
from db import DatabaseDep

servers_router = APIRouter(prefix="/servers")


class _Body(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)

    pagination: Pagination
    filters: Optional[ServersFilters] = None


@servers_router.post("", response_model=None)
def servers(body: _Body, db: DatabaseDep) -> PaginatedDataResponse[Server]:
    cursor = db.cursor(row_factory=dict_row)

    base_query = pypika.Query.from_("servers")
    if body.filters is not None:
        base_query = body.filters.apply(base_query)

    total = get_total(cursor, base_query)

    ordered_query = base_query.orderby("updated_at", order=pypika.Order.desc)

    paginated_query = body.pagination.apply(
        ordered_query.select(
            "ip",
            "created_at",
            "updated_at",
            "version",
            "type",
            "online_players",
            "max_players",
        )
    )
    cursor.execute(paginated_query.get_sql())
    results = cursor.fetchall()

    servers = [Server(**row) for row in results]

    return PaginatedDataResponse(data=servers, total=total)
