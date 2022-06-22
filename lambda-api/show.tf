# # aws_acm_certificate.earthly-tools-com:
# resource "aws_acm_certificate" "earthly-tools-com" {
#     arn                       = "arn:aws:acm:us-east-1:459018586415:certificate/c1cbfa57-36d2-423a-b815-fdb8585ba629"
#     domain_name               = "earthly-tools.com"
#     domain_validation_options = [
#         {
#             domain_name           = "earthly-tools.com"
#             resource_record_name  = "_12d1c97843134b6f6600cf18899b5ec5.earthly-tools.com."
#             resource_record_type  = "CNAME"
#             resource_record_value = "_f8939ef8e157a027322128878499a5e8.mntkzmhvxg.acm-validations.aws."
#         },
#     ]
#     id                        = "arn:aws:acm:us-east-1:459018586415:certificate/c1cbfa57-36d2-423a-b815-fdb8585ba629"
#     status                    = "ISSUED"
#     subject_alternative_names = [
#         "earthly-tools.com",
#     ]
#     tags                      = {}
#     tags_all                  = {}
#     validation_emails         = []
#     validation_method         = "DNS"

#     options {
#         certificate_transparency_logging_preference = "ENABLED"
#     }
# }

# # aws_api_gateway_domain_name.earthly-tools-com:
# resource "aws_api_gateway_domain_name" "earthly-tools-com" {
#     arn                      = "arn:aws:apigateway:us-east-1::/domainnames/earthly-tools.com"
#     certificate_upload_date  = "2022-06-21T20:06:44Z"
#     cloudfront_zone_id       = "Z2FDTNDATAQYW2"
#     domain_name              = "earthly-tools.com"
#     id                       = "earthly-tools.com"
#     regional_certificate_arn = "arn:aws:acm:us-east-1:459018586415:certificate/c1cbfa57-36d2-423a-b815-fdb8585ba629"
#     regional_domain_name     = "d-pcjjelkqsh.execute-api.us-east-1.amazonaws.com"
#     regional_zone_id         = "Z1UJRXOUMOOFQ8"
#     security_policy          = "TLS_1_2"
#     tags                     = {}
#     tags_all                 = {}

#     endpoint_configuration {
#         types = [
#             "REGIONAL",
#         ]
#     }
# }

# # aws_apigatewayv2_api.earthly-tools-com:
# resource "aws_apigatewayv2_api" "earthly-tools-com" {
#     api_endpoint                 = "https://yr255kt190.execute-api.us-east-1.amazonaws.com"
#     api_key_selection_expression = "$request.header.x-api-key"
#     arn                          = "arn:aws:apigateway:us-east-1::/apis/yr255kt190"
#     description                  = "Created by AWS Lambda"
#     disable_execute_api_endpoint = true
#     execution_arn                = "arn:aws:execute-api:us-east-1:459018586415:yr255kt190"
#     id                           = "yr255kt190"
#     name                         = "text-mode-API"
#     protocol_type                = "HTTP"
#     route_selection_expression   = "$request.method $request.path"
#     tags                         = {}
#     tags_all                     = {}
# }

# # aws_apigatewayv2_api_mapping.earthly-tools-com:
# resource "aws_apigatewayv2_api_mapping" "earthly-tools-com" {
#     api_id      = "yr255kt190"
#     domain_name = "earthly-tools.com"
#     id          = "a09jn5"
#     stage       = "default"
# }

# # aws_apigatewayv2_integration.earthly-tools-com2:
# resource "aws_apigatewayv2_integration" "earthly-tools-com2" {
#     api_id                 = "yr255kt190"
#     connection_type        = "INTERNET"
#     id                     = "9ze0cc0"
#     integration_method     = "POST"
#     integration_type       = "AWS_PROXY"
#     integration_uri        = "arn:aws:lambda:us-east-1:459018586415:function:lambda-api"
#     payload_format_version = "2.0"
#     request_parameters     = {}
#     request_templates      = {}
#     timeout_milliseconds   = 30000
# }

# # aws_apigatewayv2_route.earthly-tools-com2:
# resource "aws_apigatewayv2_route" "earthly-tools-com2" {
#     api_id               = "yr255kt190"
#     api_key_required     = false
#     authorization_scopes = []
#     authorization_type   = "NONE"
#     id                   = "ehno9tf"
#     request_models       = {}
#     route_key            = "ANY /{path+}"
#     target               = "integrations/9ze0cc0"
# }

# # aws_apigatewayv2_stage.earthly-tools-com:
# resource "aws_apigatewayv2_stage" "earthly-tools-com" {
#     api_id          = "yr255kt190"
#     arn             = "arn:aws:apigateway:us-east-1::/apis/yr255kt190/stages/default"
#     auto_deploy     = true
#     deployment_id   = "97u9ip"
#     execution_arn   = "arn:aws:execute-api:us-east-1:459018586415:yr255kt190/default"
#     id              = "default"
#     invoke_url      = "https://yr255kt190.execute-api.us-east-1.amazonaws.com/default"
#     name            = "default"
#     stage_variables = {}
#     tags            = {}
#     tags_all        = {}

#     default_route_settings {
#         data_trace_enabled       = false
#         detailed_metrics_enabled = false
#         throttling_burst_limit   = 0
#         throttling_rate_limit    = 0
#     }
# }

# # aws_ecr_repository.lambda-api:
# resource "aws_ecr_repository" "lambda-api" {
#     arn                  = "arn:aws:ecr:us-east-1:459018586415:repository/lambda-api"
#     id                   = "lambda-api"
#     image_tag_mutability = "MUTABLE"
#     name                 = "lambda-api"
#     registry_id          = "459018586415"
#     repository_url       = "459018586415.dkr.ecr.us-east-1.amazonaws.com/lambda-api"
#     tags                 = {}
#     tags_all             = {}

#     encryption_configuration {
#         encryption_type = "AES256"
#     }

#     image_scanning_configuration {
#         scan_on_push = false
#     }

#     timeouts {}
# }

# # aws_ecr_repository_policy.lambda-api:
# resource "aws_ecr_repository_policy" "lambda-api" {
#     id          = "lambda-api"
#     policy      = jsonencode(
#         {
#             Statement = [
#                 {
#                     Action    = [
#                         "ecr:BatchGetImage",
#                         "ecr:GetDownloadUrlForLayer",
#                         "ecr:SetRepositoryPolicy",
#                         "ecr:DeleteRepositoryPolicy",
#                         "ecr:GetRepositoryPolicy",
#                     ]
#                     Condition = {
#                         StringLike = {
#                             "aws:sourceArn" = "arn:aws:lambda:us-east-1:459018586415:function:*"
#                         }
#                     }
#                     Effect    = "Allow"
#                     Principal = {
#                         Service = "lambda.amazonaws.com"
#                     }
#                     Sid       = "LambdaECRImageRetrievalPolicy"
#                 },
#             ]
#             Version   = "2008-10-17"
#         }
#     )
#     registry_id = "459018586415"
#     repository  = "lambda-api"
# }

# # aws_lambda_function.lambda-api:
# resource "aws_lambda_function" "lambda-api" {
#     architectures                  = [
#         "x86_64",
#     ]
#     arn                            = "arn:aws:lambda:us-east-1:459018586415:function:lambda-api"
#     function_name                  = "lambda-api"
#     id                             = "lambda-api"
#     image_uri                      = "459018586415.dkr.ecr.us-east-1.amazonaws.com/lambda-api:latest"
#     invoke_arn                     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:459018586415:function:lambda-api/invocations"
#     last_modified                  = "2022-06-22T15:40:35.444+0000"
#     layers                         = []
#     memory_size                    = 500
#     package_type                   = "Image"
#     publish                        = false
#     qualified_arn                  = "arn:aws:lambda:us-east-1:459018586415:function:lambda-api:$LATEST"
#     reserved_concurrent_executions = -1
#     role                           = "arn:aws:iam::459018586415:role/service-role/lambda-api-role-hb6fczbh"
#     source_code_hash               = "a7c7b12ba3e328b54f6dda759adaea7bb30e3f9f3b2927c3a6064e9b55f3200a"
#     source_code_size               = 0
#     tags                           = {}
#     tags_all                       = {}
#     timeout                        = 120
#     version                        = "$LATEST"

#     environment {
#         variables = {
#             "HOME" = "/tmp"
#         }
#     }

#     ephemeral_storage {
#         size = 512
#     }

#     tracing_config {
#         mode = "PassThrough"
#     }
# }

# # aws_s3_bucket.text-mode:
# resource "aws_s3_bucket" "text-mode" {
#     arn                         = "arn:aws:s3:::text-mode"
#     bucket                      = "text-mode"
#     bucket_domain_name          = "text-mode.s3.amazonaws.com"
#     bucket_regional_domain_name = "text-mode.s3.amazonaws.com"
#     force_destroy               = false
#     hosted_zone_id              = "Z3AQBSTGFYJSTF"
#     id                          = "text-mode"
#     object_lock_enabled         = false
#     region                      = "us-east-1"
#     request_payer               = "BucketOwner"
#     tags                        = {}
#     tags_all                    = {}

#     grant {
#         id          = "cf64d8f33542c10c2190b0f9f3f651508d2fc7b8896fe916df79fac18f417c63"
#         permissions = [
#             "FULL_CONTROL",
#         ]
#         type        = "CanonicalUser"
#     }

#     lifecycle_rule {
#         abort_incomplete_multipart_upload_days = 0
#         enabled                                = true
#         id                                     = "delete_files_after_14_days"
#         tags                                   = {}

#         expiration {
#             days                         = 14
#             expired_object_delete_marker = false
#         }
#     }

#     versioning {
#         enabled    = false
#         mfa_delete = false
#     }
# }

# # aws_s3_bucket_lifecycle_configuration.text-mode:
# resource "aws_s3_bucket_lifecycle_configuration" "text-mode" {
#     bucket = "text-mode"
#     id     = "text-mode"

#     rule {
#         id     = "delete_files_after_14_days"
#         status = "Enabled"

#         expiration {
#             days                         = 14
#             expired_object_delete_marker = false
#         }

#         filter {
#         }
#     }
# }

# # data.aws_canonical_user_id.current:
# data "aws_canonical_user_id" "current" {
#     display_name = "tiffany"
#     id           = "cf64d8f33542c10c2190b0f9f3f651508d2fc7b8896fe916df79fac18f417c63"
# }
