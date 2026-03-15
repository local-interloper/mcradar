from typing import Optional

from pydantic import BaseModel, ConfigDict
from pydantic.alias_generators import to_camel
from pypika import Table
from app_types.server_type import ServerType
from pypika.queries import QueryBuilder


class Filters(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)

    version: Optional[str] = None
    type: Optional[ServerType] = None

    def apply(self, query_builder: QueryBuilder) -> QueryBuilder:
        servers_table = Table("servers")

        if self.version is not None:
            query_builder = query_builder.where(
                servers_table.version.like(f"%{self.version}%")
            )

        if self.type is not None:
            query_builder = query_builder.where(servers_table.type.like(self.type))

        return query_builder
