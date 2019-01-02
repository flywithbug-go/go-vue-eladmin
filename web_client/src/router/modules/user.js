import Layout from '@/views/layout/Layout'


const userRouter = {
  path: '/user',
  component: Layout,
  redirect: 'index',
  name: 'userManager',
  meta: {
    title: 'userManager',
    icon: 'user'
  },
  children: [
    {
      path: 'user',
      component: () => import('../../views/user/user'),
      name: 'userManager',
      meta: { title: 'userManager', noCache: true, icon: 'user' }
    },
    {
      path: 'user',
      component: () => import('../../views/user/user'),
      name: 'userManager',
      meta: { title: 'userManager', noCache: true, icon: 'user' }
    },
  ]
}

export  default userRouter
