<template>
  <div class="analytics-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Analytics</h1>
          <p class="text-muted">Overall statistics for all your links</p>
        </div>
      </div>

      <div v-if="loading" class="text-center py-5">
        <div class="spinner-border text-primary" role="status">
          <span class="visually-hidden">Loading...</span>
        </div>
      </div>

      <div v-else>
        <!-- Summary Cards -->
        <div class="row g-3 mb-4">
          <div class="col-md-4">
            <div class="card stat-card">
              <div class="card-body">
                <div class="d-flex justify-content-between align-items-center">
                  <div>
                    <p class="text-muted mb-1">Total Links</p>
                    <h2 class="mb-0">{{ stats.total_links || 0 }}</h2>
                  </div>
                  <div class="stat-icon bg-primary">
                    <i class="bi bi-link-45deg"></i>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="col-md-4">
            <div class="card stat-card">
              <div class="card-body">
                <div class="d-flex justify-content-between align-items-center">
                  <div>
                    <p class="text-muted mb-1">Total Clicks</p>
                    <h2 class="mb-0">{{ stats.total_clicks || 0 }}</h2>
                  </div>
                  <div class="stat-icon bg-success">
                    <i class="bi bi-cursor-fill"></i>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="col-md-4">
            <div class="card stat-card">
              <div class="card-body">
                <div class="d-flex justify-content-between align-items-center">
                  <div>
                    <p class="text-muted mb-1">Clicks This Month</p>
                    <h2 class="mb-0">{{ stats.clicks_this_month || 0 }}</h2>
                  </div>
                  <div class="stat-icon bg-info">
                    <i class="bi bi-calendar-check-fill"></i>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Info Cards -->
        <div class="row">
          <div class="col-md-6 mb-4">
            <div class="card">
              <div class="card-body text-center py-5">
                <i class="bi bi-graph-up text-primary mb-3" style="font-size: 3rem"></i>
                <h4>Link Performance</h4>
                <p class="text-muted">
                  Track individual link performance by viewing each link's detailed analytics
                </p>
                <RouterLink to="/links" class="btn btn-primary">
                  View All Links
                </RouterLink>
              </div>
            </div>
          </div>

          <div class="col-md-6 mb-4">
            <div class="card">
              <div class="card-body text-center py-5">
                <i class="bi bi-speedometer2 text-success mb-3" style="font-size: 3rem"></i>
                <h4>Real-time Tracking</h4>
                <p class="text-muted">
                  Every click is tracked in real-time with detailed information about location and device
                </p>
                <button class="btn btn-success" @click="loadStats">
                  <i class="bi bi-arrow-clockwise"></i> Refresh Data
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Usage Information -->
        <div class="row">
          <div class="col">
            <div class="card">
              <div class="card-header bg-white">
                <h5 class="mb-0">Account Usage</h5>
              </div>
              <div class="card-body">
                <div class="row">
                  <div class="col-md-6 mb-3">
                    <div class="d-flex justify-content-between mb-2">
                      <span>Links Created</span>
                      <strong>{{ stats.total_links || 0 }} / {{ getLinkLimit() }}</strong>
                    </div>
                    <div class="progress">
                      <div
                        class="progress-bar"
                        :style="{ width: getLinkUsagePercent() + '%' }"
                      ></div>
                    </div>
                  </div>

                  <div class="col-md-6 mb-3">
                    <div class="d-flex justify-content-between mb-2">
                      <span>Clicks This Month</span>
                      <strong>{{ stats.clicks_this_month || 0 }} / {{ getClickLimit() }}</strong>
                    </div>
                    <div class="progress">
                      <div
                        class="progress-bar bg-success"
                        :style="{ width: getClickUsagePercent() + '%' }"
                      ></div>
                    </div>
                  </div>
                </div>

                <div class="alert alert-info mb-0 mt-3">
                  <i class="bi bi-info-circle"></i>
                  You are currently on the <strong>{{ subscriptionTier }}</strong> plan.
                  <RouterLink to="/settings" class="alert-link">Upgrade your plan</RouterLink>
                  to unlock more features.
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
import { ref, computed, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import api from '@/services/api'

const authStore = useAuthStore()
const stats = ref({})
const loading = ref(true)

const subscriptionTier = computed(() => {
  return authStore.currentUser?.subscription_tier || 'free'
})

const getLinkLimit = () => {
  const limits = { free: 50, pro: 500, business: 5000 }
  return limits[subscriptionTier.value] || 50
}

const getClickLimit = () => {
  const limits = { free: 1000, pro: 10000, business: 100000 }
  return limits[subscriptionTier.value] || 1000
}

const getLinkUsagePercent = () => {
  const usage = (stats.value.total_links || 0) / getLinkLimit() * 100
  return Math.min(usage, 100)
}

const getClickUsagePercent = () => {
  const usage = (stats.value.clicks_this_month || 0) / getClickLimit() * 100
  return Math.min(usage, 100)
}

const loadStats = async () => {
  loading.value = true
  try {
    const response = await api.analytics.getUserAnalytics()
    stats.value = response.data
  } catch (error) {
    console.error('Failed to load analytics:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadStats()
})
</script>

<style scoped>
.analytics-page {
  background-color: #F9FAFB;
  min-height: 100vh;
}

.stat-card {
  border: 1px solid #E5E7EB;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
  transition: all 0.3s ease;
}

.stat-card:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  transform: translateY(-2px);
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 0.75rem;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  color: white;
}

.progress {
  height: 10px;
  border-radius: 0.5rem;
  background-color: #E5E7EB;
}

.progress-bar {
  background-color: #7C3AED;
  border-radius: 0.5rem;
}

.progress-bar.bg-success {
  background-color: #10B981;
}

.alert-info {
  background-color: #DBEAFE;
  color: #1E40AF;
  border: none;
}
</style>
