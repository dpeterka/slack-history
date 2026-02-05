# EventBridge Scheduler to run the ECS task on weekdays at 9:58 AM EST
resource "aws_scheduler_schedule" "slack_history_bot" {
  name       = var.project_name
  group_name = "default"

  flexible_time_window {
    mode = "OFF"
  }

  # Cron expression in UTC (9:58 AM EST = 14:58 UTC during standard time)
  # Format: cron(Minutes Hours Day-of-month Month Day-of-week Year)
  schedule_expression          = var.schedule_expression
  schedule_expression_timezone = "America/New_York"

  description = "Runs the Slack daily history bot on weekdays at 9:58 AM EST"

  target {
    arn      = aws_ecs_cluster.main.arn
    role_arn = aws_iam_role.eventbridge_scheduler_role.arn

    ecs_parameters {
      task_definition_arn = aws_ecs_task_definition.slack_history_bot.arn
      launch_type         = "FARGATE"

      network_configuration {
        subnets          = data.aws_subnets.default.ids
        security_groups  = [aws_security_group.ecs_task.id]
        assign_public_ip = true
      }
    }

    retry_policy {
      maximum_event_age_in_seconds = 3600
      maximum_retry_attempts       = 2
    }
  }
}
