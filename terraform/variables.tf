variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-2"
}

variable "aws_profile" {
  description = "AWS CLI profile name"
  type        = string
  default     = "pbloc"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "production"
}

variable "project_name" {
  description = "Project name"
  type        = string
  default     = "slack-daily-history"
}

variable "schedule_expression" {
  description = "EventBridge schedule expression (cron in UTC)"
  type        = string
  # 9:58 AM EST = 14:58 UTC (during standard time)
  # Runs Monday-Friday at 9:58 AM EST
  default = "cron(58 14 ? * MON-FRI *)"
}

variable "slack_webhook_url" {
  description = "Slack webhook URL"
  type        = string
  sensitive   = true
}

variable "claude_api_key" {
  description = "Anthropic Claude API key"
  type        = string
  sensitive   = true
}

variable "giphy_api_key" {
  description = "Giphy API key"
  type        = string
  sensitive   = true
}

variable "task_cpu" {
  description = "CPU units for the ECS task (256 = 0.25 vCPU)"
  type        = string
  default     = "256"
}

variable "task_memory" {
  description = "Memory for the ECS task in MB"
  type        = string
  default     = "512"
}
