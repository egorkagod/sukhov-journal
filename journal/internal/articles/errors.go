package articles

import "errors"

var NoPermissionsErr = errors.New("Недостаточно прав")
var ArticleNotFoundErr = errors.New("Статья не найдена")
