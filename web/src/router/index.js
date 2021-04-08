import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    redirect: { name: 'layout' },
    component: () => import('@/views/redirect'),
    children: [
      {
        path: '/',
        name: 'layout',
        component: () => import('@/views/Layout/'),
        redirect: { name: 'personal' },
        children: [
          {
            path: 'personal',
            name: 'personal',
            component: () => import('@/views/personal/')
          },
          {
            path: 'article',
            name: 'article',
            component: () => import('@/views/redirect'),
            redirect: { name: 'article.index' },
            children: [
              {
                path: 'index',
                name: 'article.index',
                component: () => import('@/views/article/')
              },
              {
                path: 'detail/:id',
                name: 'article.detail',
                component: () => import('@/views/article/detail')
              }
            ]
          },
          {
            path: 'example',
            name: 'example',
            component: () => import('@/views/redirect'),
            redirect: { name: 'example.index' },
            children: [
              {
                path: 'index',
                name: 'example.index',
                component: () => import('@/views/example/')
              },
              {
                path: 'spusku',
                name: 'example.spusku',
                component: () => import('@/views/examples/spusku/')
              }
            ]
          },
          {
            path: 'project',
            name: 'project',
            component: () => import('@/views/project/')
          },
          {
            name: 'manager',
            path: 'manager',
            component: () => import('@/views/redirect'),
            redirect: { name: 'manager.article' },
            children: [
              {
                name: 'manager.article',
                path: 'article',
                component: () => import('@/views/redirect'),
                redirect: { name: 'manager.article.list' },
                children: [
                  {
                    name: 'manager.article.list',
                    path: 'list',
                    component: () => import('@/views/manager/article/list')
                  },
                  {
                    name: 'manager.article.add',
                    path: 'add',
                    component: () => import('@/views/manager/article/add')
                  },
                  {
                    name: 'manager.article.update',
                    path: 'update/:id',
                    component: () => import('@/views/manager/article/add')
                  }
                ]
              },
              {
                name: 'manager.category',
                path: 'category',
                component: () => import('@/views/manager/category/')
              },
              {
                name: 'manager.tag',
                path: 'tag',
                component: () => import('@/views/manager/tag/')
              }
            ]
          }
        ]
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
