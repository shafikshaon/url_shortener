<template>
  <nav class="navbar navbar-expand-lg navbar-dark">
    <div class="container-fluid">
      <RouterLink class="navbar-brand" to="/dashboard">
        <i class="bi bi-link-45deg"></i> URL Shortener
      </RouterLink>

      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
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
  background-color: #7C3AED;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
  padding: 1rem 0;
  border-bottom: 1px solid rgba(167, 139, 250, 0.2);
}

.navbar-brand {
  font-weight: 600;
  font-size: 1.25rem;
  color: white;
  transition: all 0.2s ease;
}

.navbar-brand:hover {
  color: #E9D5FF;
}

.nav-link {
  padding: 0.5rem 1rem;
  transition: all 0.2s ease;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 500;
  border-radius: 0.5rem;
  margin: 0 0.25rem;
}

.nav-link:hover {
  background-color: rgba(255, 255, 255, 0.15);
  color: white;
}

.nav-link.router-link-active {
  background-color: rgba(255, 255, 255, 0.2);
  color: white;
}

.dropdown-menu {
  border: 1px solid #E5E7EB;
  border-radius: 0.5rem;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  padding: 0.5rem;
}

.dropdown-item {
  border-radius: 0.375rem;
  padding: 0.5rem 1rem;
  transition: all 0.2s ease;
  color: #374151;
}

.dropdown-item:hover {
  background-color: #F3F4F6;
  color: #7C3AED;
}

.dropdown-divider {
  margin: 0.5rem 0;
  border-color: #E5E7EB;
}
</style>
