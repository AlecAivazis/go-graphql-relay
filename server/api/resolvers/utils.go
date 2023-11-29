package resolvers

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type GlobalID struct {
	Type string
	ID   int
}

func ToGlobalID(id GlobalID) (string, error) {
	payload := fmt.Sprintf("%s:%v", id.Type, id.ID)
	// encode the payload
	return base64.StdEncoding.EncodeToString([]byte(payload)), nil
}

func FromGlobalID(from string) (GlobalID, error) {
	// assume the payload is a base64 encoded string
	decoded, err := base64.StdEncoding.DecodeString(from)
	if err != nil {
		return GlobalID{}, err
	}

	// assume the two values are separated by :
	parts := strings.Split(string(decoded), ":")

	// the id needs to be an int
	id, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return GlobalID{}, err
	}

	return GlobalID{
		ID:   int(id),
		Type: parts[0],
	}, nil
}
