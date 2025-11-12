<template>
  <div class="org-settings-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Organization Settings</h1>
          <p class="text-muted">Manage your organization profile and preferences</p>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Organization Information</h6>
            </div>
            <div class="card-body">
              <form>
                <div class="mb-3">
                  <label class="form-label">Organization Name</label>
                  <input type="text" class="form-control" placeholder="Enter organization name" value="My Organization">
                </div>
                <div class="mb-3">
                  <label class="form-label">Organization ID</label>
                  <input type="text" class="form-control" value="org_abc123" disabled>
                  <small class="text-muted">This ID is used for API integrations</small>
                </div>
                <div class="mb-3">
                  <label class="form-label">Website</label>
                  <input type="url" class="form-control" placeholder="https://example.com">
                </div>
                <div class="mb-3">
                  <label class="form-label">Industry</label>
                  <select class="form-select">
                    <option>Technology</option>
                    <option>Marketing</option>
                    <option>E-commerce</option>
                    <option>Education</option>
                    <option>Other</option>
                  </select>
                </div>
                <button type="submit" class="btn btn-primary">Save Changes</button>
              </form>
            </div>
          </div>

          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Organization Logo</h6>
            </div>
            <div class="card-body">
              <div class="d-flex align-items-center gap-3">
                <div class="org-logo">
                  <i class="bi bi-building"></i>
                </div>
                <div>
                  <button class="btn btn-outline-primary btn-sm">Upload Logo</button>
                  <p class="text-muted mb-0 mt-2" style="font-size: 13px;">PNG or SVG. Max size 2MB.</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="col-lg-4">
          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Organization Stats</h6>
            </div>
            <div class="card-body">
              <div class="mb-3">
                <small class="text-muted">Total Members</small>
                <div class="h5 mb-0">1</div>
              </div>
              <div class="mb-3">
                <small class="text-muted">Total Links</small>
                <div class="h5 mb-0">0</div>
              </div>
              <div>
                <small class="text-muted">Created</small>
                <div>January 2024</div>
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
.org-settings-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.org-logo {
  width: 80px;
  height: 80px;
  border-radius: 12px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-hover) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 36px;
}

/* Responsive Styles */
@media (max-width: 991px) {
  .container-fluid {
    padding: 1.5rem !important;
  }

  .col-lg-8,
  .col-lg-4 {
    margin-bottom: 1.5rem;
  }
}

@media (max-width: 768px) {
  .container-fluid {
    padding: 1rem !important;
  }

  .org-settings-page h1 {
    font-size: 1.5rem;
  }

  .card-header h5,
  .card-header h6 {
    font-size: 0.95rem;
  }

  .card-body {
    padding: 1rem;
  }

  .col-md-4,
  .col-md-6 {
    margin-bottom: 1rem;
  }

  .d-flex.justify-content-between {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 0.5rem;
  }

  .org-logo {
    width: 60px;
    height: 60px;
    font-size: 28px;
  }
}

@media (max-width: 576px) {
  .container-fluid {
    padding: 0.75rem !important;
  }

  .org-settings-page h1 {
    font-size: 1.25rem;
  }

  .org-settings-page .text-muted {
    font-size: 0.9rem;
  }

  .card {
    border-radius: 8px;
    margin-bottom: 1rem !important;
  }

  .card-header {
    padding: 0.75rem 1rem;
  }

  .card-header h5,
  .card-header h6 {
    font-size: 0.9rem;
  }

  .card-body {
    padding: 0.75rem;
  }

  .form-label {
    font-size: 0.9rem;
    margin-bottom: 0.25rem;
  }

  .form-control,
  .form-select {
    padding: 0.5rem 0.75rem;
    font-size: 0.9rem;
  }

  .mb-3 {
    margin-bottom: 0.75rem !important;
  }

  .btn {
    font-size: 0.9rem;
    padding: 0.5rem 0.75rem;
  }

  .btn-sm {
    font-size: 0.8rem;
    padding: 0.35rem 0.6rem;
  }

  .org-logo {
    width: 50px;
    height: 50px;
    font-size: 24px;
  }

  small {
    font-size: 0.8rem;
  }
}
</style>
