# SSM Parameters for sensitive configuration
resource "aws_ssm_parameter" "slack_webhook_url" {
  name        = "/${var.project_name}/slack-webhook-url"
  description = "Slack webhook URL for posting messages"
  type        = "SecureString"
  value       = var.slack_webhook_url
  tier        = "Standard"

  tags = {
    Name = "${var.project_name}-slack-webhook-url"
  }
}

resource "aws_ssm_parameter" "claude_api_key" {
  name        = "/${var.project_name}/claude-api-key"
  description = "Anthropic Claude API key"
  type        = "SecureString"
  value       = var.claude_api_key
  tier        = "Standard"

  tags = {
    Name = "${var.project_name}-claude-api-key"
  }
}

resource "aws_ssm_parameter" "giphy_api_key" {
  name        = "/${var.project_name}/giphy-api-key"
  description = "Giphy API key"
  type        = "SecureString"
  value       = var.giphy_api_key
  tier        = "Standard"

  tags = {
    Name = "${var.project_name}-giphy-api-key"
  }
}
