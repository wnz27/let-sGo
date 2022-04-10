/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/9/4 22:47 9月
 **/
package factory

import (
	"reflect"
	"testing"
)

func TestJsonConfigParserFactory_CreateRuleParser(t *testing.T) {
	tests := []struct{
		name string
		want IRuleConfigParser
	}{
		{
			name: "json",
			want: jsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleparser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonConfigParserFactory_CreateSystemParser(t *testing.T) {
	tests := []struct{
		name string
		want ISystemConfigParser
	}{
		{
			name: "json",
			want: jsonSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
