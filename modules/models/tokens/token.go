package models

type TokenType string

const (
	ID_TOKEN      TokenType = "id_token"
	ACCESS_TOKEN  TokenType = "access_token"
	REFRESH_TOKEN TokenType = "refresh_token"
)

type Tokens struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type DefaultClaim struct {
	Expired   int       `json:"exp"`
	NotBefore int       `json:"nbf"`
	IssuedAt  int       `json:"iat"`
	Issuer    string    `json:"iss"`
	Audience  string    `json:"aud"`
	JTI       string    `json:"jti"`
	Type      TokenType `json:"typ"`
}

type IdClaim struct {
	UserId string `json:"preffend_user_id"`
	Name   string `json:"preffend_name"`
	Email  string `json:"preffend_email"`
}
type AccessClaim struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
