import Layout from '@/views/layout/Layout'


const userRouter = {
  path: '/user',
  component: Layout,
  redirect: 'index',
  meta: {
    title: 'organizationStruct',
    icon: 'user'
  },
  children: [
    {
      path: 'user',
      component: () => import('../../views/user/user'),
      meta: { title: 'userManager', noCache: true, icon: 'user' }
    },
    {
      path: 'permission',
      component: () => import('../../views/user/user'),
      meta: { title: 'permissionManager', noCache: true, icon: 'user' }
    },
  ]
}

export  default userRouter
