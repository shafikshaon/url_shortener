<template>
  <div class="billing-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Billing & Subscription</h1>
          <p class="text-muted">Manage your subscription and billing information</p>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Current Plan</h6>
            </div>
            <div class="card-body">
              <div class="d-flex justify-content-between align-items-center mb-3">
                <div>
                  <h4 class="mb-1">{{ subscriptionPlan }} Plan</h4>
                  <p class="text-muted mb-0">{{ getPlanPrice() }} / month</p>
                </div>
                <button class="btn btn-primary">Upgrade Plan</button>
              </div>
              <hr>
              <div class="row">
                <div class="col-md-4">
                  <small class="text-muted">Links Limit</small>
                  <div class="h6">{{ getLinkLimit() }}</div>
                </div>
                <div class="col-md-4">
                  <small class="text-muted">Clicks / Month</small>
                  <div class="h6">{{ getClickLimit() }}</div>
                </div>
                <div class="col-md-4">
                  <small class="text-muted">Custom Domains</small>
                  <div class="h6">{{ getCustomDomains() }}</div>
                </div>
              </div>
            </div>
          </div>

          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Available Plans</h6>
            </div>
            <div class="card-body">
              <div class="row">
                <div class="col-md-4 mb-3">
                  <div class="plan-card">
                    <h6>Free</h6>
                    <div class="plan-price">$0</div>
                    <ul class="plan-features">
                      <li>50 Links</li>
                      <li>1,000 Clicks/mo</li>
                      <li>Basic Analytics</li>
                    </ul>
                    <button class="btn btn-outline-secondary btn-sm w-100">Current Plan</button>
                  </div>
                </div>
                <div class="col-md-4 mb-3">
                  <div class="plan-card featured">
                    <h6>Pro</h6>
                    <div class="plan-price">$9</div>
                    <ul class="plan-features">
                      <li>500 Links</li>
                      <li>10,000 Clicks/mo</li>
                      <li>Advanced Analytics</li>
                      <li>Custom Domains</li>
                    </ul>
                    <button class="btn btn-primary btn-sm w-100">Upgrade</button>
                  </div>
                </div>
                <div class="col-md-4 mb-3">
                  <div class="plan-card">
                    <h6>Business</h6>
                    <div class="plan-price">$29</div>
                    <ul class="plan-features">
                      <li>5,000 Links</li>
                      <li>100,000 Clicks/mo</li>
                      <li>Full Analytics</li>
                      <li>Unlimited Domains</li>
                      <li>API Access</li>
                    </ul>
                    <button class="btn btn-outline-primary btn-sm w-100">Upgrade</button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Billing History</h6>
            </div>
            <div class="card-body">
              <div class="table-responsive">
                <table class="table">
                  <thead>
                    <tr>
                      <th>Date</th>
                      <th>Description</th>
                      <th>Amount</th>
                      <th>Status</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>Nov 1, 2024</td>
                      <td>Free Plan</td>
                      <td>$0.00</td>
                      <td><span class="badge bg-success">Paid</span></td>
                    </tr>
                  </tbody>
                </table>
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
const subscriptionPlan = computed(() => {
  const tier = authStore.currentUser?.subscription_tier || 'free'
  return tier.charAt(0).toUpperCase() + tier.slice(1)
})

const getPlanPrice = () => {
  const prices = { free: '$0', pro: '$9', business: '$29' }
  return prices[authStore.currentUser?.subscription_tier || 'free']
}

const getLinkLimit = () => {
  const limits = { free: '50', pro: '500', business: '5,000' }
  return limits[authStore.currentUser?.subscription_tier || 'free']
}

const getClickLimit = () => {
  const limits = { free: '1,000', pro: '10,000', business: '100,000' }
  return limits[authStore.currentUser?.subscription_tier || 'free']
}

const getCustomDomains = () => {
  const domains = { free: 'None', pro: '1', business: 'Unlimited' }
  return domains[authStore.currentUser?.subscription_tier || 'free']
}
</script>

<style scoped>
.billing-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.plan-card {
  border: 2px solid var(--border-light);
  border-radius: 8px;
  padding: 20px;
  text-align: center;
  transition: all 0.2s ease;
}

.plan-card.featured {
  border-color: var(--primary-color);
  box-shadow: var(--shadow-md);
}

.plan-card:hover {
  box-shadow: var(--shadow-lg);
}

.plan-price {
  font-size: 32px;
  font-weight: 700;
  color: var(--primary-color);
  margin: 12px 0;
}

.plan-features {
  list-style: none;
  padding: 0;
  margin: 20px 0;
  text-align: left;
}

.plan-features li {
  padding: 8px 0;
  font-size: 14px;
  color: var(--text-secondary);
}

.plan-features li::before {
  content: '✓ ';
  color: var(--success-color);
  font-weight: bold;
  margin-right: 8px;
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

  .billing-page h1 {
    font-size: 1.5rem;
  }

  .card-header h6 {
    font-size: 0.95rem;
  }

  .card-body {
    padding: 1rem;
  }

  .d-flex.justify-content-between {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 1rem;
  }

  .d-flex.justify-content-between .btn {
    width: 100%;
  }

  h4 {
    font-size: 1.35rem;
  }

  .col-md-4 {
    margin-bottom: 1rem;
  }

  .plan-card {
    padding: 16px;
  }

  .plan-price {
    font-size: 28px;
  }

  .table {
    font-size: 0.9rem;
  }
}

@media (max-width: 576px) {
  .container-fluid {
    padding: 0.75rem !important;
  }

  .billing-page h1 {
    font-size: 1.25rem;
  }

  .billing-page .text-muted {
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

  h4 {
    font-size: 1.2rem;
  }

  .h6 {
    font-size: 1rem;
  }

  .mb-3 {
    margin-bottom: 0.75rem !important;
  }

  .btn {
    font-size: 0.9rem;
    padding: 0.5rem 0.75rem;
  }

  .btn-sm {
    font-size: 0.8rem;
    padding: 0.35rem 0.6rem;
  }

  .plan-card {
    padding: 12px;
  }

  .plan-price {
    font-size: 24px;
    margin: 8px 0;
  }

  .plan-features {
    margin: 16px 0;
  }

  .plan-features li {
    padding: 6px 0;
    font-size: 13px;
  }

  .table {
    font-size: 0.85rem;
  }

  .table th,
  .table td {
    padding: 0.5rem 0.4rem;
  }

  small {
    font-size: 0.8rem;
  }
}
</style>
