package login

type (
	Response struct {
		Token string `json:"token"`
		Error string `json:"error"`
	}
)
