package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	taskGroupID := uuid.New().String()
	fmt.Println(taskGroupID)

}

//ffdb7e3d-461f-4ed5-b87e-746e8f41fc89
