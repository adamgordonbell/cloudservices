# resource "aws_iam_role" "tfer--AWSServiceRoleForAPIGateway" {
#   assume_role_policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": "sts:AssumeRole",
#       "Effect": "Allow",
#       "Principal": {
#         "Service": "ops.apigateway.amazonaws.com"
#       }
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY

#   description          = "The Service Linked Role is used by Amazon API Gateway."
#   managed_policy_arns  = ["arn:aws:iam::aws:policy/aws-service-role/APIGatewayServiceRolePolicy"]
#   max_session_duration = "3600"
#   name                 = "AWSServiceRoleForAPIGateway"
#   path                 = "/aws-service-role/ops.apigateway.amazonaws.com/"
# }

# resource "aws_iam_role" "tfer--AWSServiceRoleForSupport" {
#   assume_role_policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": "sts:AssumeRole",
#       "Effect": "Allow",
#       "Principal": {
#         "Service": "support.amazonaws.com"
#       }
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY

#   description          = "Enables resource access for AWS to provide billing, administrative and support services"
#   managed_policy_arns  = ["arn:aws:iam::aws:policy/aws-service-role/AWSSupportServiceRolePolicy"]
#   max_session_duration = "3600"
#   name                 = "AWSServiceRoleForSupport"
#   path                 = "/aws-service-role/support.amazonaws.com/"
# }

# resource "aws_iam_role" "tfer--AWSServiceRoleForTrustedAdvisor" {
#   assume_role_policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": "sts:AssumeRole",
#       "Effect": "Allow",
#       "Principal": {
#         "Service": "trustedadvisor.amazonaws.com"
#       }
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY

#   description          = "Access for the AWS Trusted Advisor Service to help reduce cost, increase performance, and improve security of your AWS environment."
#   managed_policy_arns  = ["arn:aws:iam::aws:policy/aws-service-role/AWSTrustedAdvisorServiceRolePolicy"]
#   max_session_duration = "3600"
#   name                 = "AWSServiceRoleForTrustedAdvisor"
#   path                 = "/aws-service-role/trustedadvisor.amazonaws.com/"
# }

# resource "aws_iam_role" "tfer--lambda-api-role-hb6fczbh" {
#   assume_role_policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": "sts:AssumeRole",
#       "Effect": "Allow",
#       "Principal": {
#         "Service": "lambda.amazonaws.com"
#       }
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY

#   managed_policy_arns  = ["arn:aws:iam::459018586415:policy/service-role/AWSLambdaBasicExecutionRole-604dedf0-83f1-4b74-a1fe-45e05a13292e", "arn:aws:iam::aws:policy/AmazonS3FullAccess"]
#   max_session_duration = "3600"
#   name                 = "lambda-api-role-hb6fczbh"
#   path                 = "/service-role/"
# }

# resource "aws_iam_role" "tfer--text-mode-go-role-mnce0npc" {
#   assume_role_policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": "sts:AssumeRole",
#       "Effect": "Allow",
#       "Principal": {
#         "Service": "lambda.amazonaws.com"
#       }
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY

#   managed_policy_arns  = ["arn:aws:iam::459018586415:policy/service-role/AWSLambdaBasicExecutionRole-d2b2ac5b-5288-40a3-b284-7ece3b1748e2"]
#   max_session_duration = "3600"
#   name                 = "text-mode-go-role-mnce0npc"
#   path                 = "/service-role/"
# }

# resource "aws_iam_role" "tfer--text-mode-role-236astbr" {
#   assume_role_policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": "sts:AssumeRole",
#       "Effect": "Allow",
#       "Principal": {
#         "Service": "lambda.amazonaws.com"
#       }
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY

#   managed_policy_arns  = ["arn:aws:iam::459018586415:policy/service-role/AWSLambdaBasicExecutionRole-937470c8-4aac-4db2-9a22-8be27d98acb4"]
#   max_session_duration = "3600"
#   name                 = "text-mode-role-236astbr"
#   path                 = "/service-role/"
# }
