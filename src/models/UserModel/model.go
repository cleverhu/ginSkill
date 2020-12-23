package UserModel

import "time"

//用户包含 用户ID 用户名  用户密码  用户手机  用户邮箱  入库时间 omitempty
type UserModelImpl struct {
	ID         int       `json:"uid,omitempty"  gorm:"column:u_id;primary_key"`
	Username   string    `binding:"UserName" json:"username,omitempty" gorm:"column:u_name" `
	Password   string    `json:"password" gorm:"column:u_password"`
	Telephone  string    `json:"tel,omitempty" gorm:"column:u_tel"`
	Email      string    `json:"email,omitempty" gorm:"column:u_email"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"column:u_update_time;type:datetime"`
}

func (UserModelImpl) TableName() string {
	return "t_user"
}

func New(attrs ...UserModelAttrFunc) *UserModelImpl {
	u := &UserModelImpl{}
	UserModelAttrFuncs(attrs).Apply(u) //强转 attrs 到 UserModelAttrFuncs
	return u
}

func (u *UserModelImpl) Mutate(attrs ...UserModelAttrFunc) *UserModelImpl {
	UserModelAttrFuncs(attrs).Apply(u)
	return u
}
