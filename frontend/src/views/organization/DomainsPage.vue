<template>
  <div class="domains-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <div class="d-flex justify-content-between align-items-center">
            <div>
              <h1 class="h3 mb-0">Custom Domains</h1>
              <p class="text-muted mb-0">Use your own domain for short links</p>
            </div>
            <button class="btn btn-primary" :disabled="!canAddDomain">
              <i class="bi bi-plus-circle"></i> Add Domain
            </button>
          </div>
        </div>
      </div>

      <div class="row">
        <div class="col">
          <div v-if="!canAddDomain" class="alert alert-info">
            <i class="bi bi-info-circle"></i>
            Custom domains are available for Pro and Business plans. <RouterLink to="/account/billing">Upgrade your plan</RouterLink> to use this feature.
          </div>

          <div class="card">
            <div class="card-body">
              <div class="text-center py-5">
                <i class="bi bi-globe" style="font-size: 64px; color: #ccc;"></i>
                <h4 class="mt-3">No Custom Domains Yet</h4>
                <p class="text-muted">Add your own domain to create branded short links</p>
                <button class="btn btn-primary" :disabled="!canAddDomain">
                  <i class="bi bi-plus-circle"></i> Add Your First Domain
                </button>
              </div>

              <hr class="my-4">

              <div class="row">
                <div class="col-md-4">
                  <div class="feature-card">
                    <i class="bi bi-shield-check text-primary" style="font-size: 32px;"></i>
                    <h6 class="mt-3">SSL Included</h6>
                    <p class="text-muted small">Automatic SSL certificates for all custom domains</p>
                  </div>
                </div>
                <div class="col-md-4">
                  <div class="feature-card">
                    <i class="bi bi-lightning-charge text-warning" style="font-size: 32px;"></i>
                    <h6 class="mt-3">Fast Setup</h6>
                    <p class="text-muted small">Connect your domain in minutes with simple DNS configuration</p>
                  </div>
                </div>
                <div class="col-md-4">
                  <div class="feature-card">
                    <i class="bi bi-graph-up text-success" style="font-size: 32px;"></i>
                    <h6 class="mt-3">Full Analytics</h6>
                    <p class="text-muted small">Track all metrics for links using your custom domain</p>
                  </div>
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
import { computed } from 'vue'
import { RouterLink } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const authStore = useAuthStore()
const canAddDomain = computed(() => {
  const tier = authStore.currentUser?.subscription_tier || 'free'
  return tier === 'pro' || tier === 'business'
})
</script>

<style scoped>
.domains-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.feature-card {
  text-align: center;
  padding: 20px;
}
</style>
