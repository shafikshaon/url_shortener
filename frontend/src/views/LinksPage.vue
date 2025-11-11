<template>
  <div class="links-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <div class="d-flex justify-content-between align-items-center">
            <div>
              <h1 class="h3 mb-0">My Links</h1>
              <p class="text-muted mb-0">Manage and track all your shortened links</p>
            </div>
            <button class="btn btn-primary" @click="showCreateModal = true">
              <i class="bi bi-plus-circle"></i> Create Link
            </button>
          </div>
        </div>
      </div>

      <!-- Filters -->
      <div class="row mb-4">
        <div class="col">
          <div class="card">
            <div class="card-body">
              <div class="row g-3">
                <div class="col-md-6">
                  <input
                    type="text"
                    class="form-control"
                    placeholder="Search links..."
                    v-model="searchQuery"
                    @input="debouncedSearch"
                  />
                </div>
                <div class="col-md-3">
                  <select class="form-select" v-model="sortBy" @change="loadLinks">
                    <option value="created_desc">Newest First</option>
                    <option value="created_asc">Oldest First</option>
                    <option value="clicks">Most Clicks</option>
                  </select>
                </div>
                <div class="col-md-3">
                  <select class="form-select" v-model="perPage" @change="loadLinks">
                    <option :value="20">20 per page</option>
                    <option :value="50">50 per page</option>
                    <option :value="100">100 per page</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Links Grid -->
      <div class="row">
        <div class="col">
          <div class="card">
            <div class="card-body p-0">
              <div v-if="loading" class="text-center py-5">
                <div class="spinner-border text-primary" role="status">
                  <span class="visually-hidden">Loading...</span>
                </div>
              </div>

              <div v-else-if="links.length === 0" class="empty-state">
                <i class="bi bi-link-45deg"></i>
                <h3>{{ searchQuery ? 'No links found' : 'No links yet' }}</h3>
                <p>
                  {{ searchQuery ? 'Try adjusting your search' : 'Create your first shortened URL to get started!' }}
                </p>
                <button v-if="!searchQuery" class="btn btn-primary" @click="showCreateModal = true">
                  <i class="bi bi-plus-circle"></i> Create Your First Link
                </button>
              </div>

              <div v-else class="links-grid">
                <div v-for="link in links" :key="link.id" class="link-card">
                  <div class="link-card-header">
                    <div class="link-title-section">
                      <h6 class="link-title">{{ link.title || 'Untitled Link' }}</h6>
                      <div class="link-stats-badges">
                        <span class="stat-badge">
                          <i class="bi bi-cursor-fill"></i> {{ formatNumber(link.click_count || 0) }}
                        </span>
                        <span v-if="link.expires_at" class="stat-badge expires">
                          <i class="bi bi-clock"></i> Expires {{ formatDateRelative(link.expires_at) }}
                        </span>
                      </div>
                    </div>
                  </div>

                  <div class="link-card-body">
                    <div class="url-section">
                      <div class="short-url-row">
                        <i class="bi bi-link-45deg"></i>
                        <code class="short-code">{{ link.short_code }}</code>
                        <button
                          class="btn-copy-inline"
                          @click="copyShortURL(link.short_url || `${getBaseUrl()}/${link.short_code}`)"
                          title="Copy short URL"
                        >
                          <i class="bi bi-clipboard"></i>
                        </button>
                      </div>
                      <div class="destination-url-row">
                        <i class="bi bi-arrow-right"></i>
                        <a :href="link.destination_url" target="_blank" rel="noopener" class="destination-link">
                          {{ truncateUrl(link.destination_url, 60) }}
                          <i class="bi bi-box-arrow-up-right"></i>
                        </a>
                      </div>
                    </div>

                    <div v-if="link.tags && link.tags.length > 0" class="tags-section">
                      <i class="bi bi-tags"></i>
                      <div class="tags-list">
                        <span v-for="tag in link.tags" :key="tag" class="tag">{{ tag }}</span>
                      </div>
                    </div>

                    <div class="link-meta-row">
                      <span class="meta-text">
                        <i class="bi bi-calendar3"></i> {{ formatDate(link.created_at) }}
                      </span>
                    </div>
                  </div>

                  <div class="link-card-actions">
                    <RouterLink
                      :to="`/links/${link.id}`"
                      class="btn btn-sm btn-outline-primary"
                    >
                      <i class="bi bi-bar-chart"></i> View Analytics
                    </RouterLink>
                    <button
                      class="btn btn-sm btn-outline-danger"
                      @click="confirmDelete(link)"
                      title="Delete link"
                    >
                      <i class="bi bi-trash"></i>
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- Pagination -->
            <div v-if="links.length > 0" class="card-footer bg-white">
              <div class="d-flex justify-content-between align-items-center">
                <div class="text-muted">
                  Showing {{ links.length }} links
                </div>
                <div class="btn-group">
                  <button
                    class="btn btn-outline-primary"
                    :disabled="offset === 0"
                    @click="previousPage"
                  >
                    <i class="bi bi-chevron-left"></i> Previous
                  </button>
                  <button
                    class="btn btn-outline-primary"
                    :disabled="links.length < perPage"
                    @click="nextPage"
                  >
                    Next <i class="bi bi-chevron-right"></i>
                  </button>
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

const links = ref([])
const loading = ref(true)
const showCreateModal = ref(false)
const searchQuery = ref('')
const sortBy = ref('created_desc')
const perPage = ref(20)
const offset = ref(0)

let searchTimeout = null

const loadLinks = async () => {
  loading.value = true
  try {
    const response = await api.links.list({
      limit: perPage.value,
      offset: offset.value,
      search: searchQuery.value,
      sort: sortBy.value
    })
    links.value = response.data.links || []
  } catch (error) {
    console.error('Failed to load links:', error)
  } finally {
    loading.value = false
  }
}

const debouncedSearch = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    offset.value = 0
    loadLinks()
  }, 500)
}

const handleLinkCreated = () => {
  showCreateModal.value = false
  loadLinks()
}

const copyShortURL = async (url) => {
  try {
    await navigator.clipboard.writeText(url)
    alert('Short URL copied to clipboard!')
  } catch (err) {
    console.error('Failed to copy:', err)
  }
}

const confirmDelete = async (link) => {
  if (confirm(`Are you sure you want to delete the link "${link.short_code}"?`)) {
    try {
      await api.links.delete(link.id)
      loadLinks()
    } catch (error) {
      alert('Failed to delete link')
    }
  }
}

const nextPage = () => {
  offset.value += perPage.value
  loadLinks()
}

const previousPage = () => {
  offset.value = Math.max(0, offset.value - perPage.value)
  loadLinks()
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
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

const getBaseUrl = () => {
  return window.location.origin
}

onMounted(() => {
  loadLinks()
})
</script>

<style scoped>
.links-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.links-page h1 {
  font-size: 1.75rem;
  font-weight: 600;
  color: var(--text-primary);
}

.links-page p {
  color: var(--text-tertiary);
  font-size: 14px;
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
  margin-bottom: 20px;
}

/* Links Grid */
.links-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 20px;
  padding: 20px;
}

.link-card {
  background-color: var(--bg-white);
  border: 1px solid var(--border-light);
  border-radius: 8px;
  overflow: hidden;
  transition: all 0.2s ease;
  display: flex;
  flex-direction: column;
  box-shadow: var(--shadow-xs);
}

.link-card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--border-color);
  transform: translateY(-2px);
}

.link-card-header {
  padding: 16px;
  border-bottom: 1px solid var(--border-light);
  background-color: var(--bg-secondary);
}

.link-title-section {
  display: flex;
  justify-content: space-between;
  align-items: start;
  gap: 12px;
}

.link-title {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.link-stats-badges {
  display: flex;
  gap: 8px;
  flex-shrink: 0;
}

.stat-badge {
  background-color: var(--bg-white);
  color: var(--text-secondary);
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
  display: inline-flex;
  align-items: center;
  gap: 4px;
  white-space: nowrap;
  border: 1px solid var(--border-light);
}

.stat-badge i {
  font-size: 11px;
  color: var(--success-color);
}

.stat-badge.expires i {
  color: var(--warning-color);
}

.link-card-body {
  padding: 16px;
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.url-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.short-url-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.short-url-row > i {
  color: var(--primary-color);
  font-size: 18px;
  flex-shrink: 0;
}

.short-code {
  font-size: 14px;
  font-weight: 600;
  color: var(--primary-color);
  background-color: var(--primary-subtle);
  padding: 4px 8px;
  border-radius: 4px;
  border: 1px solid rgba(99, 91, 255, 0.2);
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
}

.btn-copy-inline {
  background: none;
  border: none;
  color: var(--text-tertiary);
  cursor: pointer;
  padding: 4px 6px;
  border-radius: 4px;
  transition: all 0.15s ease;
  font-size: 14px;
  flex-shrink: 0;
}

.btn-copy-inline:hover {
  color: var(--primary-color);
  background-color: var(--primary-subtle);
}

.destination-url-row {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
}

.destination-url-row > i {
  color: var(--text-tertiary);
  font-size: 12px;
  flex-shrink: 0;
}

.destination-link {
  color: var(--text-tertiary);
  text-decoration: none;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
  display: flex;
  align-items: center;
  gap: 4px;
}

.destination-link:hover {
  color: var(--primary-color);
  text-decoration: underline;
}

.destination-link i {
  font-size: 11px;
  flex-shrink: 0;
}

.tags-section {
  display: flex;
  align-items: center;
  gap: 8px;
  padding-top: 8px;
  border-top: 1px solid var(--border-light);
}

.tags-section > i {
  color: var(--text-tertiary);
  font-size: 14px;
  flex-shrink: 0;
}

.tags-list {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.tag {
  background-color: var(--bg-tertiary);
  color: var(--text-secondary);
  padding: 3px 8px;
  border-radius: 4px;
  font-size: 11px;
  font-weight: 500;
}

.link-meta-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.meta-text {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: var(--text-muted);
}

.meta-text i {
  font-size: 12px;
}

.link-card-actions {
  padding: 12px 16px;
  border-top: 1px solid var(--border-light);
  background-color: var(--bg-secondary);
  display: flex;
  gap: 8px;
}

.link-card-actions .btn {
  flex: 1;
}

.card-footer {
  background-color: var(--bg-white);
  border-top: 1px solid var(--border-light);
  padding: 16px 24px;
}

.card-footer .text-muted {
  color: var(--text-tertiary);
  font-size: 14px;
}

/* Responsive */
@media (max-width: 768px) {
  .links-grid {
    grid-template-columns: 1fr;
  }

  .link-title-section {
    flex-direction: column;
    align-items: flex-start;
  }

  .link-stats-badges {
    width: 100%;
  }
}
</style>
