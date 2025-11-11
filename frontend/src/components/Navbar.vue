<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <div class="container-fluid">
      <RouterLink class="navbar-brand" to="/dashboard">
        <i class="bi bi-link-45deg"></i> URL Shortener
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
              <i class="bi bi-speedometer2"></i> Dashboard
            </RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink class="nav-link" to="/links">
              <i class="bi bi-link"></i> Links
            </RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink class="nav-link" to="/analytics">
              <i class="bi bi-bar-chart"></i> Analytics
            </RouterLink>
          </li>
        </ul>

        <ul class="navbar-nav">
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              role="button"
              data-bs-toggle="dropdown"
            >
              <i class="bi bi-person-circle"></i> {{ userEmail }}
            </a>
            <ul class="dropdown-menu dropdown-menu-end">
              <li>
                <RouterLink class="dropdown-item" to="/settings">
                  <i class="bi bi-gear"></i> Settings
                </RouterLink>
              </li>
              <li><hr class="dropdown-divider" /></li>
              <li>
                <a class="dropdown-item" href="#" @click.prevent="handleLogout">
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

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.navbar {
  background-color: var(--bg-white);
  box-shadow: 0 1px 0 rgba(10, 37, 64, 0.1);
  padding: 0;
  border-bottom: 1px solid var(--border-light);
  height: 64px;
}

.navbar .container-fluid {
  padding: 0 24px;
  height: 100%;
  display: flex;
  align-items: center;
}

.navbar-brand {
  font-weight: 600;
  font-size: 18px;
  color: var(--text-primary);
  transition: all 0.15s ease;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 0;
  margin-right: 32px;
}

.navbar-brand:hover {
  color: var(--primary-color);
}

.navbar-brand i {
  font-size: 24px;
  color: var(--primary-color);
}

.nav-link {
  padding: 8px 12px;
  transition: all 0.15s ease;
  color: var(--text-tertiary);
  font-weight: 500;
  border-radius: 6px;
  margin: 0 2px;
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 6px;
}

.nav-link:hover {
  background-color: var(--bg-secondary);
  color: var(--text-primary);
}

.nav-link.router-link-active {
  background-color: var(--primary-subtle);
  color: var(--primary-color);
}

.nav-link i {
  font-size: 16px;
}

.navbar-nav {
  align-items: center;
}

.dropdown-toggle {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  border-radius: 6px;
  transition: all 0.15s ease;
  color: var(--text-secondary);
  font-size: 14px;
}

.dropdown-toggle:hover {
  background-color: var(--bg-secondary);
  color: var(--text-primary);
}

.dropdown-toggle::after {
  margin-left: 4px;
}

.dropdown-menu {
  border: 1px solid var(--border-light);
  border-radius: 8px;
  box-shadow: var(--shadow-lg);
  padding: 8px;
  margin-top: 8px;
  min-width: 200px;
}

.dropdown-item {
  border-radius: 6px;
  padding: 8px 12px;
  transition: all 0.15s ease;
  color: var(--text-secondary);
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.dropdown-item:hover {
  background-color: var(--bg-secondary);
  color: var(--text-primary);
}

.dropdown-item:active {
  background-color: var(--bg-tertiary);
  color: var(--text-primary);
}

.dropdown-item i {
  font-size: 16px;
  color: var(--text-tertiary);
}

.dropdown-divider {
  margin: 6px 0;
  border-color: var(--border-light);
}

/* Mobile responsive */
@media (max-width: 991px) {
  .navbar {
    height: auto;
    min-height: 64px;
  }

  .navbar .container-fluid {
    padding: 12px 16px;
  }

  .navbar-collapse {
    margin-top: 12px;
    padding-top: 12px;
    border-top: 1px solid var(--border-light);
  }

  .navbar-nav {
    align-items: flex-start;
  }

  .nav-link {
    margin: 2px 0;
  }
}
</style>
