package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const ArraySize = 65536

func getFileContent() string {
	body, err := os.ReadFile("../helloWorld.bf")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	outString := strings.ReplaceAll(string(body), "\n", "")
	outString = strings.ReplaceAll(string(outString), " ", "")
	outString = strings.ReplaceAll(string(outString), "\t", "")
	outString = strings.ReplaceAll(string(outString), "\r", "")
	return outString
}

func main() {
	// bytes for program
	var byteArray [ArraySize]byte

	// "pointer" to program-array
	var cellPointer uint16 = 0

	var fileIndex uint16 = 0

	// content from file
	var contentArray [ArraySize]string

	fileContent := getFileContent()

	for i := 0; i < len(fileContent); i++ {
		contentArray[i] = string([]rune(fileContent)[i])
	}
	contentArray[len(fileContent)] = "EOF"

	for i := range ArraySize {
		byteArray[i] = 0
	}

	running := true

	jumpIndices := []uint8{}

	for running {

		switch contentArray[fileIndex] {

		case "+":
			byteArray[cellPointer]++
		case "-":
			byteArray[cellPointer]--
		case ">":
			cellPointer++
		case "<":
			cellPointer--
		case ".":
			fmt.Print(string(byteArray[cellPointer]))
		case ",":
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			byteArray[cellPointer] = byte([]rune(text)[0])
		case "EOF":
			running = false
			continue

		case "[":
			//set reset point in array of reset point
			jumpIndices = append(jumpIndices, uint8(fileIndex))
		case "]":
			//check for 0 --> jump to last reset point || remove last reset point
			if len(jumpIndices) == 0 {
				break
			}

			if byteArray[cellPointer] != 0 {
				fileIndex = uint16(jumpIndices[len(jumpIndices)-1])
			}
			//if len(jumpIndices) > 0 {
			//	jumpIndices = jumpIndices[:len(jumpIndices)-1]
			//}
		default:
		}

		fileIndex++
	}
}
