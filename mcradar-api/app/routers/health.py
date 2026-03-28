from http.client import OK

from fastapi import APIRouter, Response


health_router = APIRouter(prefix="/health")

@health_router.get("")
def health() -> Response:
    return Response(status_code=OK)