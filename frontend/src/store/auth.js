import { defineStore } from 'pinia'
import api from '@/services/api'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')) || null,
    token: localStorage.getItem('auth_token') || null,
    isAuthenticated: !!localStorage.getItem('auth_token')
  }),

  getters: {
    currentUser: (state) => state.user,
    isLoggedIn: (state) => state.isAuthenticated,
    subscriptionTier: (state) => state.user?.subscription_tier || 'free'
  },

  actions: {
    async login(email, password) {
      try {
        const response = await api.auth.login(email, password)
        const { token, user } = response.data

        this.token = token
        this.user = user
        this.isAuthenticated = true

        localStorage.setItem('auth_token', token)
        localStorage.setItem('user', JSON.stringify(user))

        return { success: true }
      } catch (error) {
        return {
          success: false,
          error: error.response?.data?.error || 'Login failed'
        }
      }
    },

    async signup(email, password) {
      try {
        const response = await api.auth.signup(email, password)
        const { token, user } = response.data

        this.token = token
        this.user = user
        this.isAuthenticated = true

        localStorage.setItem('auth_token', token)
        localStorage.setItem('user', JSON.stringify(user))

        return { success: true }
      } catch (error) {
        return {
          success: false,
          error: error.response?.data?.error || 'Signup failed'
        }
      }
    },

    async logout() {
      this.token = null
      this.user = null
      this.isAuthenticated = false

      localStorage.removeItem('auth_token')
      localStorage.removeItem('user')
    },

    async refreshProfile() {
      try {
        const response = await api.auth.getProfile()
        this.user = response.data
        localStorage.setItem('user', JSON.stringify(response.data))
      } catch (error) {
        console.error('Failed to refresh profile:', error)
      }
    },

    updateUser(userData) {
      this.user = userData
      localStorage.setItem('user', JSON.stringify(userData))
    }
  }
})
