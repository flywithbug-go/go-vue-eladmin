import Layout from '@/views/layout/Layout'


const userRouter = {
  path: '/system',
  component: Layout,
  redirect: 'index',
  name:"systemManager",
  meta: {
    title: 'systemManager',
    icon: 'system'
  },
  children: [
    {
      path: 'user',
      name:"user",
      component: () => import('../../views/system/user/'),
      meta: { title: 'userManager', noCache: true, icon: 'user' }
    },
    {
      path: 'permission',
      name: 'permission',
      component: () => import('../../views/system/permission'),
      meta: { title: 'permissionManager', noCache: true, icon: 'permission' },
    },
    {
      path: 'role',
      name: 'role',
      component: () => import('../../views/system/role'),
      meta: { title: 'roleManager', noCache: true, icon: 'role' }
    },
  ]
}

export  default userRouter
