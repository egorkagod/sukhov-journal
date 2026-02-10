import torch
import uuid

from app.voice_models import tts_model_path


def text_to_speech(text, model_path=tts_model_path, speaker='eugene', sample_rate=48000, folder='media/generated') -> str | None:
    device = torch.device('cpu')
    torch.set_num_threads(4)

    model = torch.package.PackageImporter(model_path).load_pickle("tts_models", "model")
    model.to(device)

    path = f'{folder}/{uuid.uuid4()}.wav'

    try:
        audio_path = model.save_wav(
            text=text,
            speaker=speaker,
            sample_rate=sample_rate,
            audio_path=path,
        )
    except:
        audio_path = None
        
    return audio_path
