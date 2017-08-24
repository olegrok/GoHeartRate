package protocol

// AuthData is structure that contains user's login and password to send it to server
type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
