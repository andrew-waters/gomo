package moltin

// AuthResponse contains the repsonse from the auth call
type AuthResponse struct {
	Expires     int    `json:"expires"`
	ExpiresIn   int    `json:"expires_in"`
	Identifier  string `json:"identifier"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}
