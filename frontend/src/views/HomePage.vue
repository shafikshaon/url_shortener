<template>
  <div class="home-page">
    <div class="hero-section">
      <div class="container">
        <div class="row align-items-center">
          <div class="col-lg-6">
            <h1 class="display-3 fw-bold mb-4">
              Shorten URLs.<br />Track Clicks.<br />Grow Smarter.
            </h1>
            <p class="lead mb-4">
              Simple, fast, and powerful URL shortening with advanced analytics.
              Perfect for marketers, developers, and businesses of all sizes.
            </p>
            <div class="d-flex gap-3">
              <RouterLink to="/signup" class="btn btn-primary btn-lg px-4">
                Get Started Free
              </RouterLink>
              <RouterLink to="/login" class="btn btn-outline-light btn-lg px-4">
                Sign In
              </RouterLink>
            </div>
          </div>
          <div class="col-lg-6">
            <div class="hero-image">
              <i class="bi bi-link-45deg"></i>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="features-section py-5">
      <div class="container">
        <h2 class="text-center mb-5">Why Choose Us?</h2>
        <div class="row g-4">
          <div class="col-md-4">
            <div class="feature-card text-center p-4">
              <i class="bi bi-lightning-charge-fill text-primary mb-3"></i>
              <h3>Lightning Fast</h3>
              <p>Create short links in seconds with our optimized infrastructure</p>
            </div>
          </div>
          <div class="col-md-4">
            <div class="feature-card text-center p-4">
              <i class="bi bi-bar-chart-fill text-success mb-3"></i>
              <h3>Detailed Analytics</h3>
              <p>Track clicks, locations, devices, and referers in real-time</p>
            </div>
          </div>
          <div class="col-md-4">
            <div class="feature-card text-center p-4">
              <i class="bi bi-shield-check-fill text-info mb-3"></i>
              <h3>Secure & Reliable</h3>
              <p>Your links are safe with enterprise-grade security</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="pricing-section py-5 bg-light">
      <div class="container">
        <h2 class="text-center mb-5">Simple Pricing</h2>
        <div class="row g-4">
          <div class="col-md-4">
            <div class="card pricing-card">
              <div class="card-body text-center p-4">
                <h3 class="mb-3">Free</h3>
                <div class="display-4 mb-3">$0</div>
                <ul class="list-unstyled mb-4">
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> 50 links</li>
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> 1,000 clicks/month</li>
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> 30-day analytics</li>
                </ul>
                <RouterLink to="/signup" class="btn btn-outline-primary w-100">
                  Get Started
                </RouterLink>
              </div>
            </div>
          </div>
          <div class="col-md-4">
            <div class="card pricing-card featured">
              <div class="badge-featured">Popular</div>
              <div class="card-body text-center p-4">
                <h3 class="mb-3">Pro</h3>
                <div class="display-4 mb-3">
                  $9<span class="fs-6">/month</span>
                </div>
                <ul class="list-unstyled mb-4">
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> 500 links</li>
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> 10,000 clicks/month</li>
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> 1-year analytics</li>
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> API access</li>
                </ul>
                <RouterLink to="/signup" class="btn btn-primary w-100">
                  Get Started
                </RouterLink>
              </div>
            </div>
          </div>
          <div class="col-md-4">
            <div class="card pricing-card">
              <div class="card-body text-center p-4">
                <h3 class="mb-3">Business</h3>
                <div class="display-4 mb-3">
                  $29<span class="fs-6">/month</span>
                </div>
                <ul class="list-unstyled mb-4">
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> 5,000 links</li>
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> 100,000 clicks/month</li>
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> Unlimited analytics</li>
                  <li class="mb-2"><i class="bi bi-check-circle-fill text-success"></i> Custom domain</li>
                </ul>
                <RouterLink to="/signup" class="btn btn-outline-primary w-100">
                  Get Started
                </RouterLink>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { RouterLink } from 'vue-router'
</script>

<style scoped>
.hero-section {
  background-color: #7C3AED;
  color: white;
  padding: 5rem 0;
  min-height: 70vh;
  display: flex;
  align-items: center;
  position: relative;
  overflow: hidden;
}

.hero-section::before {
  content: '';
  position: absolute;
  top: -50%;
  right: -10%;
  width: 500px;
  height: 500px;
  background-color: rgba(167, 139, 250, 0.1);
  border-radius: 50%;
}

.hero-section::after {
  content: '';
  position: absolute;
  bottom: -20%;
  left: -5%;
  width: 400px;
  height: 400px;
  background-color: rgba(109, 40, 217, 0.1);
  border-radius: 50%;
}

.hero-section .container {
  position: relative;
  z-index: 1;
}

.hero-image {
  text-align: center;
  font-size: 15rem;
  opacity: 0.15;
}

.features-section {
  background-color: white;
}

.feature-card {
  background: white;
  border-radius: 1rem;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
  transition: all 0.3s ease;
  border: 1px solid #E5E7EB;
  height: 100%;
}

.feature-card:hover {
  transform: translateY(-0.5rem);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  border-color: #7C3AED;
}

.feature-card i {
  font-size: 3rem;
}

.pricing-section {
  background-color: #F9FAFB;
}

.pricing-card {
  border-radius: 1rem;
  border: 2px solid #E5E7EB;
  transition: all 0.3s ease;
  position: relative;
  background: white;
  height: 100%;
}

.pricing-card:hover {
  transform: translateY(-0.5rem);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.pricing-card.featured {
  border-color: #7C3AED;
  border-width: 3px;
  transform: scale(1.05);
}

.pricing-card.featured:hover {
  transform: scale(1.05) translateY(-0.5rem);
}

.badge-featured {
  position: absolute;
  top: -14px;
  right: 24px;
  background: #7C3AED;
  color: white;
  padding: 0.375rem 1.25rem;
  border-radius: 1.5rem;
  font-weight: 600;
  font-size: 0.875rem;
  box-shadow: 0 4px 6px -1px rgba(124, 58, 237, 0.3);
}
</style>
