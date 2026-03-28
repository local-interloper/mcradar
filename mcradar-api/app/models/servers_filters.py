from typing import List, Optional

from pydantic import BaseModel, ConfigDict
from pydantic.alias_generators import to_camel
from pypika import Criterion, Table
from app_types.server_type import ServerType
from pypika.queries import QueryBuilder


class ServersFilters(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)

    version: Optional[str] = None
    types: Optional[List[ServerType]] = None

    def apply(self, query_builder: QueryBuilder) -> QueryBuilder:
        servers_table = Table("servers")

        if self.version is not None:
            query_builder = query_builder.where(
                servers_table.version.ilike(f"%{self.version}%")
            )

        if self.types is not None:
            query_builder = query_builder.where(
                Criterion.any(
                    [servers_table.type.like(target_type) for target_type in self.types]
                )
            )

        return query_builder
