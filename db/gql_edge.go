// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (u *User) Friends(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedFriends(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.FriendsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryFriends().All(ctx)
	}
	return result, err
}

func (u *User) BestFriend(ctx context.Context) (*User, error) {
	result, err := u.Edges.BestFriendOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryBestFriend().Only(ctx)
	}
	return result, err
}

func (u *User) AdminGroups(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, where *UserGroupWhereInput,
) (*UserGroupConnection, error) {
	opts := []UserGroupPaginateOption{
		WithUserGroupFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[2][alias]
	if nodes, err := u.NamedAdminGroups(alias); err == nil || hasTotalCount {
		pager, err := newUserGroupPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &UserGroupConnection{Edges: []*UserGroupEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryAdminGroups().Paginate(ctx, after, first, before, last, opts...)
}

func (ug *UserGroup) Admin(ctx context.Context) (*User, error) {
	result, err := ug.Edges.AdminOrErr()
	if IsNotLoaded(err) {
		result, err = ug.QueryAdmin().Only(ctx)
	}
	return result, err
}
