#!/bin/bash

# Check if Go is already installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing now..."
    # Install Go
    sudo apt update
    sudo apt install -y golang-go
else
    echo "Go is already installed."
fi

# Configure Go environment
export GOPATH="$HOME/go"
export PATH="$GOPATH/bin:$PATH"

# Install necessary libraries using go get or go mod
echo "Installing dependencies..."
go get -u github.com/gocolly/colly/v2
go get -u github.com/xuri/excelize/v2
go get go.mongodb.org/mongo-driver/mongo
go get github.com/dgrijalva/jwt-go


# Build your tool or install additional dependencies if needed
# go build or go install

echo "Installation complete."
