package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func WRITE() error {
	fmt.Println()
	fmt.Println("Укажите, пожалуйста, номер OUT порта")
	fmt.Println("- - - - - - - - - - - - - - - - ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	portNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("ОШИБКА перевода инпута в число: %v", err)
	}
	outPortNumber := OutPorts[portNumber-1].number
	outPortValue := OutPorts[portNumber-1].value
	fmt.Printf("OUT порт номер: %d\nЗначение порта: %d\n", outPortNumber, outPortValue)

	fmt.Println()
	fmt.Println("Укажите, пожалуйста, значение OUT порта")
	fmt.Println("- - - - - - - - - - - - - - - - ")
	scanner.Scan()
	portValueInput, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("ОШИБКА перевода инпута в число: %v", err)
	}
	OutPorts[portNumber-1].value = portValueInput

	fmt.Printf("OUT порт новое значение: %d\n", OutPorts[portNumber - 1].value)
	return nil
}
