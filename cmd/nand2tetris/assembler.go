package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var linesInput []string

var symbol_table = make(map[string]int)

var dest = make(map[string]string)

var comp_a = make(map[string]string)

var comp_m = make(map[string]string)

var jump = make(map[string]string)

func init() {
	log.Printf("Initializing program")
	//built in constants
	symbol_table["R0"] = 0
	symbol_table["R1"] = 1
	symbol_table["R2"] = 2
	symbol_table["R3"] = 3
	symbol_table["R4"] = 4
	symbol_table["R5"] = 5
	symbol_table["R6"] = 6
	symbol_table["R7"] = 7
	symbol_table["R8"] = 8
	symbol_table["R9"] = 9
	symbol_table["R10"] = 10
	symbol_table["R11"] = 11
	symbol_table["R12"] = 12
	symbol_table["R13"] = 13
	symbol_table["R14"] = 14
	symbol_table["R15"] = 15
	symbol_table["SCREEN"] = 16384
	symbol_table["KBD"] = 24576
	//not needed now
	symbol_table["SP"] = 0
	symbol_table["LCL"] = 1
	symbol_table["ARG"] = 2
	symbol_table["THIS"] = 3
	symbol_table["THAT"] = 4

	dest["null"] = "000"
	dest["M"] = "001"
	dest["D"] = "010"
	dest["MD"] = "011"
	dest["A"] = "100"
	dest["AM"] = "101"
	dest["AD"] = "110"
	dest["AMD"] = "111"

	comp_a["0"] = "101010"
	comp_a["1"] = "111111"
	comp_a["-1"] = "111010"
	comp_a["D"] = "001100"
	comp_a["A"] = "110000"
	comp_a["!D"] = "001101"
	comp_a["!A"] = "110001"
	comp_a["-D"] = "001111"
	comp_a["-A"] = "110011"
	comp_a["D+1"] = "011111"
	comp_a["A+1"] = "110111"
	comp_a["D-1"] = "001110"
	comp_a["A-1"] = "110010"
	comp_a["D+A"] = "000010"
	comp_a["D-A"] = "010011"
	comp_a["A-D"] = "000111"
	comp_a["D&A"] = "000000"
	comp_a["D|A"] = "010101"

	comp_m["M"] = "110000"
	comp_m["!M"] = "110001"
	comp_m["-M"] = "110011"
	comp_m["M+1"] = "110111"
	comp_m["M-1"] = "110010"
	comp_m["D+M"] = "000010"
	comp_m["D-M"] = "010011"
	comp_m["M-D"] = "000111"
	comp_m["D&M"] = "000000"
	comp_m["D|M"] = "010101"

	jump["null"] = "000"
	jump["JGT"] = "001"
	jump["JEQ"] = "010"
	jump["JGE"] = "011"
	jump["JLT"] = "100"
	jump["JNE"] = "101"
	jump["JLE"] = "110"
	jump["JMP"] = "111"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Need at-least one argument as filename\n")
		return
	}

	fileName := os.Args[1]
	fmt.Printf("Input file: %s\n", fileName)

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesInput = append(linesInput, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	instructionLineNumber := 0

	for _, line := range linesInput {
		//fmt.Printf("%d:\t%s\n", lnum, line)
		buildSymbols(instructionLineNumber, strings.TrimSpace(strings.Split(line, "//")[0]))

		if isAnInstruction(line) {
			instructionLineNumber++
		}
	}

	printSymbolTable()

	outFileName := strings.Split(fileName, ".asm")[0] + ".hack"
	f, err := os.Create(outFileName)

	if err != nil {
		log.Fatal(err)
		return
	}

	n := 16
	instruction := ""

	for lnum, line := range linesInput {
		instruction = process(lnum, line, &n)

		if instruction == "" {
			continue
		}

		_, err := f.WriteString(instruction + "\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func isAnInstruction(input string) bool {
	return !(strings.HasPrefix(input, "//") || len(input) <= 0 || strings.HasPrefix(input, "("))
}

func buildSymbols(instrNum int, input string) {
	if strings.HasPrefix(input, "//") || len(input) <= 0 {
		return
	}
	if strings.HasPrefix(input, "(") {
		symbol := input[1:strings.Index(input, ")")] //todo handle errors
		symbol_table[symbol] = instrNum
	}
}

func printSymbolTable() {
	fmt.Println("Symbol table:")
	for k, v := range symbol_table {
		fmt.Printf("%s:%d\n", k, v)
	}
}

func process(lnum int, input string, memLoc *int) string {
	if strings.HasPrefix(input, "//") || len(input) <= 0 || strings.HasPrefix(input, "(") {
		return ""
	}

	instruction := ""

	input = strings.TrimSpace(strings.Split(input, "//")[0])

	if strings.HasPrefix(input, "@") {
		instruction = handleAInstruction(input, memLoc)
	} else {
		instruction = handleCInstruction(input)
	}

	//fmt.Println(instruction)
	return instruction
}

func handleAInstruction(input string, memLoc *int) string {
	address := input[1:]

	if val, exists := symbol_table[address]; exists {
		foundAddress := strconv.FormatInt(int64(val), 2)
		return paddA(foundAddress)
	} else {
		//either variable or a direct addressing
		i, err := strconv.ParseInt(address, 10, 64)
		if err == nil {
			foundAddress := strconv.FormatInt(int64(i), 2)
			return paddA(foundAddress)
		} else {
			log.Printf("Mairu - %s\n", input)
			//variable
			foundAddress := strconv.FormatInt(int64(*memLoc), 2)
			symbol_table[address] = *memLoc
			log.Printf("Pooru - %s\n", foundAddress)
			*memLoc++
			return paddA(foundAddress)
		}
	}
}

func paddA(input string) string {
	prefix := ""

	for i := len(input); i < 16; i++ {
		prefix += "0"
	}

	return prefix + input
}

func handleCInstruction(input string) string {
	instruction := ""
	prefix := "111"
	a := "0"

	if strings.Contains(input, "=") {
		parts := strings.Split(input, "=")
		destPart := parts[0]
		right := parts[1]
		a := "0"

		if _, exists := comp_m[right]; exists {
			a = "1"
		}

		if a == "1" {
			instruction = prefix + a + comp_m[right] + dest[destPart] + jump["null"]
		} else {
			instruction = prefix + a + comp_a[right] + dest[destPart] + jump["null"]
		}
	} else if strings.Contains(input, ";") {
		parts := strings.Split(input, ";")
		comp := parts[0]
		jumpPart := parts[1]

		if _, exists := comp_m[comp]; exists {
			a = "1"
		}

		if a == "1" {
			instruction = prefix + a + comp_m[comp] + dest["null"] + jump[jumpPart]
		} else {
			instruction = prefix + a + comp_a[comp] + dest["null"] + jump[jumpPart]
		}
	}

	return instruction
}
