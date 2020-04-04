package types

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalUint(i uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf("%d", i))
	})
}

func UnmarshalUint(v interface{}) (uint, error) {
	switch v := v.(type) {
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("string failed to be parsed: %v", err)
		} else {
			return uint(i), nil
		}

	case int:
		return uint(v), nil
	case int64:
		return uint(v), nil
	case json.Number:
		i, err := strconv.Atoi(string(v))
		if err != nil {
			return 0, fmt.Errorf("json.Number failed to be parsed: %v", err)
		} else {
			return uint(i), nil
		}
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
