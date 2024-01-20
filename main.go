package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const ARRAY_SIZE = 65536

func getFileContent() string {
	body, err := ioutil.ReadFile("./helloWorld.bf")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var outString string = strings.Replace(string(body), "\n", "", -1)
	outString = strings.Replace(string(outString), " ", "", -1)
	outString = strings.Replace(string(outString), "\t", "", -1)
	outString = strings.Replace(string(outString), "\r", "", -1)
	return outString
}

func main() {
	// bytes for program
	var byteArray [ARRAY_SIZE]byte

	// "pointer" to program-array
	var cellPointer uint16 = 0

	var fileIndex uint16 = 0

	// content from file
	var contentArray [ARRAY_SIZE]string

	fileContent := getFileContent()

	for i := 0; i < len(fileContent); i++ {
		contentArray[i] = string([]rune(fileContent)[i])
	}
	contentArray[len(fileContent)] = "EOF"

	for i := 0; i < ARRAY_SIZE; i++ {
		byteArray[i] = 0
	}

	var running bool = true

	var jumpIndices []uint8 = []uint8{}

	for running {

		switch contentArray[fileIndex] {

		case "+":
			byteArray[cellPointer]++
			break

		case "-":
			byteArray[cellPointer]--
			break

		case ">":
			cellPointer++
			break

		case "<":
			cellPointer--
			break

		case ".":
			fmt.Print(string(byteArray[cellPointer]))
			break

		case ",":
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			byteArray[cellPointer] = byte([]rune(text)[0])
			break

		case "EOF":
			running = false
			continue

		case "[":
			//set reset point in array of reset point
			jumpIndices = append(jumpIndices, uint8(fileIndex))
			break

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
			break

		default:
			break

		}

		fileIndex++
	}
}
