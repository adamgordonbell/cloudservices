terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

## Account ID
data "aws_canonical_user_id" "current" {}

# ## Domain Name

# resource "aws_acm_certificate" "tfer--c1cbfa57-36d2-423a-b815-fdb8585ba629_earthly-tools-002E-com" {
#   domain_name = "earthly-tools.com"

#   options {
#     certificate_transparency_logging_preference = "ENABLED"
#   }

#   subject_alternative_names = ["earthly-tools.com"]
#   validation_method         = "DNS"
# }

# ## ECR

resource "aws_ecr_repository" "lambda-api" {
  encryption_configuration {
    encryption_type = "AES256"
  }

  image_scanning_configuration {
    scan_on_push = "false"
  }

  image_tag_mutability = "MUTABLE"
  name                 = "lambda-api"
}

resource "aws_ecr_repository_policy" "lambda-api" {
  policy = <<POLICY
{
  "Statement": [
    {
      "Action": [
        "ecr:BatchGetImage",
        "ecr:GetDownloadUrlForLayer",
        "ecr:SetRepositoryPolicy",
        "ecr:DeleteRepositoryPolicy",
        "ecr:GetRepositoryPolicy"
      ],
      "Condition": {
        "StringLike": {
          "aws:sourceArn": "arn:aws:lambda:us-east-1:459018586415:function:*"
        }
      },
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Sid": "LambdaECRImageRetrievalPolicy"
    }
  ],
  "Version": "2008-10-17"
}
POLICY

  repository = aws_ecr_repository.lambda-api.name
}

## S3 
resource "aws_s3_bucket" "text-mode" {
  arn           = "arn:aws:s3:::text-mode"
  bucket        = "text-mode"
  force_destroy = "false"
  hosted_zone_id = "Z3AQBSTGFYJSTF"
}

resource "aws_s3_bucket_lifecycle_configuration" "text-mode" {
  bucket = aws_s3_bucket.text-mode.id
  rule {
    id     = "delete_files_after_14_days"
    status = "Enabled"

    expiration {
      days = 14
    }
  }
}

### Lambda 

resource "aws_lambda_function" "lambda-api" {
  architectures = ["x86_64"]

  environment {
    variables = {
      HOME = "/tmp"
    }
  }

  ephemeral_storage {
    size = "512"
  }

  function_name                  = "lambda-api"
  image_uri                      = "459018586415.dkr.ecr.us-east-1.amazonaws.com/lambda-api:latest"
  memory_size                    = "500"
  package_type                   = "Image"
  reserved_concurrent_executions = "-1"
  role                           = "arn:aws:iam::459018586415:role/service-role/lambda-api-role-hb6fczbh"
  timeout                        = "120"

  tracing_config {
    mode = "PassThrough"
  }
}

## API GATEWAY

# resource "aws_api_gateway_rest_api" "text-mode-API"{
# #     # arn                          = "arn:aws:apigateway:us-east-1::/apis/fu3kk58jj5"
#     name = "text-mode-API"
# }