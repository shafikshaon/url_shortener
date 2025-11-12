<template>
  <div id="app" :class="{ 'has-sidebar': isAuthenticated }">
    <Sidebar v-if="isAuthenticated" />
    <TopBar v-if="isAuthenticated" />
    <div class="main-wrapper" :class="{ 'with-sidebar': isAuthenticated, 'with-topbar': isAuthenticated }">
      <main>
        <RouterView />
      </main>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { RouterView } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import Sidebar from '@/components/Sidebar.vue'
import TopBar from '@/components/TopBar.vue'

const authStore = useAuthStore()
const isAuthenticated = computed(() => authStore.isLoggedIn)
</script>

<style>
:root {
  /* Stripe-Inspired Professional Colors */
  --primary-color: #635BFF;
  --primary-hover: #5851DF;
  --primary-active: #4F46E5;
  --primary-light: #8B85FF;
  --primary-dark: #3730A3;
  --primary-subtle: #F4F3FF;
  --primary-subtle-hover: #EEEEFF;

  /* Supporting Colors */
  --secondary-color: #697386;
  --success-color: #00D924;
  --success-hover: #00C020;
  --danger-color: #DF1B41;
  --danger-hover: #C61638;
  --warning-color: #F5A623;
  --info-color: #0073E6;

  /* Backgrounds - Stripe-style subtle grays */
  --bg-primary: #FAFBFC;
  --bg-secondary: #F6F9FC;
  --bg-tertiary: #E3E8EE;
  --bg-hover: #F7FAFC;
  --bg-white: #FFFFFF;

  /* Text Colors - Professional hierarchy */
  --text-primary: #0A2540;
  --text-secondary: #425466;
  --text-tertiary: #697386;
  --text-muted: #8792A2;
  --text-disabled: #A9B4C2;

  /* Borders - Subtle and professional */
  --border-color: #E3E8EE;
  --border-hover: #CBD5E0;
  --border-light: #F0F4F8;
  --border-subtle: #EDF2F7;
  --border-focus: #635BFF;

  /* Shadows - Very subtle Stripe-style */
  --shadow-xs: 0 1px 2px 0 rgba(10, 37, 64, 0.03);
  --shadow-sm: 0 1px 3px 0 rgba(10, 37, 64, 0.05), 0 1px 2px 0 rgba(10, 37, 64, 0.03);
  --shadow-md: 0 4px 6px -1px rgba(10, 37, 64, 0.08), 0 2px 4px -1px rgba(10, 37, 64, 0.04);
  --shadow-lg: 0 10px 15px -3px rgba(10, 37, 64, 0.1), 0 4px 6px -2px rgba(10, 37, 64, 0.05);
  --shadow-xl: 0 20px 25px -5px rgba(10, 37, 64, 0.1), 0 10px 10px -5px rgba(10, 37, 64, 0.04);
  --shadow-2xl: 0 25px 50px -12px rgba(10, 37, 64, 0.15);

  /* Focus ring */
  --focus-ring: 0 0 0 3px rgba(99, 91, 255, 0.1);
  --focus-ring-offset: 0 0 0 2px var(--bg-white);

  /* Transitions */
  --transition-fast: 0.1s ease;
  --transition-base: 0.15s ease;
  --transition-slow: 0.2s ease;

  /* Border Radius */
  --radius-sm: 4px;
  --radius-md: 6px;
  --radius-lg: 8px;
  --radius-xl: 12px;
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
  display: flex;
}

#app.has-sidebar {
  padding-left: 0;
}

.main-wrapper {
  flex: 1;
  min-height: 100vh;
  transition: margin-left var(--transition-base);
}

.main-wrapper.with-sidebar {
  margin-left: 260px;
}

.main-wrapper.with-topbar {
  padding-top: 64px;
}

main {
  min-height: calc(100vh - 64px);
  width: 100%;
}

/* Responsive layout */
@media (max-width: 991px) {
  .main-wrapper.with-sidebar {
    margin-left: 0;
  }

  main {
    min-height: calc(100vh - 64px);
  }
}

/* Mobile app bar (for opening sidebar on mobile) */
.mobile-header {
  display: none;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 64px;
  background-color: var(--bg-white);
  border-bottom: 1px solid var(--border-light);
  padding: 0 20px;
  align-items: center;
  justify-content: space-between;
  z-index: 998;
  box-shadow: var(--shadow-sm);
}

@media (max-width: 991px) {
  .mobile-header {
    display: flex;
  }

  main {
    padding-top: 64px;
  }
}

/* Buttons - Stripe-inspired professional style */
.btn {
  border-radius: var(--radius-md);
  font-weight: 500;
  padding: 10px 16px;
  transition: all var(--transition-base);
  border: 1px solid transparent;
  font-size: 14px;
  letter-spacing: -0.01em;
  cursor: pointer;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  position: relative;
  text-decoration: none;
  line-height: 1.5;
  white-space: nowrap;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none !important;
}

.btn-sm {
  padding: 7px 12px;
  font-size: 13px;
  gap: 4px;
}

.btn-lg {
  padding: 12px 20px;
  font-size: 15px;
  gap: 8px;
}

.btn-primary {
  background-color: var(--primary-color);
  color: white;
  box-shadow: var(--shadow-xs);
  border: none;
}

.btn-primary:hover:not(:disabled) {
  background-color: var(--primary-hover);
  box-shadow: var(--shadow-sm);
  transform: translateY(-1px);
  color: white;
}

.btn-primary:active:not(:disabled) {
  background-color: var(--primary-active);
  transform: translateY(0);
  box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.1);
}

.btn-primary:focus-visible {
  outline: none;
  box-shadow: var(--focus-ring);
}

.btn-outline-primary {
  border: 1px solid var(--border-color);
  color: var(--text-primary);
  background-color: var(--bg-white);
  box-shadow: var(--shadow-xs);
}

.btn-outline-primary:hover:not(:disabled) {
  background-color: var(--bg-hover);
  border-color: var(--border-hover);
  color: var(--text-primary);
  box-shadow: var(--shadow-sm);
  transform: translateY(-1px);
}

.btn-outline-primary:active:not(:disabled) {
  background-color: var(--bg-secondary);
  transform: translateY(0);
}

.btn-outline-primary:focus-visible {
  outline: none;
  border-color: var(--border-focus);
  box-shadow: var(--focus-ring);
}

.btn-secondary {
  background-color: var(--bg-white);
  color: var(--text-secondary);
  border: 1px solid var(--border-color);
  box-shadow: var(--shadow-xs);
}

.btn-secondary:hover:not(:disabled) {
  background-color: var(--bg-hover);
  border-color: var(--border-hover);
  color: var(--text-primary);
  box-shadow: var(--shadow-sm);
}

.btn-secondary:active:not(:disabled) {
  background-color: var(--bg-secondary);
}

.btn-success {
  background-color: var(--success-color);
  color: white;
  box-shadow: var(--shadow-xs);
  border: none;
}

.btn-success:hover:not(:disabled) {
  background-color: var(--success-hover);
  box-shadow: var(--shadow-sm);
  transform: translateY(-1px);
  color: white;
}

.btn-success:active:not(:disabled) {
  transform: translateY(0);
}

.btn-danger {
  background-color: var(--danger-color);
  color: white;
  box-shadow: var(--shadow-xs);
  border: none;
}

.btn-danger:hover:not(:disabled) {
  background-color: var(--danger-hover);
  box-shadow: var(--shadow-sm);
  transform: translateY(-1px);
  color: white;
}

.btn-danger:active:not(:disabled) {
  transform: translateY(0);
}

.btn-outline-danger {
  border: 1px solid var(--danger-color);
  color: var(--danger-color);
  background-color: var(--bg-white);
  box-shadow: var(--shadow-xs);
}

.btn-outline-danger:hover:not(:disabled) {
  background-color: var(--danger-color);
  color: white;
  box-shadow: var(--shadow-sm);
}

.btn-outline-light {
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: white;
  background-color: rgba(255, 255, 255, 0.1);
}

.btn-outline-light:hover:not(:disabled) {
  background-color: white;
  color: var(--primary-color);
  border-color: white;
}

.btn-link {
  background: none;
  border: none;
  color: var(--primary-color);
  padding: 4px 8px;
  text-decoration: none;
  font-weight: 500;
}

.btn-link:hover:not(:disabled) {
  color: var(--primary-hover);
  background-color: var(--primary-subtle);
  text-decoration: none;
}

/* Cards - Clean Stripe-style */
.card {
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--border-light);
  background-color: var(--bg-white);
  overflow: hidden;
  transition: all var(--transition-base);
  position: relative;
}

.card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--border-color);
}

.card-header {
  background-color: var(--bg-white);
  border-bottom: 1px solid var(--border-light);
  padding: 20px 24px;
  font-weight: 600;
  color: var(--text-primary);
  font-size: 15px;
}

.card-header h5,
.card-header h6 {
  margin: 0;
  font-weight: 600;
}

.card-body {
  padding: 24px;
}

.card-title {
  color: var(--text-primary);
  font-weight: 600;
  font-size: 15px;
  margin-bottom: 4px;
  line-height: 1.4;
}

.card-subtitle {
  color: var(--text-tertiary);
  font-size: 13px;
  font-weight: 400;
  line-height: 1.5;
}

.card-footer {
  background-color: var(--bg-secondary);
  border-top: 1px solid var(--border-light);
  padding: 16px 24px;
}

/* Forms - Professional Stripe-style inputs */
.form-control,
.form-select,
input[type="text"],
input[type="email"],
input[type="password"],
input[type="url"],
input[type="datetime-local"],
textarea,
select {
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  padding: 10px 12px;
  font-size: 14px;
  transition: all var(--transition-base);
  background-color: var(--bg-white);
  color: var(--text-primary);
  line-height: 1.5;
  width: 100%;
  font-family: inherit;
}

.form-control:hover:not(:disabled):not(:focus),
.form-select:hover:not(:disabled):not(:focus),
input:hover:not(:disabled):not(:focus) {
  border-color: var(--border-hover);
}

.form-control:focus,
.form-select:focus,
input:focus,
textarea:focus,
select:focus {
  border-color: var(--border-focus);
  box-shadow: var(--focus-ring);
  outline: none;
  background-color: var(--bg-white);
}

.form-control:disabled,
.form-select:disabled,
input:disabled {
  background-color: var(--bg-secondary);
  color: var(--text-disabled);
  cursor: not-allowed;
  border-color: var(--border-light);
}

.form-control::placeholder,
input::placeholder,
textarea::placeholder {
  color: var(--text-muted);
  opacity: 1;
}

.form-control-lg {
  padding: 12px 16px;
  font-size: 15px;
  border-radius: var(--radius-lg);
}

.form-label {
  font-weight: 500;
  color: var(--text-secondary);
  margin-bottom: 8px;
  font-size: 13px;
  letter-spacing: -0.01em;
  display: block;
  line-height: 1.4;
}

.form-text {
  font-size: 12px;
  color: var(--text-tertiary);
  margin-top: 6px;
  line-height: 1.5;
  display: block;
}

.form-check-input {
  border-color: var(--border-color);
  transition: all var(--transition-base);
}

.form-check-input:checked {
  background-color: var(--primary-color);
  border-color: var(--primary-color);
}

.form-check-input:focus {
  border-color: var(--border-focus);
  box-shadow: var(--focus-ring);
}

/* Tables - Clean Stripe-style */
.table {
  margin-bottom: 0;
  border-collapse: separate;
  border-spacing: 0;
  width: 100%;
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
  border-bottom: 1px solid var(--border-color);
  border-top: none;
  white-space: nowrap;
  text-align: left;
}

.table thead th:first-child {
  padding-left: 24px;
  border-radius: var(--radius-lg) 0 0 0;
}

.table thead th:last-child {
  padding-right: 24px;
  border-radius: 0 var(--radius-lg) 0 0;
}

.table tbody td {
  padding: 16px;
  color: var(--text-secondary);
  font-size: 14px;
  border-bottom: 1px solid var(--border-light);
  vertical-align: middle;
  background-color: var(--bg-white);
}

.table tbody td:first-child {
  padding-left: 24px;
}

.table tbody td:last-child {
  padding-right: 24px;
}

.table tbody tr {
  transition: background-color var(--transition-base);
}

.table-hover tbody tr:hover {
  background-color: var(--bg-hover);
}

.table-hover tbody tr:hover td {
  background-color: var(--bg-hover);
}

.table tbody tr:last-child td {
  border-bottom: none;
}

.table tbody tr:last-child td:first-child {
  border-radius: 0 0 0 var(--radius-lg);
}

.table tbody tr:last-child td:last-child {
  border-radius: 0 0 var(--radius-lg) 0;
}

/* Badges - Refined Stripe-style */
.badge {
  font-weight: 500;
  padding: 4px 10px;
  border-radius: var(--radius-sm);
  font-size: 12px;
  letter-spacing: -0.01em;
  display: inline-flex;
  align-items: center;
  line-height: 1.4;
  gap: 4px;
}

.badge-sm {
  padding: 2px 6px;
  font-size: 11px;
}

.badge-lg {
  padding: 6px 12px;
  font-size: 13px;
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

.bg-warning {
  background-color: #FFF5E6 !important;
  color: #8C5B00 !important;
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
  border-radius: var(--radius-lg);
  border: 1px solid;
  padding: 14px 16px;
  font-size: 14px;
  line-height: 1.5;
  display: flex;
  align-items: flex-start;
  gap: 10px;
}

.alert i {
  flex-shrink: 0;
  margin-top: 2px;
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

.alert-link {
  color: inherit;
  font-weight: 600;
  text-decoration: underline;
}

.alert-link:hover {
  opacity: 0.8;
}

/* Spinners */
.spinner-border {
  border-color: var(--primary-light);
  border-right-color: transparent;
  animation: spinner-border 0.75s linear infinite;
}

.spinner-border-sm {
  width: 1rem;
  height: 1rem;
  border-width: 0.15em;
}

@keyframes spinner-border {
  to {
    transform: rotate(360deg);
  }
}

/* Code blocks */
code {
  background-color: var(--bg-secondary);
  color: var(--primary-color);
  padding: 3px 6px;
  border-radius: var(--radius-sm);
  font-size: 13px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', 'Courier New', monospace;
  border: 1px solid var(--border-light);
  font-weight: 500;
}

pre {
  background-color: var(--bg-secondary);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  padding: 16px;
  overflow-x: auto;
}

pre code {
  background: none;
  border: none;
  padding: 0;
  font-size: 13px;
  line-height: 1.6;
}

/* Dropdown Menus */
.dropdown-menu {
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-lg);
  padding: 8px;
  margin-top: 8px;
  background-color: var(--bg-white);
}

.dropdown-item {
  border-radius: var(--radius-md);
  padding: 10px 12px;
  transition: all var(--transition-base);
  color: var(--text-secondary);
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 500;
  cursor: pointer;
}

.dropdown-item:hover {
  background-color: var(--bg-secondary);
  color: var(--text-primary);
  text-decoration: none;
}

.dropdown-item:active {
  background-color: var(--bg-tertiary);
  color: var(--text-primary);
}

.dropdown-divider {
  margin: 8px 0;
  border-color: var(--border-light);
}

.dropdown-header {
  padding: 8px 12px;
  font-size: 12px;
  font-weight: 600;
  color: var(--text-tertiary);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

/* Modal backdrop */
.modal-backdrop {
  background-color: rgba(10, 37, 64, 0.5);
  backdrop-filter: blur(2px);
}

.modal-backdrop.show {
  opacity: 1;
}

/* Progress bars */
.progress {
  height: 8px;
  border-radius: var(--radius-sm);
  background-color: var(--bg-tertiary);
  overflow: hidden;
}

.progress-bar {
  background-color: var(--primary-color);
  border-radius: var(--radius-sm);
  transition: width var(--transition-slow);
}

/* Utilities */
.text-decoration-none {
  text-decoration: none !important;
}

.cursor-pointer {
  cursor: pointer;
}

.user-select-none {
  user-select: none;
}
</style>
