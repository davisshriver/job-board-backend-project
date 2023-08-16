package outputs

type loginResponse struct {
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}
