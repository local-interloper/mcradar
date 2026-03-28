from http.client import UNAUTHORIZED
import logging
import os

from fastapi import Request, Response
from starlette.middleware.base import BaseHTTPMiddleware

API_KEY = os.environ.get("API_KEY")


class AuthMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        if request.method == "OPTIONS" or request.url.path == "/api/health":
            return await call_next(request)

        token = request.headers.get("Authorization")

        if token is None:
            return Response(status_code=UNAUTHORIZED)

        token = token.removeprefix("Bearer ")

        if token != API_KEY:
            return Response(status_code=UNAUTHORIZED)

        return await call_next(request)
