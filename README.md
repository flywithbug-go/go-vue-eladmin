
## 简介
使用PanJiaChen的:[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin/blob/master/README.zh-CN.md)搭好的基础框架


Backend(go):
Frontend(vue)

- **go主要框架** 
``` 
  github.com/gin-gonic/gin
  github.com/dgrijalva/jwt-go
  github.com/flywithbug/log4go
  gopkg.in/mgo.v2
```
  
  
- **Vue主要框架**   
 ``` 
  vue-router
  vuex
  axios
  vue-i18n
  element-ui
 ```

``` bash
# govendor add  dependencies
go run main.go

# serve with hot reload at localhost:6201
open  http://localhost:6203 

cd web_client

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

-------------------------------
## 功能说明 (vue & go)
- [x] 登录 / 注销

- [x] 权限验证
  - [x] 页面权限  

- [x] 全局功能
  - [x] 国际化多语言
  - [x] 动态侧边栏（支持多级路由嵌套）
  - [x] 快捷导航(标签页)
  - [x] ScreenFull全屏
  - [x] 自适应收缩侧边栏


- [x] 错误页面 
  - [x] 401 
  - [x] 404 

- [ ] 数据管理 
  - [x] App管理
  - [ ] 版本管理-doing

- [ ] 开发工具
  - [ ] 数据模型 
  - [ ] API管理
  
  
  
### 页面示例
![frontend](frontend/10EF2717-74AB-4175-8FFF-324D7A8204E7.png)   



### Server项目布局

```
.
├── common
│   ├── com_definition.go
│   ├── common.go
│   ├── compare.go
│   └── compare_test.go
├── config
│   └── config.go
├── core
│   ├── errors
│   │   ├── errors.go
│   │   ├── errors_test.go
│   │   ├── reporter.go
│   │   └── reporter_test.go
│   ├── jwt
│   │   └── jwt.go
│   └── mongo
│       ├── Increment.go
│       ├── db.go
│       └── db_test.go
├── key_sources
│   ├── private_key
│   └── public_key.pub
├── model
│   ├── app_version.go
│   ├── application.go
│   ├── login.go
│   ├── model_func.go
│   ├── model_test.go
│   ├── response.go
│   ├── role.go
│   ├── user.go
│   └── user_role.go
└── server
    ├── handler
    │   ├── app_handler.go
    │   ├── app_version_handler.go
    │   ├── file_handler.go
    │   ├── html_handler.go
    │   ├── index.go
    │   ├── para_model.go
    │   ├── router.go
    │   └── user_handler.go
    ├── middleware
    │   ├── authentication.go
    │   ├── cookie.go
    │   └── logger.go
    ├── server.go
    └── web_server.go

11 directories, 37 files

``` 

