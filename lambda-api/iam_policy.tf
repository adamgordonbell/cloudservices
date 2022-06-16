# resource "aws_iam_policy" "tfer--AWSLambdaBasicExecutionRole-604dedf0-83f1-4b74-a1fe-45e05a13292e" {
#   name = "AWSLambdaBasicExecutionRole-604dedf0-83f1-4b74-a1fe-45e05a13292e"
#   path = "/service-role/"

#   policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": "logs:CreateLogGroup",
#       "Effect": "Allow",
#       "Resource": "arn:aws:logs:us-east-1:459018586415:*"
#     },
#     {
#       "Action": [
#         "logs:CreateLogStream",
#         "logs:PutLogEvents"
#       ],
#       "Effect": "Allow",
#       "Resource": [
#         "arn:aws:logs:us-east-1:459018586415:log-group:/aws/lambda/lambda-api:*"
#       ]
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY
# }

# resource "aws_iam_policy" "tfer--AWSLambdaBasicExecutionRole-937470c8-4aac-4db2-9a22-8be27d98acb4" {
#   name = "AWSLambdaBasicExecutionRole-937470c8-4aac-4db2-9a22-8be27d98acb4"
#   path = "/service-role/"

#   policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": "logs:CreateLogGroup",
#       "Effect": "Allow",
#       "Resource": "arn:aws:logs:us-east-1:459018586415:*"
#     },
#     {
#       "Action": [
#         "logs:CreateLogStream",
#         "logs:PutLogEvents"
#       ],
#       "Effect": "Allow",
#       "Resource": [
#         "arn:aws:logs:us-east-1:459018586415:log-group:/aws/lambda/text-mode:*"
#       ]
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY
# }

# resource "aws_iam_policy" "tfer--AWSLambdaBasicExecutionRole-d2b2ac5b-5288-40a3-b284-7ece3b1748e2" {
#   name = "AWSLambdaBasicExecutionRole-d2b2ac5b-5288-40a3-b284-7ece3b1748e2"
#   path = "/service-role/"

#   policy = <<POLICY
# {
#   "Statement": [
#     {
#       "Action": [
#         "s3:ListBucket",
#         "logs:CreateLogGroup"
#       ],
#       "Effect": "Allow",
#       "Resource": [
#         "arn:aws:s3:::text-mode",
#         "arn:aws:s3:::*",
#         "arn:aws:s3:*:459018586415:job/*",
#         "arn:aws:logs:us-east-1:459018586415:*"
#       ],
#       "Sid": "VisualEditor0"
#     },
#     {
#       "Action": [
#         "s3:PutObject",
#         "s3:GetObjectAcl",
#         "s3:GetObject",
#         "logs:CreateLogStream",
#         "s3:GetObjectRetention",
#         "s3:PutObjectRetention",
#         "s3:GetObjectAttributes",
#         "s3:GetObjectTagging",
#         "s3:GetObjectLegalHold",
#         "s3:DeleteObject",
#         "logs:PutLogEvents",
#         "s3:GetObjectVersion"
#       ],
#       "Effect": "Allow",
#       "Resource": [
#         "arn:aws:logs:us-east-1:459018586415:log-group:/aws/lambda/text-mode-go:*",
#         "arn:aws:s3:::*/*"
#       ],
#       "Sid": "VisualEditor1"
#     },
#     {
#       "Action": [
#         "s3:ListStorageLensConfigurations",
#         "s3:ListAccessPointsForObjectLambda",
#         "s3:GetAccessPoint",
#         "s3:GetAccountPublicAccessBlock",
#         "s3:ListAllMyBuckets",
#         "s3:ListAccessPoints",
#         "s3:ListJobs",
#         "s3:PutStorageLensConfiguration",
#         "s3:ListMultiRegionAccessPoints",
#         "s3:CreateJob"
#       ],
#       "Effect": "Allow",
#       "Resource": "*",
#       "Sid": "VisualEditor2"
#     }
#   ],
#   "Version": "2012-10-17"
# }
# POLICY
# }
