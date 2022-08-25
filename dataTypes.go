package gocache

type API_Response struct {
	Status_code    int         `json:"status_code,omitempty"`
	Msg            string      `json:"msg,omitempty"`
	Response       interface{} `json:"response,omitempty"`
	HTTPStatusCode int         `json:"-"`
}

type DomainList struct {
	Domains     []string    `json:"domains"`
	Nameservers interface{} `json:"nameservers,omitempty"`
}

type DNSResult struct {
	Records []map[string]interface{} `json:"records"`
}

type DNSGetResult struct {
	Records map[string]interface{} `json:"records"`
}

type CertificateResult struct {
	Msg          string                   `json:"msg,omitempty"`
	Certificates []map[string]interface{} `json:"certificates,omitempty"`
}

type CertificateInput struct {
	Certificate string `json:"certificate,omitempty"`
	Privatekey  string `json:"privatekey,omitempty"`
	Type        string `json:"type,omitempty"`
	Enable      bool   `json:"enable"`
}

type SmartRuleResult struct {
	Bulk_id int                    `json:"bulk_id,omitempty"`
	Bulk    []SmartRule            `json:"bulk,omitempty"`
	Rules   []SmartRule            `json:"rules,omitempty"`
	Bulks   map[string]interface{} `json:"bulks,omitempty"`
}

type SmartRule struct {
	Match    map[string]interface{} `json:"match,omitempty"`
	Action   map[string]interface{} `json:"action,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Id       string                 `json:"id,omitempty"`
}

type fieldType struct {
	Mode    string
	Default interface{}
	Allowed []string
}

var domainFields = map[string]fieldType{
	"cache_ttl":                        fieldType{Mode: "int", Default: 86400, Allowed: []string{"300", "600", "900", "1800", "3600", "7200", "14400", "86400", "172800", "604800"}},
	"deploy_mode":                      fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"smart_status":                     fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"smart_tpl":                        fieldType{Mode: "string", Default: "custom", Allowed: []string{"custom", "wordpress", "magento", "joomla"}},
	"smart_ttl":                        fieldType{Mode: "int", Default: 14400, Allowed: []string{"300", "600", "900", "1800", "3600", "7200", "14400", "86400", "172800", "604800"}},
	"cdn_mode":                         fieldType{Mode: "string", Default: "ns", Allowed: []string{"ns", "cname"}},
	"gzip_status":                      fieldType{Mode: "boolean", Default: true, Allowed: []string{"true", "false"}},
	"expires_ttl":                      fieldType{Mode: "int", Default: 14400, Allowed: []string{"-1", "3600", "7200", "14400", "43200", "86400", "172800", "345600", "604800", "1296000", "2592000", "15552000", "31536000"}},
	"ignore_vary":                      fieldType{Mode: "boolean", Default: true, Allowed: []string{"true", "false"}},
	"ignore_cache_control":             fieldType{Mode: "boolean", Default: true, Allowed: []string{"true", "false"}},
	"ignore_expires":                   fieldType{Mode: "boolean", Default: true, Allowed: []string{"true", "false"}},
	"ssl_mode":                         fieldType{Mode: "string", Default: "full", Allowed: []string{"full", "partial"}},
	"cache_301":                        fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"cache_302":                        fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"cache_404":                        fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"header_device_type":               fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"header_geoip_continent":           fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"header_geoip_country":             fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"header_geoip_org":                 fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"caching_behavior":                 fieldType{Mode: "string", Default: "default", Allowed: []string{"default", "ignore_query_string"}},
	"waf_status":                       fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"waf_level":                        fieldType{Mode: "string", Default: "high", Allowed: []string{"low", "medium", "high"}},
	"waf_mode":                         fieldType{Mode: "string", Default: "simulate", Allowed: []string{"simulate", "challenge", "block"}},
	"expire_bypass_sec":                fieldType{Mode: "int", Default: -1, Allowed: []string{"-1", "300", "600", "900", "1800", "3600", "7200", "14400", "43200", "86400", "172800", "345600", "604800", "1296000", "2592000", "15552000", "31536000"}},
	"tls10":                            fieldType{Mode: "boolean", Default: true, Allowed: []string{"true", "false"}},
	"tls11":                            fieldType{Mode: "boolean", Default: true, Allowed: []string{"true", "false"}},
	"tls12":                            fieldType{Mode: "boolean", Default: true, Allowed: []string{"true", "false"}},
	"image_optimize":                   fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"image_optimize_webp":              fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"image_optimize_progressive":       fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"image_optimize_metadata":          fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"image_optimize_level":             fieldType{Mode: "int", Default: 0, Allowed: []string{"0-100"}},
	"rate_limit_status":                fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"rate_limit_ignore_known_bots":     fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"rate_limit_ignore_static_content": fieldType{Mode: "boolean", Default: false, Allowed: []string{"true", "false"}},
	"public_log":                       fieldType{Mode: "string", Default: "off", Allowed: []string{"off", "v1", "v2"}},
	"log_freq":                         fieldType{Mode: "int", Default: 60, Allowed: []string{"1", "5", "10", "15", "30", "60"}},
}

var smartRuleMatchFields = map[string]fieldType{
	"device_type":      fieldType{Mode: "list", Allowed: []string{"mobile", "desktop", "bot", "na"}},
	"http_version":     fieldType{Mode: "list", Allowed: []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}},
	"request_method":   fieldType{Mode: "list", Allowed: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}},
	"scheme":           fieldType{Mode: "string", Allowed: []string{"http", "https", "http*"}},
	"hostname":         fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"request_uri":      fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"http_user_agent":  fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"cookie":           fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"cookie_content":   fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"http_referer":     fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"remote_address":   fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"header":           fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"origin_country":   fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"origin_continent": fieldType{Mode: "list", Allowed: []string{"SA", "NA", "OC", "EU", "AF", "AS"}},
	"bots":             fieldType{Mode: "string", Allowed: []string{"known", "others"}},
}

var smartRuleSettingsActions = map[string]fieldType{
	"set_host":                   fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"backend":                    fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"set_uri":                    fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"custom_cache_key":           fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"expires_ttl":                fieldType{Mode: "int", Allowed: []string{"-1", "3600", "7200", "14400", "43200", "86400", "172800", "345600", "604800", "1296000", "2592000", "15552000", "31536000"}},
	"caching_behaviour":          fieldType{Mode: "string", Allowed: []string{"default", "ignore_query_string"}},
	"cache_mode":                 fieldType{Mode: "string", Allowed: []string{"off", "default", "full"}},
	"cache_ttl":                  fieldType{Mode: "int", Allowed: []string{"300", "600", "900", "1800", "3600", "7200", "14400", "86400", "172800", "604800", "1296000"}},
	"cors":                       fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"ssl_mode":                   fieldType{Mode: "string", Allowed: []string{"partial", "full"}},
	"hide":                       fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"encoding":                   fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"set":                        fieldType{Mode: "map", Allowed: []string{"ANY"}},
	"set_req_header":             fieldType{Mode: "map", Allowed: []string{"ANY"}},
	"signed_url_key":             fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"signed_url_type":            fieldType{Mode: "string", Allowed: []string{"s3qs", "off"}},
	"cache_301":                  fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"cache_302":                  fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"cache_404":                  fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"gzip_status":                fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"ignore_cache_control":       fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"ignore_expires":             fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"ignore_vary":                fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"waf_status":                 fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"waf_mode":                   fieldType{Mode: "string", Allowed: []string{"block", "simulate", "challenge"}},
	"waf_level":                  fieldType{Mode: "string", Allowed: []string{"low", "medium", "high"}},
	"ratelimit_status":           fieldType{Mode: "string", Allowed: []string{"on", "off"}},
	"image_optimize":             fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize_webp":        fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize_progressive": fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize_metadata":    fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize_level":       fieldType{Mode: "int", Allowed: []string{"0-100"}},
	"waf_rule_action":            fieldType{Mode: "string", Allowed: []string{"ANY"}},
}

var smartRuleRewriteActions = map[string]fieldType{
	"redirect_type": fieldType{Mode: "string", Allowed: []string{"301", "302"}},
	"redirect_to":   fieldType{Mode: "string", Allowed: []string{"ANY"}},
}

var smartRuleFirewallActions = map[string]fieldType{
	"action": fieldType{Mode: "string", Allowed: []string{"block", "challenge", "accept"}},
}

var smartRuleMetadataFields = map[string]fieldType{
	"status": fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"notes":  fieldType{Mode: "string", Allowed: []string{"ANY"}},
	"name":   fieldType{Mode: "string", Allowed: []string{"ANY"}},
}

var domainConvert = map[string]string{
	"smart_cache_status":          "smart_status",
	"smart_cache_template":        "smart_tpl",
	"smart_cache_ttl":             "smart_ttl",
	"dns_mode":                    "cdn_mode",
	"rate_limit_ignore_good_bots": "rate_limit_ignore_known_bots",
}
var domainConvertReversed = map[string]string{
	"smart_status":                 "smart_cache_status",
	"smart_tpl":                    "smart_cache_template",
	"smart_ttl":                    "smart_cache_ttl",
	"cdn_mode":                     "dns_mode",
	"rate_limit_ignore_known_bots": "rate_limit_ignore_good_bots",
}

var smartRuleMatchConvert = map[string]string{
	"hostname": "host",
}

var smartRuleMatchConvertReversed = map[string]string{
	"host": "hostname",
}

var smartRuleActionSettingsCovert = map[string]string{
	"hide_header":          "hide",
	"encoding_header":      "encoding",
	"set_response_headers": "set",
	"set_request_headers":  "set_req_header",
}

var smartRuleActionSettingsCovertReversed = map[string]string{
	"hide":           "hide_header",
	"encoding":       "encoding_header",
	"set":            "set_response_headers",
	"set_req_header": "set_request_headers",
}

var smartRuleActionFirewallCovert = map[string]string{
	"firewall_action": "action",
}

var smartRuleActionFirewallCovertReversed = map[string]string{
	"action": "firewall_action",
}

var recordConvert = map[string]string{
	"proxied": "cloud",
}

var rateLimitMatchConvertReversed = map[string]string{
	"host": "hostname",
}

var rateLimitActions = map[string]fieldType{
	"rate_limit_action":    fieldType{Mode: "string", Allowed: []string{"block", "challenge", "simulate"}},
	"rate_limit_block_ttl": fieldType{Mode: "string", Allowed: []string{"10", "30", "60", "120", "300", "600", "1800", "3600", "14400", "43200", "86400"}},
	"rate_limit_amount":    fieldType{Mode: "int", Default: 10, Allowed: []string{"0-N"}},
	"rate_limit_period":    fieldType{Mode: "string", Allowed: []string{"10", "30", "60", "120", "300", "600", "1800", "3600", "14400", "43200", "86400"}},
}

var rateLimitActionReversed = map[string]string{}
