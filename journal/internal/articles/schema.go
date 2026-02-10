package articles

type ArticleCreateSchema struct {
	Title string `json:"title"`
	Body  string `json:"password"`
}

type ArticleEditSchema struct {
	ArticleID int64  `json:"articleId"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}
