package policy

import (
	"errors"
	"fmt"
)

type policyType string

var KONP_POLICY = policyType("KONP")
var NP_POLICY = policyType("NP")
var GPP_POLICY = policyType("GPP")
var LC_POLICY = policyType("LC")

var POLICY_NOT_FOUND_ERROR = errors.New("POLICY_NOT_FOUND_ERROR")

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
