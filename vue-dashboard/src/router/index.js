import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import DashboardLayout from '../components/DashboardLayout.vue'
import SystemOverview from '../views/SystemOverview.vue'
import IoTDevices from '../views/IoTDevices.vue'
import AddDevice from '../views/AddDevice.vue'
import AdminManagement from '../views/AdminManagement.vue'
import AddAdmin from '../views/AddAdmin.vue'
import { useAuthStore } from '../store/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: { public: true }
    },
    {
      path: '/',
      component: DashboardLayout,
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'overview',
          component: SystemOverview
        },
        {
          path: 'devices',
          name: 'devices',
          component: IoTDevices
        },
        {
          path: 'devices/provision',
          name: 'provision-device',
          component: AddDevice
        },
        {
          path: 'devices/configure/:id',
          name: 'configure-device',
          component: AddDevice
        },
        {
          path: 'admins',
          name: 'admins',
          component: AdminManagement
        },
        {
          path: 'admins/add',
          name: 'add-admin',
          component: AddAdmin
        },
        {
          path: 'admins/edit/:id',
          name: 'edit-admin',
          component: AddAdmin
        },
        {
          path: 'sectors',
          name: 'sectors',
          component: () => import('../views/SectorManagement.vue')
        },
        {
          path: 'sectors/:id',
          name: 'sector-details',
          component: () => import('../views/SectorDetails.vue') // New Sector Details View
        }
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const isAuth = authStore.isAuthenticated

  if (requiresAuth && !isAuth) {
    next('/login')
  } else if (to.name === 'login' && isAuth) {
    next('/')
  } else {
    next()
  }
})

export default router
