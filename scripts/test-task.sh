#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Running ECS task manually for testing...${NC}"

AWS_PROFILE="pbloc"
AWS_REGION="us-east-2"
CLUSTER_NAME="slack-daily-history"
TASK_DEFINITION="slack-daily-history"

# Get VPC and subnet information
echo -e "\n${YELLOW}Getting VPC information...${NC}"
VPC_ID=$(aws ec2 describe-vpcs --filters "Name=isDefault,Values=true" --profile $AWS_PROFILE --region $AWS_REGION --query 'Vpcs[0].VpcId' --output text)
SUBNET_IDS=$(aws ec2 describe-subnets --filters "Name=vpc-id,Values=$VPC_ID" --profile $AWS_PROFILE --region $AWS_REGION --query 'Subnets[0:2].SubnetId' --output text | tr '\t' ',')
SECURITY_GROUP=$(aws ec2 describe-security-groups --filters "Name=vpc-id,Values=$VPC_ID" "Name=group-name,Values=slack-daily-history-ecs-task-sg" --profile $AWS_PROFILE --region $AWS_REGION --query 'SecurityGroups[0].GroupId' --output text)

echo "VPC ID: $VPC_ID"
echo "Subnet IDs: $SUBNET_IDS"
echo "Security Group: $SECURITY_GROUP"

# Run task
echo -e "\n${YELLOW}Running task...${NC}"
TASK_ARN=$(aws ecs run-task \
    --cluster $CLUSTER_NAME \
    --task-definition $TASK_DEFINITION \
    --launch-type FARGATE \
    --network-configuration "awsvpcConfiguration={subnets=[$SUBNET_IDS],securityGroups=[$SECURITY_GROUP],assignPublicIp=ENABLED}" \
    --profile $AWS_PROFILE \
    --region $AWS_REGION \
    --query 'tasks[0].taskArn' \
    --output text)

echo -e "${GREEN}Task started!${NC}"
echo "Task ARN: $TASK_ARN"
echo -e "\nTo view logs:"
echo "  aws logs tail /ecs/slack-daily-history --follow --profile $AWS_PROFILE --region $AWS_REGION"
echo -e "\nTo check task status:"
echo "  aws ecs describe-tasks --cluster $CLUSTER_NAME --tasks $TASK_ARN --profile $AWS_PROFILE --region $AWS_REGION"
