package model

import "gorm.io/gorm"

type User struct {
	BaseModelNoDelete
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;uniqueIndex:name_delete_unique;uniqueIndex:mobile_delete_unique;comment:软删除"`
	Username  string         `json:"userName" gorm:"column:username;type:varchar(50);not null;uniqueIndex:name_delete_unique;comment:用户名"`
	Password  string         `json:"-" gorm:"column:password;type:varchar(100);not null;comment:密码"`
	NickName  string         `json:"nickName" gorm:"column:nick_name;type:varchar(100);default:'';comment:昵称"`
	Mobile    string         `json:"phone" gorm:"column:mobile;type:varchar(11);not null;uniqueIndex:mobile_delete_unique;comment:电话"`
	Email     string         `json:"email" gorm:"column:email;type:varchar(50);default:'';comment:邮箱"`
	HeaderImg string         `json:"headerImg" gorm:"column:header_img;type:varchar(100);default:'';comment:头像"`
	Roles     []Role         `json:"authorities" gorm:"many2many:user_roles;"`
}
