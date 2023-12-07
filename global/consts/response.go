package consts

type response struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
}

var (
	CurdOK = response{20000, "success"}

	ServerError     = response{40000, "server error"}
	ParamsCheckFail = response{40001, "invalid params"}

	UserNotList      = response{40100, "user table not exist"}
	UserRegisterFail = response{40101, "register fail"}
	UserDeleteFail   = response{40102, "this user does not exist"}
	UserLoginFail    = response{40103, "user login fail"}
	PasswordFail     = response{40104, "password error"}
)
