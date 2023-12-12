package model

import (
	"ginchat/http/api_param"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	BaseModel
	Name          string    `gorm:"coloum:name" json:"name"`
	Password      string    `gorm:"coloum:password" json:"password"`
	Phone         string    `gorm:"coloum:phome" json:"phome"`
	Email         string    `gorm:"coloum:email" json:"email"`
	Identity      string    `gorm:"coloum:identity" json:"identity"`
	ClientIp      string    `gorm:"coloum:client_ip" json:"client_ip"`
	ClientPort    string    `gorm:"coloum:client_port" json:"client_port"`
	LoginTime     time.Time `gorm:"coloum:login_time" json:"login_time"`
	HeartbeatTime time.Time `gorm:"coloum:heartbeat_time" json:"heartbeat_time"`
	LogoutTime    time.Time `gorm:"coloum:logout_time" json:"logout_time"`
	IsLogout      bool      `gorm:"coloum:is_logout" json:"is_logout"`
	DeviceInfo    string    `gorm:"coloum:device_info" json:"device_info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func CreateUserModelFactory(dbConn *gorm.DB) *UserBasic {
	if dbConn == nil {
		dbConn = UseDbConn()
	}
	return &UserBasic{BaseModel: BaseModel{DB: dbConn}}
}

func (db *UserBasic) List() (result []*api_param.ListUserResp, err error) {
	paramList := []interface{}{}
	sql := `SELECT
				t.id,
				t.name,
				t.password
			FROM
				user_basic AS t`
	err = db.Raw(sql, paramList...).Scan(&result).Error
	return
}

func (db *UserBasic) Create(req api_param.CreateUserReq) (err error) {
	sql := `INSERT INTO user_basic (
		name,
		password,
		phone
	) VALUES (?,?,?)`
	ex := db.Exec(sql, req.Name, req.Password, req.Phone)
	return ex.Error
}

func (db *UserBasic) Delete(id int64) (err error) {
	sql := `DELETE FROM tb_todo_list WHERE id = ?`

	ex := db.Exec(sql, id)

	return ex.Error
}

func (db *UserBasic) Update(req api_param.UpdateUserReq) (err error) {
	sql := `UPDATE user_basic 
	        SET 
			    phone =? ,
				email= ? ,
				password = ?,
				updated_at = now()
		    WHERE
			    id = ?`
	res := db.Exec(sql, req.Phone, req.Email, req.Password, req.Id)
	return res.Error
}

func (db *UserBasic) Login(req api_param.UserLoginReq) (user api_param.UserResp, err error) {
	sql := `SELECT
	            id ,
				name ,
				password ,
				email
			FROM 
				user_basic 
			WHERE 
				name= ?`
	res := db.Raw(sql, req.Name).First(&user)
	return user, res.Error
}

// token 验证
func (db *UserBasic) Token(id int64, token string) (err error) {
	sql := `UPDATE user_basic 
	        SET 
				identity = ?
		    WHERE
			    id = ?`
	res := db.Exec(sql, token, id)
	return res.Error
}

func (db *UserBasic) GetTokenUserInfo() {

}
