package UserModel

import "time"

//用户包含 用户ID 用户名  用户密码  用户手机  用户邮箱  入库时间
type UserModelImpl struct {
	ID         int
	Username   string
	Password   string `json:"-"`
	Telephone  string
	Email      string
	UpdateTime time.Time
}

func New(attrs ...UserModelAttrFunc) *UserModelImpl {
	u := &UserModelImpl{}
	UserModelAttrFuncs(attrs).Apply(u)
	return u
}

func (u *UserModelImpl) Mutate(attrs ...UserModelAttrFunc) *UserModelImpl {
	UserModelAttrFuncs(attrs).Apply(u)
	return u
}
