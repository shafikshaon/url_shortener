<template>
  <div class="api-keys-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <div class="d-flex justify-content-between align-items-center">
            <div>
              <h1 class="h3 mb-0">API Keys</h1>
              <p class="text-muted mb-0">Manage your API keys for programmatic access</p>
            </div>
            <button class="btn btn-primary" @click="showGenerateModal = true">
              <i class="bi bi-plus-circle"></i> Generate API Key
            </button>
          </div>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Your API Keys</h6>
            </div>
            <div class="card-body">
              <div v-if="apiKeys.length === 0" class="text-center py-5">
                <i class="bi bi-key" style="font-size: 48px; color: #ccc;"></i>
                <p class="text-muted mt-3">No API keys yet</p>
                <button class="btn btn-primary" @click="showGenerateModal = true">
                  Generate Your First API Key
                </button>
              </div>

              <div v-else class="api-keys-list">
                <div v-for="key in apiKeys" :key="key.id" class="api-key-item">
                  <div class="d-flex justify-content-between align-items-start">
                    <div class="flex-grow-1">
                      <div class="d-flex align-items-center gap-2 mb-2">
                        <h6 class="mb-0">{{ key.name }}</h6>
                        <span v-if="key.active" class="badge bg-success">Active</span>
                        <span v-else class="badge bg-secondary">Inactive</span>
                      </div>
                      <div class="api-key-value">
                        <code>{{ key.key }}</code>
                        <button class="btn btn-sm btn-link" @click="copyApiKey(key.key)">
                          <i class="bi bi-clipboard"></i>
                        </button>
                      </div>
                      <small class="text-muted">Created {{ key.created }}</small>
                    </div>
                    <div class="btn-group">
                      <button class="btn btn-sm btn-outline-danger" @click="revokeKey(key.id)">
                        <i class="bi bi-trash"></i> Revoke
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">API Documentation</h6>
            </div>
            <div class="card-body">
              <p class="text-muted">Learn how to use the LinkShort API to create and manage links programmatically.</p>
              <a href="#" class="btn btn-outline-primary">
                <i class="bi bi-book"></i> View API Documentation
              </a>
            </div>
          </div>
        </div>

        <div class="col-lg-4">
          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">API Usage</h6>
            </div>
            <div class="card-body">
              <div class="mb-3">
                <small class="text-muted">API Calls This Month</small>
                <div class="h5 mb-0">0</div>
              </div>
              <div>
                <small class="text-muted">Rate Limit</small>
                <div>100 requests/minute</div>
              </div>
            </div>
          </div>

          <div class="card mt-3">
            <div class="card-header bg-white">
              <h6 class="mb-0">Security Tips</h6>
            </div>
            <div class="card-body">
              <ul class="small mb-0">
                <li class="mb-2">Never share your API keys publicly</li>
                <li class="mb-2">Rotate keys regularly</li>
                <li class="mb-2">Use environment variables</li>
                <li>Revoke unused keys immediately</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Generate API Key Modal -->
    <div
      v-if="showGenerateModal"
      class="modal fade show d-block"
      tabindex="-1"
      style="background-color: rgba(0, 0, 0, 0.5)"
      @click.self="showGenerateModal = false"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Generate New API Key</h5>
            <button type="button" class="btn-close" @click="showGenerateModal = false"></button>
          </div>
          <div class="modal-body">
            <div class="mb-3">
              <label class="form-label">Key Name</label>
              <input type="text" class="form-control" placeholder="e.g., Production API" v-model="newKeyName">
              <small class="text-muted">Give your API key a descriptive name</small>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" @click="showGenerateModal = false">Cancel</button>
            <button type="button" class="btn btn-primary" @click="generateKey">Generate Key</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const showGenerateModal = ref(false)
const newKeyName = ref('')
const apiKeys = ref([])

const generateKey = () => {
  const newKey = {
    id: Date.now(),
    name: newKeyName.value || 'Unnamed Key',
    key: 'sk_live_' + Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15),
    created: new Date().toLocaleDateString(),
    active: true
  }
  apiKeys.value.push(newKey)
  newKeyName.value = ''
  showGenerateModal.value = false
}

const copyApiKey = async (key) => {
  try {
    await navigator.clipboard.writeText(key)
    alert('API key copied to clipboard!')
  } catch (err) {
    console.error('Failed to copy:', err)
  }
}

const revokeKey = (id) => {
  if (confirm('Are you sure you want to revoke this API key? This action cannot be undone.')) {
    apiKeys.value = apiKeys.value.filter(k => k.id !== id)
  }
}
</script>

<style scoped>
.api-keys-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.api-keys-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.api-key-item {
  padding: 20px;
  border: 1px solid var(--border-light);
  border-radius: 8px;
  background-color: var(--bg-white);
}

.api-key-value {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px;
  background-color: var(--bg-secondary);
  border-radius: 6px;
  margin: 12px 0;
}

.api-key-value code {
  flex: 1;
  font-size: 13px;
  background: none;
  padding: 0;
  color: var(--text-primary);
}
</style>
