from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from routers import main_router
import middleware

origins = ["*"]

app = FastAPI(docs_url=None, redoc_url=None, openapi_url=None)

app.add_middleware(middleware.AuthMiddleware)

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(main_router)
