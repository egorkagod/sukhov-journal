from contextlib import asynccontextmanager
from fastapi import FastAPI

from .api import voice_router


@asynccontextmanager
async def lifespan(app: FastAPI):
    print("TTS приложение запущено")

    yield

    print("TTS приложение остановлено")


app = FastAPI(lifespan=lifespan)


app.include_router(voice_router)
