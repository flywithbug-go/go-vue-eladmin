import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/views/layout/Layout'


Vue.use(Router)

export  const constantRouterMap = [
  {
    path: '/redirect',
    component: Layout,
    name: 'redirect',
    hidden: true,
    children: [
      {
        path: '/redirect/:path*',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/auth/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    name: 'auth-redirect',
    component: () => import('@/views/auth/authredirect'),
    hidden: true
  },
  {
    path: '/404',
    name: '404',
    component: () => import('@/views/errorPage/404'),
    hidden: true
  },
  {
    path: '/401',
    name: '401',
    component: () => import('@/views/errorPage/401'),
    hidden: true
  },
  {
    path: '',
    component: Layout,
    name: 'dashboard',
    redirect: 'dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'Dashboard',
        meta: { title: 'dashboard', icon: 'dashboard', noCache: true }
      }
    ]
  },
]


export default new Router({
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap,
  mode: 'history'
})

export const asyncRouterMap = [

]

