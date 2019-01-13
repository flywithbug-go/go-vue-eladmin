
## Getting started

```bash
# clone the project 到go的 src目录下
git clone https://github.com/flywithbug-docmanager/docmanager.git doc-manager

启动本地Mongo服务
//go版本管理使用的是govendor 更新依赖
go run main.go

# install dependency
cd web_client
npm install

# develop
npm run dev



```

This will automatically open http://localhost:9527.

## Build

```bash
# build for test environment
npm run build:sit

# build for production environment
npm run build:prod
```

## Advanced

```bash
# --report to build with bundle size analytics
npm run build:prod --report

# --generate a bundle size analytics. default: bundle-report.html
npm run build:prod --generate_report

# --preview to start a server in local to preview
npm run build:prod --preview

# lint code
npm run lint

# auto fix
npm run lint -- --fix
```


#### 项目结构

``` 
src
├── App.vue
├── api
│   ├── app.js
│   ├── data.js
│   ├── permission.js
│   ├── role.js
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
│   ├── ThemePicker
│   │   └── index.vue
│   └── TreeTable
│       ├── eval.js
│       ├── index.vue
│       └── readme.md
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
│   └── svgo.yml
├── lang
│   ├── en.js
│   ├── index.js
│   └── zh.js
├── main.js
├── mixins
│   └── initData.js
├── permission.js
├── router
│   ├── index.js
│   └── modules
│       ├── metadata.js
│       └── system.js
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
├── vendor
│   ├── Export2Excel.js
│   └── Export2Zip.js
└── views
    ├── application
    │   ├── app.vue
    │   └── version.vue
    ├── dashboard
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
    └── system
        ├── permission
        │   ├── index.vue
        │   └── module
        │       ├── edit.vue
        │       ├── form.vue
        │       └── header.vue
        ├── role
        │   ├── index.vue
        │   └── module
        │       ├── edit.vue
        │       ├── form.vue
        │       └── header.vue
        └── user
            ├── center
            ├── center.vue
            ├── index.vue
            └── module
                ├── edit.vue
                ├── form.vue
                └── header.vue
2 [error opening dir]

58 directories, 164 files


```
