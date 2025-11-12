<template>
  <div class="security-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Security Settings</h1>
          <p class="text-muted">Manage your account security and authentication</p>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Change Password</h6>
            </div>
            <div class="card-body">
              <form>
                <div class="mb-3">
                  <label class="form-label">Current Password</label>
                  <input type="password" class="form-control" placeholder="Enter current password">
                </div>
                <div class="mb-3">
                  <label class="form-label">New Password</label>
                  <input type="password" class="form-control" placeholder="Enter new password">
                </div>
                <div class="mb-3">
                  <label class="form-label">Confirm New Password</label>
                  <input type="password" class="form-control" placeholder="Confirm new password">
                </div>
                <button type="submit" class="btn btn-primary">Update Password</button>
              </form>
            </div>
          </div>

          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Two-Factor Authentication</h6>
            </div>
            <div class="card-body">
              <p class="text-muted">Add an extra layer of security to your account by enabling two-factor authentication.</p>
              <div class="alert alert-warning">
                <i class="bi bi-shield-exclamation"></i> Two-factor authentication is currently <strong>disabled</strong>
              </div>
              <button class="btn btn-primary">Enable 2FA</button>
            </div>
          </div>

          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Active Sessions</h6>
            </div>
            <div class="card-body">
              <div class="session-item">
                <div class="d-flex justify-content-between align-items-center">
                  <div>
                    <div class="fw-bold"><i class="bi bi-laptop"></i> Current Session</div>
                    <small class="text-muted">Chrome on MacOS • United States</small>
                  </div>
                  <span class="badge bg-success">Active</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
.security-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.session-item {
  padding: 16px;
  border: 1px solid var(--border-light);
  border-radius: 8px;
}
</style>
