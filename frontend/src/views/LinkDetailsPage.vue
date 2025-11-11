<template>
  <div class="link-details-page">
    <div class="container-fluid py-4">
      <div v-if="loading" class="text-center py-5">
        <div class="spinner-border text-primary" role="status">
          <span class="visually-hidden">Loading...</span>
        </div>
      </div>

      <div v-else-if="link">
        <!-- Header -->
        <div class="row mb-4">
          <div class="col">
            <RouterLink to="/links" class="text-decoration-none mb-2 d-inline-block">
              <i class="bi bi-arrow-left"></i> Back to Links
            </RouterLink>
            <h1 class="h3 mb-0">Link Details</h1>
          </div>
        </div>

        <!-- Link Info Card -->
        <div class="row mb-4">
          <div class="col-lg-8">
            <div class="card">
              <div class="card-body">
                <div class="mb-3">
                  <label class="form-label text-muted">Short Code</label>
                  <div class="d-flex align-items-center gap-2">
                    <code class="fs-5 text-primary">{{ link.short_code }}</code>
                    <button
                      class="btn btn-sm btn-outline-primary"
                      @click="copyShortURL"
                    >
                      <i class="bi bi-clipboard"></i> Copy
                    </button>
                  </div>
                </div>

                <div class="mb-3">
                  <label class="form-label text-muted">Short URL</label>
                  <div>
                    <a :href="link.short_url" target="_blank" class="text-decoration-none">
                      {{ link.short_url }}
                      <i class="bi bi-box-arrow-up-right"></i>
                    </a>
                  </div>
                </div>

                <div class="mb-3">
                  <label class="form-label text-muted">Destination URL</label>
                  <div class="text-break">
                    <a :href="link.destination_url" target="_blank" class="text-decoration-none">
                      {{ link.destination_url }}
                      <i class="bi bi-box-arrow-up-right"></i>
                    </a>
                  </div>
                </div>

                <div class="mb-3" v-if="link.title">
                  <label class="form-label text-muted">Title</label>
                  <div>{{ link.title }}</div>
                </div>

                <div class="mb-3" v-if="link.tags && link.tags.length > 0">
                  <label class="form-label text-muted">Tags</label>
                  <div>
                    <span
                      v-for="tag in link.tags"
                      :key="tag"
                      class="badge bg-secondary me-1"
                    >
                      {{ tag }}
                    </span>
                  </div>
                </div>

                <div class="mb-3">
                  <label class="form-label text-muted">Created</label>
                  <div>{{ formatDate(link.created_at) }}</div>
                </div>

                <div class="mb-3" v-if="link.expires_at">
                  <label class="form-label text-muted">Expires</label>
                  <div>{{ formatDate(link.expires_at) }}</div>
                </div>
              </div>
            </div>
          </div>

          <!-- Stats Summary -->
          <div class="col-lg-4">
            <div class="card mb-3">
              <div class="card-body text-center">
                <h3 class="display-4 text-primary mb-0">{{ stats.total_clicks || 0 }}</h3>
                <p class="text-muted mb-0">Total Clicks</p>
              </div>
            </div>

            <div class="card">
              <div class="card-body text-center">
                <h3 class="display-4 text-success mb-0">{{ stats.last_30_days || 0 }}</h3>
                <p class="text-muted mb-0">Last 30 Days</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Analytics Charts -->
        <div class="row mb-4">
          <div class="col-12">
            <div class="card">
              <div class="card-header bg-white">
                <h5 class="mb-0">Click History (Last 30 Days)</h5>
              </div>
              <div class="card-body">
                <canvas ref="clicksChart"></canvas>
              </div>
            </div>
          </div>
        </div>

        <!-- Breakdown Stats -->
        <div class="row">
          <div class="col-md-4 mb-4">
            <div class="card h-100">
              <div class="card-header bg-white">
                <h5 class="mb-0">Top Countries</h5>
              </div>
              <div class="card-body">
                <div v-if="stats.countries && stats.countries.length > 0">
                  <div
                    v-for="country in stats.countries"
                    :key="country.country_code"
                    class="d-flex justify-content-between mb-2"
                  >
                    <span>{{ country.country_code || 'Unknown' }}</span>
                    <strong>{{ country.count }}</strong>
                  </div>
                </div>
                <div v-else class="text-muted">No data available</div>
              </div>
            </div>
          </div>

          <div class="col-md-4 mb-4">
            <div class="card h-100">
              <div class="card-header bg-white">
                <h5 class="mb-0">Top Referers</h5>
              </div>
              <div class="card-body">
                <div v-if="stats.referers && stats.referers.length > 0">
                  <div
                    v-for="referer in stats.referers"
                    :key="referer.referer"
                    class="d-flex justify-content-between mb-2"
                  >
                    <span class="text-truncate me-2">{{ referer.referer || 'Direct' }}</span>
                    <strong>{{ referer.count }}</strong>
                  </div>
                </div>
                <div v-else class="text-muted">No data available</div>
              </div>
            </div>
          </div>

          <div class="col-md-4 mb-4">
            <div class="card h-100">
              <div class="card-header bg-white">
                <h5 class="mb-0">Device Types</h5>
              </div>
              <div class="card-body">
                <div v-if="stats.device_types && stats.device_types.length > 0">
                  <div
                    v-for="device in stats.device_types"
                    :key="device.device_type"
                    class="d-flex justify-content-between mb-2"
                  >
                    <span>{{ device.device_type || 'Unknown' }}</span>
                    <strong>{{ device.count }}</strong>
                  </div>
                </div>
                <div v-else class="text-muted">No data available</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, RouterLink } from 'vue-router'
import { Chart, registerables } from 'chart.js'
import api from '@/services/api'

Chart.register(...registerables)

const route = useRoute()
const link = ref(null)
const stats = ref({})
const loading = ref(true)
const clicksChart = ref(null)

const loadLinkDetails = async () => {
  loading.value = true
  try {
    const linkId = route.params.id
    const [linkRes, statsRes] = await Promise.all([
      api.links.get(linkId),
      api.links.getStats(linkId)
    ])

    link.value = linkRes.data
    stats.value = statsRes.data

    // Render chart after data is loaded
    setTimeout(() => {
      renderChart()
    }, 100)
  } catch (error) {
    console.error('Failed to load link details:', error)
  } finally {
    loading.value = false
  }
}

const renderChart = () => {
  if (!clicksChart.value || !stats.value.daily_clicks) return

  const ctx = clicksChart.value.getContext('2d')
  new Chart(ctx, {
    type: 'line',
    data: {
      labels: stats.value.daily_clicks.map(d => d.date),
      datasets: [{
        label: 'Clicks',
        data: stats.value.daily_clicks.map(d => d.count),
        borderColor: '#7C3AED',
        backgroundColor: 'rgba(124, 58, 237, 0.1)',
        tension: 0.4,
        fill: true
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: true,
      plugins: {
        legend: {
          display: false
        }
      },
      scales: {
        y: {
          beginAtZero: true,
          ticks: {
            precision: 0
          }
        }
      }
    }
  })
}

const copyShortURL = async () => {
  try {
    await navigator.clipboard.writeText(link.value.short_url)
    alert('Short URL copied to clipboard!')
  } catch (err) {
    console.error('Failed to copy:', err)
  }
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

onMounted(() => {
  loadLinkDetails()
})
</script>

<style scoped>
.link-details-page {
  background-color: #F9FAFB;
  min-height: 100vh;
}

code {
  padding: 0.4em 0.6em;
  background-color: #F3F4F6;
  color: #7C3AED;
  border-radius: 0.375rem;
  font-weight: 500;
}

.text-primary {
  color: #7C3AED !important;
}

.text-success {
  color: #10B981 !important;
}
</style>
