package api

import (
	"github.com/vfn2002/hdt-response/database"

	"context"
)

// response struct
type Response struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Response int                `bson:"response" json:"response,omitempty"`
}

// GetResponses returns list of all Responses in db
func GetResponses(client *mongo.Client, filter bson.M) ([]*Response, error) 
	var responses []*Response
	return responses, err
}
