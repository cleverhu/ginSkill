package UserModel

import "time"

type UserModelAttrFunc func(impl *UserModelImpl)
type UserModelAttrFuncs []UserModelAttrFunc

func WithUserID(id int) UserModelAttrFunc{
	return func(u *UserModelImpl) {
		u.ID = id
	}
}

func WithUsername(username string) UserModelAttrFunc{
	return func(u *UserModelImpl) {
		u.Username = username
	}
}

func WithPassword(pwd string) UserModelAttrFunc{
	return func(u *UserModelImpl) {
		u.Password = pwd
	}
}

func WithTelephone(tel string) UserModelAttrFunc{
	return func(u *UserModelImpl) {
		u.Telephone = tel
	}
}

func WithEmail(mail string) UserModelAttrFunc{
	return func(u *UserModelImpl) {
		u.Email = mail
	}
}

func WithUpdateTime(updateTime time.Time) UserModelAttrFunc{
	return func(u *UserModelImpl) {
		u.UpdateTime = updateTime
	}
}

func (this UserModelAttrFuncs) Apply(u *UserModelImpl) {
	for _, f := range this {
		f(u)
	}
}
