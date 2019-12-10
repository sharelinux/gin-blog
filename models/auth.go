package models

type Auth struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 从DB检查用户和密码是否存在和一致
func CheckAuth(username, password string) bool  {
	var auth Auth
	db.Select("id").Where(Auth{Username:username, Password:password}).First(&auth)

	if auth.ID > 0 {
		return true
	}

	return false
}