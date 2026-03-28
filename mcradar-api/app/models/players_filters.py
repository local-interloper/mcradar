from typing import List, Optional

from pydantic import BaseModel, ConfigDict
from pydantic.alias_generators import to_camel
from pypika import Criterion, Table
from app_types.server_type import ServerType
from pypika.queries import QueryBuilder


class PlayersFilters(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)

    username: Optional[str] = None

    def apply(self, query_builder: QueryBuilder) -> QueryBuilder:
        players_table = Table("players")

        if self.username is not None:
            query_builder = query_builder.where(
                players_table.name.ilike(f"%{self.username}%")
            )

        return query_builder
