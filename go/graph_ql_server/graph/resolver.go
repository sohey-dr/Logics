//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"graph_ql_server/graph/model"
)

type Resolver struct{
	todos []*model.Todo
}
