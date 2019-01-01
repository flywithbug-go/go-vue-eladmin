.
├── LICENSE
├── README.md
├── Tree.md
├── build
│   ├── build.js
│   ├── check-versions.js
│   ├── logo.png
│   ├── utils.js
│   ├── vue-loader.conf.js
│   ├── webpack.base.conf.js
│   ├── webpack.dev.conf.js
│   └── webpack.prod.conf.js
├── config
│   ├── dev.env.js
│   ├── index.js
│   ├── prod.env.js
│   └── sit.env.js
├── dist
│   ├── favicon.ico
│   ├── index.html
│   └── static
│       ├── css
│       │   ├── app.eaa2858f.css
│       │   ├── chunk-35cf.15ac22f5.css
│       │   ├── chunk-3fb2.b803fe38.css
│       │   ├── chunk-401d.fb2cb1da.css
│       │   ├── chunk-4da9.8df2d56d.css
│       │   ├── chunk-5c89.9f9ca2e8.css
│       │   ├── chunk-626b.64862291.css
│       │   ├── chunk-7691.895f3241.css
│       │   ├── chunk-elementUI.c5f2f1e8.css
│       │   └── chunk-libs.6c5fa19c.css
│       ├── fonts
│       │   └── element-icons.6f0a763.ttf
│       ├── img
│       │   ├── 401.089007e.gif
│       │   └── 404.a57b6f3.png
│       ├── js
│       │   ├── 7zzA.acd9b7e4.js
│       │   ├── app.c688458b.js
│       │   ├── chunk-35cf.aad5aff3.js
│       │   ├── chunk-3fb2.d2ad125f.js
│       │   ├── chunk-401d.efd1960a.js
│       │   ├── chunk-4da9.56b533b0.js
│       │   ├── chunk-5c89.b3bbce99.js
│       │   ├── chunk-626b.2b8c8f6e.js
│       │   ├── chunk-7691.6e26e542.js
│       │   ├── chunk-elementUI.eda1a8bc.js
│       │   └── chunk-libs.4d0b8f2d.js
│       └── tinymce4.7.5
│           ├── langs
│           │   └── zh_CN.js
│           ├── plugins
│           │   ├── codesample
│           │   ├── emoticons
│           │   └── visualblocks
│           ├── skins
│           │   └── lightgray
│           └── tinymce.min.js
├── favicon.ico
├── index.html
├── package-lock.json
├── package.json
├── src
│   ├── App.vue
│   ├── api
│   │   ├── app.js
│   │   └── user.js
│   ├── assets
│   │   ├── 401_images
│   │   │   └── 401.gif
│   │   ├── 404_images
│   │   │   ├── 404.png
│   │   │   └── 404_cloud.png
│   │   ├── custom-theme
│   │   │   ├── fonts
│   │   │   │   ├── element-icons.ttf
│   │   │   │   └── element-icons.woff
│   │   │   └── index.css
│   │   ├── image_placeholder.png
│   │   └── logo.png
│   ├── components
│   │   ├── Breadcrumb
│   │   │   └── index.vue
│   │   ├── ErrorLog
│   │   │   └── index.vue
│   │   ├── FixedButton
│   │   │   └── index.vue
│   │   ├── Hamburger
│   │   │   └── index.vue
│   │   ├── HelloWorld.vue
│   │   ├── LangSelect
│   │   │   └── index.vue
│   │   ├── Pagination
│   │   │   └── index.vue
│   │   ├── Screenfull
│   │   │   └── index.vue
│   │   ├── ScrollPane
│   │   │   └── index.vue
│   │   ├── Share
│   │   │   └── dropdownMenu.vue
│   │   ├── SizeSelect
│   │   │   └── index.vue
│   │   ├── SvgIcon
│   │   │   └── index.vue
│   │   └── ThemePicker
│   │       └── index.vue
│   ├── config
│   │   └── index.js
│   ├── directive
│   │   ├── clipboard
│   │   │   ├── clipboard.js
│   │   │   └── index.js
│   │   ├── el-dragDialog
│   │   │   ├── drag.js
│   │   │   └── index.js
│   │   ├── permission
│   │   │   ├── index.js
│   │   │   └── permission.js
│   │   ├── sticky.js
│   │   └── waves
│   │       ├── index.js
│   │       ├── waves.css
│   │       └── waves.js
│   ├── errorLog.js
│   ├── filters
│   │   └── index.js
│   ├── icons
│   │   ├── index.js
│   │   ├── svg
│   │   │   ├── 404.svg
│   │   │   ├── add.svg
│   │   │   ├── application.svg
│   │   │   ├── bug.svg
│   │   │   ├── chart.svg
│   │   │   ├── clipboard.svg
│   │   │   ├── component.svg
│   │   │   ├── dashboard.svg
│   │   │   ├── documentation.svg
│   │   │   ├── drag.svg
│   │   │   ├── edit.svg
│   │   │   ├── email.svg
│   │   │   ├── example.svg
│   │   │   ├── excel.svg
│   │   │   ├── eye.svg
│   │   │   ├── eye_open.svg
│   │   │   ├── form.svg
│   │   │   ├── guide\ 2.svg
│   │   │   ├── guide.svg
│   │   │   ├── icon.svg
│   │   │   ├── image_placeholder.svg
│   │   │   ├── international.svg
│   │   │   ├── language.svg
│   │   │   ├── link.svg
│   │   │   ├── list.svg
│   │   │   ├── lock.svg
│   │   │   ├── message.svg
│   │   │   ├── metadata.svg
│   │   │   ├── money.svg
│   │   │   ├── nested.svg
│   │   │   ├── password.svg
│   │   │   ├── people.svg
│   │   │   ├── peoples.svg
│   │   │   ├── qq.svg
│   │   │   ├── shopping.svg
│   │   │   ├── size.svg
│   │   │   ├── star.svg
│   │   │   ├── tab.svg
│   │   │   ├── table.svg
│   │   │   ├── theme.svg
│   │   │   ├── tree.svg
│   │   │   ├── user.svg
│   │   │   ├── version.svg
│   │   │   ├── wechat.svg
│   │   │   └── zip.svg
│   │   └── svgo.yml
│   ├── lang
│   │   ├── en.js
│   │   ├── index.js
│   │   └── zh.js
│   ├── main.js
│   ├── permission.js
│   ├── router
│   │   ├── index.js
│   │   └── modules
│   │       └── metadata.js
│   ├── store
│   │   ├── getters.js
│   │   ├── index.js
│   │   └── modules
│   │       ├── app.js
│   │       ├── errorLog.js
│   │       ├── permission.js
│   │       ├── tagsView.js
│   │       └── user.js
│   ├── styles
│   │   ├── btn.scss
│   │   ├── element-ui.scss
│   │   ├── index.scss
│   │   ├── mixin.scss
│   │   ├── sidebar.scss
│   │   ├── transition.scss
│   │   └── variables.scss
│   ├── utils
│   │   ├── auth.js
│   │   ├── clipboard.js
│   │   ├── createUniqueString.js
│   │   ├── date.js
│   │   ├── fetch
│   │   │   └── index.js
│   │   ├── i18n.js
│   │   ├── index.js
│   │   ├── openWindow.js
│   │   ├── permission.js
│   │   ├── scrollTo.js
│   │   └── validate.js
│   └── views
│       ├── application
│       │   ├── app.vue
│       │   └── version.vue
│       ├── dashboard
│       │   └── index.vue
│       ├── documentation
│       │   └── index.vue
│       ├── errorPage
│       │   ├── 401.vue
│       │   └── 404.vue
│       ├── layout
│       │   ├── Layout.vue
│       │   ├── components
│       │   │   ├── AppMain.vue
│       │   │   ├── Navbar.vue
│       │   │   ├── Sidebar
│       │   │   ├── TagsView.vue
│       │   │   └── index.js
│       │   └── mixin
│       │       └── ResizeHandler.js
│       ├── login
│       │   ├── authredirect.vue
│       │   └── index.vue
│       ├── redirect
│       │   └── index.vue
│       ├── svg-icons
│       │   ├── index.vue
│       │   └── requireIcons.js
│       └── user
│           └── user.vue
└── static
    └── tinymce4.7.5
        ├── langs
        │   └── zh_CN.js
        ├── plugins
        │   ├── codesample
        │   │   └── css
        │   ├── emoticons
        │   │   └── img
        │   └── visualblocks
        │       └── css
        ├── skins
        │   └── lightgray
        │       ├── content.inline.min.css
        │       ├── content.min.css
        │       ├── fonts
        │       ├── img
        │       ├── skin.min.css
        │       └── skin.min.css.map
        └── tinymce.min.js

80 directories, 187 files
