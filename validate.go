package u

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type vMap struct{
	Errors []string
	MapRules map[string]interface{}
	MapData map[string]interface{}
	Val interface{}
}
type vStruct struct{
	Errors []string
	Val interface{}
}

func ValidateMap(val interface{}) *vMap {
	return &vMap{
		MapRules: make(map[string]interface{}),
		MapData:  val.(map[string]interface{}),
	}
}

func ValidateStruct(val interface{}) *vStruct {
	return &vStruct{Val:val}
}

// Validate validate struct
func (receiver *vStruct)Validate() bool {
	var errors []string
	vo := reflect.ValueOf(receiver.Val)
	t := vo.Type()

	for i:=0; i<t.NumField(); i++ {
		tagRule := t.Field(i).Tag.Get("v")
		rules := strings.Split(tagRule, "|")
		field := t.Field(i).Name
		val := vo.Field(i).String()

		// ergodic
		for _, ruleItem := range rules {
			rule := strings.Split(ruleItem, "#")[0]

			// required rule
			if rule == "required" {
				ok, errMsg := required(field, ruleItem, val)
				if !ok { errors = append(errors, errMsg) }
			}

			// minLen:n
			if matched, _ := regexp.MatchString(`minLen:(\d+)`, rule); matched == true{
				ok, errMsg := minLen(field, ruleItem, val)
				if !ok { errors = append(errors, errMsg) }
			}

			// maxLen:n
			if matched, _ := regexp.MatchString(`maxLen:(\d+)`, rule); matched == true{
				ok, errMsg := maxLen(field, ruleItem, val)
				if !ok { errors = append(errors, errMsg) }
			}

		}
	}

	receiver.Errors = errors
	if len(errors) == 0 {
		return true
	} else {
		return false
	}
}

// Validate validate map
func (receiver *vMap) Validate() bool {
	var errors []string
	for field, val := range receiver.MapData {
		if receiver.MapRules[field] != nil {
			tagRule := receiver.MapRules[field].(string)
			rules := strings.Split(tagRule, "|")

			for _, ruleItem := range rules {
				rule := strings.Split(ruleItem, "#")[0]

				// required rule
				if rule == "required" {
					ok, errMsg := required(field, ruleItem, val.(string))
					if !ok { errors = append(errors, errMsg) }
				}

				// minLen:n
				if matched, _ := regexp.MatchString(`minLen:(\d+)`, rule); matched == true{
					ok, errMsg := minLen(field, ruleItem, val.(string))
					if !ok { errors = append(errors, errMsg) }
				}

				// maxLen:n
				if matched, _ := regexp.MatchString(`maxLen:(\d+)`, rule); matched == true{
					ok, errMsg := maxLen(field, ruleItem, val.(string))
					if !ok { errors = append(errors, errMsg) }
				}
			}
		}
	}

	receiver.Errors = errors
	if len(errors) == 0 {
		return true
	} else {
		return false
	}
}

// AddRule add map rule
func (receiver *vMap)AddRule(field string, rule string) {
	receiver.MapRules[field] = rule
}

// FirstError return first err msg
func (receiver *vMap) FirstError() string {
	if len(receiver.Errors) == 0 {
		return ""
	}
	return receiver.Errors[0]
}

// FirstError return first err msg
func (receiver *vStruct) FirstError() string {
	if len(receiver.Errors) == 0 {
		return ""
	}
	return receiver.Errors[0]
}

// maxLen:n
func maxLen(fieldName ,rule string, val string) (bool, string) {
	var customMsg string
	if len(strings.Split(rule, "#")) > 1 {
		customMsg = strings.Split(rule, "#")[1]
	}

	compile, _ := regexp.Compile(`maxLen:(\d+)`)
	if compile.MatchString(rule) {
		ruleNum := ToInt(compile.FindStringSubmatch(rule)[1])
		if len(val) > ruleNum {
			if len(customMsg) == 0 {
				return false, fmt.Sprintf("%s %s", fieldName, "maxLen")
			}else {
				return false, fmt.Sprintf("%s %s", fieldName, customMsg)
			}
		}
	}
	return true, ""
}

// minLen:n
func minLen(fieldName , rule string, val string) (bool, string) {
	var customMsg string
	if len(strings.Split(rule, "#")) > 1 {
		customMsg = strings.Split(rule, "#")[0]
	}

	compile, _ := regexp.Compile(`minLen:(\d+)`)
	if compile.MatchString(rule) {
		ruleNum := ToInt(compile.FindStringSubmatch(rule)[1])
		if len(val) < ruleNum {
			if len(customMsg) == 0 {
				return false, fmt.Sprintf("%s %s", fieldName, "minLen")
			}else {
				return false, fmt.Sprintf("%s %s", fieldName, customMsg)
			}
		}
	}
	return true, ""
}

// required
func required(fieldName ,customMsg string, val string) (bool, string) {
	if len(val) == 0 {
		if len(customMsg) == 0 {
			return false, fmt.Sprintf("%s %s", fieldName, "required")
		}else {
			return false, fmt.Sprintf("%s %s", fieldName, customMsg)
		}
	}
	return true, ""
}
