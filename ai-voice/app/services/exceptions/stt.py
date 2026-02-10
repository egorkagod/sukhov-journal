class SttException(Exception):
    message = 'Ошибка'

    def __init__(self, message=None):
        self.message = message or self.message

    def __str__(self) -> str:
        return self.message


class ConversionError(SttException):
    message = 'Ошибка преобразования'


class BadExtensionError(SttException):
    message = 'Не поддерживающиеся расширение файла'


class NoFileError(SttException):
    message = 'Не найден сохраненный файл'


class NoModelError(SttException):
    message = 'Не подгружена модель'