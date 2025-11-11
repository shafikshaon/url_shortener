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
  background-color: rgba(0, 0, 0, 0.5);
}

.modal-content {
  border-radius: 0.5rem;
}
</style>
