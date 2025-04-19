#!/bin/bash

# Build assets
echo "Building assets..."
npm install
npm run build

# Build Go binary
echo "Building Go binary..."
GOOS=linux GOARCH=amd64 go build -o wedding-app cmd/main.go

# Deploy to VPS (replace these values with your actual VPS details)
VPS_USER="your_user"
VPS_HOST="your_vps_ip"
APP_DIR="/opt/wedding-app"

echo "Deploying to VPS..."
# Create directory structure
ssh $VPS_USER@$VPS_HOST "mkdir -p $APP_DIR/{templates,static}"

# Copy files
scp wedding-app $VPS_USER@$VPS_HOST:$APP_DIR/
scp -r templates/* $VPS_USER@$VPS_HOST:$APP_DIR/templates/
scp -r static/* $VPS_USER@$VPS_HOST:$APP_DIR/static/

# Set permissions and restart service
ssh $VPS_USER@$VPS_HOST "chmod +x $APP_DIR/wedding-app && sudo systemctl restart wedding-app"

echo "Deployment complete!" 