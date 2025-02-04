package main

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/alex-arraga/apple_store/hello"
)

func main() {
	fmt.Print(hello.Greet())
	fmt.Print("New uuid:" + uuid.New().String())
}
