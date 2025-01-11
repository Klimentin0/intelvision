package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type InPort struct {
	number int
	value  int
}

type OutPort struct {
	number int
	value  int
}

var InPorts []InPort
var OutPorts []OutPort

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if len(InPorts) == 0 {
			fmt.Print("Введите число IN портов \n")
			scanner.Scan()
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatalf("ОШИБКА перевода инпута в число: %v", err)
			}

			for i := 0; i < number; i++ {
				InPorts = append(InPorts, InPort{
					number: i + 1,
					value:  rand.Intn(2),
				})
			}

			fmt.Printf("IN портов: %d\n", len(InPorts))
		}
		if len(OutPorts) == 0 {
			fmt.Print("Введите число OUT портов \n")
			scanner.Scan()
			number2, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatalf("ОШИБКА перевода инпута в число: %v", err)
			}

			for i := 0; i < number2; i++ {
				OutPorts = append(OutPorts, OutPort{
					number: i + 1,
					value:  rand.Intn(2),
				})
			}
			fmt.Printf("OUT портов: %d\n", len(OutPorts))
		}

		fmt.Print("Введите команду('read', 'write' иди 'exit')>>>")
		scanner.Scan()
		strSlice := cleanInput(scanner.Text())
		if len(strSlice) == 0 {
			continue
		}

		commandName := strSlice[0]

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)
	for i, word := range strSlice {
		strSlice[i] = strings.ToLower(word)
	}
	return strSlice
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the repl",
			callback:    quitPrompt,
		},
		"read": {
			name:        "read",
			description: "Read the INport",
			callback:    Read,
		},
		"write": {
			name:        "write",
			description: "Write into OUT port",
			callback:    WRITE,
		},
	}
}
