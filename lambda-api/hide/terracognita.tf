provider "aws" {
}


terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }

  }

  required_version = ">= 1.0"
}

resource "aws_iam_policy" "arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_604dedf0_83f1_4b74_a1fe_45e05a13292e" {
  name   = "AWSLambdaBasicExecutionRole-604dedf0-83f1-4b74-a1fe-45e05a13292e"
  path   = "/service-role/"
  policy = "{\"Statement\":[{\"Action\":\"logs:CreateLogGroup\",\"Effect\":\"Allow\",\"Resource\":\"arn:aws:logs:us-east-1:459018586415:*\"},{\"Action\":[\"logs:CreateLogStream\",\"logs:PutLogEvents\"],\"Effect\":\"Allow\",\"Resource\":[\"arn:aws:logs:us-east-1:459018586415:log-group:/aws/lambda/lambda-api:*\"]}],\"Version\":\"2012-10-17\"}"
}

resource "aws_iam_policy" "arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_937470c8_4aac_4db2_9a22_8be27d98acb4" {
  name   = "AWSLambdaBasicExecutionRole-937470c8-4aac-4db2-9a22-8be27d98acb4"
  path   = "/service-role/"
  policy = "{\"Statement\":[{\"Action\":\"logs:CreateLogGroup\",\"Effect\":\"Allow\",\"Resource\":\"arn:aws:logs:us-east-1:459018586415:*\"},{\"Action\":[\"logs:CreateLogStream\",\"logs:PutLogEvents\"],\"Effect\":\"Allow\",\"Resource\":[\"arn:aws:logs:us-east-1:459018586415:log-group:/aws/lambda/text-mode:*\"]}],\"Version\":\"2012-10-17\"}"
}

resource "aws_iam_policy" "arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_d2b2ac5b_5288_40a3_b284_7ece3b1748e2" {
  name   = "AWSLambdaBasicExecutionRole-d2b2ac5b-5288-40a3-b284-7ece3b1748e2"
  path   = "/service-role/"
  policy = "{\"Statement\":[{\"Action\":[\"s3:ListBucket\",\"logs:CreateLogGroup\"],\"Effect\":\"Allow\",\"Resource\":[\"arn:aws:s3:::text-mode\",\"arn:aws:s3:::*\",\"arn:aws:s3:*:459018586415:job/*\",\"arn:aws:logs:us-east-1:459018586415:*\"],\"Sid\":\"VisualEditor0\"},{\"Action\":[\"s3:PutObject\",\"s3:GetObjectAcl\",\"s3:GetObject\",\"logs:CreateLogStream\",\"s3:GetObjectRetention\",\"s3:PutObjectRetention\",\"s3:GetObjectAttributes\",\"s3:GetObjectTagging\",\"s3:GetObjectLegalHold\",\"s3:DeleteObject\",\"logs:PutLogEvents\",\"s3:GetObjectVersion\"],\"Effect\":\"Allow\",\"Resource\":[\"arn:aws:logs:us-east-1:459018586415:log-group:/aws/lambda/text-mode-go:*\",\"arn:aws:s3:::*/*\"],\"Sid\":\"VisualEditor1\"},{\"Action\":[\"s3:ListStorageLensConfigurations\",\"s3:ListAccessPointsForObjectLambda\",\"s3:GetAccessPoint\",\"s3:GetAccountPublicAccessBlock\",\"s3:ListAllMyBuckets\",\"s3:ListAccessPoints\",\"s3:ListJobs\",\"s3:PutStorageLensConfiguration\",\"s3:ListMultiRegionAccessPoints\",\"s3:CreateJob\"],\"Effect\":\"Allow\",\"Resource\":\"*\",\"Sid\":\"VisualEditor2\"}],\"Version\":\"2012-10-17\"}"
}

resource "aws_iam_role" "awsserviceroleforapigateway" {
  assume_role_policy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ops.apigateway.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"
  description        = "The Service Linked Role is used by Amazon API Gateway."
  inline_policy {
  }

  managed_policy_arns  = ["arn:aws:iam::aws:policy/aws-service-role/APIGatewayServiceRolePolicy"]
  max_session_duration = 3600
  name                 = "AWSServiceRoleForAPIGateway"
  path                 = "/aws-service-role/ops.apigateway.amazonaws.com/"
}

resource "aws_iam_role" "awsserviceroleforsupport" {
  assume_role_policy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"support.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"
  description        = "Enables resource access for AWS to provide billing, administrative and support services"
  inline_policy {
  }

  managed_policy_arns  = ["arn:aws:iam::aws:policy/aws-service-role/AWSSupportServiceRolePolicy"]
  max_session_duration = 3600
  name                 = "AWSServiceRoleForSupport"
  path                 = "/aws-service-role/support.amazonaws.com/"
}

resource "aws_iam_role" "awsservicerolefortrustedadvisor" {
  assume_role_policy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"trustedadvisor.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"
  description        = "Access for the AWS Trusted Advisor Service to help reduce cost, increase performance, and improve security of your AWS environment."
  inline_policy {
  }

  managed_policy_arns  = ["arn:aws:iam::aws:policy/aws-service-role/AWSTrustedAdvisorServiceRolePolicy"]
  max_session_duration = 3600
  name                 = "AWSServiceRoleForTrustedAdvisor"
  path                 = "/aws-service-role/trustedadvisor.amazonaws.com/"
}

resource "aws_iam_role" "lambda_api_role_hb6fczbh" {
  assume_role_policy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"lambda.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"
  inline_policy {
  }

  managed_policy_arns  = [aws_iam_policy.arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_604dedf0_83f1_4b74_a1fe_45e05a13292e.id, "arn:aws:iam::aws:policy/AmazonS3FullAccess"]
  max_session_duration = 3600
  name                 = "lambda-api-role-hb6fczbh"
  path                 = "/service-role/"
}

resource "aws_iam_role" "text_mode_go_role_mnce0npc" {
  assume_role_policy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"lambda.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"
  inline_policy {
  }

  managed_policy_arns  = [aws_iam_policy.arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_d2b2ac5b_5288_40a3_b284_7ece3b1748e2.id]
  max_session_duration = 3600
  name                 = "text-mode-go-role-mnce0npc"
  path                 = aws_iam_policy.arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_604dedf0_83f1_4b74_a1fe_45e05a13292e.path
}

resource "aws_iam_role" "text_mode_role_236astbr" {
  assume_role_policy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"lambda.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"
  inline_policy {
  }

  managed_policy_arns  = [aws_iam_policy.arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_937470c8_4aac_4db2_9a22_8be27d98acb4.id]
  max_session_duration = 3600
  name                 = "text-mode-role-236astbr"
  path                 = aws_iam_policy.arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_604dedf0_83f1_4b74_a1fe_45e05a13292e.path
}

resource "aws_iam_role_policy_attachment" "awsserviceroleforapigateway_arn_aws_iam__aws_policy_aws_service_role_apigatewayservicerolepolicy" {
  policy_arn = "arn:aws:iam::aws:policy/aws-service-role/APIGatewayServiceRolePolicy"
  role       = aws_iam_role.awsserviceroleforapigateway.id
}

resource "aws_iam_role_policy_attachment" "awsserviceroleforsupport_arn_aws_iam__aws_policy_aws_service_role_awssupportservicerolepolicy" {
  policy_arn = "arn:aws:iam::aws:policy/aws-service-role/AWSSupportServiceRolePolicy"
  role       = aws_iam_role.awsserviceroleforsupport.id
}

resource "aws_iam_role_policy_attachment" "awsservicerolefortrustedadvisor_arn_aws_iam__aws_policy_aws_service_role_awstrustedadvisorservicerolepolicy" {
  policy_arn = "arn:aws:iam::aws:policy/aws-service-role/AWSTrustedAdvisorServiceRolePolicy"
  role       = aws_iam_role.awsservicerolefortrustedadvisor.id
}

resource "aws_iam_role_policy_attachment" "lambda_api_role_hb6fczbh_arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_604dedf0_83f1_4b74_a1fe_45e05a13292e" {
  policy_arn = aws_iam_policy.arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_604dedf0_83f1_4b74_a1fe_45e05a13292e.id
  role       = aws_iam_role.lambda_api_role_hb6fczbh.id
}

resource "aws_iam_role_policy_attachment" "lambda_api_role_hb6fczbh_arn_aws_iam__aws_policy_amazons3fullaccess" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
  role       = aws_iam_role.lambda_api_role_hb6fczbh.id
}

resource "aws_iam_role_policy_attachment" "text_mode_go_role_mnce0npc_arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_d2b2ac5b_5288_40a3_b284_7ece3b1748e2" {
  policy_arn = aws_iam_policy.arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_d2b2ac5b_5288_40a3_b284_7ece3b1748e2.id
  role       = aws_iam_role.text_mode_go_role_mnce0npc.id
}

resource "aws_iam_role_policy_attachment" "text_mode_role_236astbr_arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_937470c8_4aac_4db2_9a22_8be27d98acb4" {
  policy_arn = aws_iam_policy.arn_aws_iam__459018586415_policy_service_role_awslambdabasicexecutionrole_937470c8_4aac_4db2_9a22_8be27d98acb4.id
  role       = aws_iam_role.text_mode_role_236astbr.id
}

resource "aws_internet_gateway" "igw_00ba9e9efaf200476" {
  vpc_id = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_route_table" "rtb_0639e2d57d80e35eb" {
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw_00ba9e9efaf200476.id
  }

  vpc_id = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_security_group" "sg_0ffad359b81810e7c" {
  description = "default VPC security group"
  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
  }

  ingress {
    from_port = 0
    protocol  = "-1"
    self      = true
    to_port   = 0
  }

  name   = aws_vpc.vpc_0a3f76748f449c35a.instance_tenancy
  vpc_id = "vpc-0a3f76748f449c35a"
}

resource "aws_subnet" "subnet_0220733da80608b88" {
  availability_zone                   = "us-east-1c"
  cidr_block                          = "172.31.80.0/20"
  map_public_ip_on_launch             = true
  private_dns_hostname_type_on_launch = "ip-name"
  vpc_id                              = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_subnet" "subnet_08420e5d667ea3af5" {
  availability_zone_id                = "use1-az5"
  cidr_block                          = "172.31.64.0/20"
  map_public_ip_on_launch             = true
  private_dns_hostname_type_on_launch = "ip-name"
  vpc_id                              = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_subnet" "subnet_097640716d19251da" {
  availability_zone                   = "us-east-1a"
  cidr_block                          = "172.31.32.0/20"
  map_public_ip_on_launch             = true
  private_dns_hostname_type_on_launch = "ip-name"
  vpc_id                              = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_subnet" "subnet_0c57ee31848edbeb2" {
  availability_zone                   = "us-east-1e"
  cidr_block                          = "172.31.48.0/20"
  map_public_ip_on_launch             = true
  private_dns_hostname_type_on_launch = "ip-name"
  vpc_id                              = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_subnet" "subnet_0d6fb7fc1c74034f8" {
  availability_zone                   = "us-east-1d"
  cidr_block                          = "172.31.16.0/20"
  map_public_ip_on_launch             = true
  private_dns_hostname_type_on_launch = "ip-name"
  vpc_id                              = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_subnet" "subnet_0df1b2a23cd34288d" {
  availability_zone                   = "us-east-1b"
  cidr_block                          = "172.31.0.0/20"
  map_public_ip_on_launch             = true
  private_dns_hostname_type_on_launch = "ip-name"
  vpc_id                              = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_vpc" "vpc_0a3f76748f449c35a" {
  cidr_block           = "172.31.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  instance_tenancy     = "default"
}

resource "aws_lambda_function" "lambda_api" {
  architectures = ["x86_64"]
  environment {
    variables = {
      HOME = "/tmp"
    }

  }

  ephemeral_storage {
    size = 512
  }

  function_name                  = "lambda-api"
  image_uri                      = "459018586415.dkr.ecr.us-east-1.amazonaws.com/lambda-api:latest"
  memory_size                    = 500
  package_type                   = "Image"
  reserved_concurrent_executions = -1
  role                           = "arn:aws:iam::459018586415:role/service-role/lambda-api-role-hb6fczbh"
  source_code_hash               = "aa8708407f7371afd5525295659eba4a9272444b74bcb203536e94c8d34333ae"
  timeout                        = 120
  tracing_config {
    mode = "PassThrough"
  }

}

resource "aws_route53_record" "_hostedzone_z0907636hdo135djdt7g__12d1c97843134b6f6600cf18899b5ec5_earthly_tools_com__cname" {
  name    = "_12d1c97843134b6f6600cf18899b5ec5.earthly-tools.com"
  records = ["_f8939ef8e157a027322128878499a5e8.mntkzmhvxg.acm-validations.aws."]
  ttl     = 300
  type    = "CNAME"
  zone_id = aws_route53_zone._hostedzone_z0907636hdo135djdt7g.id
}

resource "aws_route53_record" "_hostedzone_z0907636hdo135djdt7g_earthly_tools_com__a" {
  alias {
    evaluate_target_health = true
    name                   = "d-gax6xknw2e.execute-api.us-east-1.amazonaws.com"
    zone_id                = "Z1UJRXOUMOOFQ8"
  }

  name    = "earthly-tools.com"
  type    = "A"
  zone_id = aws_route53_zone._hostedzone_z0907636hdo135djdt7g.id
}

resource "aws_route53_record" "_hostedzone_z0907636hdo135djdt7g_earthly_tools_com__ns" {
  name    = "earthly-tools.com"
  records = ["ns-277.awsdns-34.com.", "ns-1917.awsdns-47.co.uk.", "ns-1433.awsdns-51.org.", "ns-760.awsdns-31.net."]
  ttl     = 172800
  type    = "NS"
  zone_id = aws_route53_zone._hostedzone_z0907636hdo135djdt7g.id
}

resource "aws_route53_record" "_hostedzone_z0907636hdo135djdt7g_earthly_tools_com__soa" {
  name    = "earthly-tools.com"
  records = ["ns-1917.awsdns-47.co.uk. awsdns-hostmaster.amazon.com. 1 7200 900 1209600 86400"]
  ttl     = 900
  type    = "SOA"
  zone_id = aws_route53_zone._hostedzone_z0907636hdo135djdt7g.id
}

resource "aws_route53_zone" "_hostedzone_z0907636hdo135djdt7g" {
  comment = "HostedZone created by Route53 Registrar"
  name    = "earthly-tools.com"
}

resource "aws_route53_resolver_rule_association" "rslvr_autodefined_assoc_vpc_0a3f76748f449c35a_internet_resolver" {
  name             = "System Rule Association"
  resolver_rule_id = "rslvr-autodefined-rr-internet-resolver"
  vpc_id           = aws_vpc.vpc_0a3f76748f449c35a.id
}

resource "aws_s3_bucket" "text_mode" {
  arn            = "arn:aws:s3:::text-mode"
  bucket         = "text-mode"
  hosted_zone_id = "Z3AQBSTGFYJSTF"
}

