# AWS Deployment Guide

This guide walks you through deploying the Slack Daily History Bot to AWS using ECS Fargate with scheduled execution via EventBridge.

## Architecture Overview

The bot runs as a containerized application on AWS ECS Fargate, triggered by EventBridge Scheduler at 9:58 AM EST on weekdays. The infrastructure includes:

- **ECR Repository**: Stores the Docker images
- **ECS Cluster & Task Definition**: Runs the bot on Fargate (serverless containers)
- **EventBridge Scheduler**: Triggers the task at 9:58 AM EST Monday-Friday
- **SSM Parameter Store**: Securely stores secrets (Slack webhook, Claude API key, Giphy API key)
- **CloudWatch Logs**: Captures container logs for debugging
- **VPC & Security Group**: Uses default VPC with outbound-only traffic

## Prerequisites

Before you begin, ensure you have:

1. **AWS CLI** installed and configured with profile `pbloc`
2. **Terraform** >= 1.0 installed
3. **Docker** installed and running
4. **API Keys and Webhook**:
   - Slack incoming webhook URL
   - Anthropic Claude API key
   - Giphy API key

## Step-by-Step Deployment

### Step 1: Configure Terraform Variables

1. Copy the example variables file:
   ```bash
   cd terraform
   cp terraform.tfvars.example terraform.tfvars
   ```

2. Edit `terraform/terraform.tfvars` with your actual values:
   ```hcl
   # These should already be set correctly
   aws_region  = "us-east-2"
   aws_profile = "pbloc"
   environment = "production"
   project_name = "slack-daily-history"

   # Schedule: 9:58 AM EST on weekdays
   schedule_expression = "cron(58 9 ? * MON-FRI *)"

   # YOUR SECRETS - Replace these!
   slack_webhook_url = "https://hooks.slack.com/services/YOUR/WEBHOOK/URL"
   claude_api_key    = "sk-ant-api03-YOUR_CLAUDE_API_KEY"
   giphy_api_key     = "YOUR_GIPHY_API_KEY"

   # Task sizing (adjust if needed)
   task_cpu    = "256"  # 0.25 vCPU
   task_memory = "512"  # 512 MB
   ```

### Step 2: Deploy Infrastructure

1. Initialize Terraform:
   ```bash
   cd terraform
   terraform init
   ```

2. Review the deployment plan:
   ```bash
   terraform plan
   ```

3. Apply the configuration:
   ```bash
   terraform apply
   ```

   Type `yes` when prompted to confirm.

4. Note the outputs:
   ```bash
   terraform output
   ```

   Save the ECR repository URL - you'll need it for the next step.

### Step 3: Build and Push Docker Image

Use the deployment script for a streamlined process:

```bash
cd /home/dpeterka/src/slack-daily-history
./scripts/deploy.sh
```

This script will:
1. Build the Docker image from the Dockerfile
2. Tag it for ECR
3. Authenticate with ECR
4. Push the image to ECR
5. Update the ECS task definition

**Or** deploy manually:

```bash
# Get ECR repository URL
cd terraform
ECR_REPO=$(terraform output -raw ecr_repository_url)
cd ..

# Build Docker image
docker build -t slack-daily-history:latest .

# Tag for ECR
docker tag slack-daily-history:latest $ECR_REPO:latest

# Login to ECR
aws ecr get-login-password --region us-east-2 --profile pbloc | \
    docker login --username AWS --password-stdin $(echo $ECR_REPO | cut -d'/' -f1)

# Push to ECR
docker push $ECR_REPO:latest
```

### Step 4: Verify Deployment

1. **Test the task manually**:
   ```bash
   ./scripts/test-task.sh
   ```

2. **View logs**:
   ```bash
   aws logs tail /ecs/slack-daily-history --follow --profile pbloc --region us-east-2
   ```

3. **Check scheduler status**:
   ```bash
   aws scheduler get-schedule --name slack-daily-history --profile pbloc --region us-east-2
   ```

4. **Verify in Slack**:
   Check your configured Slack channel for the bot's message.

## Schedule Configuration

The bot runs at **9:58 AM EST on weekdays (Monday-Friday)**.

- Schedule expression: `cron(58 9 ? * MON-FRI *)`
- Timezone: `America/New_York`
- Automatically handles daylight saving time

To change the schedule, update `schedule_expression` in `terraform.tfvars` and run `terraform apply`.

## Updating the Application

### Update Code

When you make changes to the bot code:

```bash
# Test locally first
RUN_ONCE=true go run cmd/bot/main.go

# Deploy to AWS
./scripts/deploy.sh
```

The next scheduled run will use the new image automatically.

### Update Infrastructure

When you change Terraform configuration:

```bash
cd terraform
terraform plan    # Review changes
terraform apply   # Apply changes
```

### Update Secrets

To update API keys or webhook URL:

```bash
cd terraform
# Edit terraform.tfvars with new values
terraform apply
```

The next task run will use the updated secrets.

## Monitoring

### CloudWatch Logs

Real-time log streaming:
```bash
aws logs tail /ecs/slack-daily-history --follow --profile pbloc --region us-east-2
```

View specific time range:
```bash
aws logs tail /ecs/slack-daily-history \
  --since 1h \
  --profile pbloc \
  --region us-east-2
```

### ECS Task History

List recent task runs:
```bash
aws ecs list-tasks \
  --cluster slack-daily-history \
  --profile pbloc \
  --region us-east-2
```

Get task details:
```bash
aws ecs describe-tasks \
  --cluster slack-daily-history \
  --tasks TASK_ARN \
  --profile pbloc \
  --region us-east-2
```

### Scheduler Status

Check schedule configuration:
```bash
aws scheduler get-schedule \
  --name slack-daily-history \
  --profile pbloc \
  --region us-east-2
```

## Troubleshooting

### Task Fails to Start

1. **Check CloudWatch Logs**:
   ```bash
   aws logs tail /ecs/slack-daily-history --follow --profile pbloc --region us-east-2
   ```

2. **Verify ECR image exists**:
   ```bash
   aws ecr describe-images \
     --repository-name slack-daily-history \
     --profile pbloc \
     --region us-east-2
   ```

3. **Check IAM roles**:
   Ensure the ECS task execution role has permissions to:
   - Pull images from ECR
   - Write to CloudWatch Logs
   - Read SSM parameters

### Secrets Not Loading

1. **Verify SSM parameters exist**:
   ```bash
   aws ssm get-parameter \
     --name /slack-daily-history/slack-webhook-url \
     --profile pbloc \
     --region us-east-2
   ```

2. **Check parameter values** (be careful with sensitive data):
   ```bash
   aws ssm get-parameter \
     --name /slack-daily-history/slack-webhook-url \
     --with-decryption \
     --profile pbloc \
     --region us-east-2 \
     --query Parameter.Value \
     --output text
   ```

### Schedule Not Triggering

1. **Check scheduler state**:
   ```bash
   aws scheduler get-schedule \
     --name slack-daily-history \
     --profile pbloc \
     --region us-east-2 \
     --query State \
     --output text
   ```

2. **View scheduler target**:
   ```bash
   aws scheduler get-schedule \
     --name slack-daily-history \
     --profile pbloc \
     --region us-east-2
   ```

3. **Check EventBridge Scheduler IAM role** has permissions to run ECS tasks

### Network Issues

1. **Verify security group** allows outbound traffic:
   ```bash
   aws ec2 describe-security-groups \
     --filters "Name=group-name,Values=slack-daily-history-ecs-task-sg" \
     --profile pbloc \
     --region us-east-2
   ```

2. **Ensure subnets have internet access** (default VPC should have IGW)

3. **Verify task has public IP** (required for Fargate tasks to access internet)

## Cost Estimates

Estimated monthly AWS costs:

| Service | Cost |
|---------|------|
| ECS Fargate | ~$0.50/month |
| CloudWatch Logs (7 day retention) | ~$0.50/month |
| ECR Storage | ~$0.10/month |
| SSM Parameters (Standard) | Free |
| EventBridge Scheduler | Free |
| **Total** | **~$1-2/month** |

Cost breakdown:
- **Fargate**: 22 weekdays × 1 minute × $0.04048/hour ÷ 60 = ~$0.015/month
- Minimal costs for logging and storage
- No NAT Gateway needed (using public subnets)

## Cleanup

To remove all AWS resources:

```bash
cd terraform
terraform destroy
```

Type `yes` when prompted. This will:
- Delete the ECS cluster and task definition
- Delete the EventBridge scheduler
- Delete the ECR repository (and all images)
- Delete SSM parameters
- Delete CloudWatch log groups
- Delete IAM roles and security groups

**Note**: CloudWatch logs are deleted immediately, but log data may be retained for up to 7 days based on retention settings.

## Next Steps

After deployment:

1. **Monitor the first few runs** to ensure everything works correctly
2. **Adjust the schedule** if needed in `terraform.tfvars`
3. **Tune task CPU/memory** if the task is slow or uses too much resources
4. **Set up CloudWatch alarms** (optional) for task failures
5. **Review logs regularly** to catch any issues early

## Support

For issues or questions:
- Check the [main README](README.md) for bot configuration
- Review [Terraform README](terraform/README.md) for infrastructure details
- Check AWS CloudWatch Logs for runtime errors
- Review GitHub Issues for known problems
