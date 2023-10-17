package model

import (
	"reflect"
	"time"
)

// 权限控制 - 7 张表

// 角色
type Role struct {
	Universal
	Name      string `gorm:"type:varchar(20);comment:角色名" json:"name"`
	Label     string `gorm:"type:varchar(50);comment:角色描述" json:"label"`
	IsDisable int    `gorm:"type:tinyint(1);comment:是否禁用(0-否 1-是)" json:"is_disable"`
}

func (r *Role) IsEmpty() bool {
	return reflect.DeepEqual(r, &Role{})
}

// 角色-资源 关联
type RoleResource struct {
	RoleId     int `json:"role_id"`
	ResourceId int `json:"resource_id"`
}

// 角色-菜单 关联
type RoleMenu struct {
	RoleId int `json:"role_id"`
	MenuId int `json:"menu_id"`
}

// 用户账户信息
type UserAuth struct {
	Universal     `mapstructure:",squash"`
	UserInfoId    int       `gorm:"comment:用户信息ID" json:"user_info_id"`
	Username      string    `gorm:"type:varchar(50);comment:用户名" json:"username"`
	Password      string    `gorm:"type:varchar(100);comment:密码" json:"password"`
	LoginType     int       `gorm:"type:tinyint(1);comment:登录类型" json:"login_type"`
	IpAddress     string    `gorm:"type:varchar(20);comment:登录IP地址" json:"ip_address"`
	IpSource      string    `gorm:"type:varchar(50);comment:IP来源" json:"ip_source"`
	LastLoginTime time.Time `gorm:"comment:上次登录时间" json:"last_login_time"`
}

func (u *UserAuth) IsEmpty() bool {
	return reflect.DeepEqual(u, &UserAuth{})
}
