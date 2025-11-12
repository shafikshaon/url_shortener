<template>
  <div class="topbar">
    <!-- Mobile Menu Button -->
    <button class="mobile-menu-btn" @click="toggleMobileMenu">
      <i class="bi bi-list"></i>
    </button>

    <!-- Search Bar (Desktop) -->
    <div class="search-container desktop-search">
      <i class="bi bi-search search-icon"></i>
      <input
        type="text"
        class="search-input"
        placeholder="Search links..."
        v-model="searchQuery"
        @focus="showSearchResults = true"
        @blur="hideSearchResults"
      />
      <kbd class="search-kbd">⌘K</kbd>

      <!-- Search Results Dropdown -->
      <div class="search-results" v-if="showSearchResults && searchQuery">
        <div class="search-results-header">
          <span class="text-muted">Recent searches</span>
        </div>
        <div class="search-result-item">
          <i class="bi bi-clock-history"></i>
          <span>Example search</span>
        </div>
      </div>
    </div>

    <!-- Right Section -->
    <div class="topbar-right">
      <!-- Create Button -->
      <button class="btn btn-primary btn-sm create-btn" @click="showCreateModal">
        <i class="bi bi-plus-lg"></i>
        <span class="create-text">Create</span>
      </button>

      <!-- Notifications -->
      <div class="topbar-item dropdown-container">
        <button class="icon-btn" @click="toggleNotifications" title="Notifications">
          <i class="bi bi-bell"></i>
          <span class="notification-badge" v-if="unreadCount > 0">{{ unreadCount }}</span>
        </button>

        <!-- Notifications Dropdown -->
        <div class="dropdown-menu notifications-dropdown" v-if="showNotifications">
          <div class="dropdown-header">
            <h6>Notifications</h6>
            <button class="btn-link">Mark all as read</button>
          </div>
          <div class="notifications-list">
            <div class="notification-item">
              <div class="notification-icon success">
                <i class="bi bi-check-circle"></i>
              </div>
              <div class="notification-content">
                <div class="notification-title">Link created successfully</div>
                <div class="notification-time">2 minutes ago</div>
              </div>
            </div>
            <div class="notification-item unread">
              <div class="notification-icon info">
                <i class="bi bi-info-circle"></i>
              </div>
              <div class="notification-content">
                <div class="notification-title">Your link reached 100 clicks</div>
                <div class="notification-time">1 hour ago</div>
              </div>
            </div>
          </div>
          <div class="dropdown-footer">
            <a href="#" class="view-all-link">View all notifications</a>
          </div>
        </div>
      </div>

      <!-- Help -->
      <div class="topbar-item">
        <button class="icon-btn" title="Help & Documentation">
          <i class="bi bi-question-circle"></i>
        </button>
      </div>

      <!-- User Menu -->
      <div class="topbar-item user-menu-container">
        <button class="user-btn" @click="toggleUserMenu">
          <div class="user-avatar">
            <i class="bi bi-person-fill"></i>
          </div>
          <span class="user-name">{{ userName }}</span>
          <i class="bi bi-chevron-down"></i>
        </button>

        <!-- User Dropdown -->
        <div class="dropdown-menu user-dropdown" v-if="showUserMenu">
          <div class="dropdown-header">
            <div class="user-info">
              <div class="user-info-name">{{ userName }}</div>
              <div class="user-info-email">{{ userEmail }}</div>
            </div>
          </div>
          <div class="dropdown-divider"></div>
          <RouterLink to="/settings" class="dropdown-item">
            <i class="bi bi-person"></i>
            <span>Profile</span>
          </RouterLink>
          <RouterLink to="/settings" class="dropdown-item">
            <i class="bi bi-gear"></i>
            <span>Settings</span>
          </RouterLink>
          <a href="#" class="dropdown-item" v-if="false">
            <i class="bi bi-credit-card"></i>
            <span>Billing</span>
          </a>
          <div class="dropdown-divider"></div>
          <a href="#" class="dropdown-item" @click.prevent="handleLogout">
            <i class="bi bi-box-arrow-right"></i>
            <span>Logout</span>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const authStore = useAuthStore()

const searchQuery = ref('')
const showSearchResults = ref(false)
const showNotifications = ref(false)
const showUserMenu = ref(false)
const unreadCount = ref(2)

const userName = computed(() => {
  const email = authStore.currentUser?.email || 'User'
  return email.split('@')[0]
})

const userEmail = computed(() => {
  return authStore.currentUser?.email || 'user@example.com'
})

const toggleMobileMenu = () => {
  window.dispatchEvent(new CustomEvent('toggle-mobile-menu'))
}

const showCreateModal = () => {
  window.dispatchEvent(new CustomEvent('show-create-modal'))
}

const toggleNotifications = () => {
  showNotifications.value = !showNotifications.value
  showUserMenu.value = false
}

const toggleUserMenu = () => {
  showUserMenu.value = !showUserMenu.value
  showNotifications.value = false
}

const hideSearchResults = () => {
  setTimeout(() => {
    showSearchResults.value = false
  }, 200)
}

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.topbar {
  position: fixed;
  top: 0;
  left: 260px;
  right: 0;
  height: 64px;
  background-color: var(--bg-white);
  border-bottom: 1px solid var(--border-light);
  padding: 0 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  z-index: 100;
  box-shadow: var(--shadow-xs);
  backdrop-filter: blur(8px);
  background-color: rgba(255, 255, 255, 0.98);
  transition: left var(--transition-base);
  gap: 24px;
}

/* Mobile Menu Button */
.mobile-menu-btn {
  display: none;
  width: 40px;
  height: 40px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-light);
  background-color: var(--bg-white);
  color: var(--text-primary);
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all var(--transition-base);
  font-size: 20px;
  flex-shrink: 0;
}

.mobile-menu-btn:hover {
  background-color: var(--bg-hover);
  border-color: var(--border-hover);
}

/* Search Container */
.search-container {
  position: relative;
  flex: 1;
  max-width: 480px;
}

.search-container .search-icon {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--text-muted);
  font-size: 16px;
  pointer-events: none;
}

.search-input {
  width: 100%;
  height: 40px;
  padding: 0 60px 0 40px;
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  background-color: var(--bg-secondary);
  color: var(--text-primary);
  font-size: 14px;
  transition: all var(--transition-base);
}

.search-input:hover {
  border-color: var(--border-hover);
  background-color: var(--bg-white);
}

.search-input:focus {
  outline: none;
  border-color: var(--border-focus);
  background-color: var(--bg-white);
  box-shadow: var(--focus-ring);
}

.search-kbd {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  padding: 2px 6px;
  background-color: var(--bg-tertiary);
  color: var(--text-muted);
  font-size: 11px;
  font-weight: 600;
  border-radius: var(--radius-sm);
  border: 1px solid var(--border-color);
  font-family: inherit;
  pointer-events: none;
}

/* Search Results */
.search-results {
  position: absolute;
  top: calc(100% + 8px);
  left: 0;
  right: 0;
  background-color: var(--bg-white);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-xl);
  padding: 8px;
  z-index: 1000;
}

.search-results-header {
  padding: 8px 12px;
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--text-muted);
}

.search-result-item {
  padding: 10px 12px;
  display: flex;
  align-items: center;
  gap: 12px;
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all var(--transition-base);
  color: var(--text-secondary);
  font-size: 14px;
}

.search-result-item:hover {
  background-color: var(--bg-hover);
  color: var(--text-primary);
}

.search-result-item i {
  color: var(--text-muted);
  font-size: 16px;
}

/* Right Section */
.topbar-right {
  display: flex;
  align-items: center;
  gap: 8px;
}

.create-btn {
  padding: 8px 16px;
  gap: 6px;
}

.topbar-item {
  position: relative;
}

/* Icon Buttons */
.icon-btn {
  width: 40px;
  height: 40px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-light);
  background-color: var(--bg-white);
  color: var(--text-secondary);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all var(--transition-base);
  position: relative;
}

.icon-btn:hover {
  background-color: var(--bg-hover);
  border-color: var(--border-hover);
  color: var(--text-primary);
}

.icon-btn i {
  font-size: 18px;
}

/* Notification Badge */
.notification-badge {
  position: absolute;
  top: -4px;
  right: -4px;
  min-width: 18px;
  height: 18px;
  padding: 0 4px;
  background-color: var(--danger-color);
  color: white;
  font-size: 11px;
  font-weight: 600;
  border-radius: 9px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid var(--bg-white);
}

/* User Button */
.user-btn {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 6px 12px 6px 6px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-light);
  background-color: var(--bg-white);
  cursor: pointer;
  transition: all var(--transition-base);
  height: 40px;
}

.user-btn:hover {
  background-color: var(--bg-hover);
  border-color: var(--border-hover);
}

.user-avatar {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 14px;
  flex-shrink: 0;
}

.user-name {
  font-size: 14px;
  font-weight: 500;
  color: var(--text-primary);
  max-width: 120px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-btn i.bi-chevron-down {
  font-size: 12px;
  color: var(--text-tertiary);
  transition: transform var(--transition-base);
}

.user-btn:hover i.bi-chevron-down {
  transform: translateY(1px);
}

/* Dropdown Menus */
.dropdown-menu {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  background-color: var(--bg-white);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-xl);
  padding: 8px;
  z-index: 1000;
  animation: dropdownSlideIn 0.2s ease;
  min-width: 200px;
}

@keyframes dropdownSlideIn {
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
  padding: 12px 12px 8px 12px;
}

.dropdown-header h6 {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0 0 4px 0;
}

.user-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.user-info-name {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-primary);
}

.user-info-email {
  font-size: 13px;
  color: var(--text-tertiary);
  overflow: hidden;
  text-overflow: ellipsis;
}

/* Notifications Dropdown */
.notifications-dropdown {
  width: 360px;
  max-height: 480px;
  overflow-y: auto;
}

.notifications-list {
  max-height: 320px;
  overflow-y: auto;
}

.notification-item {
  display: flex;
  gap: 12px;
  padding: 12px;
  border-radius: var(--radius-md);
  transition: all var(--transition-base);
  cursor: pointer;
  margin-bottom: 4px;
}

.notification-item:hover {
  background-color: var(--bg-hover);
}

.notification-item.unread {
  background-color: var(--primary-subtle);
}

.notification-icon {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.notification-icon.success {
  background-color: #E6FAF0;
  color: var(--success-color);
}

.notification-icon.info {
  background-color: #E6F2FF;
  color: var(--info-color);
}

.notification-icon i {
  font-size: 16px;
}

.notification-content {
  flex: 1;
  min-width: 0;
}

.notification-title {
  font-size: 14px;
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: 2px;
  line-height: 1.4;
}

.notification-time {
  font-size: 12px;
  color: var(--text-tertiary);
}

.dropdown-footer {
  padding: 8px 12px;
  border-top: 1px solid var(--border-light);
  margin-top: 4px;
}

.view-all-link {
  display: block;
  text-align: center;
  font-size: 13px;
  font-weight: 500;
  color: var(--primary-color);
  text-decoration: none;
  padding: 4px;
  border-radius: var(--radius-md);
  transition: all var(--transition-base);
}

.view-all-link:hover {
  background-color: var(--primary-subtle);
}

/* User Dropdown */
.user-dropdown {
  min-width: 240px;
}

/* Responsive */
@media (max-width: 991px) {
  .topbar {
    left: 0;
  }

  .mobile-menu-btn {
    display: flex;
  }

  .desktop-search {
    display: none;
  }

  .create-text {
    display: none;
  }

  .user-name {
    display: none;
  }

  .user-btn {
    padding: 6px;
  }
}

@media (max-width: 576px) {
  .topbar {
    padding: 0 16px;
  }

  .topbar-right {
    gap: 4px;
  }

  .create-btn span {
    display: none;
  }

  .create-btn {
    padding: 8px 12px;
  }
}
</style>
