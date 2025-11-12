<template>
  <div class="security-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Security Settings</h1>
          <p class="text-muted">Manage your account security and authentication</p>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Change Password</h6>
            </div>
            <div class="card-body">
              <div v-if="successMessage" class="alert alert-success alert-dismissible fade show" role="alert">
                <i class="bi bi-check-circle"></i> {{ successMessage }}
                <button type="button" class="btn-close" @click="successMessage = ''" aria-label="Close"></button>
              </div>
              <div v-if="errorMessage" class="alert alert-danger alert-dismissible fade show" role="alert">
                <i class="bi bi-exclamation-circle"></i> {{ errorMessage }}
                <button type="button" class="btn-close" @click="errorMessage = ''" aria-label="Close"></button>
              </div>
              <form @submit.prevent="handleChangePassword">
                <div class="mb-3">
                  <label class="form-label">Current Password</label>
                  <input
                    type="password"
                    class="form-control"
                    v-model="passwordForm.currentPassword"
                    placeholder="Enter current password"
                    required
                  >
                </div>
                <div class="mb-3">
                  <label class="form-label">New Password</label>
                  <input
                    type="password"
                    class="form-control"
                    v-model="passwordForm.newPassword"
                    placeholder="Enter new password (min 8 characters)"
                    minlength="8"
                    required
                  >
                </div>
                <div class="mb-3">
                  <label class="form-label">Confirm New Password</label>
                  <input
                    type="password"
                    class="form-control"
                    v-model="passwordForm.confirmPassword"
                    placeholder="Confirm new password"
                    required
                  >
                </div>
                <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                  <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-2"></span>
                  Update Password
                </button>
              </form>
            </div>
          </div>

          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Two-Factor Authentication</h6>
            </div>
            <div class="card-body">
              <p class="text-muted">Add an extra layer of security to your account by enabling two-factor authentication.</p>
              <div class="alert alert-warning">
                <i class="bi bi-shield-exclamation"></i> Two-factor authentication is currently <strong>disabled</strong>
              </div>
              <button class="btn btn-primary">Enable 2FA</button>
            </div>
          </div>

          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Active Sessions</h6>
            </div>
            <div class="card-body">
              <div class="session-item">
                <div class="d-flex justify-content-between align-items-center">
                  <div>
                    <div class="fw-bold"><i class="bi bi-laptop"></i> Current Session</div>
                    <small class="text-muted">Chrome on MacOS • United States</small>
                  </div>
                  <span class="badge bg-success">Active</span>
                </div>
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
import api from '@/services/api'

const passwordForm = ref({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const isSubmitting = ref(false)
const successMessage = ref('')
const errorMessage = ref('')

const handleChangePassword = async () => {
  // Reset messages
  successMessage.value = ''
  errorMessage.value = ''

  // Validate passwords match
  if (passwordForm.value.newPassword !== passwordForm.value.confirmPassword) {
    errorMessage.value = 'New passwords do not match'
    return
  }

  // Validate new password length
  if (passwordForm.value.newPassword.length < 8) {
    errorMessage.value = 'New password must be at least 8 characters long'
    return
  }

  try {
    isSubmitting.value = true

    await api.auth.changePassword(
      passwordForm.value.currentPassword,
      passwordForm.value.newPassword
    )

    successMessage.value = 'Password changed successfully!'

    // Reset form
    passwordForm.value = {
      currentPassword: '',
      newPassword: '',
      confirmPassword: ''
    }
  } catch (error) {
    errorMessage.value = error.response?.data?.error || 'Failed to change password. Please try again.'
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
.security-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.session-item {
  padding: 16px;
  border: 1px solid var(--border-light);
  border-radius: 8px;
}

/* Responsive Styles */
@media (max-width: 991px) {
  .container-fluid {
    padding: 1.5rem !important;
  }

  .col-lg-8 {
    margin-bottom: 1.5rem;
  }
}

@media (max-width: 768px) {
  .container-fluid {
    padding: 1rem !important;
  }

  .security-page h1 {
    font-size: 1.5rem;
  }

  .card-header h6 {
    font-size: 0.95rem;
  }

  .card-body {
    padding: 1rem;
  }

  .session-item {
    padding: 12px;
  }

  .d-flex.justify-content-between {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 0.5rem;
  }
}

@media (max-width: 576px) {
  .container-fluid {
    padding: 0.75rem !important;
  }

  .security-page h1 {
    font-size: 1.25rem;
  }

  .security-page .text-muted {
    font-size: 0.9rem;
  }

  .card {
    border-radius: 8px;
    margin-bottom: 1rem !important;
  }

  .card-header {
    padding: 0.75rem 1rem;
  }

  .card-header h6 {
    font-size: 0.9rem;
  }

  .card-body {
    padding: 0.75rem;
  }

  .form-label {
    font-size: 0.9rem;
    margin-bottom: 0.25rem;
  }

  .form-control {
    padding: 0.5rem 0.75rem;
    font-size: 0.9rem;
  }

  .mb-3 {
    margin-bottom: 0.75rem !important;
  }

  .alert {
    padding: 0.75rem;
    font-size: 0.85rem;
  }

  .btn {
    font-size: 0.9rem;
    padding: 0.5rem 0.75rem;
  }

  .session-item {
    padding: 10px;
  }

  .fw-bold {
    font-size: 0.9rem;
  }

  small {
    font-size: 0.8rem;
  }
}
</style>
