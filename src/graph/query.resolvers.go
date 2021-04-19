package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/aodpi/hrm-go-graphql/v2/config"
	"github.com/aodpi/hrm-go-graphql/v2/graph/generated"
	"github.com/aodpi/hrm-go-graphql/v2/graph/model"
	pb "github.com/aodpi/hrm-go-graphql/v2/grpc/generated/grpc/proto"
	"google.golang.org/grpc"
)

func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	config := config.GetConfig()

	conn, err := grpc.Dial(config.GetString("hrm.grpc.url"), grpc.WithInsecure())

	if err != nil {
		log.Printf("did not connect: %v", err)
		return nil, err
	}

	defer conn.Close()

	c := pb.NewTagsGServiceClient(conn)

	response, err := c.GetTags(context.Background(), &pb.GetTagsRequest{})

	if err != nil {
		log.Printf("Could not get tags: %v", err)
		return nil, err
	}

	var tags []*model.Tag

	for _, s := range response.Tags {
		tags = append(tags, &model.Tag{
			ID:   s.Id,
			Name: s.Name,
		})
	}

	return tags, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
