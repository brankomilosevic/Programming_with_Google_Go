package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var people []Name
	var file_name string

	fmt.Println("Enter file name:")
	fmt.Scanln(&file_name)

	f, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	my_read := bufio.NewReader(f)
	for {
		line, _, err := my_read.ReadLine()

		if err != nil || io.EOF == err {
			break
		}
		tokens := strings.Split(string(line), " ")
		tmp_name := Name{
			trim_to_20(tokens[0]),
			trim_to_20(tokens[1]),
		}
		people = append(people, tmp_name)
	}

	fmt.Println("#\tFirst Name:\tLast Name")
	for i := 0; i < len(people); i++ {
		fmt.Printf("%03d:\t%s\t%s\n", i+1, people[i].fname, people[i].lname)
	}
}

// trim the string if the name is longer than 20 characters (runes)
func trim_to_20(str string) string {

	if len(str) > 20 {
		return string(str[0:20])
	} else {
		return str
	}
}
