from fastapi import APIRouter
from fastapi.responses import FileResponse
from pydantic import BaseModel

from app.services.tts import text_to_speech


router = APIRouter(prefix='/voice')


class TtsSchema(BaseModel):
    text: str
    

@router.post("/tts")
def tts_view(dto: TtsSchema):
    audio_path = text_to_speech(dto.text)
    return FileResponse(
        path=audio_path,
        filename='audio.wav',
        media_type='audio/wav',
    )
