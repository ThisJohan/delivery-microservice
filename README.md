# Delivery Microservice

This is the Delivery Microservice, part of a suite of applications including `delivery`, `core` and `tpl`. This microservice handles the delivery operations within our system.

## Prerequisites

Before running the application, ensure you have the following:

- Golang installed
- PostgreSQL database and redis running (you can run them with docker compose) 

## Database Setup

Before starting the application, you need to set up the database with the required type for shipment status. Run the following SQL command in your PostgreSQL database:

```sql
CREATE TYPE shipment_status AS ENUM ('PENDING', 'ASSIGNED', 'DELIVERED', 'CANCELED', 'QUEUED');
```

## Project Running

You can run any of services with run command: `make run appName`
ex: `make run delivery` 