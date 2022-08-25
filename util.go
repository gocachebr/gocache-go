package gocache

import (
	"fmt"
	"strconv"
	"strings"
)

// Convert api response map keys to terraform variable names
func responseConvert(data map[string]interface{}, fields map[string]string) map[string]interface{} {
	output := make(map[string]interface{})
	for key, field := range data {
		modes, found := domainFields[key]
		//If the key is changed~
		newKey := key
		value, ok := fields[key]
		if ok {
			newKey = value
		}

		if found {
			if modes.Mode == "boolean" {
				if field == "true" {
					output[newKey] = true
				} else {
					output[newKey] = false
				}
			} else if modes.Mode == "int" {
				i, err := strconv.Atoi(field.(string))
				if err == nil {
					output[newKey] = i
				} else {
					output[newKey] = 0
				}
			} else {
				output[newKey] = field
			}
		} else {
			output[newKey] = field
		}
	}
	return output
}

// Convert map to formData for request body
func formData(data map[string]interface{}, fields map[string]string) string {
	result := ""

	for key, value := range data {
		key = strings.ToLower(key)
		field, ok := fields[key]
		name := key
		if ok {
			name = field
		}
		result += fmt.Sprintf("%v=%v&", name, value)
	}

	if result[len(result)-1] == '&' {
		result = result[:len(result)-1]
	}
	return result
}

func GetAllDomainSettings() map[string]fieldType {
	return domainFields
}

func GetFieldName(str string, resource string) string {
	var convert map[string]string
	switch resource {
	case "domain":
		convert = domainConvert
	case "domain_reversed":
		convert = domainConvertReversed
	case "smart_rule_match":
		convert = smartRuleMatchConvert
	case "smart_rule_match_reversed":
		convert = smartRuleMatchConvertReversed
	case "smart_rule_settings_action":
		convert = smartRuleActionSettingsCovert
	case "smart_rule_settings_action_reversed":
		convert = smartRuleActionSettingsCovertReversed
	case "smart_rule_firewall_action":
		convert = smartRuleActionFirewallCovert
	case "smart_rule_firewall_action_reversed":
		convert = smartRuleActionFirewallCovertReversed
	case "smart_rule_ratelimit_action":
		convert = rateLimitActionReversed
	case "smart_rule_ratelimit_match":
		convert = rateLimitMatchConvertReversed
	}
	converted, ok := convert[str]
	if !ok {
		return str
	}
	return converted
}

func GetAllFieldsAdjusted(resource string) map[string]fieldType {
	var aux map[string]fieldType
	var convert map[string]string
	switch resource {
	case "domain":
		aux = domainFields
		convert = domainConvertReversed
	case "smart_rule_match":
		aux = smartRuleMatchFields
		convert = smartRuleMatchConvertReversed
	case "smart_rule_settings_action":
		aux = smartRuleSettingsActions
		convert = smartRuleActionSettingsCovertReversed
	case "smart_rule_rewrite_action":
		aux = smartRuleRewriteActions
	case "smart_rule_firewall_action":
		aux = smartRuleFirewallActions
		convert = smartRuleActionFirewallCovertReversed
	case "smart_rule_metadata":
		aux = smartRuleMetadataFields
	case "smart_rule_ratelimit_match":
		aux = smartRuleMatchFields
		convert = rateLimitMatchConvertReversed
	case "smart_rule_ratelimit_action":
		aux = rateLimitActions
		convert = rateLimitActionReversed
	case "smart_rule_ratelimit_metadata":
		aux = smartRuleMetadataFields

	default:
		return nil
	}
	fields := make(map[string]fieldType)
	for key, value := range aux {
		alias, ok := convert[key]
		if ok {
			fields[alias] = value
		} else {
			fields[key] = value
		}
	}
	return fields
}

func inList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func CheckDomainValue(field string, value string) bool {
	allowed, ok := GetAllFieldsAdjusted("domain")[field]
	if ok {
		if allowed.Allowed[0] == "0-100" {
			i, err := strconv.Atoi(field)
			if err != nil {
				return i >= 0 && i <= 100
			}
		} else if allowed.Allowed[0] == "0-N" {
			i, err := strconv.Atoi(field)
			if err != nil {
				return i >= 0
			}
		} else if allowed.Allowed[0] == "ANY" {
			return true
		} else if inList(value, allowed.Allowed) {
			return true
		}
	}
	return false
}

func FieldExists(field string, resource string) bool {
	for k := range GetAllFieldsAdjusted(resource) {
		if field == k {
			return true
		}
	}
	return false
}
