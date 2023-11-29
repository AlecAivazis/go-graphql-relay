package resolvers

import (
	"context"
	"errors"

	"projectname/db"

	"entgo.io/contrib/entgql"
)

type Query struct{ Client *db.Client }

func (q *Query) Users(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *db.UserWhereInput) (*db.UserConnection, error) {
	return q.Client.User.Query().Paginate(ctx, after, first, before, last, db.WithUserFilter(where.Filter))
}

func (q *Query) UserGroups(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *db.UserGroupWhereInput) (*db.UserGroupConnection, error) {
	return q.Client.UserGroup.Query().Paginate(ctx, after, first, before, last, db.WithUserGroupFilter(where.Filter))
}

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
