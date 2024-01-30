package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/liamg/gifwrap/pkg/ascii"
)

func main() {
	//* GET csv file
	/// csvFileName is a pointer to a String
	csvFileName := flag.String("csv", "problems.csv", "Đưa ra tất cả câu hỏi trong file, format: 'problem,answer'. Ví dụ: -csv=\"abc.csv\"")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Không thể mở file: %s\n", *csvFileName))
		os.Exit(1)
	}

	//* READ csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Không thể đọc file")
	}

	//* FETCH csv file data to problems array
	problems := getProblems(lines)

	//* Show questions and get input answer
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Câu hỏi %d: %s\n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}
	point := float64(10) * float64(correct) / float64(len(problems))
	printRes(point)
}

func getProblems(lines [][]string) []problem {
	//Initialize an array of problems([problem]) type and the length of the array
	result := make([]problem, len(lines))
	for i, line := range lines {
		// A line : [5+5, 10]
		result[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return result
}

func printRes(point float64) {
	if point <= 5 {
		fmt.Printf("Bạn CHỈ được %.3f điểm. Quá kém! ( ͡° ͜ʖ ͡°)\n", point)
		return
	}
	if point > 5 && point < 8 {
		fmt.Printf("Bạn được %.3f điểm. Có cố gắng! (ง ͠° ͟ل͜ ͡°)ง\n", point)
		return
	}
	if point >= 8 && point < 10 {
		fmt.Printf("Bạn được %.3f điểm (°ロ°)☝. Được 10 là sẽ có phần thưởng siêu ngầu luôn nha! ( ⚆ _ ⚆ )\n", point)
	}
	if point == 10 {
		 displayGIF("dancing-gopher.gif")
	}
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

var enableFill bool
func displayGIF(gifFileName string) {
	var renderer *ascii.Renderer
	var err error

	renderer, err = ascii.FromFile(gifFileName)

	if err != nil {
		exit(fmt.Sprintln("Lỗi không thể lấy file phần thưởng :(("))
	}
	renderer.SetFill(enableFill)

	if err := renderer.PlayOnce(); err != nil && err != ascii.ErrQuit {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
