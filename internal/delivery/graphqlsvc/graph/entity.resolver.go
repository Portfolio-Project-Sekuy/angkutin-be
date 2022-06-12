package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/delivery/graphqlsvc/graph/generated"
	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/model"
)

func (r *entityResolver) FindOrderByID(ctx context.Context, id int64) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
