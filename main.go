package main

import (
	"fmt"

	"github.com/alex-arraga/apple_store/hello"
	"github.com/google/uuid"
)

func main() {
	fmt.Print(hello.Greet())
	fmt.Print("New uuid:" + uuid.New().String())
}
