import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1'

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Request interceptor to add auth token
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('auth_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor to handle errors
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Unauthorized - clear token and redirect to login
      localStorage.removeItem('auth_token')
      localStorage.removeItem('user')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export default {
  // Auth endpoints
  auth: {
    signup(email, password, fullName = '') {
      return apiClient.post('/auth/signup', { email, password, full_name: fullName })
    },
    login(email, password) {
      return apiClient.post('/auth/login', { email, password })
    },
    getProfile() {
      return apiClient.get('/profile')
    },
    updateProfile(fullName) {
      return apiClient.put('/profile', { full_name: fullName })
    },
    changePassword(currentPassword, newPassword) {
      return apiClient.post('/auth/change-password', {
        current_password: currentPassword,
        new_password: newPassword
      })
    },
    generateAPIKey() {
      return apiClient.post('/api-key')
    }
  },

  // Link endpoints
  links: {
    create(linkData) {
      return apiClient.post('/links', linkData)
    },
    list(params = {}) {
      return apiClient.get('/links', { params })
    },
    get(id) {
      return apiClient.get(`/links/${id}`)
    },
    update(id, linkData) {
      return apiClient.patch(`/links/${id}`, linkData)
    },
    delete(id) {
      return apiClient.delete(`/links/${id}`)
    },
    getStats(id) {
      return apiClient.get(`/links/${id}/stats`)
    }
  },

  // Analytics endpoints
  analytics: {
    getUserAnalytics() {
      return apiClient.get('/analytics')
    }
  },

  // Tags endpoints
  tags: {
    list() {
      return apiClient.get('/tags')
    }
  }
}
