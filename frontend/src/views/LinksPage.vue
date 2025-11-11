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

      <!-- Links Table -->
      <div class="row">
        <div class="col">
          <div class="card">
            <div class="card-body p-0">
              <div v-if="loading" class="text-center py-5">
                <div class="spinner-border text-primary" role="status">
                  <span class="visually-hidden">Loading...</span>
                </div>
              </div>

              <div v-else-if="links.length === 0" class="text-center py-5">
                <i class="bi bi-link-45deg text-muted" style="font-size: 3rem"></i>
                <p class="text-muted mt-2">
                  {{ searchQuery ? 'No links found' : 'No links yet. Create your first link!' }}
                </p>
                <button v-if="!searchQuery" class="btn btn-primary" @click="showCreateModal = true">
                  <i class="bi bi-plus-circle"></i> Create Link
                </button>
              </div>

              <div v-else class="table-responsive">
                <table class="table table-hover align-middle mb-0">
                  <thead class="table-light">
                    <tr>
                      <th>Short Code</th>
                      <th>Destination</th>
                      <th>Title</th>
                      <th>Tags</th>
                      <th>Created</th>
                      <th>Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="link in links" :key="link.id">
                      <td>
                        <div class="d-flex align-items-center gap-2">
                          <code class="text-primary">{{ link.short_code }}</code>
                          <button
                            class="btn btn-sm btn-link p-0"
                            @click="copyShortURL(link.short_url)"
                            title="Copy"
                          >
                            <i class="bi bi-clipboard"></i>
                          </button>
                        </div>
                      </td>
                      <td class="text-truncate" style="max-width: 300px">
                        <a :href="link.destination_url" target="_blank" class="text-decoration-none">
                          {{ link.destination_url }}
                          <i class="bi bi-box-arrow-up-right small"></i>
                        </a>
                      </td>
                      <td>{{ link.title || '-' }}</td>
                      <td>
                        <span
                          v-for="tag in link.tags"
                          :key="tag"
                          class="badge bg-secondary me-1"
                        >
                          {{ tag }}
                        </span>
                        <span v-if="!link.tags || link.tags.length === 0">-</span>
                      </td>
                      <td>{{ formatDate(link.created_at) }}</td>
                      <td>
                        <div class="btn-group btn-group-sm">
                          <RouterLink
                            :to="`/links/${link.id}`"
                            class="btn btn-outline-primary"
                            title="View Details"
                          >
                            <i class="bi bi-eye"></i>
                          </RouterLink>
                          <button
                            class="btn btn-outline-danger"
                            @click="confirmDelete(link)"
                            title="Delete"
                          >
                            <i class="bi bi-trash"></i>
                          </button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
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

.table-hover tbody tr:hover {
  cursor: default;
  background-color: var(--bg-secondary);
}

.btn-link {
  color: var(--text-tertiary);
  text-decoration: none;
  transition: all 0.15s ease;
}

.btn-link:hover {
  color: var(--primary-color);
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
</style>
