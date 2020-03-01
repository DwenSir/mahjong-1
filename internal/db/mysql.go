package db

//UserInfo 用户信息
type UserInfo struct {
	OpenId string `json:"open_id" db:"open_id"`
}
