<template>
  <div class="dashboard-page">
    <div class="container-fluid py-4">
      <!-- Header -->
      <div class="row mb-4">
        <div class="col">
          <h1 class="page-title">Create & Manage Links</h1>
          <p class="page-subtitle">Shorten, customize, and track your URLs</p>
        </div>
      </div>

      <!-- URL Shortener Form -->
      <div class="row mb-4">
        <div class="col">
          <div class="card shortener-card">
            <div class="card-body p-4">
              <form @submit.prevent="handleQuickCreate">
                <div class="row g-3 align-items-end">
                  <div class="col-md-8">
                    <label for="quickUrl" class="form-label">
                      <i class="bi bi-link-45deg"></i> Paste long URL
                    </label>
                    <input
                      type="url"
                      class="form-control form-control-lg"
                      id="quickUrl"
                      v-model="quickUrl"
                      placeholder="https://example.com/your-very-long-url"
                      required
                    />
                  </div>
                  <div class="col-md-4">
                    <button type="submit" class="btn btn-primary btn-lg w-100" :disabled="creating">
                      <span v-if="creating" class="spinner-border spinner-border-sm me-2"></span>
                      <i v-else class="bi bi-scissors"></i>
                      {{ creating ? 'Shortening...' : 'Shorten URL' }}
                    </button>
                  </div>
                </div>

                <div v-if="quickError" class="alert alert-danger mt-3 mb-0">
                  <i class="bi bi-exclamation-circle"></i> {{ quickError }}
                </div>

                <div v-if="newLink" class="result-box mt-3">
                  <div class="d-flex align-items-center justify-content-between">
                    <div class="flex-grow-1">
                      <label class="result-label">Your shortened URL:</label>
                      <div class="short-url-display">
                        <code>{{ newLink.short_url }}</code>
                      </div>
                    </div>
                    <div class="d-flex gap-2">
                      <button
                        type="button"
                        class="btn btn-outline-primary"
                        @click="copyToClipboard(newLink.short_url)"
                        title="Copy to clipboard"
                      >
                        <i class="bi bi-clipboard"></i> Copy
                      </button>
                      <RouterLink
                        :to="`/links/${newLink.id}`"
                        class="btn btn-outline-primary"
                        title="View details"
                      >
                        <i class="bi bi-bar-chart"></i> Stats
                      </RouterLink>
                    </div>
                  </div>
                </div>
              </form>

              <div class="mt-3 text-center">
                <button class="btn btn-link btn-sm" @click="showCreateModal = true">
                  <i class="bi bi-sliders"></i> Advanced options (custom short code, title, tags, expiry)
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="row g-3 mb-4">
        <div class="col-md-4">
          <div class="card stat-card">
            <div class="card-body">
              <div class="d-flex justify-content-between align-items-center">
                <div>
                  <p class="stat-label">Total Links</p>
                  <h2 class="stat-value">{{ formatNumber(stats.total_links || 0) }}</h2>
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
                  <p class="stat-label">Total Clicks</p>
                  <h2 class="stat-value">{{ formatNumber(stats.total_clicks || 0) }}</h2>
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
                  <p class="stat-label">This Month</p>
                  <h2 class="stat-value">{{ formatNumber(stats.clicks_this_month || 0) }}</h2>
                </div>
                <div class="stat-icon bg-info">
                  <i class="bi bi-graph-up-arrow"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Links -->
      <div class="row">
        <div class="col">
          <div class="card">
            <div class="card-header d-flex justify-content-between align-items-center">
              <h5 class="mb-0"><i class="bi bi-clock-history"></i> Recent Links</h5>
              <RouterLink to="/links" class="btn btn-sm btn-outline-primary">
                View All <i class="bi bi-arrow-right"></i>
              </RouterLink>
            </div>
            <div class="card-body p-0">
              <div v-if="loading" class="text-center py-5">
                <div class="spinner-border text-primary" role="status">
                  <span class="visually-hidden">Loading...</span>
                </div>
              </div>

              <div v-else-if="links.length === 0" class="empty-state">
                <i class="bi bi-link-45deg"></i>
                <h3>No links yet</h3>
                <p>Create your first shortened URL above to get started!</p>
              </div>

              <div v-else class="links-list">
                <div v-for="link in links" :key="link.id" class="link-item">
                  <div class="link-content">
                    <div class="link-main">
                      <div class="link-header">
                        <h6 class="link-title">{{ link.title || 'Untitled Link' }}</h6>
                        <div class="link-stats-inline">
                          <span class="stat-badge">
                            <i class="bi bi-cursor-fill"></i> {{ formatNumber(link.click_count || 0) }} clicks
                          </span>
                        </div>
                      </div>
                      <div class="link-urls">
                        <div class="short-url">
                          <i class="bi bi-link-45deg"></i>
                          <code>{{ link.short_code }}</code>
                          <button
                            class="btn-copy"
                            @click="copyToClipboard(link.short_url || `${getBaseUrl()}/${link.short_code}`)"
                            title="Copy short URL"
                          >
                            <i class="bi bi-clipboard"></i>
                          </button>
                        </div>
                        <div class="destination-url">
                          <i class="bi bi-arrow-right"></i>
                          <a :href="link.destination_url" target="_blank" rel="noopener">
                            {{ truncateUrl(link.destination_url) }}
                          </a>
                        </div>
                      </div>
                      <div class="link-meta">
                        <span class="meta-item">
                          <i class="bi bi-calendar3"></i> Created {{ formatDateRelative(link.created_at) }}
                        </span>
                        <span v-if="link.tags && link.tags.length > 0" class="meta-item">
                          <i class="bi bi-tags"></i>
                          <span v-for="tag in link.tags.slice(0, 2)" :key="tag" class="tag-badge">{{ tag }}</span>
                        </span>
                      </div>
                    </div>
                  </div>
                  <div class="link-actions">
                    <RouterLink
                      :to="`/links/${link.id}`"
                      class="btn btn-sm btn-outline-primary"
                      title="View details & analytics"
                    >
                      <i class="bi bi-bar-chart"></i> Analytics
                    </RouterLink>
                  </div>
                </div>
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

// Quick create form
const quickUrl = ref('')
const creating = ref(false)
const quickError = ref('')
const newLink = ref(null)

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

const handleQuickCreate = async () => {
  if (!quickUrl.value) return

  creating.value = true
  quickError.value = ''
  newLink.value = null

  try {
    const response = await api.links.create({
      destination_url: quickUrl.value
    })
    newLink.value = response.data
    quickUrl.value = ''
    loadData() // Refresh the list
  } catch (error) {
    quickError.value = error.response?.data?.error || 'Failed to create short link'
  } finally {
    creating.value = false
  }
}

const handleLinkCreated = () => {
  showCreateModal.value = false
  newLink.value = null
  loadData()
}

const copyToClipboard = async (text) => {
  try {
    await navigator.clipboard.writeText(text)
    // You could add a toast notification here
    alert('Copied to clipboard!')
  } catch (err) {
    console.error('Failed to copy:', err)
  }
}

const getBaseUrl = () => {
  return window.location.origin
}

const formatNumber = (num) => {
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  }
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'K'
  }
  return num.toString()
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

const formatDateRelative = (dateString) => {
  const date = new Date(dateString)
  const now = new Date()
  const diffMs = now - date
  const diffMins = Math.floor(diffMs / 60000)
  const diffHours = Math.floor(diffMs / 3600000)
  const diffDays = Math.floor(diffMs / 86400000)

  if (diffMins < 1) return 'just now'
  if (diffMins < 60) return `${diffMins}m ago`
  if (diffHours < 24) return `${diffHours}h ago`
  if (diffDays < 7) return `${diffDays}d ago`

  return formatDate(dateString)
}

const truncateUrl = (url, maxLength = 50) => {
  if (url.length <= maxLength) return url
  return url.substring(0, maxLength) + '...'
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

/* Header */
.page-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 0.25rem;
  letter-spacing: -0.02em;
}

.page-subtitle {
  color: var(--text-tertiary);
  font-size: 1rem;
  margin-bottom: 0;
}

/* Shortener Card */
.shortener-card {
  border: 2px solid var(--primary-color);
  box-shadow: var(--shadow-lg);
  background: linear-gradient(to bottom, var(--bg-white), var(--primary-subtle));
}

.shortener-card .form-label {
  font-weight: 600;
  color: var(--text-primary);
  font-size: 14px;
  margin-bottom: 8px;
}

.shortener-card .form-label i {
  color: var(--primary-color);
  font-size: 18px;
  margin-right: 4px;
}

.result-box {
  background-color: var(--bg-white);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  padding: 20px;
  box-shadow: var(--shadow-sm);
}

.result-label {
  font-size: 13px;
  font-weight: 500;
  color: var(--text-tertiary);
  margin-bottom: 8px;
  display: block;
}

.short-url-display {
  background-color: var(--bg-secondary);
  padding: 12px 16px;
  border-radius: 6px;
  border: 1px solid var(--border-light);
}

.short-url-display code {
  font-size: 16px;
  color: var(--primary-color);
  font-weight: 600;
  background: none;
  padding: 0;
  border: none;
}

/* Stats Cards */
.stat-card {
  border: 1px solid var(--border-light);
  box-shadow: var(--shadow-sm);
  transition: all 0.2s ease;
  height: 100%;
}

.stat-card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--border-color);
  transform: translateY(-2px);
}

.stat-card .card-body {
  padding: 24px;
}

.stat-label {
  color: var(--text-tertiary);
  font-size: 13px;
  font-weight: 600;
  margin: 0 0 8px 0;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.stat-value {
  font-size: 36px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0;
  line-height: 1;
  letter-spacing: -0.02em;
}

.stat-icon {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  color: white;
  box-shadow: var(--shadow-sm);
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

/* Card Header */
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
  display: flex;
  align-items: center;
  gap: 8px;
}

.card-header h5 i {
  color: var(--primary-color);
}

/* Empty State */
.empty-state {
  text-align: center;
  padding: 60px 20px;
}

.empty-state i {
  font-size: 64px;
  color: var(--text-muted);
  margin-bottom: 16px;
  display: block;
}

.empty-state h3 {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-secondary);
  margin-bottom: 8px;
}

.empty-state p {
  color: var(--text-tertiary);
  margin-bottom: 0;
}

/* Links List */
.links-list {
  padding: 0;
}

.link-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  border-bottom: 1px solid var(--border-light);
  transition: background-color 0.15s ease;
}

.link-item:last-child {
  border-bottom: none;
}

.link-item:hover {
  background-color: var(--bg-secondary);
}

.link-content {
  flex: 1;
  min-width: 0;
}

.link-main {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.link-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
}

.link-title {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
}

.link-stats-inline {
  display: flex;
  gap: 12px;
}

.stat-badge {
  background-color: var(--bg-secondary);
  color: var(--text-secondary);
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 13px;
  font-weight: 500;
  display: inline-flex;
  align-items: center;
  gap: 4px;
}

.stat-badge i {
  font-size: 12px;
  color: var(--success-color);
}

.link-urls {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.short-url {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
}

.short-url i {
  color: var(--primary-color);
  font-size: 18px;
}

.short-url code {
  font-size: 14px;
  font-weight: 600;
  color: var(--primary-color);
  background-color: var(--primary-subtle);
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid rgba(99, 91, 255, 0.2);
}

.btn-copy {
  background: none;
  border: none;
  color: var(--text-tertiary);
  cursor: pointer;
  padding: 4px 6px;
  border-radius: 4px;
  transition: all 0.15s ease;
  font-size: 14px;
}

.btn-copy:hover {
  color: var(--primary-color);
  background-color: var(--primary-subtle);
}

.destination-url {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: var(--text-tertiary);
}

.destination-url i {
  font-size: 12px;
}

.destination-url a {
  color: var(--text-tertiary);
  text-decoration: none;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.destination-url a:hover {
  color: var(--primary-color);
  text-decoration: underline;
}

.link-meta {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: var(--text-muted);
}

.meta-item i {
  font-size: 12px;
}

.tag-badge {
  background-color: var(--bg-tertiary);
  color: var(--text-secondary);
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 11px;
  margin-left: 4px;
  font-weight: 500;
}

.link-actions {
  margin-left: 20px;
  flex-shrink: 0;
}

/* Responsive */
@media (max-width: 768px) {
  .link-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .link-actions {
    margin-left: 0;
    width: 100%;
  }

  .link-actions .btn {
    width: 100%;
  }

  .page-title {
    font-size: 1.5rem;
  }
}
</style>
