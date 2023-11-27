package resolvers

import (
	"context"
	"errors"

	"projectname/db"
)

type Query struct{ Client *db.Client }

func (q *Query) Node(ctx context.Context, id string) (db.Noder, error) {
	// turn our string into the globally unique id from ent
	target, err := FromGlobalID(id)
	if err != nil {
		return nil, errors.New("You provided an invalid id")
	}
	node, err := q.Client.Noder(ctx, target.ID)
	if err != nil {
		return nil, errors.New("unable to find node")
	}

	return node, nil
}

func (q *Query) Nodes(ctx context.Context, ids []string) ([]db.Noder, error) {
	globalIDs := []int{}

	// turn our string into the globally unique id from ent
	for _, id := range ids {
		target, err := FromGlobalID(id)
		if err != nil {
			return nil, errors.New("You provided an invalid id")
		}

		globalIDs = append(globalIDs, target.ID)

	}

	node, err := q.Client.Noders(ctx, globalIDs)
	if err != nil {
		return nil, errors.New("unable to find node")
	}

	return node, nil
}
