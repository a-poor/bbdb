package main

import (
	"fmt"

	pb "github.com/a-poor/bbdb/gen/proto/go/bbdb/v1alpha"
)

func main() {
	fmt.Println("Hello, world!")

	f := pb.Field{
		Name:  "user",
		Dtype: pb.DataType_STRING,
	}
	fmt.Println(f.String())
	// fmt.Printf("%+v\n", f)
}
