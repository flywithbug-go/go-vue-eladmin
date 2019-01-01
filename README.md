
<p align="center">
  <a href="https://github.com/vuejs/vue">
    <img src="https://img.shields.io/badge/vue-2.5.17-brightgreen.svg" alt="vue">
  </a>
  <a href="https://github.com/ElemeFE/element">
    <img src="https://img.shields.io/badge/element--ui-2.4.6-brightgreen.svg" alt="element-ui">
  </a>
  <a href="https://travis-ci.org/PanJiaChen/vue-element-admin" rel="nofollow">
    <img src="https://github.com/PanJiaChen/vue-element-admin/blob/master/favicon.ico" alt="vue-element-admin" style="height="20px">
  </a>
</p>

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



# 项目布局

```
.
├── App.vue
├── api
│   ├── app.js
│   └── user.js
├── assets
│   ├── 401_images
│   │   └── 401.gif
│   ├── 404_images
│   │   ├── 404.png
│   │   └── 404_cloud.png
│   ├── custom-theme
│   │   ├── fonts
│   │   │   ├── element-icons.ttf
│   │   │   └── element-icons.woff
│   │   └── index.css
│   ├── image_placeholder.png
│   └── logo.png
├── components
│   ├── Breadcrumb
│   │   └── index.vue
│   ├── ErrorLog
│   │   └── index.vue
│   ├── FixedButton
│   │   └── index.vue
│   ├── Hamburger
│   │   └── index.vue
│   ├── HelloWorld.vue
│   ├── LangSelect
│   │   └── index.vue
│   ├── Pagination
│   │   └── index.vue
│   ├── Screenfull
│   │   └── index.vue
│   ├── ScrollPane
│   │   └── index.vue
│   ├── Share
│   │   └── dropdownMenu.vue
│   ├── SizeSelect
│   │   └── index.vue
│   ├── SvgIcon
│   │   └── index.vue
│   └── ThemePicker
│       └── index.vue
├── config
│   └── index.js
├── directive
│   ├── clipboard
│   │   ├── clipboard.js
│   │   └── index.js
│   ├── el-dragDialog
│   │   ├── drag.js
│   │   └── index.js
│   ├── permission
│   │   ├── index.js
│   │   └── permission.js
│   ├── sticky.js
│   └── waves
│       ├── index.js
│       ├── waves.css
│       └── waves.js
├── errorLog.js
├── filters
│   └── index.js
├── icons
│   ├── index.js
│   ├── svg
│   │   ├── 404.svg
│   │   ├── add.svg
│   │   ├── application.svg
│   │   ├── bug.svg
│   │   ├── chart.svg
│   │   ├── clipboard.svg
│   │   ├── component.svg
│   │   ├── dashboard.svg
│   │   ├── documentation.svg
│   │   ├── drag.svg
│   │   ├── edit.svg
│   │   ├── email.svg
│   │   ├── example.svg
│   │   ├── excel.svg
│   │   ├── eye.svg
│   │   ├── eye_open.svg
│   │   ├── form.svg
│   │   ├── guide\ 2.svg
│   │   ├── guide.svg
│   │   ├── icon.svg
│   │   ├── image_placeholder.svg
│   │   ├── international.svg
│   │   ├── language.svg
│   │   ├── link.svg
│   │   ├── list.svg
│   │   ├── lock.svg
│   │   ├── message.svg
│   │   ├── metadata.svg
│   │   ├── money.svg
│   │   ├── nested.svg
│   │   ├── password.svg
│   │   ├── people.svg
│   │   ├── peoples.svg
│   │   ├── qq.svg
│   │   ├── shopping.svg
│   │   ├── size.svg
│   │   ├── star.svg
│   │   ├── tab.svg
│   │   ├── table.svg
│   │   ├── theme.svg
│   │   ├── tree.svg
│   │   ├── user.svg
│   │   ├── version.svg
│   │   ├── wechat.svg
│   │   └── zip.svg
│   └── svgo.yml
├── lang
│   ├── en.js
│   ├── index.js
│   └── zh.js
├── main.js
├── permission.js
├── router
│   ├── index.js
│   └── modules
│       └── metadata.js
├── store
│   ├── getters.js
│   ├── index.js
│   └── modules
│       ├── app.js
│       ├── errorLog.js
│       ├── permission.js
│       ├── tagsView.js
│       └── user.js
├── styles
│   ├── btn.scss
│   ├── element-ui.scss
│   ├── index.scss
│   ├── mixin.scss
│   ├── sidebar.scss
│   ├── transition.scss
│   └── variables.scss
├── utils
│   ├── auth.js
│   ├── clipboard.js
│   ├── createUniqueString.js
│   ├── date.js
│   ├── fetch
│   │   └── index.js
│   ├── i18n.js
│   ├── index.js
│   ├── openWindow.js
│   ├── permission.js
│   ├── scrollTo.js
│   └── validate.js
└── views
    ├── application
    │   ├── app.vue
    │   └── version.vue
    ├── dashboard
    │   └── index.vue
    ├── documentation
    │   └── index.vue
    ├── errorPage
    │   ├── 401.vue
    │   └── 404.vue
    ├── layout
    │   ├── Layout.vue
    │   ├── components
    │   │   ├── AppMain.vue
    │   │   ├── Navbar.vue
    │   │   ├── Sidebar
    │   │   │   ├── FixiOSBug.js
    │   │   │   ├── Item.vue
    │   │   │   ├── Link.vue
    │   │   │   ├── SidebarItem.vue
    │   │   │   └── index.vue
    │   │   ├── TagsView.vue
    │   │   └── index.js
    │   └── mixin
    │       └── ResizeHandler.js
    ├── login
    │   ├── authredirect.vue
    │   └── index.vue
    ├── redirect
    │   └── index.vue
    ├── svg-icons
    │   ├── index.vue
    │   └── requireIcons.js
    └── user
        └── user.vue

49 directories, 139 files


``` 

