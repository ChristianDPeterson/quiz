package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	//
	filenamePtr := flag.String("filename", "problems.csv", "Quiz")
	seconds := flag.Int("time", 30, "Time allowed in seconds")
	flag.Parse()

	//
	f, err := os.Open(*filenamePtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	correct := 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Press enter to begin the quiz...")
	reader.ReadString('\n')

	timer := time.NewTimer(time.Second * time.Duration(*seconds))

	// when time runs out, print score and exit
	go func() {
		<-timer.C
		fmt.Println("Oh no, time ran out!")
		endGame(correct, len(data))
		os.Exit(0)
	}()

	for _, v := range data {
		question := v[0]
		answer := v[1]

		fmt.Println(question)

		text, _ := reader.ReadString('\n')

		text = strings.Trim(text, "\n")

		if strings.Compare(text, answer) == 0 {
			correct++
			fmt.Println("Right!")
		} else {
			fmt.Println("Wrong!")
		}
	}

	endGame(correct, len(data))
}

func endGame(correct int, total int) {
	fmt.Println("\n\n--- TOTAL ---")
	fmt.Printf("    %d / %d \n", correct, total)
}
