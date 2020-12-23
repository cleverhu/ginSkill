package UserService

import (
	"ginSkill/src/dbs"
	"ginSkill/src/models/UserModel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db *gorm.DB
)

func init() {
	db = dbs.Orm
}

//insert
func Insert(u *UserModel.UserModelImpl) bool {
	u.UpdateTime = time.Now()
	return db.Save(&u).RowsAffected == 1
}

//delete
func DeleteByID(id int) bool {
	u := UserModel.New(UserModel.WithUserID(id))
	return db.First(&u).RecordNotFound() == false && db.Delete(&u).RowsAffected == 1
}

//queryb
func QueryByID(id int) *UserModel.UserModelImpl {
	u := UserModel.New()
	db.First(&u, id)
	return u
}

//分页查询
func QueryPage(page, size int) []*UserModel.UserModelImpl {
	us := make([]*UserModel.UserModelImpl, 0)
	db.Limit(size).Offset(page*size - size).Find(&us)
	return us
}

//update 成功返回修改后的数据,失败返回nil
func Update(u *UserModel.UserModelImpl) *UserModel.UserModelImpl {
	if db.First(UserModel.New(), u.ID).RecordNotFound() == false {
		db.Table("t_user").Where("u_id = ?", u.ID).Update(&u)
		return u
	}
	return nil
}
