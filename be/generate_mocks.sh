#!/bin/bash

# Generate mocks for service interfaces
echo "Generating mocks..."

# User Service
mockgen -source=internal/service/user_service.go -destination=internal/mocks/mock_user_repository.go -package=mocks UserRepository

# Payment Service  
mockgen -source=internal/service/payment_service.go -destination=internal/mocks/mock_payment_repository.go -package=mocks PaymentRepository

# Controller interfaces (from controller files)
mockgen -source=internal/controller/user_controller.go -destination=internal/mocks/mock_user_service.go -package=mocks UserService
mockgen -source=internal/controller/payment_controller.go -destination=internal/mocks/mock_payment_service.go -package=mocks PaymentService
mockgen -source=internal/controller/admin_controller.go -destination=internal/mocks/mock_admin_service.go -package=mocks AdminService

echo "Mocks generated successfully!"