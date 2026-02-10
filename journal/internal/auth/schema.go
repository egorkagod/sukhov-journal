package auth


type CredentialsSchema struct {
	Login string `json:"login"`
	Password string `json:"password"`
}