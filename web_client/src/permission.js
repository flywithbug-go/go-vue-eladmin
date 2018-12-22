import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css'// progress bar style
import { getToken } from '@/utils/auth' // getToken from cookie

NProgress.configure({ showSpinner: false })// NProgress Configuration



//permission judge
const whiteList = ['/login','/auth-redirect']


router.beforeEach((to, from ,next) => {
  NProgress.start()
  console.log(getToken())
  if (getToken()){  //判断是否登录
    if (to.name === 'login'){
      next({path:'/'})
      NProgress.done()
    }else {
      next()
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

