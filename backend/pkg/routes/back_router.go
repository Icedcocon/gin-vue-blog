package routes

import (
	"backend/pkg/config"
	"backend/pkg/routes/middleware"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func BackRouter() http.Handler {
	// 设置 gin 模式
	gin.SetMode(config.GlobalConfig.SERVER.Mode)

	// 创建路由
	r := gin.New()
	// 设置可信源 Origin 白名单，用于 CORS
	r.SetTrustedProxies([]string{"*"})

	// 使用本地文件上传, 需要静态文件服务, 使用七牛云不需要
	if config.GetConfig().UPLOAD.OssType == "local" {
		r.Static("/public", "./public")
		// 将 public 目录内的文件列举展示
		r.StaticFS("/dir", http.Dir("./public"))
	}

	// 日志中间件
	r.Use(middleware.Logger())
	// 自定义错误处理中间件
	r.Use(middleware.ErrorRecovery(false))
	// 跨域中间件
	r.Use(middleware.Cors())

	// 基于 cookie 存储 session
	store := cookie.NewStore([]byte(config.GlobalConfig.SESSION.Salt))
	// session 存储时间跟 JWT 过期时间一致
	store.Options(sessions.Options{MaxAge: config.GlobalConfig.SESSION.MaxAge * 3600})
	r.Use(sessions.Sessions(config.GlobalConfig.SESSION.Name, store)) // Session 中间件

	// 无需鉴权的接口
	base := r.Group("/api")
	{
		// TODO: 用户注册 和 后台登录 应该记录到 日志
		base.POST("/login", userAuthAPI.Login)   // 后台登录
		base.POST("/report", blogInfoAPI.Report) // 上报信息
	}

	return r
}
