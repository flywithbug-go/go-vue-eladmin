import router from '@/router'
import store from '@/store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css'// progress bar style
import { getToken } from './utils/auth' // getToken from cookie

NProgress.configure({ showSpinner: false })// NProgress Configuration

//permission judge
// permission judge function
function hasPermission(role, permissionRoles) {
  if (role === 1) return true // admin permission passed directly
  if (permissionRoles == 0) return true
  return role <= permissionRoles
}


const whiteList = ['/login','/auth-redirect']

router.beforeEach((to, from ,next) => {
  NProgress.start()
  if (getToken()){  //判断是否登录
    if (to.path === '/login'){
      next({path:'/'})
      NProgress.done()
    }else {
      if (store.getters.role < 1) {
        store.dispatch('GetUserInfo').then(user => {
          store.dispatch('GenerateRoutes',user.role).then(() => {
            router.addRoutes(store.getters.addRouters) // 动态添加可访问路由表
            next()
          })
        }).catch((err) => {
          store.dispatch('FedLogOut').then(() => {
            Message.error(err || 'Verification failed, please login again')
            next({ path: '/' })
          })
        })
      } else {
        if (hasPermission(store.getters.role, to.meta.role)){
          console.log("hasPermission")
          next()
        } else {
          next({ path: '/401', replace: true, query: { noGoBack: true }})
        }
      }
    }
  }else {
    /* has no token*/
    if (whiteList.indexOf(to.path) !== -1) { // 在免登录白名单，直接进入
      next()
    } else {
      next(`/login?redirect=${to.path}`) // 否则全部重定向到登录页
      NProgress.done() // if current page is login will not trigger afterEach hook, so manually handle it
    }
  }
})



router.afterEach(() => {
  NProgress.done() // finish progress bar
})

