package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var arguments []string
	for scanner.Scan() {
		line := scanner.Text()
		arguments = append(arguments, line)
	}

	if len(arguments) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	x := len(arguments[0])
	y := len(arguments)

	var sl []string
	for a := 'A'; a <= 'E'; a++ {
		cmd := exec.Command("./quad"+string([]rune{a}), strconv.Itoa(x), strconv.Itoa(y))
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}
		outputStr := string(output)
		sl = append(sl, outputStr)
	}

	var match []bool
	joinedArgs := strings.Join(arguments, "\n")
	//fmt.Println(joinedArgs)
	for _, shape := range sl {
		// fmt.Println(joinedArgs)
		// fmt.Println(shape)
		
		if strings.TrimSpace(shape) == strings.TrimSpace(joinedArgs) {
			match = append(match, true)
		} else {
			match = append(match, false)
		}
	}

	// for _, word := range match{
	// 	if word {
	// 		fmt.Print("true")
	// 	}else{
	// 		fmt.Println("false")
	// 	}
	// }

	X := strconv.Itoa(x)
	Y := strconv.Itoa(y)

	output := ""
	for i, isMatch := range match {
		if isMatch {
			if output != "" {
				output += " || "
			}
			output += fmt.Sprintf("[quad%c] [%s] [%s]", 'A'+i, X, Y)
		}
	}
	
	if output == "" {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(output)
	}
}
