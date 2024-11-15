# Order Management System

## Overview
The **Order Management System** is a full-stack application with a **Golang** backend, a **Vue.js** Admin dashboard, and a **Vue.js** User interface. It allows users to place and track orders, while the Admin can manage orders and view detailed information.

---

## Requirements

- **Go** (v1.18+)
- **PostgreSQL**
- **Node.js** (v14+)
- **npm** (v6+)

---

## Setup Instructions

1. **Clone the repository**:

    ```bash
    git clone <repository-url>
    cd <project-folder>
    ```

2. **Backend (Golang)**:

    - Install **PostgreSQL** and create a `order_management` database.
    - Set up `.env` in the **Backend** folder with PostgreSQL credentials:

    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=store_db
    DB_USER=your_username
    DB_PASSWORD=your_password
    ```

    - Install dependencies:

    ```bash
    cd Backend
    go mod tidy
    ```

3. **Admin (Vue.js)**:

    - Install dependencies:

    ```bash
    cd Admin
    npm install
    ```

4. **User (Vue.js)**:

    - Install dependencies:

    ```bash
    cd User
    npm install
    ```

---

## Running the Project

1. **Backend**:

    ```bash
    cd Backend
    go run main.go
    ```

    The backend will be accessible at `http://localhost:5000`.

2. **Admin Dashboard**:

    ```bash
    cd Admin
    npm run serve
    ```

    Access at `http://localhost:8080`.

3. **User Interface**:

    ```bash
    cd User
    npm run serve
    ```

    Access at `http://localhost:8081`.

---

## Testing

1. **Backend (Go)**:

    ```bash
    cd Backend
    go test ./...
    ```

2. **Admin Dashboard (Vue.js)**:

    ```bash
    cd Admin
    npm run test:unit
    ```

---

## Features

- **Admin Interface**:
    - View detailed order information.
    

- **User Interface**:
    - Place new orders.


---

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Make changes and commit (`git commit -m 'Add new feature'`).
4. Push changes to your fork (`git push origin feature-name`).
5. Open a pull request.

---


