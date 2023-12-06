package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func fetch_input(day string) string {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%v/input", day)
	cookieName := "session"
	cookieValue := os.Getenv("AUTH_TOKEN")

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookieName, cookieValue))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	return string(body)
}

func do_setup(day string) {
	input_file := fetch_input(day)
	prefix := fmt.Sprintf("day_%v", day)
	if err := os.Mkdir(prefix, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s/test.txt", prefix), []byte(""), 0644); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(fmt.Sprintf("%s/input.txt", prefix), []byte(input_file), 0644); err != nil {
		log.Fatal(err)
	}

	fileString := fmt.Sprintf(`package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}

func main() {
	data, err := os.ReadFile("%s/test.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Printf("Result for part1: %s\n", part1(lines))
	fmt.Printf("Result for part2: %s\n", part2(lines))
}`, prefix, "%d", "%d")

	d := []byte(fileString)

	if err := os.WriteFile(fmt.Sprintf("%s/%s.go", prefix, prefix), d, 0644); err != nil {
		log.Fatal(err)
	}
}

func do_execute(day string) {
	cmd := exec.Command("go", "run", filepath.Join(".", fmt.Sprintf("day_%s", day), fmt.Sprintf("day_%s.go", day)))
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		// Print stderr output
		fmt.Printf("Command stderr: \n%s\n", output)
		// Print exit status
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Exit status: %d\n", exitErr.ExitCode())
		}
		os.Exit(1)
	}

	// Print stdout output
	fmt.Printf("%s", output)
}

func main() {
	osArgs := os.Args[1:]

	for _, arg := range osArgs {
		switch {
		case strings.HasPrefix(arg, "--day="):
			day := arg[len("--day="):]

			if day == "" {
				log.Fatal("You need to pass in a day flag, --day=N")
			}

			if _, err := os.Stat(fmt.Sprintf("day_%v", day)); os.IsNotExist(err) {
				do_setup(day)
			} else {
				do_execute(day)
			}

		default:
			log.Fatal("You need to pass in a day flag, --day=N")
		}
	}
}
