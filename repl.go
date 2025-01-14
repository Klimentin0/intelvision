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

type Port[v any] struct {
	number int
	value  v
	typeOf string
}

var InPorts []Port[int]
var OutPorts []Port[int]
var availablePorts [][]Port[int]

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
				InPorts = append(InPorts, Port[int]{
					number: i + 1,
					value:  rand.Intn(2),
					typeOf: "InPort",
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
				OutPorts = append(OutPorts, Port[int]{
					number: i + 1,
					value:  rand.Intn(2),
					typeOf: "OutPort",
				})
			}
			fmt.Printf("OUT портов: %d\n", len(OutPorts))
		}
		// availablePorts только тут я наполняю, потому что ссылка(чем является слайс по определению)
		// на лежащий под ними ряд(array), должна создаваться на новые наполненные уже слайсы in и out
		//
		availablePorts = [][]Port[int]{InPorts, OutPorts}
		//testing
		//fmt.Printf("test slice %s", availablePorts[1][1].typeOf)
		//testing

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
