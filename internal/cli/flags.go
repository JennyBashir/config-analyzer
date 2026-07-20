package cli

import (
	"flag"
	"fmt"
)

type Options struct {
	Silent bool
	Stdin  bool
	Path   string
}

func ParseFlags() (Options, error) {
	silent := flag.Bool("s", false, "не выходить с ошибкой")

	stdin := flag.Bool("stdin", false, "прочитать конфигурацию из стандартного потока ввода вместо файла")

	flag.Parse()

	args := flag.Args()

	var path string

	if !*stdin {
		if len(args) == 0 {
			return Options{}, fmt.Errorf("usage: checker <config-file>")
		}
		path = args[0]
	}

	return Options{
		Silent: *silent,
		Stdin:  *stdin,
		Path:   path,
	}, nil
}
