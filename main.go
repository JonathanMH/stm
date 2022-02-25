package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateEmail(initialEmail string) string {
	id := uuid.New()
	splitEmail := strings.Split(initialEmail, "@")
	currentTime := time.Now();
	email := fmt.Sprintf("%s+%s+%s@%s\n", splitEmail[0],currentTime.Format("2006-01-02"), id, splitEmail[1])
	return email;
}

func main() {
	email := os.Args[1]
	fmt.Printf(GenerateEmail(email))
}
