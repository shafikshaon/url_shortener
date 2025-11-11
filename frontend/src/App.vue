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
  /* Stripe-Inspired Professional Colors */
  --primary-color: #635BFF;
  --primary-hover: #4F46E5;
  --primary-light: #8B85FF;
  --primary-dark: #3730A3;
  --primary-subtle: #F4F3FF;

  /* Supporting Colors */
  --secondary-color: #697386;
  --success-color: #00D924;
  --danger-color: #DF1B41;
  --warning-color: #F5A623;
  --info-color: #0073E6;

  /* Backgrounds - Stripe-style subtle grays */
  --bg-primary: #FAFBFC;
  --bg-secondary: #F6F9FC;
  --bg-tertiary: #E3E8EE;
  --bg-white: #FFFFFF;

  /* Text Colors - Professional hierarchy */
  --text-primary: #0A2540;
  --text-secondary: #425466;
  --text-tertiary: #697386;
  --text-muted: #8792A2;

  /* Borders - Subtle and professional */
  --border-color: #E3E8EE;
  --border-light: #F0F4F8;
  --border-subtle: #EDF2F7;

  /* Shadows - Very subtle Stripe-style */
  --shadow-xs: 0 1px 2px 0 rgba(10, 37, 64, 0.03);
  --shadow-sm: 0 1px 3px 0 rgba(10, 37, 64, 0.05);
  --shadow-md: 0 4px 6px -1px rgba(10, 37, 64, 0.08), 0 2px 4px -1px rgba(10, 37, 64, 0.03);
  --shadow-lg: 0 10px 15px -3px rgba(10, 37, 64, 0.08), 0 4px 6px -2px rgba(10, 37, 64, 0.03);
  --shadow-xl: 0 20px 25px -5px rgba(10, 37, 64, 0.08), 0 10px 10px -5px rgba(10, 37, 64, 0.02);

  /* Hover states */
  --hover-overlay: rgba(0, 0, 0, 0.02);
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Helvetica Neue', Arial, sans-serif;
  background-color: var(--bg-primary);
  color: var(--text-primary);
  line-height: 1.5;
  font-size: 15px;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

h1, h2, h3, h4, h5, h6 {
  color: var(--text-primary);
  font-weight: 600;
  line-height: 1.3;
  letter-spacing: -0.02em;
}

h1 {
  font-size: 2rem;
}

h2 {
  font-size: 1.5rem;
}

h3 {
  font-size: 1.25rem;
}

p {
  line-height: 1.6;
  color: var(--text-secondary);
}

#app {
  min-height: 100vh;
}

main {
  min-height: calc(100vh - 64px);
}

/* Buttons - Stripe-inspired professional style */
.btn {
  border-radius: 6px;
  font-weight: 500;
  padding: 10px 16px;
  transition: all 0.15s ease;
  border: 1px solid transparent;
  font-size: 14px;
  letter-spacing: -0.01em;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
}

.btn-sm {
  padding: 6px 12px;
  font-size: 13px;
}

.btn-lg {
  padding: 12px 20px;
  font-size: 15px;
}

.btn-primary {
  background-color: var(--primary-color);
  color: white;
  box-shadow: 0 1px 2px rgba(10, 37, 64, 0.1);
}

.btn-primary:hover {
  background-color: var(--primary-hover);
  box-shadow: 0 2px 4px rgba(10, 37, 64, 0.12);
  transform: translateY(-1px);
  color: white;
}

.btn-primary:active {
  transform: translateY(0);
  box-shadow: 0 1px 2px rgba(10, 37, 64, 0.1);
}

.btn-outline-primary {
  border: 1px solid var(--border-color);
  color: var(--text-primary);
  background-color: var(--bg-white);
  box-shadow: 0 1px 2px rgba(10, 37, 64, 0.05);
}

.btn-outline-primary:hover {
  background-color: var(--bg-secondary);
  border-color: var(--border-color);
  color: var(--text-primary);
  box-shadow: 0 2px 4px rgba(10, 37, 64, 0.08);
}

.btn-secondary {
  background-color: var(--bg-tertiary);
  color: var(--text-primary);
  border: 1px solid var(--border-color);
}

.btn-secondary:hover {
  background-color: #D1D9E0;
  color: var(--text-primary);
}

.btn-success {
  background-color: var(--success-color);
  color: white;
  box-shadow: 0 1px 2px rgba(0, 217, 36, 0.15);
}

.btn-success:hover {
  background-color: #00C020;
  color: white;
}

.btn-danger {
  background-color: var(--danger-color);
  color: white;
  box-shadow: 0 1px 2px rgba(223, 27, 65, 0.15);
}

.btn-danger:hover {
  background-color: #C61638;
  color: white;
}

.btn-outline-danger {
  border: 1px solid var(--danger-color);
  color: var(--danger-color);
  background-color: transparent;
}

.btn-outline-danger:hover {
  background-color: var(--danger-color);
  color: white;
}

.btn-outline-light {
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: white;
  background-color: transparent;
}

.btn-outline-light:hover {
  background-color: white;
  color: var(--primary-color);
  border-color: white;
}

/* Cards - Clean Stripe-style */
.card {
  border-radius: 8px;
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--border-light);
  background-color: var(--bg-white);
  overflow: hidden;
  transition: all 0.2s ease;
}

.card:hover {
  box-shadow: var(--shadow-md);
}

.card-header {
  background-color: var(--bg-white);
  border-bottom: 1px solid var(--border-light);
  padding: 20px 24px;
  font-weight: 600;
  color: var(--text-primary);
  font-size: 15px;
}

.card-body {
  padding: 24px;
}

.card-title {
  color: var(--text-primary);
  font-weight: 600;
  font-size: 15px;
  margin-bottom: 4px;
}

.card-subtitle {
  color: var(--text-tertiary);
  font-size: 13px;
  font-weight: 400;
}

/* Forms - Professional Stripe-style inputs */
.form-control,
.form-select {
  border: 1px solid var(--border-color);
  border-radius: 6px;
  padding: 10px 12px;
  font-size: 14px;
  transition: all 0.15s ease;
  background-color: var(--bg-white);
  color: var(--text-primary);
  line-height: 1.5;
}

.form-control:hover,
.form-select:hover {
  border-color: #C7CCD1;
}

.form-control:focus,
.form-select:focus {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(99, 91, 255, 0.08);
  outline: none;
  background-color: var(--bg-white);
}

.form-control::placeholder {
  color: var(--text-muted);
}

.form-label {
  font-weight: 500;
  color: var(--text-secondary);
  margin-bottom: 8px;
  font-size: 13px;
  letter-spacing: -0.01em;
  display: block;
}

.form-text {
  font-size: 13px;
  color: var(--text-tertiary);
  margin-top: 4px;
}

/* Tables - Clean Stripe-style */
.table {
  margin-bottom: 0;
  border-collapse: separate;
  border-spacing: 0;
}

.table thead {
  background-color: var(--bg-secondary);
}

.table thead th {
  font-weight: 600;
  color: var(--text-tertiary);
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  padding: 12px 16px;
  border-bottom: 1px solid var(--border-light);
  border-top: none;
}

.table tbody td {
  padding: 16px;
  color: var(--text-secondary);
  font-size: 14px;
  border-bottom: 1px solid var(--border-light);
  vertical-align: middle;
}

.table tbody tr {
  transition: background-color 0.15s ease;
}

.table-hover tbody tr:hover {
  background-color: var(--bg-secondary);
}

.table tbody tr:last-child td {
  border-bottom: none;
}

/* Badges - Refined Stripe-style */
.badge {
  font-weight: 500;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
  letter-spacing: -0.01em;
  display: inline-flex;
  align-items: center;
}

.bg-primary {
  background-color: var(--primary-subtle) !important;
  color: var(--primary-color) !important;
}

.bg-success {
  background-color: #E6FAF0 !important;
  color: #00A86B !important;
}

.bg-danger {
  background-color: #FCEDEF !important;
  color: #C61638 !important;
}

.bg-secondary {
  background-color: var(--bg-tertiary) !important;
  color: var(--text-secondary) !important;
}

.bg-info {
  background-color: #E6F2FF !important;
  color: #0073E6 !important;
}

/* Text Colors */
.text-primary {
  color: var(--primary-color) !important;
}

.text-secondary {
  color: var(--text-secondary) !important;
}

.text-muted {
  color: var(--text-muted) !important;
}

.text-tertiary {
  color: var(--text-tertiary) !important;
}

/* Alerts - Professional messaging */
.alert {
  border-radius: 6px;
  border: 1px solid;
  padding: 14px 16px;
  font-size: 14px;
  line-height: 1.5;
}

.alert-success {
  background-color: #E6FAF0;
  border-color: #B3E6D1;
  color: #006644;
}

.alert-danger {
  background-color: #FCEDEF;
  border-color: #F5B8C5;
  color: #9E1030;
}

.alert-info {
  background-color: #E6F2FF;
  border-color: #B3D9FF;
  color: #004C99;
}

.alert-warning {
  background-color: #FFF5E6;
  border-color: #FFE0B3;
  color: #8C5B00;
}

/* Spinners */
.spinner-border {
  border-color: var(--primary-light);
  border-right-color: transparent;
}

/* Code blocks */
code {
  background-color: var(--bg-secondary);
  color: var(--primary-color);
  padding: 3px 6px;
  border-radius: 4px;
  font-size: 13px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  border: 1px solid var(--border-light);
}
</style>
