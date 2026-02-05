output "ecr_repository_url" {
  description = "ECR repository URL"
  value       = aws_ecr_repository.slack_history_bot.repository_url
}

output "ecs_cluster_name" {
  description = "ECS cluster name"
  value       = aws_ecs_cluster.main.name
}

output "ecs_task_definition_arn" {
  description = "ECS task definition ARN"
  value       = aws_ecs_task_definition.slack_history_bot.arn
}

output "cloudwatch_log_group" {
  description = "CloudWatch log group name"
  value       = aws_cloudwatch_log_group.slack_history_bot.name
}

output "scheduler_name" {
  description = "EventBridge scheduler name"
  value       = aws_scheduler_schedule.slack_history_bot.name
}

output "scheduler_expression" {
  description = "EventBridge scheduler cron expression"
  value       = aws_scheduler_schedule.slack_history_bot.schedule_expression
}

output "scheduler_timezone" {
  description = "EventBridge scheduler timezone"
  value       = aws_scheduler_schedule.slack_history_bot.schedule_expression_timezone
}
