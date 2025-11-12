<template>
  <div class="integrations-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <h1 class="h3 mb-0">Integrations</h1>
          <p class="text-muted">Connect LinkShort with your favorite tools and services</p>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card mb-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Available Integrations</h6>
            </div>
            <div class="card-body">
              <div class="row">
                <div class="col-md-6 mb-3">
                  <div class="integration-card">
                    <div class="d-flex gap-3">
                      <div class="integration-icon">
                        <i class="bi bi-slack"></i>
                      </div>
                      <div class="flex-grow-1">
                        <h6 class="mb-1">Slack</h6>
                        <p class="text-muted small mb-2">Receive notifications in Slack channels</p>
                        <button class="btn btn-sm btn-outline-primary">Connect</button>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="col-md-6 mb-3">
                  <div class="integration-card">
                    <div class="d-flex gap-3">
                      <div class="integration-icon zapier">
                        <i class="bi bi-lightning-charge"></i>
                      </div>
                      <div class="flex-grow-1">
                        <h6 class="mb-1">Zapier</h6>
                        <p class="text-muted small mb-2">Automate workflows with 5000+ apps</p>
                        <button class="btn btn-sm btn-outline-primary">Connect</button>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="col-md-6 mb-3">
                  <div class="integration-card">
                    <div class="d-flex gap-3">
                      <div class="integration-icon analytics">
                        <i class="bi bi-google"></i>
                      </div>
                      <div class="flex-grow-1">
                        <h6 class="mb-1">Google Analytics</h6>
                        <p class="text-muted small mb-2">Track link clicks in Google Analytics</p>
                        <button class="btn btn-sm btn-outline-primary">Connect</button>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="col-md-6 mb-3">
                  <div class="integration-card">
                    <div class="d-flex gap-3">
                      <div class="integration-icon discord">
                        <i class="bi bi-discord"></i>
                      </div>
                      <div class="flex-grow-1">
                        <h6 class="mb-1">Discord</h6>
                        <p class="text-muted small mb-2">Post updates to Discord servers</p>
                        <button class="btn btn-sm btn-outline-primary">Connect</button>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="col-md-6 mb-3">
                  <div class="integration-card">
                    <div class="d-flex gap-3">
                      <div class="integration-icon notion">
                        <i class="bi bi-journal-text"></i>
                      </div>
                      <div class="flex-grow-1">
                        <h6 class="mb-1">Notion</h6>
                        <p class="text-muted small mb-2">Sync links to Notion databases</p>
                        <button class="btn btn-sm btn-outline-primary">Connect</button>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="col-md-6 mb-3">
                  <div class="integration-card">
                    <div class="d-flex gap-3">
                      <div class="integration-icon github">
                        <i class="bi bi-github"></i>
                      </div>
                      <div class="flex-grow-1">
                        <h6 class="mb-1">GitHub</h6>
                        <p class="text-muted small mb-2">Create links from GitHub releases</p>
                        <button class="btn btn-sm btn-outline-primary">Connect</button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="alert alert-info">
            <i class="bi bi-lightbulb"></i>
            <strong>Need a custom integration?</strong> Contact our team to discuss building a custom integration for your needs.
          </div>
        </div>

        <div class="col-lg-4">
          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Connected Integrations</h6>
            </div>
            <div class="card-body">
              <div class="text-center py-4">
                <i class="bi bi-plug" style="font-size: 48px; color: #ccc;"></i>
                <p class="text-muted mt-3 mb-0">No integrations connected yet</p>
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
.integrations-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.integration-card {
  padding: 20px;
  border: 1px solid var(--border-light);
  border-radius: 8px;
  background-color: var(--bg-white);
  transition: all 0.2s ease;
}

.integration-card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--border-color);
}

.integration-icon {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  background-color: #5865F2;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
  flex-shrink: 0;
}

.integration-icon.zapier {
  background-color: #FF4A00;
}

.integration-icon.analytics {
  background-color: #4285F4;
}

.integration-icon.discord {
  background-color: #5865F2;
}

.integration-icon.notion {
  background-color: #000000;
}

.integration-icon.github {
  background-color: #181717;
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

  .integrations-page h1 {
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

  .integration-card {
    padding: 15px;
  }

  .integration-card h6 {
    font-size: 0.95rem;
  }

  .integration-icon {
    width: 42px;
    height: 42px;
    font-size: 20px;
  }
}

@media (max-width: 576px) {
  .container-fluid {
    padding: 0.75rem !important;
  }

  .integrations-page h1 {
    font-size: 1.25rem;
  }

  .integrations-page .text-muted {
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

  .form-control {
    padding: 0.5rem 0.75rem;
    font-size: 0.9rem;
  }

  .mb-3 {
    margin-bottom: 0.75rem !important;
  }

  .alert {
    padding: 0.75rem;
    font-size: 0.85rem;
  }

  .btn {
    font-size: 0.9rem;
    padding: 0.5rem 0.75rem;
  }

  .btn-sm {
    font-size: 0.8rem;
    padding: 0.35rem 0.6rem;
  }

  .integration-card {
    padding: 12px;
  }

  .integration-card h6 {
    font-size: 0.9rem;
    margin-bottom: 0.5rem;
  }

  .integration-card .small {
    font-size: 0.8rem;
    margin-bottom: 0.5rem !important;
  }

  .integration-icon {
    width: 38px;
    height: 38px;
    font-size: 18px;
  }

  .text-center i[style*="font-size: 48px"] {
    font-size: 36px !important;
  }

  small {
    font-size: 0.8rem;
  }
}
</style>
