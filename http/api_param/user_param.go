package api_param

type ListUserResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateUserReq struct {
	Name       string `form:"name" json:"name"`
	Password   string `form:"password" json:"password"`
	Phone      string `form:"phome" json:"phome"`
	Email      string `form:"email" json:"email"`
	Identity   string `form:"identity" json:"identity"`
	ClientIp   string `form:"client_ip" json:"client_ip"`
	ClientPort string `form:"client_port" json:"client_port"`
}

type UpdateUserReq struct {
	Id       int64  `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Phone    string `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email" valid:"email"`
}

type UserLoginReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Identity string `form:"identity" json:"identity"`
}

type UserResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
