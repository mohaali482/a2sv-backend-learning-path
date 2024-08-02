# Task Manager

This is a Go project that uses the Gin web framework to create a task management API.

## Getting Started

### Prerequisites

- Go programming language installed on your system

### Installation

1. Clone the repository:

```shell
git clone https://github.com/mohaali482/a2sv-backend-learning-path.git
```

2. Navigate to the project directory:

```shell
cd task-5/task_manager
```

3. Install the required dependencies:

```shell
go get
```

4. Copy the environment variable example to a new `.env` file.

```shell
cp .env.example .env
```

5. Fill the correct information in the `.env` file.

### Usage

1. Build the application:

```shell
go build -o task_manager
```

2. Run the compiled binary:

```shell
./task_manager
```

This will start the task manager api and it can be accessed through the provided `HOST_URL` or `http://localhost:8000`.

### Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

### License

This project is licensed under the MIT License.
