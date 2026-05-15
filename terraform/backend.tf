terraform {
  backend "s3" {
    bucket  = "terraform-state-536726971394-us-east-2"
    key     = "slack-daily-history/terraform.tfstate"
    region  = "us-east-2"
    encrypt = true
  }
}
