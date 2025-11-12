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
              <form>
                <div class="mb-3">
                  <label class="form-label">Full Name</label>
                  <input type="text" class="form-control" placeholder="Enter your full name" value="John Doe">
                </div>
                <div class="mb-3">
                  <label class="form-label">Email Address</label>
                  <input type="email" class="form-control" :value="userEmail" disabled>
                  <small class="text-muted">Email cannot be changed</small>
                </div>
                <div class="mb-3">
                  <label class="form-label">Company / Organization</label>
                  <input type="text" class="form-control" placeholder="Enter company name">
                </div>
                <div class="mb-3">
                  <label class="form-label">Job Title</label>
                  <input type="text" class="form-control" placeholder="Enter job title">
                </div>
                <div class="mb-3">
                  <label class="form-label">Time Zone</label>
                  <select class="form-select">
                    <option>UTC</option>
                    <option>America/New_York</option>
                    <option>America/Los_Angeles</option>
                    <option>Europe/London</option>
                    <option>Asia/Tokyo</option>
                  </select>
                </div>
                <button type="submit" class="btn btn-primary">Save Changes</button>
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
import { computed } from 'vue'
import { useAuthStore } from '@/store/auth'

const authStore = useAuthStore()
const userEmail = computed(() => authStore.currentUser?.email || '')
const subscriptionPlan = computed(() => {
  const tier = authStore.currentUser?.subscription_tier || 'free'
  return tier.charAt(0).toUpperCase() + tier.slice(1)
})
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
</style>
