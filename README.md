# 🎬 CineSpot - Backend

Welcome to the backend of **CineSpot** – a movie platform API built with **Node.js**, **Express.js**, and **MongoDB**. This backend provides RESTful APIs for managing movie data, including details like title, description, director, and image.

---

## 🚀 Features

- Get all movies
- Get a single movie by ID
- Create a new movie
- Update movie details
- Delete a movie
- CORS and JSON Middleware integrated
- Clean and modular code structure using Router

---

## 🛠️ Tech Stack

- Node.js
- Express.js
- MongoDB
- Mongoose (Optional)
- Gorilla Mux (if Go is used)
- dotenv
- CORS

---

## 📁 Project Structure

---

## 🔗 API Endpoints

| Method | Endpoint        | Description              |
| ------ | --------------- | ------------------------ |
| GET    | /api/movies     | Get all movies           |
| GET    | /api/movies/:id | Get a movie by ID        |
| POST   | /api/movies     | Create a new movie       |
| PUT    | /api/movies/:id | Update an existing movie |
| DELETE | /api/movies/:id | Delete a movie by ID     |

---

## 📦 Installation & Setup

```bash
# Clone the repository
git clone https://github.com/your-username/cinespot-backend.git

# Navigate into the project
cd cinespot-backend

# Install dependencies
npm install

# Run the server
npm run dev
```
