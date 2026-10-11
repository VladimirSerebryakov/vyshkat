package console

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"vyshkat/internal/park"
	"vyshkat/internal/reports"
)

func Run(reception *park.Reception, p *park.Park, svc *reports.Service, in io.Reader, out io.Writer, showPrompt bool) error {
	s := &session{reception: reception, p: p, reports: svc, out: out}

	fmt.Fprintln(out, "ВышКат: учёт парка. Введите help для списка команд.")

	scanner := bufio.NewScanner(in)
	for {
		if showPrompt {
			fmt.Fprint(out, "> ")
		}
		if !scanner.Scan() {
			return scanner.Err()
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		args, err := splitArgs(line)
		if err != nil {
			fmt.Fprintf(out, "ошибка: %v\n", err)
			continue
		}

		exit, err := s.execute(args)
		if err != nil {
			fmt.Fprintf(out, "ошибка: %v\n", err)
		}
		if exit {
			return nil
		}
	}
}

// splitArgs делит строку по пробелам, но текст в двойных кавычках считает одним аргументом
func splitArgs(line string) ([]string, error) {
	var args []string
	var current strings.Builder
	inQuotes := false
	hasToken := false

	for _, r := range line {
		switch {
		case r == '"':
			inQuotes = !inQuotes
			hasToken = true
		case (r == ' ' || r == '\t') && !inQuotes:
			if hasToken {
				args = append(args, current.String())
				current.Reset()
				hasToken = false
			}
		default:
			current.WriteRune(r)
			hasToken = true
		}
	}

	if inQuotes {
		return nil, fmt.Errorf("незакрытая кавычка")
	}
	if hasToken {
		args = append(args, current.String())
	}
	return args, nil
}
