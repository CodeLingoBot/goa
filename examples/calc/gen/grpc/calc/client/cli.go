// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package client

import (
	"encoding/json"
	"fmt"

	calcsvc "goa.design/goa/examples/calc/gen/calc"
	calcpb "goa.design/goa/examples/calc/gen/grpc/calc"
)

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddP string) (*calcsvc.AddPayload, error) {
	var err error
	var p calcpb.AddRequest
	{
		if calcAddP != "" {
			err = json.Unmarshal([]byte(calcAddP), &p)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for p, example of valid JSON:\n%s", "'{\n      \"a\": 8399553735696626949,\n      \"b\": 360622074634248926\n   }'")
			}
		}
	}
	if err != nil {
		return nil, err
	}
	v := &calcsvc.AddPayload{
		A: int(p.A),
		B: int(p.B),
	}
	return v, nil
}
