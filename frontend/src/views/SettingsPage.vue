<template>
  <div class="settings-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Settings</h1>
          <p class="text-muted">Manage your account and preferences</p>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <!-- Profile Information -->
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h5 class="mb-0">Profile Information</h5>
            </div>
            <div class="card-body">
              <div class="mb-3">
                <label class="form-label">Email</label>
                <input
                  type="email"
                  class="form-control"
                  :value="user?.email"
                  disabled
                />
                <div class="form-text">Your email address cannot be changed</div>
              </div>

              <div class="mb-3">
                <label class="form-label">Subscription Tier</label>
                <input
                  type="text"
                  class="form-control text-capitalize"
                  :value="user?.subscription_tier"
                  disabled
                />
              </div>

              <div class="mb-3">
                <label class="form-label">Member Since</label>
                <input
                  type="text"
                  class="form-control"
                  :value="formatDate(user?.created_at)"
                  disabled
                />
              </div>
            </div>
          </div>

          <!-- API Access -->
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h5 class="mb-0">API Access</h5>
            </div>
            <div class="card-body">
              <div v-if="user?.subscription_tier === 'free'" class="alert alert-warning">
                <i class="bi bi-lock"></i>
                API access is available on Pro and Business plans.
                <RouterLink to="/" class="alert-link">Upgrade your plan</RouterLink>
                to get API access.
              </div>

              <div v-else>
                <p class="text-muted">
                  Use your API key to integrate URL shortening into your applications.
                </p>

                <div class="mb-3">
                  <label class="form-label">API Key</label>
                  <div class="input-group">
                    <input
                      type="text"
                      class="form-control"
                      :value="apiKey || 'No API key generated'"
                      :type="showApiKey ? 'text' : 'password'"
                      readonly
                    />
                    <button
                      class="btn btn-outline-secondary"
                      @click="showApiKey = !showApiKey"
                    >
                      <i :class="showApiKey ? 'bi-eye-slash' : 'bi-eye'"></i>
                    </button>
                    <button
                      v-if="apiKey"
                      class="btn btn-outline-primary"
                      @click="copyAPIKey"
                    >
                      <i class="bi bi-clipboard"></i>
                    </button>
                  </div>
                </div>

                <button
                  class="btn btn-primary"
                  @click="generateAPIKey"
                  :disabled="generating"
                >
                  <span v-if="generating" class="spinner-border spinner-border-sm me-2"></span>
                  {{ apiKey ? 'Regenerate' : 'Generate' }} API Key
                </button>

                <div class="alert alert-info mt-3 mb-0">
                  <i class="bi bi-info-circle"></i>
                  <strong>Note:</strong> Regenerating your API key will invalidate the old key.
                  Make sure to update your applications.
                </div>
              </div>
            </div>
          </div>

          <!-- Subscription Plans -->
          <div class="card">
            <div class="card-header bg-white">
              <h5 class="mb-0">Subscription Plans</h5>
            </div>
            <div class="card-body">
              <div class="row g-3">
                <div class="col-md-4">
                  <div class="card text-center" :class="{ 'border-primary': user?.subscription_tier === 'free' }">
                    <div class="card-body">
                      <h5>Free</h5>
                      <h3>$0</h3>
                      <ul class="list-unstyled text-start small">
                        <li><i class="bi bi-check text-success"></i> 50 links</li>
                        <li><i class="bi bi-check text-success"></i> 1K clicks/month</li>
                        <li><i class="bi bi-check text-success"></i> 30-day analytics</li>
                      </ul>
                      <button class="btn btn-sm btn-outline-secondary" disabled>
                        Current Plan
                      </button>
                    </div>
                  </div>
                </div>

                <div class="col-md-4">
                  <div class="card text-center" :class="{ 'border-primary': user?.subscription_tier === 'pro' }">
                    <div class="card-body">
                      <h5>Pro</h5>
                      <h3>$9<small>/mo</small></h3>
                      <ul class="list-unstyled text-start small">
                        <li><i class="bi bi-check text-success"></i> 500 links</li>
                        <li><i class="bi bi-check text-success"></i> 10K clicks/month</li>
                        <li><i class="bi bi-check text-success"></i> 1-year analytics</li>
                        <li><i class="bi bi-check text-success"></i> API access</li>
                      </ul>
                      <button class="btn btn-sm btn-primary" disabled>
                        {{ user?.subscription_tier === 'pro' ? 'Current Plan' : 'Coming Soon' }}
                      </button>
                    </div>
                  </div>
                </div>

                <div class="col-md-4">
                  <div class="card text-center" :class="{ 'border-primary': user?.subscription_tier === 'business' }">
                    <div class="card-body">
                      <h5>Business</h5>
                      <h3>$29<small>/mo</small></h3>
                      <ul class="list-unstyled text-start small">
                        <li><i class="bi bi-check text-success"></i> 5K links</li>
                        <li><i class="bi bi-check text-success"></i> 100K clicks/month</li>
                        <li><i class="bi bi-check text-success"></i> Unlimited analytics</li>
                        <li><i class="bi bi-check text-success"></i> Custom domain</li>
                      </ul>
                      <button class="btn btn-sm btn-primary" disabled>
                        {{ user?.subscription_tier === 'business' ? 'Current Plan' : 'Coming Soon' }}
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Sidebar -->
        <div class="col-lg-4">
          <div class="card mb-4">
            <div class="card-body">
              <h5 class="card-title">Account Actions</h5>
              <div class="d-grid gap-2">
                <button class="btn btn-outline-primary" disabled>
                  <i class="bi bi-key"></i> Change Password
                </button>
                <button class="btn btn-outline-danger" disabled>
                  <i class="bi bi-trash"></i> Delete Account
                </button>
              </div>
              <p class="text-muted small mt-3 mb-0">
                These features are coming soon.
              </p>
            </div>
          </div>

          <div class="card">
            <div class="card-body">
              <h5 class="card-title">Documentation</h5>
              <p class="text-muted small">
                Learn how to use the API and integrate with your applications.
              </p>
              <a href="#" class="btn btn-outline-primary btn-sm w-100" disabled>
                <i class="bi bi-book"></i> View API Docs
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import api from '@/services/api'

const authStore = useAuthStore()
const apiKey = ref(null)
const showApiKey = ref(false)
const generating = ref(false)

const user = computed(() => authStore.currentUser)

const generateAPIKey = async () => {
  if (!confirm('This will invalidate your existing API key. Continue?')) {
    return
  }

  generating.value = true
  try {
    const response = await api.auth.generateAPIKey()
    apiKey.value = response.data.api_key
    await authStore.refreshProfile()
  } catch (error) {
    alert('Failed to generate API key: ' + (error.response?.data?.error || error.message))
  } finally {
    generating.value = false
  }
}

const copyAPIKey = async () => {
  try {
    await navigator.clipboard.writeText(apiKey.value)
    alert('API key copied to clipboard!')
  } catch (err) {
    console.error('Failed to copy:', err)
  }
}

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

onMounted(() => {
  apiKey.value = user.value?.api_key || null
})
</script>

<style scoped>
.settings-page {
  background-color: #F9FAFB;
  min-height: 100vh;
}

.border-primary {
  border-width: 2px;
  border-color: #7C3AED !important;
}

.alert-warning {
  background-color: #FEF3C7;
  color: #92400E;
  border: none;
}

.alert-info {
  background-color: #DBEAFE;
  color: #1E40AF;
  border: none;
}

.input-group .btn-outline-secondary {
  border-color: #E5E7EB;
  color: #6B7280;
}

.input-group .btn-outline-secondary:hover {
  background-color: #F3F4F6;
  color: #374151;
}

.card-title {
  color: #111827;
  font-weight: 600;
  margin-bottom: 1rem;
}
</style>
