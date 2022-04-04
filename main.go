package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var notepad []string
var counter = 0

func setupStorage() {
	var size int
	fmt.Print("Enter the maximum number of notes:")
	fmt.Scan(&size)
	notepad = make([]string, size)
}

func getInput() (command string, note string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter command and data: ")
	scanner.Scan()
	inputLine := scanner.Text()
	command = strings.Split(inputLine, " ")[0]
	note = strings.TrimPrefix(strings.TrimPrefix(inputLine, command), " ")
	return
}

func isEmpty(note string) bool {
	if strings.ReplaceAll(note, " ", "") == "" {
		return true
	} else {
		return false
	}
}

func createNote(note string) {
	if counter == len(notepad) {
		fmt.Print("[Error] Notepad is full\n")
		return
	}
	if isEmpty(note) {
		fmt.Print("[Error] Missing note argument\n")
		return
	}
	notepad[counter] = note
	counter++
	fmt.Print("[OK] The note was successfully created\n")
}

func clearNotepad() {
	notepad = make([]string, len(notepad))
	counter = 0
	fmt.Print("[OK] All notes were successfully deleted\n")
}

func listNotepad() {
	if counter == 0 {
		fmt.Print("[Info] Notepad is empty\n")
		return
	}
	for i, item := range notepad {
		if notepad[i] != "" {
			fmt.Printf("[Info] %d: %s\n", i+1, item)
		}
	}
}

func exitNotepad() {
	fmt.Print("[Info] Bye!\n")
	os.Exit(0)
}

func getPosition(line string) (pos int, note string) {
	position, err := strconv.ParseInt(strings.Split(line, " ")[0], 10, 0)
	if err == nil {
		pos = int(position)
		note = strings.TrimPrefix(strings.TrimPrefix(line, strconv.Itoa(int(pos))), " ")
	} else {
		fmt.Printf("[Error] Invalid position: %s\n", strings.Split(line, " ")[0])
	}
	return
}

func isCorrectPosition(position int, note string) bool {
	if position == 0 {

		return false
	} else if position > len(notepad) {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, len(notepad))
		return false
	}
	posExists := false
	for i, _ := range notepad {
		if i == position-1 && notepad[i] != "" {
			posExists = true
		}
	}
	if !posExists {
		fmt.Print("[Error] Missing position argument\n")
		return false
	}
	return true
}

func getArguments(line string) (position int, note string) {
	if isEmpty(line) {
		fmt.Printf("[Error] Missing position argument\n")
		return
	}
	position, note = getPosition(line)
	if isCorrectPosition(position, note) {
		return
	}
	return
}

func deleteNote(note string) {
	position, _ := getArguments(note)
	if position == 0 {
		return
	}
	for i, _ := range notepad {
		if i == int(position-1) {
			notepad = append(notepad[:i], notepad[i+1:]...)
			notepad = append(notepad, "")
			counter--
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
			break
		}
	}
}

func updateNote(line string) {
	position, note := getArguments(line)
	if position == 0 {
		return
	}
	if isEmpty(note) {
		fmt.Print("[Error] Missing note argument\n")
		return
	}
	notepad[position-1] = note
	fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
}

func main() {
	setupStorage()
	for {
		command, note := getInput()
		switch command {
		case "create":
			createNote(note)
		case "delete":
			deleteNote(note)
		case "update":
			updateNote(note)
		case "clear":
			clearNotepad()
		case "list":
			listNotepad()
		case "exit":
			exitNotepad()
		default:
			fmt.Print("[Error] Unknown command\n")
		}
	}
}
