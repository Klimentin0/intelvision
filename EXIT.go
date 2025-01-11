package main

import (
	"fmt"
	"os"
)

func quitPrompt() error {
	fmt.Println("Заканчиваю работу, всего хорошего!")
	os.Exit(0)
	return nil
}
