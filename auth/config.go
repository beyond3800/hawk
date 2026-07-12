package auth

type Config struct {
	SecretKey string `json:"secret_key"`
	Issuer	string `json:"issuer"`
}