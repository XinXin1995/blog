import lazy from '@/components/lazy'

const routes = [
  {
    path: '/',
    name: 'layout',
    component: lazy(() => import('@/layout')),
    childRoutes: [
      {
        path: '',
        name: 'home',
        component: lazy(() => import('@/pages/home'))
      },
      {
        path: '/detail/:id',
        name: 'detail',
        component: lazy(() => import('@/pages/detail'))
      },
      {
        path: '/archives',
        name: 'archives',
        component: lazy(() => import('@/pages/archives'))
      },
      {
        path: '/message',
        name: 'message',
        component: lazy(() => import('@/pages/message'))
      },
      {
        path: '/note',
        name: 'note',
        component: lazy(() => import('@/pages/note'))
      },
      {
        path: '/about',
        name: 'about',
        component: lazy(() => import('@/pages/about'))
      }
    ]
  }
]
export default routes
