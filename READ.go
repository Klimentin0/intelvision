package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Read() error {
	fmt.Println()
	fmt.Println("Укажите, пожалуйста, номер IN порта")
	fmt.Println("- - - - - - - - - - - - - - - - ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	portNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("ОШИБКА перевода инпута в число: %v", err)
	}
	inPortValue := InPorts[portNumber-1].value
	inPortNumber := InPorts[portNumber-1].number
	fmt.Printf("Случайное значение %d IN порта полученное при старте программы: %d\n", inPortNumber, inPortValue)
	return nil
}
