# URL Shortener - MVP

A lean, production-ready URL shortening service built with Go (Gin) backend and Vue.js (Bootstrap) frontend.

## 🎯 Features

### Core Features
- **URL Shortening**: Random 7-character codes or custom aliases
- **Link Management**: Create, edit, delete, and organize links with tags
- **Analytics**: Track clicks with daily breakdowns, geographic data, referrers, and device types
- **User Authentication**: Secure email/password authentication with JWT
- **API Access**: RESTful API with key-based authentication (Pro/Business tiers)
- **Subscription Tiers**: Free, Pro, and Business plans with different limits

### Technical Highlights
- **Backend**: Go 1.21+ with Gin framework
- **Frontend**: Vue.js 3 with Bootstrap 5
- **Database**: PostgreSQL with optimized indexes
- **Caching**: Optional Redis support
- **Analytics**: Pre-computed daily statistics for performance
- **Security**: Password hashing, JWT tokens, SQL injection protection

## 📁 Project Structure

```
url_shortener/
├── backend/                 # Go backend
│   ├── cmd/
│   │   └── server/         # Main application entry point
│   ├── internal/
│   │   ├── api/           # HTTP handlers
│   │   ├── auth/          # Authentication & JWT
│   │   ├── database/      # Database repositories
│   │   ├── models/        # Data models
│   │   ├── service/       # Business logic
│   │   └── analytics/     # Click tracking
│   ├── migrations/        # Database migrations
│   ├── config/           # Configuration management
│   ├── Dockerfile
│   ├── Makefile
│   └── go.mod
│
├── frontend/              # Vue.js frontend
│   ├── src/
│   │   ├── components/   # Reusable components
│   │   ├── views/        # Page components
│   │   ├── services/     # API client
│   │   ├── store/        # Pinia state management
│   │   ├── router/       # Vue Router configuration
│   │   └── App.vue
│   ├── public/
│   ├── Dockerfile
│   ├── nginx.conf
│   └── package.json
│
└── docker-compose.yml     # Full stack deployment
```

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose (recommended)
- OR: Go 1.21+, Node.js 18+, PostgreSQL 15+

### Option 1: Docker Compose (Recommended)

1. **Clone the repository**
```bash
git clone https://github.com/shafikshaon/url_shortener.git
cd url_shortener
```

2. **Start all services**
```bash
docker-compose up -d
```

3. **Run database migrations**
```bash
docker-compose exec backend sh -c "migrate -path migrations -database 'postgresql://urlshortener:urlshortener_password@postgres:5432/urlshortener?sslmode=disable' up"
```

4. **Access the application**
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- Health check: http://localhost:8080/health

### Option 2: Local Development

#### Backend Setup

1. **Install dependencies**
```bash
cd backend
go mod download
```

2. **Configure environment**
```bash
cp .env.example .env
# Edit .env with your database credentials
```

3. **Start PostgreSQL**
```bash
# Using Docker
docker run -d \
  --name url-shortener-db \
  -e POSTGRES_USER=urlshortener \
  -e POSTGRES_PASSWORD=urlshortener_password \
  -e POSTGRES_DB=urlshortener \
  -p 5432:5432 \
  postgres:15-alpine
```

4. **Run migrations**
```bash
# Install migrate tool
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run migrations
export DB_USER=urlshortener
export DB_PASSWORD=urlshortener_password
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=urlshortener
export DB_SSLMODE=disable

./scripts/migrate.sh
```

5. **Start the backend**
```bash
go run cmd/server/main.go
# Or use the Makefile
make run
```

#### Frontend Setup

1. **Install dependencies**
```bash
cd frontend
npm install
```

2. **Configure environment**
```bash
cp .env.example .env
# Defaults to http://localhost:8080/api/v1
```

3. **Start development server**
```bash
npm run dev
```

Frontend will be available at http://localhost:5173

## 📊 Database Schema

### Users Table
- User authentication and subscription management
- Fields: id, email, password_hash, api_key, subscription_tier, timestamps

### Links Table
- Shortened link information
- Fields: id, user_id, short_code, destination_url, title, tags, expires_at, timestamps
- Indexes: short_code, user_id, created_at

### Clicks Table
- Raw click tracking data
- Fields: id, link_id, clicked_at, ip_address, country_code, referer, user_agent, device_type
- Indexes: link_id + clicked_at, clicked_at

### Link Analytics Daily Table
- Pre-computed daily statistics
- Fields: link_id, date, click_count
- Primary key: (link_id, date)

## 🔌 API Documentation

### Authentication Endpoints

**Signup**
```bash
POST /api/v1/auth/signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

**Login**
```bash
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword123"
}

# Response
{
  "token": "jwt_token_here",
  "user": { ... }
}
```

### Link Management Endpoints

All endpoints require authentication via JWT or API key.

**Create Link**
```bash
POST /api/v1/links
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "destination_url": "https://example.com/very-long-url",
  "short_code": "custom",  # optional
  "title": "My Link",      # optional
  "tags": ["marketing", "campaign"],  # optional
  "expires_at": "2024-12-31T23:59:59Z"  # optional
}
```

**List Links**
```bash
GET /api/v1/links?limit=20&offset=0&search=keyword&sort=created_desc
Authorization: Bearer <jwt_token>
```

**Get Link Details**
```bash
GET /api/v1/links/:id
Authorization: Bearer <jwt_token>
```

**Update Link**
```bash
PATCH /api/v1/links/:id
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "destination_url": "https://example.com/updated-url",
  "title": "Updated Title",
  "tags": ["updated", "tags"]
}
```

**Delete Link**
```bash
DELETE /api/v1/links/:id
Authorization: Bearer <jwt_token>
```

**Get Link Statistics**
```bash
GET /api/v1/links/:id/stats
Authorization: Bearer <jwt_token>

# Response
{
  "total_clicks": 1250,
  "last_30_days": 450,
  "daily_clicks": [...],
  "countries": [...],
  "referers": [...],
  "device_types": [...]
}
```

### Analytics Endpoints

**Get User Analytics**
```bash
GET /api/v1/analytics
Authorization: Bearer <jwt_token>

# Response
{
  "total_links": 15,
  "total_clicks": 5000,
  "clicks_this_month": 1200
}
```

### Redirect Endpoint

**Short URL Redirect**
```bash
GET /:shortCode
# Redirects to destination URL and tracks the click
```

## 💰 Subscription Tiers

| Feature | Free | Pro | Business |
|---------|------|-----|----------|
| Links | 50 | 500 | 5,000 |
| Clicks/month | 1,000 | 10,000 | 100,000 |
| Analytics retention | 30 days | 1 year | Unlimited |
| API access | ❌ | ✅ | ✅ |
| API requests/day | 0 | 5,000 | 50,000 |
| Custom domain | ❌ | ❌ | ✅ |
| Price | $0 | $9/mo | $29/mo |

## 🛠️ Development

### Backend Commands

```bash
# Build
make build

# Run
make run

# Run tests
make test

# Clean build artifacts
make clean

# Run migrations
make migrate-up

# Rollback migrations
make migrate-down

# Create new migration
make migrate-create name=add_new_table

# Install development tools
make install-tools
```

### Frontend Commands

```bash
# Development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

## 🔒 Security Considerations

1. **Change JWT Secret**: Update `JWT_SECRET` in production
2. **Use HTTPS**: Always use HTTPS in production
3. **Database Security**: Use strong passwords and restrict access
4. **Rate Limiting**: Implement rate limiting at the reverse proxy level
5. **CORS Configuration**: Update CORS settings for your domain
6. **Environment Variables**: Never commit `.env` files

## 📈 Performance Tips

1. **Enable Redis**: Set `REDIS_ENABLED=true` for caching hot links
2. **Database Indexes**: Already optimized, but monitor query performance
3. **CDN**: Use a CDN for frontend assets in production
4. **Connection Pooling**: Adjust `SetMaxOpenConns` based on load
5. **Horizontal Scaling**: Run multiple backend instances behind a load balancer

## 🚢 Production Deployment

### Using Docker Compose

1. Update `docker-compose.yml` with production settings
2. Set strong passwords for PostgreSQL
3. Configure proper volumes for data persistence
4. Use environment variables for sensitive data
5. Set up SSL/TLS with a reverse proxy (Nginx/Traefik)

### Environment Variables

Backend:
- `SERVER_PORT`: API server port (default: 8080)
- `BASE_URL`: Public URL for short links
- `DB_*`: Database connection settings
- `JWT_SECRET`: Secret key for JWT tokens
- `REDIS_*`: Redis configuration

Frontend:
- `VITE_API_BASE_URL`: Backend API URL

## 📝 License

MIT License - feel free to use this project for personal or commercial purposes.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 🐛 Troubleshooting

### Backend won't start
- Check PostgreSQL is running and accessible
- Verify database credentials in `.env`
- Ensure migrations have been run

### Frontend can't connect to backend
- Check `VITE_API_BASE_URL` in frontend `.env`
- Verify CORS settings in backend
- Ensure backend is running on correct port

### Links not redirecting
- Check short code exists in database
- Verify base URL is correct
- Check for expired links

## 📞 Support

For issues and questions:
- GitHub Issues: https://github.com/shafikshaon/url_shortener/issues
- Email: support@yourdomain.com

## 🎉 Acknowledgments

Built with:
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [Vue.js](https://vuejs.org/) - Progressive JavaScript framework
- [Bootstrap](https://getbootstrap.com/) - CSS framework
- [PostgreSQL](https://www.postgresql.org/) - Database
- [Chart.js](https://www.chartjs.org/) - Analytics visualization

---

**Built with ❤️ for the MVP community**