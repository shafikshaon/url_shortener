<template>
  <div class="dashboard-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Dashboard</h1>
          <p class="text-muted">Overview of your links and analytics</p>
        </div>
      </div>

      <!-- Stats Cards -->
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
                  <i class="bi bi-cursor"></i>
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
                  <i class="bi bi-calendar-check"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="row mb-4">
        <div class="col">
          <div class="card">
            <div class="card-body">
              <h5 class="card-title mb-3">Quick Actions</h5>
              <div class="d-flex gap-2 flex-wrap">
                <button class="btn btn-primary" @click="showCreateModal = true">
                  <i class="bi bi-plus-circle"></i> Create New Link
                </button>
                <RouterLink to="/links" class="btn btn-outline-primary">
                  <i class="bi bi-list"></i> View All Links
                </RouterLink>
                <RouterLink to="/analytics" class="btn btn-outline-primary">
                  <i class="bi bi-bar-chart"></i> View Analytics
                </RouterLink>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Links -->
      <div class="row">
        <div class="col">
          <div class="card">
            <div class="card-header bg-white">
              <h5 class="mb-0">Recent Links</h5>
            </div>
            <div class="card-body p-0">
              <div v-if="loading" class="text-center py-5">
                <div class="spinner-border text-primary" role="status">
                  <span class="visually-hidden">Loading...</span>
                </div>
              </div>

              <div v-else-if="links.length === 0" class="text-center py-5">
                <i class="bi bi-link-45deg text-muted" style="font-size: 3rem"></i>
                <p class="text-muted mt-2">No links yet. Create your first link!</p>
                <button class="btn btn-primary" @click="showCreateModal = true">
                  <i class="bi bi-plus-circle"></i> Create Link
                </button>
              </div>

              <div v-else class="table-responsive">
                <table class="table table-hover mb-0">
                  <thead>
                    <tr>
                      <th>Short Code</th>
                      <th>Destination</th>
                      <th>Created</th>
                      <th>Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="link in links" :key="link.id">
                      <td>
                        <code>{{ link.short_code }}</code>
                      </td>
                      <td class="text-truncate" style="max-width: 300px">
                        {{ link.destination_url }}
                      </td>
                      <td>{{ formatDate(link.created_at) }}</td>
                      <td>
                        <RouterLink
                          :to="`/links/${link.id}`"
                          class="btn btn-sm btn-outline-primary"
                        >
                          View
                        </RouterLink>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Create Link Modal -->
    <CreateLinkModal
      v-if="showCreateModal"
      @close="showCreateModal = false"
      @created="handleLinkCreated"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import api from '@/services/api'
import CreateLinkModal from '@/components/CreateLinkModal.vue'

const stats = ref({})
const links = ref([])
const loading = ref(true)
const showCreateModal = ref(false)

const loadData = async () => {
  loading.value = true
  try {
    const [analyticsRes, linksRes] = await Promise.all([
      api.analytics.getUserAnalytics(),
      api.links.list({ limit: 5, sort: 'created_desc' })
    ])

    stats.value = analyticsRes.data
    links.value = linksRes.data.links || []
  } catch (error) {
    console.error('Failed to load dashboard data:', error)
  } finally {
    loading.value = false
  }
}

const handleLinkCreated = () => {
  showCreateModal.value = false
  loadData()
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.dashboard-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.stat-card {
  border: 1px solid var(--border-light);
  box-shadow: var(--shadow-xs);
  transition: all 0.2s ease;
  height: 100%;
}

.stat-card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--border-color);
}

.stat-card .card-body {
  padding: 20px;
}

.stat-card h2 {
  font-size: 32px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
  line-height: 1.2;
  letter-spacing: -0.02em;
}

.stat-card p {
  color: var(--text-tertiary);
  font-size: 13px;
  font-weight: 500;
  margin: 0;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: white;
}

.stat-icon.bg-primary {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
}

.stat-icon.bg-success {
  background: linear-gradient(135deg, #00D924 0%, #00A86B 100%);
}

.stat-icon.bg-info {
  background: linear-gradient(135deg, #0073E6 0%, #005BB5 100%);
}

.card-header {
  border-bottom: 1px solid var(--border-light);
  background-color: var(--bg-white);
  padding: 20px 24px;
}

.card-header h5 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 16px;
}
</style>
