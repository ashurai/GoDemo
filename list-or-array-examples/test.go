package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("website_text_logs_00.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	formatString := make(map[string]int)

	for scanner.Scan() {
		linetext := strings.Split(scanner.Text(), (" "))
		LineKey := strings.Replace(linetext[0], "[", "", -1)
		formatString[LineKey] += 0

		_, exist := formatString[LineKey]

		if exist {
			formatString[LineKey] += 1
		} else {
			formatString[LineKey] = 1
		}
	}

	w, err := os.Create("output.log")
	if err != nil {
		fmt.Println(err)
	}

	for key := range formatString {
		newline := key + " " + fmt.Sprint(formatString[key]) + "\r\n"
		_, err := fmt.Fprintln(w, newline)
		//fmt.Println(key, formatString[key])
		if err != nil {
			fmt.Println(err)
			w.Close()
		}
	}

	err = w.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file appended successfully")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
