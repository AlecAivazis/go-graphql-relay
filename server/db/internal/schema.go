// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"projectname/db/schema","Package":"projectname/db","Schemas":[{"name":"User","config":{"Table":""},"edges":[{"name":"friends","type":"User","ref":{"name":"bestFriend","type":"User","unique":true,"required":true},"inverse":true},{"name":"adminGroups","type":"UserGroup","ref_name":"admin","inverse":true,"annotations":{"EntGQL":{"Mapping":["adminGroups"],"RelayConnection":true,"Unbind":true}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"annotations":{"EntGQL":{"QueryField":{},"RelayConnection":true}}},{"name":"UserGroup","config":{"Table":""},"edges":[{"name":"admin","type":"User","unique":true,"required":true}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"annotations":{"EntGQL":{"QueryField":{},"RelayConnection":true}}}],"Features":["schema/snapshot","sql/upsert","intercept","namedges"]}`
