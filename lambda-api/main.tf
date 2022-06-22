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

resource "aws_acm_certificate" "earthly-tools-com" {
    domain_name               = "earthly-tools.com"
    subject_alternative_names = [
        "earthly-tools.com",
    ]

    options {
        certificate_transparency_logging_preference = "ENABLED"
    }
}

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


## API GATEWAY

resource "aws_api_gateway_domain_name" "earthly-tools-com" {
  domain_name              = "earthly-tools.com"
  regional_certificate_arn = aws_acm_certificate.earthly-tools-com.arn

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_apigatewayv2_api" "earthly-tools-com" {
   api_key_selection_expression = "$request.header.x-api-key"
    description                  = "Created by AWS Lambda"
    disable_execute_api_endpoint = true
    name                         = "text-mode-API"
    protocol_type                = "HTTP"
    route_selection_expression   = "$request.method $request.path"
}

resource "aws_apigatewayv2_stage" "earthly-tools-com" {
  api_id = aws_apigatewayv2_api.earthly-tools-com.id
  name   = "default"
  auto_deploy = true
}

resource "aws_apigatewayv2_api_mapping" "earthly-tools-com" {
  api_id      = aws_apigatewayv2_api.earthly-tools-com.id
  domain_name = aws_api_gateway_domain_name.earthly-tools-com.domain_name
  stage       = aws_apigatewayv2_stage.earthly-tools-com.id
}


## Lambda 

resource "aws_lambda_function" "lambda-api" {
  function_name                  = "lambda-api"
  image_uri                      = "459018586415.dkr.ecr.us-east-1.amazonaws.com/lambda-api:latest"
  memory_size                    = "500"
  package_type                   = "Image"
  reserved_concurrent_executions = "-1"
  role                           = "arn:aws:iam::459018586415:role/service-role/lambda-api-role-hb6fczbh"
  timeout                        = "120"
  architectures = ["x86_64"]

  environment {
    variables = {
      HOME = "/tmp"
    }
  }

  ephemeral_storage {
    size = "512"
  }

  tracing_config {
    mode = "PassThrough"
  }
}

## Attach Lambda to API Gateway
resource "aws_apigatewayv2_integration" "earthly-tools-com" {
   api_id = aws_apigatewayv2_api.earthly-tools-com.id
   connection_type        = "INTERNET"
    integration_method     = "POST"
    integration_type       = "AWS_PROXY"
    integration_uri        = aws_lambda_function.lambda-api.arn
    payload_format_version = "2.0"
    request_parameters     = {}
    request_templates      = {}
    timeout_milliseconds   = 30000
}

resource "aws_apigatewayv2_route" "earthly-tools-com" {
  api_id = aws_apigatewayv2_api.earthly-tools-com.id
  route_key            = "ANY /{path+}"
  target               = "integrations/${aws_apigatewayv2_integration.earthly-tools-com.id}"
  api_key_required     = false
  authorization_scopes = []
  authorization_type   = "NONE"
  request_models       = {}
}

## Give API Gateway access to Lambda
resource "aws_lambda_permission" "earthly-tools-com" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda-api.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:us-east-1:459018586415:yr255kt190/*/*/{path+}"
}