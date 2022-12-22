package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	elf      int
	calories int
}

func main() {
	fmt.Println("advent day 1")

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	slice := make([]string, 0)

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		str1 := strings.Split(str, "\n")
		slice = append(slice, str1[0])
	}

	var calories int
	var elfNum int
	var prevMax int
	var elf Elf
	var summary []Elf
	var topThree int
	for _, cal := range slice {
		var c int

		if cal == "" {
			elfNum += 1

			elf := Elf{elfNum, calories}
			summary = append(summary, elf)
			calories = 0
			continue
		}
		c, err = strconv.Atoi(cal)
		calories += c

		if calories > prevMax {
			prevMax = calories

			elf.elf = elfNum + 1
			elf.calories = prevMax

		}

		if err != nil {
			panic(err)
		}

		sort.Slice(summary, func(i, j int) bool {
			return summary[i].calories > summary[j].calories
		})

	}
	topThree = summary[0].calories + summary[1].calories + summary[2].calories
	fmt.Println(summary)
	fmt.Println(elf)
	fmt.Println(prevMax)
	fmt.Println(topThree)
}
