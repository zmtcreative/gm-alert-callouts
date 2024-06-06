package main

import (
	"fmt"
	"regexp"
)

func main() {
  var regex = regexp.MustCompile("\\[!(?P<kind>[\\w]+)\\](?P<closed>-{0,1})(?P<open>\\+{0,1})($|\\s+(?P<title>.*))")

  bad_line := "> [!info]stuff"

  ok_line := "> [!info] stuff"
  ok_line_2 := "> [!info]"
  ok_line_3 := `> [!info]
`

  lines := []string{ok_line, ok_line_2, ok_line_3, bad_line}

  for i := range lines {
    if ! regex.MatchString(lines[i]) {
      fmt.Println("Line %s didn't match", lines[i])
    }
  }
}
