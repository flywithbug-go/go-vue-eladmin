import Layout from '@/views/layout/Layout'


const userRouter = {
  path: '/user',
  component: Layout,
  redirect: 'index',
  name:"organizationStruct",
  meta: {
    title: 'organizationStruct',
    icon: 'organization'
  },
  children: [
    {
      path: 'user',
      name:"user",
      component: () => import('../../views/organization/user'),
      meta: { title: 'userManager', noCache: true, icon: 'user' }
    },
    {
      path: 'permission',
      name: 'permission',
      component: () => import('../../views/organization/permission'),
      meta: { title: 'permissionManager', noCache: true, icon: 'permission' }
    },
    {
      path: 'role',
      name: 'role',
      component: () => import('../../views/organization/role'),
      meta: { title: 'roleManager', noCache: true, icon: 'role' }
    },
  ]
}

export  default userRouter
