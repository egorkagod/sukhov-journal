package articles

type ArticleCreateSchema struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ArticleEditSchema struct {
	ArticleID uint64 `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}
