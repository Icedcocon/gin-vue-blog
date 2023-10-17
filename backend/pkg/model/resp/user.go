package resp

import (
	"backend/pkg/model"
	"time"
)

// 后台列表 VO
type UserVO struct {
	ID            int          `json:"id"`
	UserInfoId    int          `json:"user_info_id"`
	Avatar        string       `json:"avatar"`
	Nickname      string       `json:"nickname"`
	Roles         []model.Role `json:"roles" gorm:"many2many:user_role;foreignKey:UserInfoId;joinForeignKey:UserId;"`
	LoginType     int          `json:"login_type"`
	IpAddress     string       `json:"ip_address"`
	IpSource      string       `json:"ip_source"`
	CreatedAt     time.Time    `json:"created_at"`
	LastLoginTime time.Time    `json:"last_login_time"`
	IsDisable     int          `json:"is_disable"`
	// Intro         string       `json:"intro"`
}

// 登录 VO
type LoginVO struct {
	ID         int `json:"id"`
	UserInfoId int `json:"user_info_id"`

	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`
	Website  string `json:"website"`

	IpAddress     string    `json:"ip_address"`
	IpSource      string    `json:"ip_source"`
	LastLoginTime time.Time `json:"last_login_time"`
	LoginType     int       `json:"login_type"`

	// 点赞 Set: 用于记录用户点赞过的文章, 评论
	ArticleLikeSet []string `json:"article_like_set"`
	CommentLikeSet []string `json:"comment_like_set"`
	// TalkLikeSet

	Token string `json:"token"`
}
