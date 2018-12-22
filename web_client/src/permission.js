import router from 'router'
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
  if (getToken()){  //判断是否登录
    if (to.name === 'login'){
      next({path:'/'})
      NProgress.done()
    }else {

    }
  }


})



router.afterEach(() => {
  NProgress.done() // finish progress bar
})

