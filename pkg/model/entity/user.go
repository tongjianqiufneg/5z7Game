package entity

const (
	TableUser = "users"
)

// User 用户
type User struct {
	BaseModel
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Status int `json:"status" gorm:"column:status"`
}

// TableName 指定模型的表名称
func (User) TableName() string {
	return TableUser
}
