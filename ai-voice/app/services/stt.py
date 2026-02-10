import uuid

from .exceptions.stt import BadExtensionError, NoModelError


extensions = ['mp3', 'wav', 'm4a']

async def save_audio(file, path='media/get') -> str | None:
    extension = file.filename.split('.')[-1]
    if extension not in extensions:
        raise BadExtensionError
    
    file_path = f'{path}/{uuid.uuid4()}.{extension}'
    with open(file_path, 'wb') as audio:
        audio.write(await file.read())
    return file_path


async def speech_to_text(file, model) -> str | None:
    if not model:
        raise NoModelError

    try:
        file_path = await save_audio(file)
    except Exception as e:
        raise e
    
    result = model.transcribe(file_path)
    return result['text']
