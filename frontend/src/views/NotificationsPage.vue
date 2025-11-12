<template>
  <div class="notifications-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Notifications</h1>
          <p class="text-muted">View and manage all your notifications</p>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card">
            <div class="card-body">
              <div class="notification-list">
                <div
                  v-for="notification in notifications"
                  :key="notification.id"
                  class="notification-item"
                  :class="{ 'unread': !notification.read }"
                >
                  <div class="notification-icon-wrapper" :class="notification.type">
                    <i :class="getNotificationIcon(notification.type)"></i>
                  </div>
                  <div class="notification-content">
                    <div class="d-flex justify-content-between align-items-start mb-2">
                      <h6 class="notification-title mb-0">{{ notification.title }}</h6>
                      <span class="notification-time">{{ notification.time }}</span>
                    </div>
                    <p class="notification-message mb-0">{{ notification.message }}</p>
                  </div>
                  <button
                    v-if="!notification.read"
                    class="btn btn-sm btn-link"
                    @click="markAsRead(notification.id)"
                  >
                    Mark as read
                  </button>
                </div>
              </div>

              <div v-if="notifications.length === 0" class="text-center py-5">
                <i class="bi bi-inbox" style="font-size: 48px; color: #ccc;"></i>
                <p class="text-muted mt-3">No notifications</p>
              </div>
            </div>
          </div>
        </div>

        <div class="col-lg-4">
          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Notification Settings</h6>
            </div>
            <div class="card-body">
              <div class="form-check form-switch mb-3">
                <input class="form-check-input" type="checkbox" id="emailNotif" checked>
                <label class="form-check-label" for="emailNotif">
                  Email Notifications
                </label>
              </div>
              <div class="form-check form-switch mb-3">
                <input class="form-check-input" type="checkbox" id="linkCreated" checked>
                <label class="form-check-label" for="linkCreated">
                  Link Created
                </label>
              </div>
              <div class="form-check form-switch mb-3">
                <input class="form-check-input" type="checkbox" id="linkExpiring" checked>
                <label class="form-check-label" for="linkExpiring">
                  Link Expiring
                </label>
              </div>
              <div class="form-check form-switch">
                <input class="form-check-input" type="checkbox" id="usageAlerts" checked>
                <label class="form-check-label" for="usageAlerts">
                  Usage Alerts
                </label>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

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

const getNotificationIcon = (type) => {
  const icons = {
    success: 'bi bi-check-circle-fill',
    info: 'bi bi-info-circle-fill',
    warning: 'bi bi-exclamation-triangle-fill',
    error: 'bi bi-x-circle-fill'
  }
  return icons[type] || 'bi bi-bell-fill'
}

const markAsRead = (id) => {
  const notification = notifications.value.find(n => n.id === id)
  if (notification) {
    notification.read = true
  }
}
</script>

<style scoped>
.notifications-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.notification-list {
  display: flex;
  flex-direction: column;
  gap: 0;
}

.notification-item {
  display: flex;
  gap: 16px;
  padding: 20px;
  border-bottom: 1px solid var(--border-light);
  transition: all 0.15s ease;
}

.notification-item:last-child {
  border-bottom: none;
}

.notification-item.unread {
  background-color: var(--primary-subtle);
}

.notification-item:hover {
  background-color: var(--bg-secondary);
}

.notification-icon-wrapper {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
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
  font-size: 15px;
  color: var(--text-primary);
}

.notification-message {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.5;
}

.notification-time {
  font-size: 13px;
  color: var(--text-tertiary);
  white-space: nowrap;
}

/* Responsive Styles */
@media (max-width: 991px) {
  .container-fluid {
    padding: 1.5rem !important;
  }

  .col-lg-8,
  .col-lg-4 {
    margin-bottom: 1.5rem;
  }
}

@media (max-width: 768px) {
  .container-fluid {
    padding: 1rem !important;
  }

  .notifications-page h1 {
    font-size: 1.5rem;
  }

  .notification-item {
    padding: 16px;
    gap: 12px;
  }

  .notification-icon-wrapper {
    width: 36px;
    height: 36px;
    font-size: 16px;
  }

  .notification-title {
    font-size: 14px;
  }

  .notification-message {
    font-size: 13px;
  }

  .notification-time {
    font-size: 12px;
  }

  .card-header h6 {
    font-size: 0.95rem;
  }
}

@media (max-width: 576px) {
  .container-fluid {
    padding: 0.75rem !important;
  }

  .notifications-page h1 {
    font-size: 1.25rem;
  }

  .notifications-page .text-muted {
    font-size: 0.9rem;
  }

  .notification-item {
    padding: 12px;
    gap: 10px;
    flex-wrap: wrap;
  }

  .notification-icon-wrapper {
    width: 32px;
    height: 32px;
    font-size: 14px;
  }

  .notification-content {
    flex: 1;
    min-width: 0;
  }

  .d-flex.justify-content-between {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 4px;
  }

  .notification-title {
    font-size: 13px;
  }

  .notification-message {
    font-size: 12px;
  }

  .notification-time {
    font-size: 11px;
  }

  .btn-sm {
    padding: 0.25rem 0.5rem;
    font-size: 0.8rem;
    margin-top: 8px;
    width: 100%;
  }

  .form-check-label {
    font-size: 14px;
  }
}
</style>
