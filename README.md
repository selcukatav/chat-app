# Chat-App

Chat-App is a Discord alternative, providing basic voice and text chat functionalities. This application supports PostgreSQL for persistent data storage and Redis for real-time messaging.

## Features
- **Text Messaging**: Real-time text chat functionality.
- **Voice Chat**: Basic voice chat capabilities.
- **User Authentication**: Secure login and registration system.
- **API Documentation**: Integrated Swagger UI for exploring and testing APIs.

---

## Prerequisites

Before starting, ensure you have the following installed on your system:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://go.dev/) (if building manually)

---

## Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/selcukatav/chat-app.git
cd chat-app
```

### 2. Setup Environment Variables
Create a `.env` file in the project root and configure the following variables:

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=123qwe
POSTGRES_DB=chat_app
POSTGRES_HOST=localhost
POSTGRES_PORT=5432

REDIS_HOST=redis
REDIS_PORT=6379

APP_PORT=3000
```

### 3. Start Services with Docker Compose

Run the following command to start PostgreSQL, Redis, and the Chat-App backend:

```bash
docker-compose up -d
```

This will:
- Start a PostgreSQL container for persistent data storage.
- Start a Redis container for real-time messaging.
- Launch the Chat-App backend.

**!!IMPORTANT!!**
chatapp_backend still doens't work with docker compose since I couldn't figure out what to use instead of POSTGRE_HOST=localhost. So after starting the docker compose up just manually start the server with go run main.go in ./server. I'm open to suggestions for to fix the issue 

### 4. Access the Application

- **API Base URL**: `http://localhost:3000`
- **Swagger UI**: `http://localhost:3000/swagger/index.html`

---

## API Documentation

Chat-App provides API documentation using Swagger. To access it:

1. Start the application.
2. Navigate to `http://localhost:3000/swagger/index.html` in your browser.
3. Explore and test the available endpoints directly in the browser.

---

## Docker Setup Details

### PostgreSQL Setup
PostgreSQL is configured via Docker Compose. The database is initialized with the following default values:

- **Username**: Defined in `.env` as `POSTGRES_USER`
- **Password**: Defined in `.env` as `POSTGRES_PASSWORD`
- **Database Name**: Defined in `.env` as `POSTGRES_DB`

The PostgreSQL service will be available on port `5432` inside the Docker network.

### Redis Setup
Redis is used for real-time messaging and is configured in the `docker-compose.yml` file. The service will be available on port `6379` inside the Docker network.

---

## Development

To run the application locally without Docker:

1. Ensure PostgreSQL and Redis are running locally.
2. Configure the `.env` file with the appropriate database connection strings.
3. Run the application:

   ```bash
   go run main.go
   ```

---

## Contribution

Contributions are welcome! Feel free to open issues or submit pull requests for new features or bug fixes.

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/your-feature`.
3. Commit your changes: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Open a pull request.

---

