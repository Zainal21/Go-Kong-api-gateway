# Kong API Gateway Docker Compose Setup

This document provides instructions for setting up Kong API Gateway along with a PostgreSQL database using Docker Compose.

## Prerequisites

- Docker and Docker Compose are installed on your system. see documentation [here](https://docs.docker.com/compose/install/)

## Getting Started

1. Clone the repository or create a directory to store the `docker-compose.yml` file.
2. Create a file named `POSTGRES_PASSWORD` in the same directory as the `docker-compose.yml` file. This file will contain the password for the PostgreSQL database.
3. Replace any environment variables enclosed in `${}` with your desired values. For example, replace `${KONG_PG_DATABASE}` with the name of your desired PostgreSQL database.

## Running the Services

Open a terminal window and navigate to the directory containing the `docker-compose.yml` file.

Run the following command to start the services:

```bash
docker compose up -d --builld
```

## Accessing Kong and Konga

- Kong Admin API: http://localhost:8001
- Kong Proxy: http://localhost:8000
- Konga Admin Interface: http://localhost:1337

## Example Backend Service with Golang

- User Service : port 8085
- Product Service : port 8084

## References

- [Go Programming Language](https://golang.org/)
- [Kong API Gateway ](https://docs.konghq.com/)
