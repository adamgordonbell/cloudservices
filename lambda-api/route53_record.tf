# resource "aws_route53_record" "tfer--Z0907636HDO135DJDT7G__12d1c97843134b6f6600cf18899b5ec5-002E-earthly-tools-002E-com-002E-_CNAME_" {
#   name    = "_12d1c97843134b6f6600cf18899b5ec5.earthly-tools.com"
#   records = ["_f8939ef8e157a027322128878499a5e8.mntkzmhvxg.acm-validations.aws."]
#   ttl     = "300"
#   type    = "CNAME"
#   zone_id = "${aws_route53_zone.tfer--Z0907636HDO135DJDT7G_earthly-tools-002E-com.zone_id}"
# }

# resource "aws_route53_record" "tfer--Z0907636HDO135DJDT7G_earthly-tools-002E-com-002E-_A_" {
#   alias {
#     evaluate_target_health = "true"
#     name                   = "d-gax6xknw2e.execute-api.us-east-1.amazonaws.com"
#     zone_id                = "Z1UJRXOUMOOFQ8"
#   }

#   name    = "earthly-tools.com"
#   type    = "A"
#   zone_id = "${aws_route53_zone.tfer--Z0907636HDO135DJDT7G_earthly-tools-002E-com.zone_id}"
# }

# resource "aws_route53_record" "tfer--Z0907636HDO135DJDT7G_earthly-tools-002E-com-002E-_NS_" {
#   name    = "earthly-tools.com"
#   records = ["ns-1433.awsdns-51.org.", "ns-1917.awsdns-47.co.uk.", "ns-277.awsdns-34.com.", "ns-760.awsdns-31.net."]
#   ttl     = "172800"
#   type    = "NS"
#   zone_id = "${aws_route53_zone.tfer--Z0907636HDO135DJDT7G_earthly-tools-002E-com.zone_id}"
# }

# resource "aws_route53_record" "tfer--Z0907636HDO135DJDT7G_earthly-tools-002E-com-002E-_SOA_" {
#   name    = "earthly-tools.com"
#   records = ["ns-1917.awsdns-47.co.uk. awsdns-hostmaster.amazon.com. 1 7200 900 1209600 86400"]
#   ttl     = "900"
#   type    = "SOA"
#   zone_id = "${aws_route53_zone.tfer--Z0907636HDO135DJDT7G_earthly-tools-002E-com.zone_id}"
# }
