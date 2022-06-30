package model

import (
	"go_bin_blog/utils/errmsg"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null " json:"username"`
	PassWord string `gorm:"type:varchar(20);not null " json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}

//查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

//新增用户
func CreateUser(data *User) int {
	//data.PassWord = ScryptPw(data.PassWord)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id,username,role,created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return users, 0
	}
	return users, total
}

//编辑用户信息
func UpdateUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.UserName
	maps["role"] = data.Role

	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// BeforeCreate 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.PassWord = ScryptPw(u.PassWord)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.PassWord = ScryptPw(u.PassWord)
	return nil
}

//密码加密

func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

//登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var Passworderr error

	db.Where("username = ?", username).First(&user)
	Passworderr = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if Passworderr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCESS

}
