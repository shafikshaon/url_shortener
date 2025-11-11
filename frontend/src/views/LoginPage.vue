<template>
  <div class="auth-page">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-md-5">
          <div class="card shadow-lg mt-5">
            <div class="card-body p-5">
              <div class="text-center mb-4">
                <i class="bi bi-link-45deg text-primary" style="font-size: 3rem"></i>
                <h2 class="mt-2">Welcome Back</h2>
                <p class="text-muted">Sign in to your account</p>
              </div>

              <div v-if="error" class="alert alert-danger" role="alert">
                {{ error }}
              </div>

              <form @submit.prevent="handleLogin">
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
                    placeholder="Enter your password"
                  />
                </div>

                <button
                  type="submit"
                  class="btn btn-primary w-100"
                  :disabled="loading"
                >
                  <span v-if="loading" class="spinner-border spinner-border-sm me-2"></span>
                  {{ loading ? 'Signing in...' : 'Sign In' }}
                </button>
              </form>

              <hr class="my-4" />

              <div class="text-center">
                <p class="mb-0">
                  Don't have an account?
                  <RouterLink to="/signup" class="text-primary fw-semibold">
                    Sign up
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
const error = ref('')
const loading = ref(false)

const handleLogin = async () => {
  error.value = ''
  loading.value = true

  try {
    const result = await authStore.login(email.value, password.value)

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
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-dark) 100%);
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
  background-color: rgba(139, 133, 255, 0.15);
  border-radius: 50%;
}

.auth-page::after {
  content: '';
  position: absolute;
  bottom: -15%;
  left: -10%;
  width: 500px;
  height: 500px;
  background-color: rgba(79, 70, 229, 0.15);
  border-radius: 50%;
}

.container {
  position: relative;
  z-index: 1;
}

.card {
  border: 1px solid var(--border-light);
  border-radius: 12px;
  box-shadow: var(--shadow-xl);
  background: var(--bg-white);
}

.card h2 {
  color: var(--text-primary);
  font-weight: 600;
  font-size: 1.75rem;
}

.card p {
  color: var(--text-tertiary);
}

.form-control {
  padding: 12px 14px;
  border-radius: 6px;
}

.btn-primary {
  padding: 12px;
  border-radius: 6px;
  font-weight: 600;
}
</style>
