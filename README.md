Car Dealership Platform
=======================

Overview
--------
The Car Dealership Platform is a full-stack web application designed to allow users to browse, search, and filter car listings. Built with a responsive Angular frontend and a Go (Golang) backend, this platform provides a seamless user experience across devices. It also implements secure data transactions using HTTPS and enables cross-origin communication with CORS for flexibility in API interactions.

Features
--------
- Responsive Design: A fully responsive user interface built with Angular, ensuring accessibility on all devices.
- Backend Services in Go: Efficient server-side logic implemented using Go for managing car listings, search filters, and pagination.
- Secure API Communication: HTTPS is configured for secure data transmission, and CORS is enabled for handling cross-origin requests.
- Advanced Search Filters: Users can search and filter cars by make, model, year, and price.
- Database Integration: MongoDB (Atlas cloud) is used as the primary database for managing car data.

Technologies Used
-----------------
- Frontend: Angular, TypeScript
- Backend: Go (Golang)
- Database: MongoDB (Atlas cloud)
- Security: HTTPS, CORS

Getting Started
---------------

Prerequisites
-------------
Make sure you have the following installed:

- Go (Golang)
- Node.js and npm
- MongoDB Atlas Account (or MongoDB locally)
- Angular CLI
- Git

Installation and Running the Service
------------------------------------

1. Clone the Repository
-----------------------

git clone https://github.com/your-username/carsite.git
cd carsite


2. Backend Setup
----------------
Configure Environment Variables:

Create a .env file in the backend directory with the following content:


MONGO_URI=mongodb+srv://<username>:<password>@cluster.mongodb.net/carsite
PORT=8080


Install Backend Dependencies:


cd backend
go mod tidy


Run the Backend:


go run main.go


3. Frontend Setup
-----------------
Navigate to the Frontend Directory:


cd src


Install Frontend Dependencies:


npm install
npm install @angular/cli


Run the Frontend:


ng serve


Access the Application
----------------------
- Frontend: Open http://localhost:4200 in your browser.
- Backend: The backend API will be available at http://localhost:8080.