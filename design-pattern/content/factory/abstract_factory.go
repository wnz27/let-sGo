/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/4 22:47 9月
 **/
package factory


// ISystemConfigParser hh
type ISystemConfigParser interface {
	ParseSystem(date []byte)
}

// jsonSystemConfigParser hh
type jsonSystemConfigParser struct {
}

// Parse
func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}


type jsonConfigParserFactory struct {
}

func(j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}


