package policy

import (
	"errors"
	"fmt"
)

type policyType string

const KONP_POLICY = policyType("KONP")
const NP_POLICY = policyType("NP")
const GPP_POLICY = policyType("GPP")
const LC_POLICY = policyType("LC")

const POLICY_NOT_FOUND_ERROR = errors.New("POLICY_NOT_FOUND_ERROR")

// New policy for buffer management
func New(name policyType) Policy {
	switch name {
	case KONP_POLICY:
		return KONP{Name: string(name)}
	case NP_POLICY:
		return NP{Name: string(name)}
	case GPP_POLICY:
		return GPP{Name: string(name)}
	case LC_POLICY:
		return LC{Name: string(name)}
	default:
		panic(fmt.Errorf("%w %s", POLICY_NOT_FOUND_ERROR, name))
	}
}
