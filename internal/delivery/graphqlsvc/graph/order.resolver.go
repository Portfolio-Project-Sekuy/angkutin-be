package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/delivery/graphqlsvc/graph/generated"
	"github.com/Portfolio-Project-Sekuy/angkutin-be/internal/model"
)

func (r *queryResolver) Order(ctx context.Context, id int64) (*model.Order, error) {
	return &model.Order{
		ID:                  123,
		SenderName:          "sender name",
		SenderPhoneNumber:   "sender number",
		ReceiverName:        "receiver name",
		ReceiverPhoneNumber: "receiver number",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
