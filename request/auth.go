package request

type AuthRequest struct {
	SecretId  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
}
