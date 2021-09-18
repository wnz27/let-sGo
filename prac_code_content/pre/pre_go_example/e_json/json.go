/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/17 11:14 上午
* Description:
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// We’ll use these two structs to demonstrate encoding and decoding of custom types below.
type response1 struct {
	Page int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.
type response2 struct {
	Page int			`json:"page"`
	Fruits []string		`json:"fruits"`
}

func main() {
	// First we’ll look at encoding basic data types to JSON strings.
	// Here are some examples for atomic values.
	bolB, _ := json.Marshal(true)
	fmt.Println("bool:true -->", string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println("int: 1 -->", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("float: 2.34 -->", string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println("string: gopher", string(strB))

	// And here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("string[]:", string(slcB), reflect.TypeOf(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("map:", string(mapB), reflect.TypeOf(mapB))

	// The JSON package can automatically encode your custom data types.
	// It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("response1:", string(res1B), reflect.TypeOf(res1B))

	//You can use tags on struct field declarations to customize the encoded JSON key names.
	// Check the definition of response2 above to see an example of such tags.
	res2D := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("response2:", string(res2B))

	// Now let’s look at decoding JSON data into Go values. Here’s an example for a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a", "b"]}`)

	// We need to provide a variable where the JSON package can put the decoded data.
	// This map[string]interface{} will hold a map of strings to arbitrary data types.
	var dat map[string]interface{}

	// Here’s the actual decoding, and a check for associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat, reflect.TypeOf(dat["num"]), reflect.TypeOf(dat["strs"]))

	// In order to use the values in the decoded map, we’ll need to convert them to their appropriate type.
	// For example here we convert the value in num to the expected float64 type.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of conversions.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// We can also decode JSON into custom data types.
	// This has the advantages of adding additional type-safety to our programs and
	// eliminating the need for type assertions when accessing the decoded data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("decode resp2:", res)
	fmt.Printf("decode resp2 with field: %+v\n", res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and strings as intermediates between the data and
	// JSON representation on standard out.
	// We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
