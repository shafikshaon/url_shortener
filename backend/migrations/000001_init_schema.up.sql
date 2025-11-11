-- Users table
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    api_key VARCHAR(64) UNIQUE,
    subscription_tier VARCHAR(20) DEFAULT 'free' CHECK (subscription_tier IN ('free', 'pro', 'business')),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_api_key ON users(api_key);

-- Links table
CREATE TABLE IF NOT EXISTS links (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    destination_url TEXT NOT NULL,
    title VARCHAR(255),
    tags TEXT[],
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_links_short_code ON links(short_code);
CREATE INDEX idx_links_user_id ON links(user_id);
CREATE INDEX idx_links_created_at ON links(created_at DESC);

-- Clicks table (raw click data)
CREATE TABLE IF NOT EXISTS clicks (
    id BIGSERIAL PRIMARY KEY,
    link_id BIGINT REFERENCES links(id) ON DELETE CASCADE,
    clicked_at TIMESTAMP DEFAULT NOW(),
    ip_address INET,
    country_code CHAR(2),
    referer TEXT,
    user_agent TEXT,
    device_type VARCHAR(20)
);

CREATE INDEX idx_clicks_link_id ON clicks(link_id, clicked_at DESC);
CREATE INDEX idx_clicks_clicked_at ON clicks(clicked_at DESC);

-- Analytics cache table (pre-computed daily stats)
CREATE TABLE IF NOT EXISTS link_analytics_daily (
    link_id BIGINT REFERENCES links(id) ON DELETE CASCADE,
    date DATE,
    click_count INTEGER DEFAULT 0,
    PRIMARY KEY (link_id, date)
);

CREATE INDEX idx_analytics_link_date ON link_analytics_daily(link_id, date DESC);

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger for users table
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger for links table
CREATE TRIGGER update_links_updated_at BEFORE UPDATE ON links
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
