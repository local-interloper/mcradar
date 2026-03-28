from fastapi import APIRouter

from .overview import overivew_router
from .servers import servers_router
from .players import players_router
from .auth import auth_router
from .health import health_router

main_router = APIRouter(prefix="/api")

main_router.include_router(overivew_router)
main_router.include_router(servers_router)
main_router.include_router(players_router)
main_router.include_router(auth_router)
main_router.include_router(health_router)