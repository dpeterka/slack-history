# CloudWatch Log Group for ECS task logs
resource "aws_cloudwatch_log_group" "slack_history_bot" {
  name              = "/ecs/${var.project_name}"
  retention_in_days = 7

  tags = {
    Name = "${var.project_name}-logs"
  }
}

# ECS Cluster
resource "aws_ecs_cluster" "main" {
  name = var.project_name

  setting {
    name  = "containerInsights"
    value = "enabled"
  }

  tags = {
    Name = var.project_name
  }
}

# ECS Task Definition
resource "aws_ecs_task_definition" "slack_history_bot" {
  family                   = var.project_name
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.task_cpu
  memory                   = var.task_memory
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn

  container_definitions = jsonencode([
    {
      name      = var.project_name
      image     = "${aws_ecr_repository.slack_history_bot.repository_url}:latest"
      essential = true

      environment = [
        {
          name  = "CLAUDE_MODEL"
          value = "claude-sonnet-4-5"
        },
        {
          name  = "RSS_FEED_URLS"
          value = "https://www.onthisday.com/rss/today-in-history.xml,https://unbelievablefactsblog.com/rss,http://feeds.feedburner.com/FutilityCloset,https://www.kickassfacts.com/feed/,https://www.mentalfloss.com/feed"
        },
        {
          name  = "HOLIDAY_FEED_URL"
          value = "https://api.checkiday.com/rss?tz=America/New_York"
        },
        {
          name  = "SCHEDULE_CRON"
          value = "0 10 * * *"
        },
        {
          name  = "RUN_ONCE"
          value = "true"
        },
        {
          name  = "MAX_EVENTS"
          value = "1"
        },
        {
          name  = "MAX_HOLIDAYS"
          value = "2"
        },
        {
          name  = "INCLUDE_QUOTE"
          value = "true"
        },
        {
          name  = "INCLUDE_PEOPLE"
          value = "true"
        },
        {
          name  = "INCLUDE_EMO_COMMENT"
          value = "true"
        },
        {
          name  = "INCLUDE_BLOBBY_FACT"
          value = "true"
        },
        {
          name  = "INCLUDE_WIKIHOW"
          value = "true"
        },
        {
          name  = "INCLUDE_WIKIHOW_QUIZZES"
          value = "true"
        },
        {
          name  = "INCLUDE_HOTTUB"
          value = "true"
        },
        {
          name  = "INCLUDE_GARDENING"
          value = "true"
        },
        {
          name  = "INCLUDE_PRINTING3D"
          value = "true"
        },
        {
          name  = "INCLUDE_EVENTS"
          value = "true"
        },
        {
          name  = "MAX_PEOPLE"
          value = "2"
        },
        {
          name  = "CACHE_DIR"
          value = ".cache"
        },
        {
          name  = "CONTENT_ROTATION_WEEKS"
          value = "6"
        },
        {
          name  = "SKIP_INITIAL_RUN"
          value = "false"
        },
        {
          name  = "TZ"
          value = "America/New_York"
        }
      ]

      secrets = [
        {
          name      = "SLACK_WEBHOOK_URL"
          valueFrom = aws_ssm_parameter.slack_webhook_url.arn
        },
        {
          name      = "CLAUDE_API_KEY"
          valueFrom = aws_ssm_parameter.claude_api_key.arn
        },
        {
          name      = "GIPHY_API_KEY"
          valueFrom = aws_ssm_parameter.giphy_api_key.arn
        }
      ]

      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.slack_history_bot.name
          "awslogs-region"        = var.aws_region
          "awslogs-stream-prefix" = "ecs"
        }
      }
    }
  ])

  tags = {
    Name = var.project_name
  }
}
