from pydantic import BaseModel, ConfigDict
from pydantic.alias_generators import to_camel
from datetime import datetime


class Server(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)

    ip: str
    created_at: datetime
    updated_at: datetime
    version: str
    type: str
    online_players: int
    max_players: int
