package auth

type AppError struct {
	code    int
	message string
}

func (e AppError) Error() string {
	return e.message
}

func NewAppError(code int, message string) AppError {
	return AppError{code: code, message: message}
}

var ErrUserNotFound = NewAppError(404, "Пользователь не наден")
var ErrUserAlreadyExists = NewAppError(409, "Пользователь с таким логином уже есть")
var ErrCredentialIsIncorrect = NewAppError(401, "Неверный логин или пароль")
