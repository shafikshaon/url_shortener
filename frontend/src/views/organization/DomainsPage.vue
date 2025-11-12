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

  .domains-page h1 {
    font-size: 1.5rem;
  }

  .domains-page h4 {
    font-size: 1.25rem;
  }

  .card-header h5,
  .card-header h6 {
    font-size: 0.95rem;
  }

  .card-body {
    padding: 1rem;
  }

  .col-md-4,
  .col-md-6 {
    margin-bottom: 1rem;
  }

  .d-flex.justify-content-between {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 0.5rem;
  }

  .feature-card {
    padding: 15px;
  }

  .feature-card h6 {
    font-size: 0.95rem;
  }

  .feature-card i {
    font-size: 28px !important;
  }
}

@media (max-width: 576px) {
  .container-fluid {
    padding: 0.75rem !important;
  }

  .domains-page h1 {
    font-size: 1.25rem;
  }

  .domains-page h4 {
    font-size: 1.1rem;
  }

  .domains-page .text-muted {
    font-size: 0.9rem;
  }

  .card {
    border-radius: 8px;
    margin-bottom: 1rem !important;
  }

  .card-header {
    padding: 0.75rem 1rem;
  }

  .card-header h5,
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

  .feature-card {
    padding: 12px;
  }

  .feature-card h6 {
    font-size: 0.9rem;
    margin-top: 0.5rem !important;
  }

  .feature-card i {
    font-size: 24px !important;
  }

  .feature-card .small {
    font-size: 0.8rem;
  }

  .text-center i[style*="font-size: 64px"] {
    font-size: 48px !important;
  }

  small {
    font-size: 0.8rem;
  }
}
</style>
