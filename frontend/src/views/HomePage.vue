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
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--primary-dark) 100%);
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
  background-color: rgba(139, 133, 255, 0.15);
  border-radius: 50%;
}

.hero-section::after {
  content: '';
  position: absolute;
  bottom: -20%;
  left: -5%;
  width: 400px;
  height: 400px;
  background-color: rgba(79, 70, 229, 0.15);
  border-radius: 50%;
}

.hero-section .container {
  position: relative;
  z-index: 1;
}

.hero-section h1 {
  color: white;
  font-weight: 700;
  letter-spacing: -0.03em;
}

.hero-section .lead {
  font-size: 1.125rem;
  opacity: 0.95;
  line-height: 1.7;
}

.hero-image {
  text-align: center;
  font-size: 15rem;
  opacity: 0.15;
}

.features-section {
  background-color: var(--bg-white);
  padding: 5rem 0;
}

.features-section h2 {
  color: var(--text-primary);
  font-weight: 600;
  font-size: 2rem;
  margin-bottom: 3rem;
}

.feature-card {
  background: var(--bg-white);
  border-radius: 8px;
  box-shadow: var(--shadow-sm);
  transition: all 0.2s ease;
  border: 1px solid var(--border-light);
  height: 100%;
}

.feature-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-lg);
  border-color: var(--primary-color);
}

.feature-card i {
  font-size: 3rem;
}

.feature-card h3 {
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 0.75rem;
}

.feature-card p {
  color: var(--text-secondary);
  line-height: 1.6;
  margin: 0;
}

.pricing-section {
  background-color: var(--bg-secondary);
  padding: 5rem 0;
}

.pricing-section h2 {
  color: var(--text-primary);
  font-weight: 600;
  font-size: 2rem;
  margin-bottom: 3rem;
}

.pricing-card {
  border-radius: 8px;
  border: 1px solid var(--border-light);
  transition: all 0.2s ease;
  position: relative;
  background: var(--bg-white);
  height: 100%;
  box-shadow: var(--shadow-sm);
}

.pricing-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-lg);
}

.pricing-card h3 {
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--text-primary);
}

.pricing-card.featured {
  border-color: var(--primary-color);
  border-width: 2px;
  transform: scale(1.05);
  box-shadow: var(--shadow-md);
}

.pricing-card.featured:hover {
  transform: scale(1.05) translateY(-4px);
  box-shadow: var(--shadow-xl);
}

.badge-featured {
  position: absolute;
  top: -14px;
  right: 24px;
  background: var(--primary-color);
  color: white;
  padding: 6px 16px;
  border-radius: 20px;
  font-weight: 600;
  font-size: 13px;
  box-shadow: 0 4px 6px -1px rgba(99, 91, 255, 0.3);
  letter-spacing: -0.01em;
}

.pricing-card ul {
  text-align: left;
  padding: 0;
}

.pricing-card ul li {
  color: var(--text-secondary);
  font-size: 14px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.pricing-card ul li i {
  color: var(--success-color);
  font-size: 16px;
}

.pricing-card .display-4 {
  color: var(--text-primary);
  font-weight: 700;
  letter-spacing: -0.03em;
}
</style>
