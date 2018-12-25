import Layout from '@/views/layout/Layout'


const metadataRouter = {
  path: '/meta',
  component: Layout,
  redirect: 'index',
  name: 'Metadata',
  meta: {
    title: 'metadata',
    icon: 'metadata'
  },
  children: [
    {
      path: 'app',
      component: () => import('@/views/application/app'),
      name: 'AppManager',
      meta: { title: 'appManager', noCache: true, icon: 'application' }
    },
    {
      path: 'version',
      component: () => import('@/views/application/version'),
      name: 'VersionManager',
      meta: { title: 'versionManager', noCache: true, icon: 'version'}
    },

  ]
}

export  default metadataRouter
