package main

import (
	"fmt"

	bbdbv1 "github.com/a-poor/bbdb/gen/proto/go/bbdb/v1"
)

func main() {
	fmt.Println("Hello, world!")

	f := bbdbv1.Field{
		Name:  "user",
		Dtype: bbdbv1.DataType_STRING,
	}
	// fmt.Println(f.String())
	fmt.Printf("%+v\n", f)
}
