from http.client import OK

from fastapi import APIRouter, Response


auth_router = APIRouter(prefix="/auth")

@auth_router.get("")
def auth() -> Response:
    return Response(status_code=OK)