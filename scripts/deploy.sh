#!/bin/bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Slack Daily History Bot - Deployment Script${NC}"
echo "=========================================="

# Check if terraform.tfvars exists
if [ ! -f "terraform/terraform.tfvars" ]; then
    echo -e "${RED}Error: terraform/terraform.tfvars not found${NC}"
    echo "Please create it from terraform/terraform.tfvars.example"
    exit 1
fi

# Get AWS account ID
echo -e "\n${YELLOW}Getting AWS account information...${NC}"
AWS_PROFILE="pbloc"
AWS_REGION="us-east-2"
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --profile $AWS_PROFILE --query Account --output text)
echo "AWS Account ID: $AWS_ACCOUNT_ID"
echo "AWS Region: $AWS_REGION"

# Get ECR repository URL
ECR_REPO="${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/slack-daily-history"

# Build Docker image
echo -e "\n${YELLOW}Building Docker image...${NC}"
docker build -t slack-daily-history:latest .

# Tag for ECR
echo -e "\n${YELLOW}Tagging image for ECR...${NC}"
docker tag slack-daily-history:latest $ECR_REPO:latest

# Login to ECR
echo -e "\n${YELLOW}Logging in to ECR...${NC}"
aws ecr get-login-password --region $AWS_REGION --profile $AWS_PROFILE | \
    docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com

# Push to ECR
echo -e "\n${YELLOW}Pushing image to ECR...${NC}"
docker push $ECR_REPO:latest

# Update ECS task definition
echo -e "\n${YELLOW}Updating ECS task definition...${NC}"
cd terraform
terraform init -reconfigure
terraform apply -auto-approve

echo -e "\n${GREEN}Deployment complete!${NC}"
echo -e "Image: ${ECR_REPO}:latest"
echo -e "\nTo view logs:"
echo -e "  aws logs tail /ecs/slack-daily-history --follow --profile $AWS_PROFILE"
