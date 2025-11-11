<template>
  <div id="app">
    <Navbar v-if="isAuthenticated" />
    <main>
      <RouterView />
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { RouterView } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import Navbar from '@/components/Navbar.vue'

const authStore = useAuthStore()
const isAuthenticated = computed(() => authStore.isLoggedIn)
</script>

<style>
:root {
  --primary-color: #0d6efd;
  --secondary-color: #6c757d;
  --success-color: #198754;
  --danger-color: #dc3545;
  --warning-color: #ffc107;
  --info-color: #0dcaf0;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background-color: #f8f9fa;
}

#app {
  min-height: 100vh;
}

main {
  min-height: calc(100vh - 56px);
}

.btn {
  border-radius: 0.375rem;
  font-weight: 500;
}

.card {
  border-radius: 0.5rem;
  box-shadow: 0 0.125rem 0.25rem rgba(0, 0, 0, 0.075);
}

.form-control:focus,
.form-select:focus {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 0.2rem rgba(13, 110, 253, 0.25);
}

.table {
  margin-bottom: 0;
}

.badge {
  font-weight: 500;
}
</style>
