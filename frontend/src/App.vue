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
  /* Purple Theme Colors */
  --primary-color: #7C3AED;
  --primary-hover: #6D28D9;
  --primary-light: #A78BFA;
  --primary-dark: #5B21B6;

  /* Supporting Colors */
  --secondary-color: #6B7280;
  --success-color: #10B981;
  --danger-color: #EF4444;
  --warning-color: #F59E0B;
  --info-color: #3B82F6;

  /* Backgrounds */
  --bg-primary: #F9FAFB;
  --bg-secondary: #F3F4F6;
  --bg-white: #FFFFFF;

  /* Text Colors */
  --text-primary: #111827;
  --text-secondary: #6B7280;
  --text-muted: #9CA3AF;

  /* Borders */
  --border-color: #E5E7EB;
  --border-light: #F3F4F6;

  /* Shadows */
  --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  --shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background-color: var(--bg-primary);
  color: var(--text-primary);
  line-height: 1.6;
}

#app {
  min-height: 100vh;
}

main {
  min-height: calc(100vh - 64px);
}

/* Buttons */
.btn {
  border-radius: 0.5rem;
  font-weight: 500;
  padding: 0.625rem 1.25rem;
  transition: all 0.2s ease;
  border: none;
  font-size: 0.875rem;
}

.btn-primary {
  background-color: var(--primary-color);
  color: white;
}

.btn-primary:hover {
  background-color: var(--primary-hover);
  transform: translateY(-1px);
  box-shadow: var(--shadow-md);
}

.btn-outline-primary {
  border: 2px solid var(--primary-color);
  color: var(--primary-color);
  background-color: transparent;
}

.btn-outline-primary:hover {
  background-color: var(--primary-color);
  color: white;
}

.btn-secondary {
  background-color: var(--secondary-color);
  color: white;
}

.btn-success {
  background-color: var(--success-color);
  color: white;
}

.btn-danger {
  background-color: var(--danger-color);
  color: white;
}

.btn-outline-danger {
  border: 2px solid var(--danger-color);
  color: var(--danger-color);
  background-color: transparent;
}

.btn-outline-danger:hover {
  background-color: var(--danger-color);
  color: white;
}

.btn-outline-light {
  border: 2px solid white;
  color: white;
  background-color: transparent;
}

.btn-outline-light:hover {
  background-color: white;
  color: var(--primary-color);
}

/* Cards */
.card {
  border-radius: 0.75rem;
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--border-color);
  background-color: var(--bg-white);
  overflow: hidden;
}

.card-header {
  background-color: var(--bg-white);
  border-bottom: 1px solid var(--border-color);
  padding: 1rem 1.5rem;
  font-weight: 600;
}

.card-body {
  padding: 1.5rem;
}

/* Forms */
.form-control,
.form-select {
  border: 1px solid var(--border-color);
  border-radius: 0.5rem;
  padding: 0.625rem 0.875rem;
  font-size: 0.875rem;
  transition: all 0.2s ease;
}

.form-control:focus,
.form-select:focus {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(124, 58, 237, 0.1);
  outline: none;
}

.form-label {
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
}

/* Tables */
.table {
  margin-bottom: 0;
}

.table thead {
  background-color: var(--bg-secondary);
  border-bottom: 2px solid var(--border-color);
}

.table th {
  font-weight: 600;
  color: var(--text-primary);
  font-size: 0.875rem;
  padding: 0.875rem 1rem;
}

.table td {
  padding: 0.875rem 1rem;
  color: var(--text-secondary);
  font-size: 0.875rem;
}

.table-hover tbody tr:hover {
  background-color: var(--bg-secondary);
}

/* Badges */
.badge {
  font-weight: 500;
  padding: 0.375rem 0.75rem;
  border-radius: 0.375rem;
  font-size: 0.75rem;
}

.bg-primary {
  background-color: var(--primary-color) !important;
}

.bg-success {
  background-color: var(--success-color) !important;
}

.bg-danger {
  background-color: var(--danger-color) !important;
}

.bg-secondary {
  background-color: var(--secondary-color) !important;
}

.bg-info {
  background-color: var(--info-color) !important;
}

/* Text Colors */
.text-primary {
  color: var(--primary-color) !important;
}

.text-muted {
  color: var(--text-muted) !important;
}

/* Alerts */
.alert {
  border-radius: 0.5rem;
  border: none;
  padding: 1rem 1.25rem;
}

.alert-success {
  background-color: #D1FAE5;
  color: #065F46;
}

.alert-danger {
  background-color: #FEE2E2;
  color: #991B1B;
}

/* Spinners */
.spinner-border {
  border-color: var(--primary-light);
  border-right-color: transparent;
}
</style>
