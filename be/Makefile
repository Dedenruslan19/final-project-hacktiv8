.PHONY: test test-verbose test-coverage generate-mocks clean-mocks help

# Default target
help:
	@echo "Available targets:"
	@echo "  test           - Run all tests"
	@echo "  test-verbose   - Run tests with verbose output"
	@echo "  test-coverage  - Run tests with coverage and generate HTML report"
	@echo "  generate-mocks - Generate all mock files"
	@echo "  clean-mocks    - Remove all generated mock files"
	@echo "  help          - Show this help message"

# Run all tests
test:
	@echo "Running tests..."
	go test ./...

# Run tests with verbose output
test-verbose:
	@echo "Running tests with verbose output..."
	go test -v ./...

# Run tests with coverage and generate HTML report
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Generate all mock files
generate-mocks:
	@echo "Generating mocks..."
	@mkdir -p internal/mocks
	go run github.com/golang/mock/mockgen -source=internal/service/user_service.go -destination=internal/mocks/mock_user_repository.go -package=mocks UserRepository
	go run github.com/golang/mock/mockgen -source=internal/service/payment_service.go -destination=internal/mocks/mock_payment_repository.go -package=mocks PaymentRepository
	go run github.com/golang/mock/mockgen -source=internal/controller/user_controller.go -destination=internal/mocks/mock_user_service.go -package=mocks UserService
	go run github.com/golang/mock/mockgen -source=internal/controller/payment_controller.go -destination=internal/mocks/mock_payment_service.go -package=mocks PaymentService
	go run github.com/golang/mock/mockgen -source=internal/controller/admin_controller.go -destination=internal/mocks/mock_admin_service.go -package=mocks AdminService
	@echo "Mocks generated successfully!"

# Clean generated mock files
clean-mocks:
	@echo "Cleaning mock files..."
	rm -rf internal/mocks/
	@echo "Mock files cleaned!"

# Install test dependencies
install-deps:
	@echo "Installing test dependencies..."
	go get github.com/golang/mock/gomock@latest
	go get github.com/golang/mock/mockgen@latest
	go get github.com/stretchr/testify/assert@latest
	go get github.com/stretchr/testify/mock@latest
	@echo "Dependencies installed!"