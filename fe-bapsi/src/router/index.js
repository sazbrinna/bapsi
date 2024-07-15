import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/authenticated/HomeView.vue'

const routes = [
  {
    path: '/TestKalmal',
    name: 'TestKalmal',
    component: () => import('../views/TestKalimalang.vue')
  },
  {
    path: '/',
    name: 'LandingView',
    component: () => import('../views/LandingView.vue')
  },
  {
    path: '/login',
    name: 'LoginView',
    component: () => import('../views/LoginView.vue')
  },
  {
    path: '/registration',
    name: 'RegistrationView',
    component: () => import('../views/RegistrationView2.vue')
  },
  // -----------------------------------------------------------------
  // authenticated
  // -----------------------------------------------------------------
  {
    path: '/home',
    name: 'HomeView',
    component: HomeView
  },
  {
    path: '/profile',
    name: 'ProfileView',
    component: () => import('../views/authenticated/ProfileView.vue')
  },
  {
    path: '/profile/:id',
    name: 'ProfileOtherView',
    component: () => import('../views/authenticated/ProfileOtherView.vue')
  },
  {
    path: '/absen',
    name: 'AbsenView',
    component: () => import('../views/authenticated/AbsenView.vue')
  },

  // -----------------------------------------------------------------
  // Admin
  // -----------------------------------------------------------------

  {
    path: '/konfigurasi',
    name: 'KonfigurasiView',
    component: () => import('../views/admin/KonfigurasiView.vue')
  },
  {
    path: '/staffs',
    name: 'DaftarStaff',
    component: () => import('../views/admin/DaftarStaff.vue')
  },
  {
    path: '/invitation',
    name: 'InvitationView',
    component: () => import('../views/admin/InvitationView.vue')
  },
  {
    path: '/divisi/add',
    name: 'TambahDivisi',
    component: () => import('../views/admin/TambahDivisi.vue')
  },
  {
    path: '/loket/add',
    name: 'TambahLoket',
    component: () => import('../views/admin/TambahLoket.vue')
  },
  {
    path: '/absen/darurat',
    name: 'AbsenDaruratView',
    component: () => import('../views/admin/AbsenDaruratView.vue')
  },

  // -----------------------------------------------------------------
  // Wildcard route for 404 - Page not found
  // -----------------------------------------------------------------
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
