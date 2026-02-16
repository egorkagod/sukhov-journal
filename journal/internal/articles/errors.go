package articles

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

var ErrNoPermision = NewAppError(403, "Недостаточно прав")
var ErrArticleNotFound = NewAppError(404, "Статья не найдена")
