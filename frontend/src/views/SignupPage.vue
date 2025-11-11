<template>
  <div class="auth-page">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-md-5">
          <div class="card shadow-lg mt-5">
            <div class="card-body p-5">
              <div class="text-center mb-4">
                <i class="bi bi-link-45deg text-primary" style="font-size: 3rem"></i>
                <h2 class="mt-2">Create Account</h2>
                <p class="text-muted">Start shortening URLs for free</p>
              </div>

              <div v-if="error" class="alert alert-danger" role="alert">
                {{ error }}
              </div>

              <form @submit.prevent="handleSignup">
                <div class="mb-3">
                  <label for="email" class="form-label">Email address</label>
                  <input
                    type="email"
                    class="form-control"
                    id="email"
                    v-model="email"
                    required
                    placeholder="you@example.com"
                  />
                </div>

                <div class="mb-3">
                  <label for="password" class="form-label">Password</label>
                  <input
                    type="password"
                    class="form-control"
                    id="password"
                    v-model="password"
                    required
                    minlength="8"
                    placeholder="At least 8 characters"
                  />
                  <div class="form-text">Password must be at least 8 characters</div>
                </div>

                <div class="mb-3">
                  <label for="confirmPassword" class="form-label">Confirm Password</label>
                  <input
                    type="password"
                    class="form-control"
                    id="confirmPassword"
                    v-model="confirmPassword"
                    required
                    placeholder="Confirm your password"
                  />
                </div>

                <button
                  type="submit"
                  class="btn btn-primary w-100"
                  :disabled="loading"
                >
                  <span v-if="loading" class="spinner-border spinner-border-sm me-2"></span>
                  {{ loading ? 'Creating account...' : 'Sign Up' }}
                </button>
              </form>

              <hr class="my-4" />

              <div class="text-center">
                <p class="mb-0">
                  Already have an account?
                  <RouterLink to="/login" class="text-primary fw-semibold">
                    Sign in
                  </RouterLink>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const error = ref('')
const loading = ref(false)

const handleSignup = async () => {
  error.value = ''

  if (password.value !== confirmPassword.value) {
    error.value = 'Passwords do not match'
    return
  }

  if (password.value.length < 8) {
    error.value = 'Password must be at least 8 characters'
    return
  }

  loading.value = true

  try {
    const result = await authStore.signup(email.value, password.value)

    if (result.success) {
      router.push('/dashboard')
    } else {
      error.value = result.error
    }
  } catch (err) {
    error.value = 'An unexpected error occurred'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.auth-page {
  min-height: 100vh;
  background-color: #7C3AED;
  padding: 2rem 0;
  position: relative;
  overflow: hidden;
}

.auth-page::before {
  content: '';
  position: absolute;
  top: -10%;
  right: -5%;
  width: 400px;
  height: 400px;
  background-color: rgba(167, 139, 250, 0.15);
  border-radius: 50%;
}

.auth-page::after {
  content: '';
  position: absolute;
  bottom: -15%;
  left: -10%;
  width: 500px;
  height: 500px;
  background-color: rgba(109, 40, 217, 0.15);
  border-radius: 50%;
}

.container {
  position: relative;
  z-index: 1;
}

.card {
  border: 1px solid #E5E7EB;
  border-radius: 1rem;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}

.form-control {
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
}

.btn-primary {
  padding: 0.75rem;
  border-radius: 0.5rem;
  font-weight: 600;
}
</style>
