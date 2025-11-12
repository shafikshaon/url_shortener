<template>
  <div class="members-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <div class="d-flex justify-content-between align-items-center">
            <div>
              <h1 class="h3 mb-0">Team Members</h1>
              <p class="text-muted mb-0">Manage your team members and their permissions</p>
            </div>
            <button class="btn btn-primary">
              <i class="bi bi-plus-circle"></i> Invite Member
            </button>
          </div>
        </div>
      </div>

      <div class="row">
        <div class="col">
          <div class="card">
            <div class="card-body">
              <div class="table-responsive">
                <table class="table">
                  <thead>
                    <tr>
                      <th>Member</th>
                      <th>Role</th>
                      <th>Status</th>
                      <th>Joined</th>
                      <th>Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>
                        <div class="d-flex align-items-center gap-2">
                          <div class="member-avatar">
                            <i class="bi bi-person-fill"></i>
                          </div>
                          <div>
                            <div class="fw-bold">{{ userEmail }}</div>
                            <small class="text-muted">Owner</small>
                          </div>
                        </div>
                      </td>
                      <td><span class="badge bg-primary">Owner</span></td>
                      <td><span class="badge bg-success">Active</span></td>
                      <td>January 2024</td>
                      <td>
                        <button class="btn btn-sm btn-outline-secondary" disabled>
                          <i class="bi bi-three-dots"></i>
                        </button>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <div class="text-center py-4">
                <i class="bi bi-people" style="font-size: 48px; color: #ccc;"></i>
                <p class="text-muted mt-3">No other team members yet</p>
                <button class="btn btn-primary">Invite Your First Member</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useAuthStore } from '@/store/auth'

const authStore = useAuthStore()
const userEmail = computed(() => authStore.currentUser?.email || '')
</script>

<style scoped>
.members-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.member-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 18px;
}
</style>
