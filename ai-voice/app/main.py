from contextlib import asynccontextmanager
from fastapi import FastAPI

from .voice_models import stt_model_load
from .api import voice_router


@asynccontextmanager
async def lifespan(app: FastAPI):
    app.state.stt_model = stt_model_load()
    print("STT модель успешно загружена")

    print("Приложение запустилось")

    yield

    print("Остановка приложения")


app = FastAPI(lifespan=lifespan)


app.include_router(voice_router)