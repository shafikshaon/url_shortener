<template>
  <div class="profile-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Profile Settings</h1>
          <p class="text-muted">Manage your account profile information</p>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Personal Information</h6>
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
              <form @submit.prevent="handleSaveProfile">
                <div class="mb-3">
                  <label class="form-label">Full Name</label>
                  <input
                    type="text"
                    class="form-control"
                    v-model="profileForm.fullName"
                    placeholder="Enter your full name"
                  >
                </div>
                <div class="mb-3">
                  <label class="form-label">Email Address</label>
                  <input type="email" class="form-control" :value="userEmail" disabled>
                  <small class="text-muted">Email cannot be changed</small>
                </div>
                <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                  <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-2"></span>
                  Save Changes
                </button>
              </form>
            </div>
          </div>

          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Profile Picture</h6>
            </div>
            <div class="card-body">
              <div class="d-flex align-items-center gap-3">
                <div class="avatar-large">
                  <i class="bi bi-person-fill"></i>
                </div>
                <div>
                  <button class="btn btn-outline-primary btn-sm">Upload Picture</button>
                  <p class="text-muted mb-0 mt-2" style="font-size: 13px;">JPG, PNG or GIF. Max size 2MB.</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="col-lg-4">
          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Account Status</h6>
            </div>
            <div class="card-body">
              <div class="mb-3">
                <small class="text-muted">Subscription Plan</small>
                <div class="h5 mb-0">{{ subscriptionPlan }}</div>
              </div>
              <div class="mb-3">
                <small class="text-muted">Member Since</small>
                <div>January 2024</div>
              </div>
              <div>
                <small class="text-muted">Account Status</small>
                <div><span class="badge bg-success">Active</span></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '@/store/auth'
import api from '@/services/api'

const authStore = useAuthStore()
const userEmail = computed(() => authStore.currentUser?.email || '')
const subscriptionPlan = computed(() => {
  const tier = authStore.currentUser?.subscription_tier || 'free'
  return tier.charAt(0).toUpperCase() + tier.slice(1)
})

const profileForm = ref({
  fullName: ''
})

const isSubmitting = ref(false)
const successMessage = ref('')
const errorMessage = ref('')

// Load profile data on mount
onMounted(async () => {
  try {
    const response = await api.auth.getProfile()
    profileForm.value.fullName = response.data.full_name || ''
  } catch (error) {
    console.error('Failed to load profile:', error)
  }
})

const handleSaveProfile = async () => {
  // Reset messages
  successMessage.value = ''
  errorMessage.value = ''

  try {
    isSubmitting.value = true

    const response = await api.auth.updateProfile(profileForm.value.fullName)

    // Update user data in store
    authStore.updateUser(response.data)

    successMessage.value = 'Profile updated successfully!'
  } catch (error) {
    errorMessage.value = error.response?.data?.error || 'Failed to update profile. Please try again.'
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
.profile-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.avatar-large {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 36px;
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

  .profile-page h1 {
    font-size: 1.5rem;
  }

  .card-header h6 {
    font-size: 0.95rem;
  }

  .card-body {
    padding: 1rem;
  }

  .d-flex.align-items-center {
    flex-direction: column;
    align-items: flex-start !important;
  }

  .avatar-large {
    width: 64px;
    height: 64px;
    font-size: 28px;
  }
}

@media (max-width: 576px) {
  .container-fluid {
    padding: 0.75rem !important;
  }

  .profile-page h1 {
    font-size: 1.25rem;
  }

  .profile-page .text-muted {
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

  .btn-sm {
    font-size: 0.8rem;
    padding: 0.35rem 0.6rem;
  }

  .avatar-large {
    width: 56px;
    height: 56px;
    font-size: 24px;
    margin-bottom: 0.75rem;
  }

  .h5 {
    font-size: 1.1rem;
  }

  small {
    font-size: 0.85rem;
  }
}
</style>
