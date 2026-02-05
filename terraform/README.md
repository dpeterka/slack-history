# Slack Daily History Bot - AWS Infrastructure

This Terraform configuration deploys the Slack Daily History Bot to AWS using ECS Fargate with EventBridge Scheduler.

## Architecture

- **ECR Repository**: Stores Docker images
- **ECS Cluster**: Runs the containerized bot on Fargate
- **EventBridge Scheduler**: Triggers the task at 9:58 AM EST on weekdays
- **SSM Parameter Store**: Securely stores secrets (Slack webhook, Claude API key, Giphy API key)
- **CloudWatch Logs**: Captures container logs
- **VPC**: Uses default VPC with security group allowing only outbound traffic

## Prerequisites

1. AWS CLI configured with profile `pbloc`
2. Terraform >= 1.0
3. Docker
4. Valid Slack webhook URL
5. Anthropic Claude API key
6. Giphy API key

## Initial Setup

1. **Create terraform.tfvars file:**
   ```bash
   cd terraform
   cp terraform.tfvars.example terraform.tfvars
   ```

2. **Edit terraform.tfvars with your values:**
   ```hcl
   slack_webhook_url = "https://hooks.slack.com/services/YOUR/WEBHOOK/URL"
   claude_api_key    = "sk-ant-api03-YOUR_API_KEY"
   giphy_api_key     = "YOUR_GIPHY_API_KEY"
   ```

3. **Initialize Terraform:**
   ```bash
   terraform init
   ```

4. **Review the plan:**
   ```bash
   terraform plan
   ```

5. **Apply the configuration:**
   ```bash
   terraform apply
   ```

6. **Note the outputs:**
   After apply completes, note the ECR repository URL for the next step.

## Deploying the Application

### Option 1: Using the deployment script (Recommended)

```bash
cd /home/dpeterka/src/slack-daily-history
chmod +x scripts/deploy.sh
./scripts/deploy.sh
```

This script will:
1. Build the Docker image
2. Tag it for ECR
3. Push to ECR
4. Update the ECS task definition

### Option 2: Manual deployment

```bash
# Get ECR repository URL from Terraform output
ECR_REPO=$(cd terraform && terraform output -raw ecr_repository_url)

# Build and tag the Docker image
docker build -t slack-daily-history:latest .
docker tag slack-daily-history:latest $ECR_REPO:latest

# Login to ECR
aws ecr get-login-password --region us-east-2 --profile pbloc | \
    docker login --username AWS --password-stdin $(echo $ECR_REPO | cut -d'/' -f1)

# Push to ECR
docker push $ECR_REPO:latest
```

## Testing

### Test the task manually

```bash
cd /home/dpeterka/src/slack-daily-history
chmod +x scripts/test-task.sh
./scripts/test-task.sh
```

### View logs

```bash
aws logs tail /ecs/slack-daily-history --follow --profile pbloc --region us-east-2
```

### Check scheduler status

```bash
aws scheduler get-schedule --name slack-daily-history --profile pbloc --region us-east-2
```

## Schedule Configuration

The bot is scheduled to run at **9:58 AM EST on weekdays (Monday-Friday)**.

The schedule uses EventBridge Scheduler with timezone support:
- Schedule expression: `cron(58 9 ? * MON-FRI *)`
- Timezone: `America/New_York`

This automatically handles daylight saving time transitions.

## Updating the Application

When you make changes to the bot code:

```bash
# Update infrastructure (if needed)
cd terraform
terraform apply

# Deploy new container image
cd /home/dpeterka/src/slack-daily-history
./scripts/deploy.sh
```

The next scheduled run will automatically use the new image.

## Updating Secrets

To update secrets (Slack webhook, API keys):

```bash
cd terraform
# Edit terraform.tfvars with new values
terraform apply
```

This will update the SSM parameters. The next task run will use the new values.

## Monitoring

### CloudWatch Logs
```bash
aws logs tail /ecs/slack-daily-history --follow --profile pbloc --region us-east-2
```

### ECS Task History
```bash
aws ecs list-tasks --cluster slack-daily-history --profile pbloc --region us-east-2
```

### Scheduler History
```bash
aws scheduler get-schedule --name slack-daily-history --profile pbloc --region us-east-2
```

## Costs

Estimated monthly costs:
- **ECS Fargate**: ~$0.50/month (assuming 1 minute runtime per day)
  - 22 weekdays × 1 minute × $0.04048/hour ÷ 60 = ~$0.015
- **CloudWatch Logs**: ~$0.50/month (7 day retention)
- **ECR Storage**: ~$0.10/month (assuming <1GB)
- **SSM Parameters**: Free (Standard tier)
- **EventBridge Scheduler**: Free (included in free tier)

**Total: ~$1-2/month**

## Troubleshooting

### Task fails to start
- Check ECS task logs in CloudWatch
- Verify ECR image exists and is accessible
- Verify IAM roles have correct permissions

### Secrets not loading
- Verify SSM parameters exist: `aws ssm get-parameter --name /slack-daily-history/slack-webhook-url --profile pbloc --region us-east-2`
- Check ECS task execution role has SSM read permissions

### Schedule not triggering
- Check EventBridge Scheduler status: `aws scheduler get-schedule --name slack-daily-history --profile pbloc --region us-east-2`
- Verify scheduler IAM role has permissions to run ECS tasks

### Container networking issues
- Verify security group allows outbound traffic
- Verify subnets have internet gateway access
- Verify task has public IP assigned

## Cleanup

To destroy all resources:

```bash
cd terraform
terraform destroy
```

Note: This will delete all infrastructure but will not delete CloudWatch logs immediately (they will expire after 7 days).

## Files

- `providers.tf` - AWS provider and version configuration
- `backend.tf` - S3 backend configuration
- `variables.tf` - Input variables
- `terraform.tfvars` - Variable values (gitignored, create from example)
- `ecr.tf` - ECR repository
- `vpc.tf` - VPC and security group configuration
- `iam.tf` - IAM roles for ECS and EventBridge
- `ssm.tf` - SSM parameters for secrets
- `ecs.tf` - ECS cluster and task definition
- `scheduler.tf` - EventBridge scheduler
- `outputs.tf` - Output values
