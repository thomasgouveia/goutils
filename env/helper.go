package env

import (
	"fmt"
	"strconv"
)

func parseFloat(v string, bitSize int) (float64, error) {
	parsed, err := strconv.ParseFloat(v, bitSize)
	if err != nil {
		return -1, fmt.Errorf("could not parse float%d: %w", bitSize, err)
	}
	return parsed, nil
}

func parseInt(v string, base, bitSize int) (int64, error) {
	parsed, err := strconv.ParseInt(v, base, bitSize)
	if err != nil {
		return -1, fmt.Errorf("could not parse int%d: %w", bitSize, err)
	}
	return parsed, nil
}
