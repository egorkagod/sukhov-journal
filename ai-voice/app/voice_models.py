# import os
# import torch
import whisper

tts_model_path = "silero-model.pt"


def stt_model_load():
    return whisper.load_model("small")


# def tts_model_load():
#     filename = 'silero-model.pt'
#     if not os.path.isfile(filename):
#         torch.hub.download_url_to_file('https://models.silero.ai/models/tts/ru/v5_ru.pt',
#                                     filename)  
