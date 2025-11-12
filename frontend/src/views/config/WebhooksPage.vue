<template>
  <div class="webhooks-page">
    <div class="container-fluid py-4">
      <div class="row mb-4">
        <div class="col">
          <div class="d-flex justify-content-between align-items-center">
            <div>
              <h1 class="h3 mb-0">Webhooks</h1>
              <p class="text-muted mb-0">Configure webhooks to receive real-time event notifications</p>
            </div>
            <button class="btn btn-primary">
              <i class="bi bi-plus-circle"></i> Add Webhook
            </button>
          </div>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-8">
          <div class="card">
            <div class="card-body">
              <div class="text-center py-5">
                <i class="bi bi-webhook" style="font-size: 64px; color: #ccc;"></i>
                <h4 class="mt-3">No Webhooks Configured</h4>
                <p class="text-muted">Set up webhooks to receive real-time notifications when events occur</p>
                <button class="btn btn-primary">
                  <i class="bi bi-plus-circle"></i> Create Your First Webhook
                </button>
              </div>
            </div>
          </div>

          <div class="card mt-4">
            <div class="card-header bg-white">
              <h6 class="mb-0">Available Events</h6>
            </div>
            <div class="card-body">
              <div class="event-list">
                <div class="event-item">
                  <div>
                    <strong>link.created</strong>
                    <p class="text-muted small mb-0">Triggered when a new link is created</p>
                  </div>
                </div>
                <div class="event-item">
                  <div>
                    <strong>link.clicked</strong>
                    <p class="text-muted small mb-0">Triggered when a link is clicked</p>
                  </div>
                </div>
                <div class="event-item">
                  <div>
                    <strong>link.deleted</strong>
                    <p class="text-muted small mb-0">Triggered when a link is deleted</p>
                  </div>
                </div>
                <div class="event-item">
                  <div>
                    <strong>link.expired</strong>
                    <p class="text-muted small mb-0">Triggered when a link expires</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="col-lg-4">
          <div class="card">
            <div class="card-header bg-white">
              <h6 class="mb-0">Webhook Stats</h6>
            </div>
            <div class="card-body">
              <div class="mb-3">
                <small class="text-muted">Active Webhooks</small>
                <div class="h5 mb-0">0</div>
              </div>
              <div class="mb-3">
                <small class="text-muted">Events Sent Today</small>
                <div class="h5 mb-0">0</div>
              </div>
              <div>
                <small class="text-muted">Success Rate</small>
                <div>-</div>
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
.webhooks-page {
  background-color: var(--bg-primary);
  min-height: 100vh;
}

.event-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.event-item {
  padding: 16px;
  border: 1px solid var(--border-light);
  border-radius: 6px;
  background-color: var(--bg-white);
}
</style>
