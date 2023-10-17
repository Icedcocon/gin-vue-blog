## 项目介绍

本项目在以**博客**这个业务为主的前提下，复现完整的全栈项目代码（前台前端 + 后台前端 + 后端）

### 技术介绍

> 这里只写一些主流的通用技术，详细第三方库：前端参考 `package.json` 文件，后端参考 `go.mod` 文件

前端技术栈: 使用 pnpm 包管理工具

- 基于 TypeScript
- Vue3
- VueUse: 服务于 Vue Composition API 的工具集
- Unocss: 原子化 CSS
- Pinia
- Vue Router 
- Axios 
- Naive UI
- ...

后端技术栈:

- Golang
- Docker
- Gin
- GORM
- Viper: 支持 TOML (默认)、YAML 等常用格式作为配置文件
- Casbin
- Zap
- MySQL
- Redis
- Nginx: 部署静态资源 + 反向代理
- ...

其他:

- ...
