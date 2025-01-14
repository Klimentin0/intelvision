package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Read() error {
	fmt.Println("Укажите, пожалуйста, тип порта:")
	fmt.Println()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	portTypeInput := strings.ToLower(scanner.Text())
	// testing
	// fmt.Printf("input %s\n", portTypeInput)
	// testing
	var desiredPortSlice []Port[int]
	var desiredPortSliceType string
	found := false
	for i := range availablePorts {
		for _, port := range availablePorts[i] {
			if portTypeInput == strings.ToLower(port.typeOf) {
				desiredPortSlice = availablePorts[i]
				desiredPortSliceType = port.typeOf
				found = true
			}
		}
		if found {
			break
		}
	}
	if !found {
		return errors.New("такого порта нет")
	}
	//testing
	fmt.Printf("Вы выбрали тип: %s\n", desiredPortSliceType)
	//testing
	fmt.Println()
	fmt.Printf("Укажите, пожалуйста, номер %s порта\n", desiredPortSliceType)
	fmt.Println("- - - - - - - - - - - - - - - - ")
	scanner.Scan()
	portNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("ОШИБКА перевода инпута в число: %v", err)
	}
	desiredPortValue := desiredPortSlice[portNumber-1].value
	desiredPortNumber := desiredPortSlice[portNumber-1].number
	fmt.Printf("Случайное значение %d  порта полученное при старте программы: %d\n", desiredPortNumber, desiredPortValue)
	return nil
}
