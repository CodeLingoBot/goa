// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/cellar/design -o
// $(GOPATH)/src/goa.design/goa/examples/cellar

package client

import (
	"encoding/json"
	"fmt"

	sommelierpb "goa.design/goa/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/goa/examples/cellar/gen/sommelier"
)

// BuildPickPayload builds the payload for the sommelier pick endpoint from CLI
// flags.
func BuildPickPayload(sommelierPickMessage string) (*sommelier.Criteria, error) {
	var err error
	var message sommelierpb.PickRequest
	{
		if sommelierPickMessage != "" {
			err = json.Unmarshal([]byte(sommelierPickMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"name\": \"Blue\\'s Cuvee\",\n      \"varietal\": [\n         \"pinot noir\",\n         \"merlot\",\n         \"cabernet franc\"\n      ],\n      \"winery\": \"longoria\"\n   }'")
			}
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &sommelier.Criteria{
		Name:   &message.Name,
		Winery: &message.Winery,
	}
	if message.Varietal != nil {
		payload.Varietal = make([]string, len(message.Varietal))
		for i, val := range message.Varietal {
			payload.Varietal[i] = val
		}
	}
	return payload, nil
}
