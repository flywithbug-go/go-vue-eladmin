
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