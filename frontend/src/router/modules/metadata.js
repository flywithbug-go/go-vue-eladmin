import Layout from '@/views/layout/Layout'


const metadataRouter = {
    path: '',
    component: Layout,
    name: 'metadata',
    redirect: 'metadata',
    children: [
      {
        path: 'metadata',
        component: () => import('@/views/metadata/index'),
        name: 'Metadata',
        meta: { title: 'metadata', icon: 'meta_data', noCache: true }
      }
    ]
}
export  default metadataRouter
