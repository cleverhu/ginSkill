package UserService

import (
	"fmt"
	"ginSkill/src/models/UserModel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

var (
	db *gorm.DB
)

func init() {
	err := fmt.Errorf("")
	db, err = gorm.Open("mysql", "root:123456@tcp(101.132.138.205:3306)/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	logger := logrus.New()
	db.SetLogger(logger)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
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
