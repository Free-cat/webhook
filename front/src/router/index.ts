import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
      },
      {
        path: '/:uuid',
        name: 'Webhook',
        component: () => import('@/views/Webhook.vue')
      }
    ],
  },
]

const router = createRouter({
  history: createWebHashHistory('/#'),
  routes,
})

export default router
