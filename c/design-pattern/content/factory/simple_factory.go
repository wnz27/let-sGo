/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/4 14:46 9月
 **/
package factory

// IRuleConfigParser
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser
type jsonRuleConfigParser struct {

}

// Parse
func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser
type yamlRuleConfigParser struct {
}

// Parse
func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// NewIRuleConfigParser
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}

