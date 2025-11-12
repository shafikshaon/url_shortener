<template>
  <div class="sidebar-wrapper" :class="{ 'collapsed': isCollapsed }">
    <!-- Sidebar -->
    <aside class="sidebar" :class="{ 'collapsed': isCollapsed }">
      <!-- Brand -->
      <div class="sidebar-header">
        <RouterLink to="/dashboard" class="sidebar-brand">
          <div class="brand-icon">
            <i class="bi bi-link-45deg"></i>
          </div>
          <span class="brand-text" v-if="!isCollapsed">LinkShort</span>
        </RouterLink>
        <button class="collapse-btn" @click="toggleSidebar" v-if="!isMobile">
          <i class="bi" :class="isCollapsed ? 'bi-chevron-right' : 'bi-chevron-left'"></i>
        </button>
      </div>

      <!-- Navigation -->
      <nav class="sidebar-nav">
        <!-- Main Section -->
        <div class="nav-section">
          <div class="nav-section-title" v-if="!isCollapsed">Main</div>
          <RouterLink to="/dashboard" class="nav-item" title="Dashboard">
            <i class="bi bi-grid-fill"></i>
            <span v-if="!isCollapsed">Dashboard</span>
          </RouterLink>
          <RouterLink to="/links" class="nav-item" title="My Links">
            <i class="bi bi-link-45deg"></i>
            <span v-if="!isCollapsed">My Links</span>
            <span class="nav-badge" v-if="!isCollapsed && linkCount > 0">{{ linkCount }}</span>
          </RouterLink>
          <RouterLink to="/analytics" class="nav-item" title="Analytics">
            <i class="bi bi-graph-up"></i>
            <span v-if="!isCollapsed">Analytics</span>
          </RouterLink>
        </div>

        <!-- Tools Section -->
        <div class="nav-section">
          <div class="nav-section-title" v-if="!isCollapsed">Tools</div>
          <a href="#" class="nav-item" @click.prevent="showCreateModal" title="Create Link">
            <i class="bi bi-plus-circle"></i>
            <span v-if="!isCollapsed">Create Link</span>
          </a>
          <RouterLink to="/qr-codes" class="nav-item" title="QR Codes" v-if="false">
            <i class="bi bi-qr-code"></i>
            <span v-if="!isCollapsed">QR Codes</span>
          </RouterLink>
          <RouterLink to="/bio-links" class="nav-item" title="Link in Bio" v-if="false">
            <i class="bi bi-person-badge"></i>
            <span v-if="!isCollapsed">Link in Bio</span>
          </RouterLink>
          <RouterLink to="/bulk" class="nav-item" title="Bulk Operations" v-if="false">
            <i class="bi bi-stack"></i>
            <span v-if="!isCollapsed">Bulk Create</span>
          </RouterLink>
        </div>

        <!-- Organization Section -->
        <div class="nav-section">
          <div class="nav-section-title" v-if="!isCollapsed">Organization</div>
          <RouterLink to="/tags" class="nav-item" title="Tags" v-if="false">
            <i class="bi bi-tags"></i>
            <span v-if="!isCollapsed">Tags</span>
          </RouterLink>
          <RouterLink to="/folders" class="nav-item" title="Folders" v-if="false">
            <i class="bi bi-folder"></i>
            <span v-if="!isCollapsed">Folders</span>
          </RouterLink>
          <RouterLink to="/archive" class="nav-item" title="Archive" v-if="false">
            <i class="bi bi-archive"></i>
            <span v-if="!isCollapsed">Archive</span>
          </RouterLink>
        </div>

        <!-- Configuration Section -->
        <div class="nav-section">
          <div class="nav-section-title" v-if="!isCollapsed">Configuration</div>
          <RouterLink to="/domains" class="nav-item" title="Custom Domains" v-if="false">
            <i class="bi bi-globe"></i>
            <span v-if="!isCollapsed">Domains</span>
          </RouterLink>
          <RouterLink to="/branded" class="nav-item" title="Branded Links" v-if="false">
            <i class="bi bi-brush"></i>
            <span v-if="!isCollapsed">Branded Links</span>
          </RouterLink>
          <RouterLink to="/api" class="nav-item" title="API & Webhooks" v-if="false">
            <i class="bi bi-code-square"></i>
            <span v-if="!isCollapsed">API & Webhooks</span>
          </RouterLink>
          <RouterLink to="/integrations" class="nav-item" title="Integrations" v-if="false">
            <i class="bi bi-puzzle"></i>
            <span v-if="!isCollapsed">Integrations</span>
          </RouterLink>
        </div>

        <!-- Account Section -->
        <div class="nav-section">
          <div class="nav-section-title" v-if="!isCollapsed">Account</div>
          <RouterLink to="/settings" class="nav-item" title="Settings">
            <i class="bi bi-gear"></i>
            <span v-if="!isCollapsed">Settings</span>
          </RouterLink>
          <RouterLink to="/team" class="nav-item" title="Team" v-if="false">
            <i class="bi bi-people"></i>
            <span v-if="!isCollapsed">Team</span>
          </RouterLink>
          <RouterLink to="/billing" class="nav-item" title="Billing & Plans" v-if="false">
            <i class="bi bi-credit-card"></i>
            <span v-if="!isCollapsed">Billing</span>
          </RouterLink>
        </div>
      </nav>

      <!-- User Section -->
      <div class="sidebar-footer">
        <div class="user-section" :class="{ 'collapsed': isCollapsed }">
          <div class="user-dropdown" @click="toggleUserMenu">
            <div class="user-avatar">
              <i class="bi bi-person-fill"></i>
            </div>
            <div class="user-info" v-if="!isCollapsed">
              <div class="user-name">{{ userName }}</div>
              <div class="user-email">{{ shortEmail }}</div>
            </div>
            <i class="bi bi-three-dots-vertical" v-if="!isCollapsed"></i>
          </div>

          <!-- User Dropdown Menu -->
          <div class="user-menu" v-if="showUserMenu && !isCollapsed">
            <RouterLink to="/settings" class="user-menu-item">
              <i class="bi bi-gear"></i>
              <span>Settings</span>
            </RouterLink>
            <a href="#" class="user-menu-item" @click.prevent="handleLogout">
              <i class="bi bi-box-arrow-right"></i>
              <span>Logout</span>
            </a>
          </div>
        </div>
      </div>
    </aside>

    <!-- Mobile Overlay -->
    <div class="sidebar-overlay" v-if="isMobileMenuOpen" @click="closeMobileMenu"></div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const authStore = useAuthStore()

const isCollapsed = ref(false)
const isMobile = ref(false)
const isMobileMenuOpen = ref(false)
const showUserMenu = ref(false)
const linkCount = ref(0)

const userName = computed(() => {
  const email = authStore.currentUser?.email || 'User'
  return email.split('@')[0]
})

const shortEmail = computed(() => {
  const email = authStore.currentUser?.email || 'user@example.com'
  if (email.length > 25) {
    return email.substring(0, 25) + '...'
  }
  return email
})

const toggleSidebar = () => {
  isCollapsed.value = !isCollapsed.value
  localStorage.setItem('sidebarCollapsed', isCollapsed.value)
}

const toggleUserMenu = () => {
  showUserMenu.value = !showUserMenu.value
}

const closeMobileMenu = () => {
  isMobileMenuOpen.value = false
}

const showCreateModal = () => {
  // This will be handled by the parent component
  window.dispatchEvent(new CustomEvent('show-create-modal'))
}

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}

const checkMobile = () => {
  isMobile.value = window.innerWidth < 992
  if (isMobile.value) {
    isCollapsed.value = false
  }
}

// Click outside to close user menu
const handleClickOutside = (event) => {
  const userSection = document.querySelector('.user-section')
  if (userSection && !userSection.contains(event.target)) {
    showUserMenu.value = false
  }
}

onMounted(() => {
  checkMobile()
  window.addEventListener('resize', checkMobile)
  document.addEventListener('click', handleClickOutside)

  // Restore sidebar state
  const saved = localStorage.getItem('sidebarCollapsed')
  if (saved !== null && !isMobile.value) {
    isCollapsed.value = saved === 'true'
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', checkMobile)
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.sidebar-wrapper {
  position: relative;
}

.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  width: 260px;
  background-color: var(--bg-white);
  border-right: 1px solid var(--border-light);
  display: flex;
  flex-direction: column;
  transition: width var(--transition-base), transform var(--transition-base);
  z-index: 1000;
  box-shadow: var(--shadow-sm);
}

.sidebar.collapsed {
  width: 72px;
}

/* Header */
.sidebar-header {
  padding: 20px 16px;
  border-bottom: 1px solid var(--border-light);
  display: flex;
  align-items: center;
  justify-content: space-between;
  min-height: 72px;
}

.sidebar-brand {
  display: flex;
  align-items: center;
  gap: 12px;
  text-decoration: none;
  transition: all var(--transition-base);
  flex: 1;
}

.sidebar-brand:hover {
  opacity: 0.85;
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
  flex-shrink: 0;
}

.sidebar-brand:hover .brand-icon {
  transform: scale(1.05);
  box-shadow: var(--shadow-md);
}

.brand-text {
  font-weight: 700;
  font-size: 18px;
  letter-spacing: -0.5px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  white-space: nowrap;
}

.collapse-btn {
  width: 28px;
  height: 28px;
  border-radius: var(--radius-md);
  border: 1px solid var(--border-light);
  background-color: var(--bg-white);
  color: var(--text-tertiary);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all var(--transition-base);
  flex-shrink: 0;
}

.collapse-btn:hover {
  background-color: var(--bg-hover);
  border-color: var(--border-hover);
  color: var(--text-primary);
}

/* Navigation */
.sidebar-nav {
  flex: 1;
  overflow-y: auto;
  overflow-x: hidden;
  padding: 16px 12px;
}

.sidebar-nav::-webkit-scrollbar {
  width: 4px;
}

.sidebar-nav::-webkit-scrollbar-thumb {
  background-color: var(--border-color);
  border-radius: 4px;
}

.nav-section {
  margin-bottom: 24px;
}

.nav-section:last-child {
  margin-bottom: 0;
}

.nav-section-title {
  font-size: 11px;
  font-weight: 600;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.08em;
  padding: 8px 12px 8px 12px;
  margin-bottom: 4px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  margin-bottom: 2px;
  border-radius: var(--radius-lg);
  color: var(--text-secondary);
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: all var(--transition-base);
  position: relative;
  cursor: pointer;
}

.sidebar.collapsed .nav-item {
  justify-content: center;
  padding: 10px 12px;
}

.nav-item i {
  font-size: 18px;
  flex-shrink: 0;
  transition: all var(--transition-base);
}

.nav-item span {
  flex: 1;
  white-space: nowrap;
}

.nav-item:hover {
  background-color: var(--bg-hover);
  color: var(--text-primary);
  transform: translateX(2px);
}

.sidebar.collapsed .nav-item:hover {
  transform: translateX(0);
}

.nav-item:hover i {
  transform: scale(1.1);
}

.nav-item.router-link-active {
  background-color: var(--primary-subtle);
  color: var(--primary-color);
  font-weight: 600;
}

.nav-item.router-link-active:hover {
  background-color: var(--primary-subtle-hover);
}

.nav-item.router-link-active::before {
  content: '';
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 3px;
  height: 20px;
  background-color: var(--primary-color);
  border-radius: 0 3px 3px 0;
}

.nav-badge {
  background-color: var(--primary-color);
  color: white;
  padding: 2px 8px;
  border-radius: var(--radius-sm);
  font-size: 11px;
  font-weight: 600;
  margin-left: auto;
}

/* Footer / User Section */
.sidebar-footer {
  border-top: 1px solid var(--border-light);
  padding: 12px;
}

.user-section {
  position: relative;
}

.user-dropdown {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  border-radius: var(--radius-lg);
  cursor: pointer;
  transition: all var(--transition-base);
}

.user-section.collapsed .user-dropdown {
  justify-content: center;
  padding: 10px;
}

.user-dropdown:hover {
  background-color: var(--bg-hover);
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 16px;
  box-shadow: var(--shadow-xs);
  flex-shrink: 0;
  transition: all var(--transition-base);
}

.user-dropdown:hover .user-avatar {
  transform: scale(1.05);
  box-shadow: var(--shadow-sm);
}

.user-info {
  flex: 1;
  min-width: 0;
}

.user-name {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-primary);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  line-height: 1.3;
}

.user-email {
  font-size: 12px;
  color: var(--text-tertiary);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  line-height: 1.3;
}

.user-dropdown > i {
  color: var(--text-tertiary);
  font-size: 16px;
  flex-shrink: 0;
}

/* User Menu Dropdown */
.user-menu {
  position: absolute;
  bottom: 100%;
  left: 12px;
  right: 12px;
  margin-bottom: 8px;
  background-color: var(--bg-white);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-xl);
  padding: 8px;
  animation: slideUp 0.2s ease;
  z-index: 10;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.user-menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  border-radius: var(--radius-md);
  color: var(--text-secondary);
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  transition: all var(--transition-base);
  cursor: pointer;
}

.user-menu-item:hover {
  background-color: var(--bg-hover);
  color: var(--text-primary);
  transform: translateX(2px);
}

.user-menu-item i {
  font-size: 16px;
  width: 20px;
  text-align: center;
}

/* Mobile Styles */
@media (max-width: 991px) {
  .sidebar {
    transform: translateX(-100%);
  }

  .sidebar.open {
    transform: translateX(0);
  }

  .sidebar-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(10, 37, 64, 0.5);
    backdrop-filter: blur(2px);
    z-index: 999;
  }

  .collapse-btn {
    display: none;
  }
}

/* Scrollbar styling */
.sidebar-nav::-webkit-scrollbar {
  width: 6px;
}

.sidebar-nav::-webkit-scrollbar-track {
  background: transparent;
}

.sidebar-nav::-webkit-scrollbar-thumb {
  background-color: var(--border-color);
  border-radius: 3px;
}

.sidebar-nav::-webkit-scrollbar-thumb:hover {
  background-color: var(--border-hover);
}
</style>
