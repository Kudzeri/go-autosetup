# go-autosetup

go-autosetup is a project generator for Go that automates the initial setup of your Go projects. It allows you to choose various options such as the web framework (none, Gin, Fiber), database (SQLite, PostgreSQL, MongoDB), and additional features like Swagger integration and CORS middleware. The tool generates a complete project structure with the necessary configuration files and code templates, so you can focus on building your application.

## Features

- **Framework Selection**  
  Choose your preferred web framework:
  - **No framework:** Standard library HTTP server.
  - **Gin:** A high-performance HTTP web framework.
  - **Fiber:** A fast web framework inspired by Express.

- **Database Integration**  
  Automatically configure your project for:
  - **SQLite:** Lightweight, file-based database.
  - **PostgreSQL:** Powerful, open-source relational database.
  - **MongoDB:** Flexible NoSQL database.

- **Swagger Integration**  
  Optionally add Swagger support for API documentation.

- **CORS Middleware**  
  Optionally include Cross-Origin Resource Sharing (CORS) middleware:
  - Provides a custom CORS implementation for Gin and Fiber.
  - Includes a basic placeholder if no framework is selected.

- **Automatic Dependency Management**  
  Generates a minimal `go.mod` file and automatically runs `go mod tidy` to fetch the latest dependencies.

- **Organized Project Structure**  
  Creates a well-organized project layout, including folders for controllers, models, services, middleware, routes, and utilities.

## Installation

Ensure you have Go 1.16 or later installed on your system.

Clone the repository:

```bash
git clone https://github.com/Kudzeri/go-autosetup.git
```

Build the binary:

```bash
cd go-autosetup
go build -o go-autosetup .
```

## Usage

Run the binary:

```bash
./go-autosetup
```

You will be prompted to provide:

- **Project Name:** The name of your new project.
- **Framework Selection:**  
  - `0` for no framework (standard library)
  - `1` for Gin Gonic
  - `2` for Fiber
- **Database Selection:**  
  - `0` for SQLite
  - `1` for PostgreSQL
  - `2` for MongoDB
- **Swagger Integration:** Yes/No
- **CORS Middleware:** Yes/No

Based on your choices, the tool will generate a new project with a structure similar to:

```
your-project/
├── config/
│   └── db.go
├── controller/
├── middleware/
│   └── cors.go         // Generated if CORS is enabled
├── model/
├── repository/
├── route/
├── service/
├── utils/
├── .env
├── go.mod
├── go.sum
└── main.go
```

After generation, the command `go mod tidy` is automatically executed in your project directory to fetch and organize dependencies.

## Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to check the [issues page](https://github.com/Kudzeri/go-autosetup/issues) if you want to contribute.

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Author

Created by [Kudzeri](https://github.com/Kudzeri)
