package resolvers

import (
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph"
)

type Resolver struct {
	LearningClient learning.LearningServiceClient
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
