package walker

import (
	"strconv"
)

func Walk(path string, value any, visit func(path string, value any)) {

	visit(path, value)

	switch v := value.(type) {

	case map[string]any:
		for key, child := range v {
			var nextPath string

			if path == "" {
				nextPath = key
			} else {
				nextPath = path + "." + key
			}
			Walk(nextPath, child, visit)
		}

	case []any:
		for ind, val := range v {
			var nextPath string
			if path == "" {
				nextPath = "[" + strconv.Itoa(ind) + "]"
			} else {
				nextPath = path + "[" + strconv.Itoa(ind) + "]"
			}
			Walk(nextPath, val, visit)
		}
	default:

	}
}
