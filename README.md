# 📦 eCommerce Microservices App

An eCommerce web application built with a **microservices architecture** using **Go (Golang)** for the backend and **React** for the frontend. Designed for scalability, maintainability, and high performance, this project demonstrates modern cloud-native development practices.

---

## 📚 Table of Contents

- [Tech Stack](#-tech-stack)
- [Architecture](#-architecture)
- [Features](#-features)
- [Getting Started](#-getting-started)
- [Microservices](#-microservices)
- [Folder Structure](#-folder-structure)
- [Contributing](#-contributing)
- [License](#-license)

---

## 🚀 Tech Stack

### Backend:
- Golang (Go)
- gRPC / REST APIs
- PostgreSQL / MongoDB
- Redis (Caching)
- RabbitMQ / Kafka (Event-Driven Architecture)
- Docker & Docker Compose
- JWT Authentication
- Prometheus & Grafana (Monitoring)
- Jaeger (Tracing)

### Frontend:
- React (with Hooks & Context)
- TailwindCSS / Material UI
- Axios for API communication
- Redux / Zustand (state management)
- React Router

---

## ⚙️ Architecture

This application is built using a microservices approach with clear boundaries for each domain service. Services communicate over gRPC or REST, and events are passed asynchronously via message brokers like RabbitMQ or Kafka.

![Microservices Architecture Diagram](https://via.placeholder.com/1000x500.png?text=Microservices+Architecture)

---

## ✨ Features

- User authentication (JWT + Role-based access)
- Product catalog & search
- Cart & Checkout system
- Order management
- Payment integration (dummy or Stripe/PayPal)
- Real-time notifications (WebSocket or SSE)
- Admin dashboard
- Service discovery (Consul or custom)
- Observability & logs

---

## 🛠 Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.22+
- Node.js 18+

### Clone the repo

```bash
git clone https://github.com/yourusername/ecommerce-microservices-go-react.git
cd ecommerce-microservices-go-react
```

### Start the stack

```bash
docker-compose up --build
```

### Run frontend separately

```bash
cd frontend
npm install
npm run dev
```

### Run backend microservices

```bash
cd backend/services/{service-name}
go run main.go
```

---

## 🔧 Microservices

| Service              | Description                          |
|----------------------|--------------------------------------|
| `auth-service`       | Handles login, registration, and JWT |
| `product-service`    | Manages product catalog              |
| `order-service`      | Handles order placement & tracking   |
| `payment-service`    | Processes payments (dummy gateway)   |
| `notification-service` | Sends emails / in-app alerts     |
| `gateway`            | API Gateway / Reverse Proxy          |

---

## 📂 Folder Structure

```
├── backend
│   ├── auth-service
│   ├── product-service
│   ├── order-service
│   └── ...
├── frontend
│   ├── public
│   ├── src
│   └── ...
├── docker-compose.yml
├── README.md
```

---

## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you'd like to change.

---

## 📄 License

This project is licensed under the GNU Public License.
