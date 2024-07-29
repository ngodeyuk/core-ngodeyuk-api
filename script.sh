#!/bin/bash

GO_BINARY="cmd/web/main.go"
GO_ENV_FILE=".env"
DOCKER_COMPOSE_FILE="./docker-compose.yml"
MIGRATE_PATH="database/migrations"

usage() {
    echo "Usage: $0 {run|build|clean} [options]"
    echo "  run        Run the Go application."
    echo "  build      Build the Go application binary."
    echo "  clean      Clean up build artifacts."
    echo "  -h         Show this help message."
    echo ""
    echo "Usage: $0 {create|up|down} [migration_name]"
    echo "  create [migration_name]  Create a new migration file with the given name."
    echo "  up                        Apply all pending migrations."
    echo "  down                      Roll back the most recent migration."
    exit 1
}

# Check and load environment variables from .env file
load_env() {
    if [ -f "$GO_ENV_FILE" ]; then
        export $(grep -v '^#' "$GO_ENV_FILE" | xargs)
    else
        echo "Error: $GO_ENV_FILE not found. Please create it or check its path."
        exit 1
    fi
}

# Check if migrate command is available
check_migrate() {
    command -v migrate >/dev/null 2>&1 || { echo >&2 "Go Migrate is not installed. Please install it first."; exit 1; }
}

# Run the Go application
run_app() {
    echo "Running Go application..."
    go run "$GO_BINARY"
}

# Build the Go application
build_app() {
    echo "Building Go application..."
    go build -o ngodeyuk "$GO_BINARY"
}

# Clean up build artifacts
clean_artifacts() {
    echo "Cleaning up build artifacts..."
    rm -f ngodeyuk
}

# Handle database migrations
handle_migrations() {
    case "$1" in
        create)
            if [ -z "$2" ]; then
                echo "Error: Migration name is required."
                usage
            fi
            migrate create -ext sql -dir "$MIGRATE_PATH" "$2"
            ;;
        up)
            migrate -path "$MIGRATE_PATH" -database "$DATABASE_URL" up
            ;;
        down)
            migrate -path "$MIGRATE_PATH" -database "$DATABASE_URL" down
            ;;
        *)
            usage
            ;;
    esac
}

# Main script logic
main() {
    if [ $# -lt 1 ]; then
        usage
    fi

    case "$1" in
        run)
            load_env
            run_app
            ;;
        build)
            load_env
            build_app
            ;;
        clean)
            clean_artifacts
            ;;
        create|up|down)
            load_env
            check_migrate
            handle_migrations "$@"
            ;;
        -h|--help)
            usage
            ;;
        *)
            usage
            ;;
    esac
}

main "$@"

