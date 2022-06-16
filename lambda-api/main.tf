terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}


## Domain Name

resource "aws_acm_certificate" "tfer--c1cbfa57-36d2-423a-b815-fdb8585ba629_earthly-tools-002E-com" {
  domain_name = "earthly-tools.com"

  options {
    certificate_transparency_logging_preference = "ENABLED"
  }

  subject_alternative_names = ["earthly-tools.com"]
  validation_method         = "DNS"
}

## ECR
resource "aws_ecr_repository_policy" "tfer--lambda-api" {
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

  repository = "lambda-api"
}

resource "aws_ecr_repository" "tfer--lambda-api" {
  encryption_configuration {
    encryption_type = "AES256"
  }

  image_scanning_configuration {
    scan_on_push = "false"
  }

  image_tag_mutability = "MUTABLE"
  name                 = "lambda-api"
}

# ## S3 
# resource "aws_s3_bucket" "tfer--text-mode" {
#   arn           = "arn:aws:s3:::text-mode"
#   bucket        = "text-mode"
#   force_destroy = "false"

#   grant {
#     id          = "cf64d8f33542c10c2190b0f9f3f651508d2fc7b8896fe916df79fac18f417c63"
#     permissions = ["FULL_CONTROL"]
#     type        = "CanonicalUser"
#   }

#   hosted_zone_id = "Z3AQBSTGFYJSTF"

#   lifecycle_rule {
#     abort_incomplete_multipart_upload_days = "0"
#     enabled                                = "true"

#     expiration {
#       days                         = "14"
#       expired_object_delete_marker = "false"
#     }

#     id = "delete_files_after_14_days"
#   }

#   object_lock_enabled = "false"
#   request_payer       = "BucketOwner"

#   versioning {
#     enabled    = "false"
#     mfa_delete = "false"
#   }
# }


# ## Lambda 

# resource "aws_lambda_function" "tfer--lambda-api" {
#   architectures = ["x86_64"]

#   environment {
#     variables = {
#       HOME = "/tmp"
#     }
#   }

#   ephemeral_storage {
#     size = "512"
#   }

#   function_name                  = "lambda-api2"
#   image_uri                      = "459018586415.dkr.ecr.us-east-1.amazonaws.com/lambda-api:latest"
#   memory_size                    = "500"
#   package_type                   = "Image"
#   reserved_concurrent_executions = "-1"
#   role                           = "arn:aws:iam::459018586415:role/service-role/lambda-api-role-hb6fczbh"
#   timeout                        = "120"

#   tracing_config {
#     mode = "PassThrough"
#   }
# }

