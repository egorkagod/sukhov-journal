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

var NoPermissionsErr = NewAppError(403, "Недостаточно прав")
var ArticleNotFoundErr = NewAppError(404, "Статья не найдена")
