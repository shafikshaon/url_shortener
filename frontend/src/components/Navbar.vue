<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <div class="container-fluid">
      <RouterLink class="navbar-brand" to="/dashboard">
        <div class="brand-icon">
          <i class="bi bi-link-45deg"></i>
        </div>
        <span class="brand-text">LinkShort</span>
      </RouterLink>

      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav me-auto">
          <li class="nav-item">
            <RouterLink class="nav-link" to="/dashboard">
              <i class="bi bi-grid-fill"></i>
              <span>Dashboard</span>
            </RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink class="nav-link" to="/links">
              <i class="bi bi-link-45deg"></i>
              <span>My Links</span>
            </RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink class="nav-link" to="/analytics">
              <i class="bi bi-graph-up"></i>
              <span>Analytics</span>
            </RouterLink>
          </li>
        </ul>

        <ul class="navbar-nav align-items-center">
          <li class="nav-item">
            <RouterLink to="/dashboard" class="btn btn-primary btn-sm create-btn">
              <i class="bi bi-plus-lg"></i>
              <span>Create Link</span>
            </RouterLink>
          </li>

          <!-- Notifications Dropdown -->
          <li class="nav-item dropdown">
            <a
              class="nav-link notification-icon"
              href="#"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              <i class="bi bi-bell-fill"></i>
              <span v-if="unreadNotifications > 0" class="notification-badge">{{ unreadNotifications }}</span>
            </a>
            <ul class="dropdown-menu dropdown-menu-end notification-dropdown">
              <li class="dropdown-header">
                <div class="d-flex justify-content-between align-items-center">
                  <strong>Notifications</strong>
                  <button class="btn btn-link btn-sm p-0" @click="markAllAsRead">Mark all as read</button>
                </div>
              </li>
              <li><hr class="dropdown-divider" /></li>
              <div class="notification-list">
                <li v-for="notification in notifications" :key="notification.id">
                  <a class="dropdown-item notification-item" :class="{ 'unread': !notification.read }" href="#">
                    <div class="notification-icon-wrapper" :class="notification.type">
                      <i :class="getNotificationIcon(notification.type)"></i>
                    </div>
                    <div class="notification-content">
                      <div class="notification-title">{{ notification.title }}</div>
                      <div class="notification-message">{{ notification.message }}</div>
                      <div class="notification-time">{{ notification.time }}</div>
                    </div>
                  </a>
                </li>
              </div>
              <li v-if="notifications.length === 0">
                <div class="dropdown-item text-center text-muted py-3">
                  <i class="bi bi-inbox"></i><br>
                  No notifications
                </div>
              </li>
              <li><hr class="dropdown-divider" /></li>
              <li>
                <RouterLink class="dropdown-item text-center" to="/notifications">
                  View all notifications
                </RouterLink>
              </li>
            </ul>
          </li>

          <!-- User Profile Dropdown -->
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle user-menu"
              href="#"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              <div class="user-avatar">
                <i class="bi bi-person-fill"></i>
              </div>
              <span class="user-email">{{ shortEmail }}</span>
            </a>
            <ul class="dropdown-menu dropdown-menu-end">
              <li class="dropdown-header">
                <div class="dropdown-user-info">
                  <div class="user-email-full">{{ userEmail }}</div>
                  <div class="user-plan-badge">{{ subscriptionPlan }} Plan</div>
                </div>
              </li>
              <li><hr class="dropdown-divider" /></li>

              <!-- Account Section -->
              <li class="dropdown-submenu-label">
                <span class="dropdown-item-text">
                  <i class="bi bi-person-circle"></i> Account
                </span>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/account/profile">
                  <i class="bi bi-person"></i> Profile
                </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/account/security">
                  <i class="bi bi-shield-lock"></i> Security
                </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/account/billing">
                  <i class="bi bi-credit-card"></i> Billing
                </RouterLink>
              </li>

              <li><hr class="dropdown-divider" /></li>

              <!-- Organization Section -->
              <li class="dropdown-submenu-label">
                <span class="dropdown-item-text">
                  <i class="bi bi-building"></i> Organization
                </span>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/organization/settings">
                  <i class="bi bi-gear"></i> Settings
                </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/organization/members">
                  <i class="bi bi-people"></i> Members
                </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/organization/domains">
                  <i class="bi bi-globe"></i> Custom Domains
                </RouterLink>
              </li>

              <li><hr class="dropdown-divider" /></li>

              <!-- Configurations Section -->
              <li class="dropdown-submenu-label">
                <span class="dropdown-item-text">
                  <i class="bi bi-sliders"></i> Configurations
                </span>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/config/api">
                  <i class="bi bi-code-square"></i> API Keys
                </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/config/webhooks">
                  <i class="bi bi-webhook"></i> Webhooks
                </RouterLink>
              </li>
              <li>
                <RouterLink class="dropdown-item dropdown-sub-item" to="/config/integrations">
                  <i class="bi bi-plug"></i> Integrations
                </RouterLink>
              </li>

              <li><hr class="dropdown-divider" /></li>
              <li>
                <a class="dropdown-item text-danger" href="#" @click.prevent="handleLogout">
                  <i class="bi bi-box-arrow-right"></i> Logout
                </a>
              </li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, computed } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const authStore = useAuthStore()

const userEmail = computed(() => authStore.currentUser?.email || 'User')

const subscriptionPlan = computed(() => {
  const tier = authStore.currentUser?.subscription_tier || 'free'
  return tier.charAt(0).toUpperCase() + tier.slice(1)
})

const shortEmail = computed(() => {
  const email = userEmail.value
  if (email.length > 20) {
    return email.substring(0, 20) + '...'
  }
  return email
})

// Sample notifications data
const notifications = ref([
  {
    id: 1,
    type: 'success',
    title: 'Link Created',
    message: 'Your short link "summer-promo" has been created successfully',
    time: '2 minutes ago',
    read: false
  },
  {
    id: 2,
    type: 'info',
    title: 'Link Performance',
    message: 'Your link "product-launch" reached 1,000 clicks',
    time: '1 hour ago',
    read: false
  },
  {
    id: 3,
    type: 'warning',
    title: 'Link Expiring Soon',
    message: 'Your link "limited-offer" will expire in 2 days',
    time: '3 hours ago',
    read: true
  },
  {
    id: 4,
    type: 'success',
    title: 'New Feature',
    message: 'Custom domains are now available for Pro users',
    time: '1 day ago',
    read: true
  },
  {
    id: 5,
    type: 'info',
    title: 'Usage Alert',
    message: 'You have used 80% of your monthly link quota',
    time: '2 days ago',
    read: true
  }
])

const unreadNotifications = computed(() => {
  return notifications.value.filter(n => !n.read).length
})

const getNotificationIcon = (type) => {
  const icons = {
    success: 'bi bi-check-circle-fill',
    info: 'bi bi-info-circle-fill',
    warning: 'bi bi-exclamation-triangle-fill',
    error: 'bi bi-x-circle-fill'
  }
  return icons[type] || 'bi bi-bell-fill'
}

const markAllAsRead = () => {
  notifications.value.forEach(n => n.read = true)
}

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.navbar {
  background-color: var(--bg-white);
  box-shadow: var(--shadow-sm);
  padding: 0;
  border-bottom: 1px solid var(--border-light);
  height: 72px;
  position: sticky;
  top: 0;
  z-index: 1000;
  backdrop-filter: blur(8px);
  background-color: rgba(255, 255, 255, 0.98);
}

.navbar .container-fluid {
  padding: 0 32px;
  height: 100%;
  display: flex;
  align-items: center;
  max-width: 1400px;
  margin: 0 auto;
}

/* Brand */
.navbar-brand {
  font-weight: 700;
  font-size: 20px;
  color: var(--text-primary);
  transition: all var(--transition-base);
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 0;
  margin-right: 48px;
  text-decoration: none;
}

.navbar-brand:hover {
  opacity: 0.85;
  transform: translateY(-1px);
}

.brand-icon {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  border-radius: var(--radius-lg);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
  box-shadow: var(--shadow-sm);
  transition: all var(--transition-base);
}

.navbar-brand:hover .brand-icon {
  box-shadow: var(--shadow-md);
  transform: scale(1.05);
}

.brand-text {
  font-weight: 700;
  letter-spacing: -0.5px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

/* Nav Links */
.navbar-nav {
  align-items: center;
  gap: 4px;
}

.nav-link {
  padding: 10px 16px;
  transition: all var(--transition-base);
  color: var(--text-secondary);
  font-weight: 500;
  border-radius: var(--radius-lg);
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 8px;
  text-decoration: none;
  position: relative;
}

.nav-link:hover {
  background-color: var(--bg-hover);
  color: var(--text-primary);
  transform: translateY(-1px);
}

.nav-link:active {
  transform: translateY(0);
  background-color: var(--bg-secondary);
}

.nav-link.router-link-active {
  background-color: var(--primary-subtle);
  color: var(--primary-color);
  font-weight: 600;
}

.nav-link.router-link-active:hover {
  background-color: var(--primary-subtle-hover);
}

.nav-link.router-link-active::after {
  content: '';
  position: absolute;
  bottom: -1px;
  left: 16px;
  right: 16px;
  height: 3px;
  background-color: var(--primary-color);
  border-radius: 3px 3px 0 0;
}

.nav-link i {
  font-size: 16px;
  transition: transform var(--transition-base);
}

.nav-link:hover i {
  transform: scale(1.1);
}

/* Create Button */
.create-btn {
  margin-left: 16px;
  margin-right: 12px;
  padding: 10px 20px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 6px;
  text-decoration: none;
  box-shadow: var(--shadow-sm);
}

.create-btn:hover {
  transform: translateY(-1px);
  box-shadow: var(--shadow-md);
}

.create-btn i {
  font-size: 14px;
}

/* User Menu */
.user-menu {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 6px 12px;
  border-radius: var(--radius-lg);
  transition: all var(--transition-base);
  color: var(--text-secondary);
  font-size: 14px;
  text-decoration: none;
  border: 1px solid var(--border-light);
  background-color: var(--bg-white);
  cursor: pointer;
}

.user-menu:hover {
  background-color: var(--bg-hover);
  border-color: var(--border-hover);
  box-shadow: var(--shadow-xs);
  transform: translateY(-1px);
}

.user-menu:active {
  transform: translateY(0);
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 16px;
  box-shadow: var(--shadow-xs);
  transition: all var(--transition-base);
}

.user-menu:hover .user-avatar {
  transform: scale(1.05);
  box-shadow: var(--shadow-sm);
}

.user-email {
  font-weight: 500;
  max-width: 150px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: var(--text-primary);
}

/* Notifications */
.notification-icon {
  position: relative;
  padding: 8px 12px;
  border-radius: 8px;
  transition: all 0.15s ease;
  color: var(--text-secondary);
  display: flex;
  align-items: center;
  font-size: 20px;
}

.notification-icon:hover {
  background-color: var(--bg-secondary);
  color: var(--text-primary);
}

.notification-badge {
  position: absolute;
  top: 4px;
  right: 6px;
  background-color: var(--danger-color);
  color: white;
  font-size: 10px;
  font-weight: 600;
  padding: 2px 5px;
  border-radius: 10px;
  min-width: 18px;
  text-align: center;
}

.notification-dropdown {
  min-width: 380px;
  max-height: 500px;
}

.notification-list {
  max-height: 350px;
  overflow-y: auto;
}

.notification-item {
  display: flex;
  gap: 12px;
  padding: 12px 16px;
  border-left: 3px solid transparent;
  transition: all 0.15s ease;
}

.notification-item.unread {
  background-color: var(--primary-subtle);
  border-left-color: var(--primary-color);
}

.notification-item:hover {
  background-color: var(--bg-secondary);
}

.notification-icon-wrapper {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  flex-shrink: 0;
}

.notification-icon-wrapper.success {
  background-color: rgba(16, 185, 129, 0.1);
  color: var(--success-color);
}

.notification-icon-wrapper.info {
  background-color: rgba(59, 130, 246, 0.1);
  color: #3B82F6;
}

.notification-icon-wrapper.warning {
  background-color: rgba(245, 158, 11, 0.1);
  color: var(--warning-color);
}

.notification-icon-wrapper.error {
  background-color: rgba(239, 68, 68, 0.1);
  color: var(--danger-color);
}

.notification-content {
  flex: 1;
  min-width: 0;
}

.notification-title {
  font-weight: 600;
  font-size: 14px;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.notification-message {
  font-size: 13px;
  color: var(--text-secondary);
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.notification-time {
  font-size: 12px;
  color: var(--text-tertiary);
  margin-top: 4px;
}

/* Dropdown */
.dropdown-menu {
  border: 1px solid var(--border-light);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-xl);
  padding: 8px;
  margin-top: 12px;
  min-width: 240px;
  background-color: var(--bg-white);
  animation: dropdownFadeIn var(--transition-base);
}

@keyframes dropdownFadeIn {
  from {
    opacity: 0;
    transform: translateY(-8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.dropdown-header {
  padding: 12px 16px;
  margin-bottom: 4px;
}

.dropdown-user-info {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.user-email-full {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-primary);
  word-break: break-all;
  line-height: 1.4;
}

.user-plan-badge {
  display: inline-block;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  color: white;
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  width: fit-content;
}

.dropdown-item {
  border-radius: var(--radius-md);
  padding: 10px 16px;
  transition: all var(--transition-base);
  color: var(--text-secondary);
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 500;
  text-decoration: none;
  cursor: pointer;
}

.dropdown-item:hover {
  background-color: var(--bg-hover);
  color: var(--text-primary);
  transform: translateX(2px);
}

.dropdown-item:active {
  background-color: var(--bg-secondary);
  color: var(--text-primary);
}

.dropdown-item.text-danger {
  color: var(--danger-color);
}

.dropdown-item.text-danger:hover {
  background-color: rgba(223, 27, 65, 0.08);
  color: var(--danger-color);
}

.dropdown-item i {
  font-size: 16px;
  width: 20px;
  text-align: center;
  transition: transform var(--transition-base);
}

.dropdown-item:hover i {
  transform: scale(1.1);
}

.dropdown-divider {
  margin: 8px 0;
  border-color: var(--border-light);
  opacity: 1;
}

.dropdown-submenu-label {
  padding: 0;
  margin: 8px 0 4px 0;
}

.dropdown-submenu-label .dropdown-item-text {
  font-size: 11px;
  font-weight: 700;
  color: var(--text-tertiary);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  padding: 8px 16px 4px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.dropdown-sub-item {
  padding: 8px 16px 8px 32px;
  font-size: 13px;
}

/* Navbar Toggler */
.navbar-toggler {
  border: 1px solid var(--border-color);
  padding: 8px 12px;
  border-radius: 6px;
}

.navbar-toggler:focus {
  box-shadow: 0 0 0 3px rgba(99, 91, 255, 0.1);
}

/* Mobile responsive */
@media (max-width: 991px) {
  .navbar {
    height: auto;
    min-height: 72px;
  }

  .navbar .container-fluid {
    padding: 16px 20px;
  }

  .navbar-brand {
    margin-right: 0;
  }

  .navbar-collapse {
    margin-top: 16px;
    padding-top: 16px;
    border-top: 1px solid var(--border-light);
  }

  .navbar-nav {
    align-items: flex-start;
    gap: 0;
    width: 100%;
  }

  .nav-item {
    width: 100%;
  }

  .nav-link {
    margin: 4px 0;
    width: 100%;
  }

  .create-btn {
    margin: 8px 0;
    width: 100%;
    justify-content: center;
  }

  .user-menu {
    width: 100%;
    justify-content: flex-start;
    margin: 4px 0;
  }

  .brand-text {
    font-size: 18px;
  }
}
</style>
