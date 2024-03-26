package dev08

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Command contains command name and arguments
type Command struct {
	Name string
	Args []string
}

// TODO string returns

// Exec runs command with args
func (c *Command) Exec() error {
	switch c.Name {
	case "cd":
		var path string
		var err error
		if strings.HasPrefix(c.Args[0], "~") {
			path, err = os.UserHomeDir()
			if err != nil {
				return err
			}
			path = path + c.Args[0][len("~"):]
		} else {
			path = c.Args[0]
		}
		return os.Chdir(path)
	case "pwd":
		path, err := os.Getwd()
		if err != nil {
			return nil
		}
		fmt.Println(path)
		return nil
	case "echo":
		fmt.Println(c.Args)
		return nil
	case "kill":
		killList := make([]*os.Process, len(c.Args))
		for i := range killList {
			pid, err := strconv.Atoi(c.Args[i])
			if err != nil {
				return fmt.Errorf("procces with %s pid number doesnt exist", c.Args[i])
			}
			p, err := os.FindProcess(pid)
			if err != nil {
				return err
			}
			killList[i] = p
		}
		for _, ps := range killList {
			err := ps.Kill()
			if err != nil {
				fmt.Printf("impossible to kill procces with %d pid number\n", ps.Pid)
			}
		}
		return nil
	case "ps":
		ps := exec.Command("ps")
		bytes, err := ps.Output()
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
		return nil
	default:
		return errors.New("command is not find")
	}
}

// ParsePipes returns list of user commands
func ParsePipes(command string) ([]Command, error) {
	var res []Command
	pipes := strings.Split(command, "|")
	for _, pipe := range pipes {
		pipe = strings.TrimSpace(pipe)
		temp := strings.Split(pipe, " ")
		cmd := Command{Name: temp[0]}
		if len(temp) > 1 {
			cmd.Args = append(cmd.Args, temp[1:]...)
		}
		if !isCommandCorrect(cmd) {
			return res, errors.New("command is not valid")
		}
		res = append(res, cmd)
	}
	return res, nil
}

func isCommandCorrect(cmd Command) bool {
	return ((cmd.Name == "cd" || cmd.Name == "echo" || cmd.Name == "kill") && len(cmd.Args) > 0) || (cmd.Name == "pwd" || cmd.Name == "ps")
}

func getCurrentPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return path, err
	}
	homePath, err := os.UserHomeDir()
	if err != nil {
		return path, err
	}
	if strings.HasPrefix(path, homePath) {
		return "~" + path[len(homePath):], nil
	}
	return path, nil
}

// PrintPrefix prints a shell path
func PrintPrefix() error {
	path, err := getCurrentPath()
	if err != nil {
		return err
	}
	fmt.Printf("%s$ ", path)
	return nil
}
