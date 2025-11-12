<template>
  <div class="modal show d-block" tabindex="-1" @click.self="$emit('close')">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Create New Link</h5>
          <button type="button" class="btn-close" @click="$emit('close')"></button>
        </div>

        <form @submit.prevent="handleSubmit">
          <div class="modal-body">
            <div v-if="error" class="alert alert-danger">
              {{ error }}
            </div>

            <div v-if="success" class="alert alert-success">
              <div class="d-flex justify-content-between align-items-center">
                <span>Link created successfully!</span>
                <button
                  type="button"
                  class="btn btn-sm btn-outline-success"
                  @click="copyShortURL"
                >
                  <i class="bi bi-clipboard"></i> Copy
                </button>
              </div>
              <div class="mt-2">
                <small><strong>Short URL:</strong> {{ createdLink?.short_url }}</small>
              </div>
            </div>

            <div class="mb-3">
              <label for="destinationUrl" class="form-label">
                Destination URL <span class="text-danger">*</span>
              </label>
              <input
                type="url"
                class="form-control"
                id="destinationUrl"
                v-model="form.destination_url"
                placeholder="https://example.com/your-long-url"
                required
              />
            </div>

            <div class="mb-3">
              <label for="shortCode" class="form-label">
                Custom Short Code (optional)
              </label>
              <input
                type="text"
                class="form-control"
                id="shortCode"
                v-model="form.short_code"
                placeholder="my-custom-code"
                pattern="[a-zA-Z0-9-]+"
                minlength="3"
                maxlength="20"
              />
              <div class="form-text">
                3-20 characters. Letters, numbers, and hyphens only. Leave blank for auto-generated code.
              </div>
            </div>

            <div class="mb-3">
              <label for="title" class="form-label">Title (optional)</label>
              <input
                type="text"
                class="form-control"
                id="title"
                v-model="form.title"
                placeholder="My Campaign Link"
              />
            </div>

            <div class="mb-3">
              <label for="tags" class="form-label">Tags (optional)</label>
              <input
                type="text"
                class="form-control"
                id="tags"
                v-model="tagsInput"
                placeholder="marketing, social, campaign"
              />
              <div class="form-text">Comma-separated tags for organization</div>
            </div>

            <div class="mb-3">
              <label for="expiresAt" class="form-label">Expires At (optional)</label>
              <input
                type="datetime-local"
                class="form-control"
                id="expiresAt"
                v-model="expiresAtInput"
              />
            </div>
          </div>

          <div class="modal-footer">
            <button
              type="button"
              class="btn btn-secondary"
              @click="$emit('close')"
              :disabled="loading"
            >
              {{ success ? 'Close' : 'Cancel' }}
            </button>
            <button
              v-if="!success"
              type="submit"
              class="btn btn-primary"
              :disabled="loading"
            >
              <span v-if="loading" class="spinner-border spinner-border-sm me-2"></span>
              {{ loading ? 'Creating...' : 'Create Link' }}
            </button>
            <button
              v-else
              type="button"
              class="btn btn-success"
              @click="createAnother"
            >
              Create Another
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
  <div class="modal-backdrop show"></div>
</template>

<script setup>
import { ref } from 'vue'
import api from '@/services/api'

const emit = defineEmits(['close', 'created'])

const form = ref({
  destination_url: '',
  short_code: '',
  title: '',
  tags: [],
  expires_at: null
})

const tagsInput = ref('')
const expiresAtInput = ref('')
const error = ref('')
const success = ref(false)
const loading = ref(false)
const createdLink = ref(null)

const handleSubmit = async () => {
  error.value = ''
  loading.value = true

  try {
    // Parse tags
    if (tagsInput.value) {
      form.value.tags = tagsInput.value.split(',').map(tag => tag.trim()).filter(tag => tag)
    }

    // Parse expiration date
    if (expiresAtInput.value) {
      form.value.expires_at = new Date(expiresAtInput.value).toISOString()
    }

    const response = await api.links.create(form.value)
    createdLink.value = response.data
    success.value = true
    emit('created', response.data)
  } catch (err) {
    error.value = err.response?.data?.error || 'Failed to create link'
  } finally {
    loading.value = false
  }
}

const createAnother = () => {
  form.value = {
    destination_url: '',
    short_code: '',
    title: '',
    tags: [],
    expires_at: null
  }
  tagsInput.value = ''
  expiresAtInput.value = ''
  success.value = false
  createdLink.value = null
  error.value = ''
}

const copyShortURL = async () => {
  if (createdLink.value?.short_url) {
    try {
      await navigator.clipboard.writeText(createdLink.value.short_url)
      alert('Short URL copied to clipboard!')
    } catch (err) {
      console.error('Failed to copy:', err)
    }
  }
}
</script>

<style scoped>
.modal {
  background-color: rgba(10, 37, 64, 0.5);
  backdrop-filter: blur(4px);
  animation: modalFadeIn var(--transition-base);
}

@keyframes modalFadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.modal-dialog {
  animation: modalSlideIn 0.2s ease-out;
}

@keyframes modalSlideIn {
  from {
    opacity: 0;
    transform: translateY(-20px) scale(0.98);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.modal-content {
  border-radius: var(--radius-xl);
  border: 1px solid var(--border-light);
  box-shadow: var(--shadow-2xl);
  background-color: var(--bg-white);
  overflow: hidden;
}

.modal-header {
  border-bottom: 1px solid var(--border-light);
  padding: 24px 32px;
  background-color: var(--bg-white);
}

.modal-title {
  color: var(--text-primary);
  font-weight: 600;
  font-size: 18px;
  line-height: 1.4;
  margin: 0;
}

.btn-close {
  opacity: 0.5;
  transition: all var(--transition-base);
  border-radius: var(--radius-md);
}

.btn-close:hover {
  opacity: 1;
  transform: scale(1.1);
}

.modal-body {
  padding: 32px;
  max-height: 70vh;
  overflow-y: auto;
}

.modal-footer {
  border-top: 1px solid var(--border-light);
  padding: 20px 32px;
  background-color: var(--bg-secondary);
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.btn-outline-success {
  border: 1px solid var(--success-color);
  color: var(--success-color);
  background-color: var(--bg-white);
}

.btn-outline-success:hover:not(:disabled) {
  background-color: var(--success-color);
  color: white;
  transform: translateY(-1px);
  box-shadow: var(--shadow-sm);
}

.modal-backdrop {
  background-color: rgba(10, 37, 64, 0.5);
  backdrop-filter: blur(4px);
}

/* Form styling within modal */
.modal-body .form-label {
  margin-bottom: 8px;
}

.modal-body .mb-3 {
  margin-bottom: 24px !important;
}

.modal-body .mb-3:last-child {
  margin-bottom: 0 !important;
}

.modal-body .alert {
  margin-bottom: 24px;
}
</style>
