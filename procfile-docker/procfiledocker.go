package procfiledocker

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func Main() {
  lines := readFile()
  commands := parseFile(lines)

  for _, command := range commands {
    fmt.Println(command)

    execute(command)
  }
}

func readFile() []string {
  data, err := os.ReadFile("./procfile-docker/Procfile.dev")

  if err != nil {
    panic(err)
  }

  fileContent := strings.Trim(string(data), "\n")

  lines := strings.Split(fileContent, "\n")

  return lines
}

func parseFile(lines []string) []Command {

  var commands []Command

  for _, line := range lines {
    _, err := isValidProcfileLine(strings.TrimSpace(line))

    if err != nil {
      panic(err)
    }

    command := parseLineToCommand(line)

    commands = append(commands, command)
  }

  return commands
}

func parseLineToCommand(line string) Command {
  tokens := strings.Split(line, ":")

  name := strings.TrimSpace(tokens[0])
  image := strings.TrimSpace(tokens[1])
  version := strings.TrimSpace(tokens[2])

  return Command{ Name: name, Image: image, Version: version }
}

func isValidProcfileLine(line string) (bool, error) {
  return regexp.MatchString("^([a-z]+):(\\s)+([a-z]+)(:?((\\d+)?)(-)?([a-z]+)?)?$", line)
}

type Command struct {
  Name string
  Image string
  Version string
}

func execute(command Command) {
  if imageOnDisk(command) {
    runImage(command)
  } else {
    downloadImageAndRun(command)
  }
}

func imageOnDisk(command Command) bool {
  cmd := exec.Command("docker", "images", command.Image + ":" + command.Version, "-q")

  out, err := cmd.Output()

  if err != nil {
    panic(err)
  }

  sizeOfResult := len(string(out))

  if sizeOfResult > 0 {
    return true
  } else {
    return false
  }
}

func runImage(command Command) {
  fmt.Println(command)
  // cmd := exec.Command("docker", "run", "-rm", command.Image + ":" + command.Version)

  // out, err := cmd.Output()

  // if err != nil {
  //   panic(err)
  // }

  // fmt.Println(string(out))
}

func downloadImageAndRun(command Command) {
  fmt.Println("Baixando imagem")
}
