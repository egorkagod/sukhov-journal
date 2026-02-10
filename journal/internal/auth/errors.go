package auth

import "errors"

var ArticleNotFoundErr = errors.New("Статья не найдена")
var UserNotFoundErr = errors.New("Пользователь не наден")
var UserAlreadyExistsErr = errors.New("Пользователь с таким логином уже есть")
var CredentialsIsIncorrect = errors.New("Неверный логин или пароль")
