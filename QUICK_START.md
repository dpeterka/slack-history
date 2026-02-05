# Quick Start Guide - AWS Deployment

## TL;DR

```bash
# 1. Configure secrets
cd terraform
cp terraform.tfvars.example terraform.tfvars
# Edit terraform.tfvars with your Slack webhook, Claude API key, and Giphy API key

# 2. Deploy infrastructure
terraform init
terraform apply

# 3. Build and push container
cd ..
./scripts/deploy.sh

# 4. Test
./scripts/test-task.sh
aws logs tail /ecs/slack-daily-history --follow --profile pbloc
```

Done! The bot will run automatically at 9:58 AM EST on weekdays.

## Essential Commands

### Deploy Infrastructure
```bash
cd terraform
terraform init
terraform apply
```

### Build & Deploy Container
```bash
./scripts/deploy.sh
```

### Test Manually
```bash
./scripts/test-task.sh
```

### View Logs
```bash
aws logs tail /ecs/slack-daily-history --follow --profile pbloc --region us-east-2
```

### Update Code
```bash
# Make your changes, then:
./scripts/deploy.sh
```

### Update Secrets
```bash
cd terraform
# Edit terraform.tfvars
terraform apply
```

### Destroy Everything
```bash
cd terraform
terraform destroy
```

## What Gets Created

- ECR repository for Docker images
- ECS Fargate cluster
- EventBridge schedule (9:58 AM EST, weekdays)
- SSM parameters (secrets)
- CloudWatch log group
- IAM roles & security group

## Configuration Files

- `terraform/terraform.tfvars` - Your secrets and config (NOT committed)
- `.env` - Local development only (NOT used in AWS)
- `Dockerfile` - Container build

## Cost

~$1-2/month (mostly logging and Fargate runtime)

## Schedule

Runs at 9:58 AM EST Monday-Friday
(Automatically handles daylight saving time)

## Full Documentation

- [DEPLOYMENT.md](DEPLOYMENT.md) - Detailed deployment guide
- [terraform/README.md](terraform/README.md) - Infrastructure documentation
- [README.md](README.md) - Bot configuration and features
