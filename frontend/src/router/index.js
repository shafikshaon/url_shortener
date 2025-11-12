import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomePage.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginPage.vue'),
      meta: { guest: true }
    },
    {
      path: '/signup',
      name: 'signup',
      component: () => import('@/views/SignupPage.vue'),
      meta: { guest: true }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/views/DashboardPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/links',
      name: 'links',
      component: () => import('@/views/LinksPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/links/:id',
      name: 'link-details',
      component: () => import('@/views/LinkDetailsPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/analytics',
      name: 'analytics',
      component: () => import('@/views/AnalyticsPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('@/views/SettingsPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/notifications',
      name: 'notifications',
      component: () => import('@/views/NotificationsPage.vue'),
      meta: { requiresAuth: true }
    },
    // Account routes
    {
      path: '/account/profile',
      name: 'account-profile',
      component: () => import('@/views/account/ProfilePage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/account/security',
      name: 'account-security',
      component: () => import('@/views/account/SecurityPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/account/billing',
      name: 'account-billing',
      component: () => import('@/views/account/BillingPage.vue'),
      meta: { requiresAuth: true }
    },
    // Organization routes
    {
      path: '/organization/settings',
      name: 'organization-settings',
      component: () => import('@/views/organization/SettingsPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/organization/members',
      name: 'organization-members',
      component: () => import('@/views/organization/MembersPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/organization/domains',
      name: 'organization-domains',
      component: () => import('@/views/organization/DomainsPage.vue'),
      meta: { requiresAuth: true }
    },
    // Configuration routes
    {
      path: '/config/api',
      name: 'config-api',
      component: () => import('@/views/config/ApiKeysPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/config/webhooks',
      name: 'config-webhooks',
      component: () => import('@/views/config/WebhooksPage.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/config/integrations',
      name: 'config-integrations',
      component: () => import('@/views/config/IntegrationsPage.vue'),
      meta: { requiresAuth: true }
    }
  ]
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  const isAuthenticated = authStore.isLoggedIn

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else if (to.meta.guest && isAuthenticated) {
    next('/dashboard')
  } else {
    next()
  }
})

export default router
