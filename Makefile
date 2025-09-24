# Target default ketika jalankan `make`
default: setup

# Setup semua dependency
setup: install-abci install-tendermint install-dep

# Install ABCI CLI (opsional, kalau kamu pakai)
install-abci:
	@echo "📦 Installing abci..."
	go install github.com/tendermint/tendermint/abci/cmd/abci-cli@latest

# Install Tendermint Core
install-tendermint:
	@echo "📦 Installing Tendermint..."
	go install github.com/tendermint/tendermint/cmd/tendermint@latest

# Install dependency Go
install-dep:
	@echo "📦 Installing Go dependencies..."
	go mod tidy

# Menjalankan node + aplikasi bersamaan
run-full:
	@echo "🚀 Starting ABCI app and Tendermint node..."
	go run main.go & tendermint node


# Menjalankan Tendermint node
run-node:
	@echo "🚀 Running Tendermint node..."
	tendermint node

# Menjalankan aplikasi ABCI
run-app:
	@echo "🚀 Running ABCI app..."
	go run main.go
