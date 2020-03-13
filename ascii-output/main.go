package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//MAAAAAAAIN!!!
func main() {

	if len(os.Args) < 4 {
		fmt.Println()
		return
	}

	output := os.Args[3]
	if output[len(output)-4] != '.' {
		output += ".txt"
	}

	if os.Args[2] != "shadow" && os.Args[2] != "thinkertoy" && os.Args[2] != "standard" {
		fmt.Println("Banner type doesn't exist")
		return
	}

	font := os.Args[2] + ".txt"

	args0 := os.Args[1]

	f, err := os.Open(font)
	if err != nil {
		fmt.Println("file cannot be opened")
		return
	}

	defer f.Close()

	var args string

	args = args0

	var newStr string

	words := strings.Split(args, "\\n")

	for _, word := range words {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				newStr = newStr + ReadLine(f, 2+((int(letter)-32)*9)+i)
				//fmt.Print(ReadLine(f, 2+((int(letter)-32)*9)+i))
			}
			//fmt.Println()
			newStr = newStr + "\n"
		}
	}
	//fmt.Println(newStr)

	out, err := os.Create(output[9:])

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = out.WriteString(newStr)
	if err != nil {
		fmt.Println(err)
		out.Close()
		return
	}

	err = out.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ReadLine(f *os.File, n int) string {
	f.Seek(0, 0) //Устанавливает файловый указатель в определенную позицию
	bf := bufio.NewReader(f)
	var line string

	for lnum := 0; lnum < n; lnum++ {
		line, _ = bf.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")

	}
	return line
}
