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
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle user-menu"
              href="#"
              role="button"
              data-bs-toggle="dropdown"
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
                </div>
              </li>
              <li><hr class="dropdown-divider" /></li>
              <li>
                <RouterLink class="dropdown-item" to="/settings">
                  <i class="bi bi-gear"></i> Settings
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
import { computed } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const authStore = useAuthStore()

const userEmail = computed(() => authStore.currentUser?.email || 'User')

const shortEmail = computed(() => {
  const email = userEmail.value
  if (email.length > 20) {
    return email.substring(0, 20) + '...'
  }
  return email
})

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
  transition: all 0.15s ease;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 0;
  margin-right: 48px;
  text-decoration: none;
}

.navbar-brand:hover {
  opacity: 0.8;
}

.brand-icon {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
  box-shadow: var(--shadow-sm);
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
  transition: all 0.15s ease;
  color: var(--text-secondary);
  font-weight: 500;
  border-radius: 8px;
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 8px;
  text-decoration: none;
  position: relative;
}

.nav-link:hover {
  background-color: var(--bg-secondary);
  color: var(--text-primary);
}

.nav-link.router-link-active {
  background-color: var(--primary-subtle);
  color: var(--primary-color);
  font-weight: 600;
}

.nav-link.router-link-active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 16px;
  right: 16px;
  height: 3px;
  background-color: var(--primary-color);
  border-radius: 3px 3px 0 0;
}

.nav-link i {
  font-size: 16px;
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
  border-radius: 8px;
  transition: all 0.15s ease;
  color: var(--text-secondary);
  font-size: 14px;
  text-decoration: none;
  border: 1px solid var(--border-light);
  background-color: var(--bg-white);
}

.user-menu:hover {
  background-color: var(--bg-secondary);
  border-color: var(--border-color);
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
}

.user-email {
  font-weight: 500;
  max-width: 150px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

/* Dropdown */
.dropdown-menu {
  border: 1px solid var(--border-light);
  border-radius: 10px;
  box-shadow: var(--shadow-lg);
  padding: 8px;
  margin-top: 12px;
  min-width: 240px;
}

.dropdown-header {
  padding: 12px 16px;
}

.dropdown-user-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.user-email-full {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-primary);
  word-break: break-all;
}

.dropdown-item {
  border-radius: 6px;
  padding: 10px 16px;
  transition: all 0.15s ease;
  color: var(--text-secondary);
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 500;
  text-decoration: none;
}

.dropdown-item:hover {
  background-color: var(--bg-secondary);
  color: var(--text-primary);
}

.dropdown-item:active {
  background-color: var(--bg-tertiary);
  color: var(--text-primary);
}

.dropdown-item.text-danger {
  color: var(--danger-color);
}

.dropdown-item.text-danger:hover {
  background-color: rgba(223, 27, 65, 0.1);
  color: var(--danger-color);
}

.dropdown-item i {
  font-size: 16px;
  width: 20px;
  text-align: center;
}

.dropdown-divider {
  margin: 8px 0;
  border-color: var(--border-light);
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
