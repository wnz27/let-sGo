/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/11 4:16 下午
* Description:
 */
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

// structValues 给任意一个结构体返回 规定字符串
// insert into order values(ordId: 456, customerId: 56)
// insert into employee values(name: "Naveen", id: 565, address: "Coimbatore", salary: 90000, country: "India")
func structValues(passStruct interface{}) string {
	builder := strings.Builder{}
	builder.WriteString("insert into ")
	sType := reflect.TypeOf(passStruct)
	sValue := reflect.ValueOf(passStruct)
	structName := sType.Name()
	builder.WriteString(structName + " values(")

	fieldCount := sValue.NumField()
	for i := 0; i < fieldCount; i++ {
		// 处理字段类型
		currField := sType.Field(i)

		var fieldValue string
		currFieldTypeStr := currField.Type.Kind().String()
		if currFieldTypeStr == "string" {
			fieldValue = sValue.Field(i).String()
		} else if currFieldTypeStr == "int" {
			fieldValue = strconv.Itoa(int(sValue.Field(i).Int()))
		}
		fieldKey := currField.Name
		builder.WriteString(fieldKey + ": " + fieldValue)

		if i != fieldCount-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")")
	return builder.String()
}

// reflect
func createQuery(q interface{}) {
	if reflect.TypeOf(q).Kind().String() != "struct" {
		fmt.Println("unsupported type")
		return
	}
	res := structValues(q)
	fmt.Println(res)
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}

/*
insert into order values(456, 56)
insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
*/
func (o order) String() string {
	return fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
}

func (e employee) String() string {
	return fmt.Sprintf("insert into employee values(%s, %d, %s, %d, %s)", e.name, e.id, e.address, e.salary, e.country)
}

// switch type assert 缺点是struct 一改，相应 String 方法就需要调整
func createQuerySwitch(q interface{}) {
	switch t := q.(type) {
	case order:
		fmt.Println(t)
	case employee:
		fmt.Println(t)
	default:
		fmt.Println("unsupported type")
	}
}
