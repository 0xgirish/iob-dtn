package policy

import (
	"errors"
	"fmt"
)

const POLICY_NOT_FOUND_ERROR = errors.New("POLICY_NOT_FOUND_ERROR")

// New policy for buffer management
func New(name string) Policy {
	switch name {
	case "KNOP":
		return KNOP{Name: name}
	case "NP":
		return NP{Name: name}
	case "GPP":
		return GPP{Name: name}
	case "LC":
		return LC{Name: name}
	default:
		panic(fmt.Errorf("%w %s", POLICY_NOT_FOUND_ERROR, name))
	}
}
