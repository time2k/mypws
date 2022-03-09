import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/realtime/:devicename',
    name: 'RealtimeData',
    component: () => import('../views/RealtimeData.vue')
  },
  {
    path: '/history/:devicename',
    name: 'HistoryData',
    component: () => import('../views/HistoryData.vue')
  }
]

const router = new VueRouter({
  routes,
  base: process.env.NODE_ENV === 'production' ? '/web/' : '/webtest/',
  mode: 'history',
})

export default router
