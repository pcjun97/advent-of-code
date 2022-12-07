package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	commands := ParseCommands(input)
	root := NewDirectory("/", nil)

	var cur *Directory
	for _, command := range commands {
		switch command.Name {
		case "ls":
			for _, item := range command.Output {
				cur.ParseContent(item)
			}
		case "cd":
			switch command.Arg {
			case "/":
				cur = root
			case "..":
				cur = cur.Parent
			default:
				if _, ok := cur.Directories[command.Arg]; !ok {
					cur.Directories[command.Arg] = NewDirectory(command.Arg, cur)
				}
				cur = cur.Directories[command.Arg]
			}
		}
	}

	sizes := root.SizeRecurse()
	sum := 0
	diff := sizes["/"] - (70000000 - 30000000)
	min := sizes["/"]

	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}

		if size >= diff && size < min {
			min = size
		}
	}

	fmt.Println(sum)
	fmt.Println(min)
}

type Directory struct {
	Name        string
	Parent      *Directory
	Directories map[string]*Directory
	Files       map[string]*File
}

func NewDirectory(name string, parent *Directory) *Directory {
	d := Directory{
		Name:        name,
		Parent:      parent,
		Directories: make(map[string]*Directory),
		Files:       make(map[string]*File),
	}
	return &d
}

func (d *Directory) ParseContent(item string) {
	fields := strings.Split(item, " ")
	name := fields[1]

	if fields[0] == "dir" {
		d.Directories[name] = NewDirectory(name, d)
		return
	}

	size, err := strconv.Atoi(fields[0])
	if err != nil {
		panic("error parsing file size: " + fields[0])
	}

	d.Files[name] = NewFile(name, size)
}

func (d *Directory) Path() string {
	if d.Name == "/" {
		return "/"
	}
	return strings.TrimRight(d.Parent.Path(), "/") + "/" + d.Name
}

func (d *Directory) SizeRecurse() map[string]int {
	sizes := make(map[string]int)

	sum := 0
	for _, child := range d.Directories {
		for path, size := range child.SizeRecurse() {
			sizes[path] = size
		}
		sum += sizes[child.Path()]
	}

	for _, f := range d.Files {
		sum += f.Size
	}

	sizes[d.Path()] = sum

	return sizes
}

type File struct {
	Name string
	Size int
}

func NewFile(name string, size int) *File {
	f := File{
		Name: name,
		Size: size,
	}
	return &f
}

type Command struct {
	Name   string
	Arg    string
	Output []string
}

func ParseCommands(s string) []Command {
	commandsRaw := strings.Split(strings.TrimSpace(s), "$")
	commands := make([]Command, len(commandsRaw))

	for _, commandRaw := range commandsRaw {
		lines := strings.Split(strings.TrimSpace(commandRaw), "\n")
		fields := strings.Split(lines[0], " ")

		name := fields[0]
		arg := ""
		if len(fields) > 1 {
			arg = fields[1]
		}

		output := lines[1:]
		for i, line := range output {
			output[i] = strings.TrimSpace(line)
		}

		command := Command{
			Name:   name,
			Arg:    arg,
			Output: output,
		}
		commands = append(commands, command)
	}

	return commands
}
