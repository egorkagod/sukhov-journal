from fastapi import APIRouter, UploadFile, File, Request
from fastapi.responses import JSONResponse, FileResponse
from pydantic import BaseModel

from app.services.stt import speech_to_text
from app.services.tts import text_to_speech
from app.services.exceptions.stt import SttException


router = APIRouter(prefix='/voice')


@router.post("/stt")
async def stt_view(req: Request, file: UploadFile = File(required=True)):
    try:
        text = await speech_to_text(file, model=req.app.state.stt_model)
    except SttException as e:
        return JSONResponse({'text': None, 'message': str(e)}, status_code=200)
    except:
        return JSONResponse({'text': None}, status_code=500)
    
    return JSONResponse({'text': text}, status_code=200)


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