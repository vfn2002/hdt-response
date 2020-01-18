package api

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Response struct
type Response struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Response string             `bson:"response" json:"response,omitempty"`
}

const (
	database string = "responses"
	collection string = "v1"
	root string = "/responses"
)

// InitRoutes adds all routes to an Echo Group
func InitRoutes(g *echo.Group, c *mongo.Client) {
	responsesGroup := g.Group(root)

	responsesGroup.GET("", func (ctx echo.Context) error {
		return onGetResponses(ctx, c)
	})

	responsesGroup.POST("", func (ctx echo.Context) error {
		return onPostResponse(ctx, c)
	})
}

// GetResponses returns list of all Responses in db
func GetResponses(client *mongo.Client, filter bson.M) ([]*Response, error) {
	var responses []*Response
	col := client.Database(database).Collection(collection)
	cur, err := col.Find(context.TODO(), filter)

	if err != nil {
		for cur.Next(context.TODO()) {
			var response Response
			cur.Decode(&response)
			responses = append(responses, &response)
		}
	}

	return responses, err
}

// AddResponse adds response document to mongo client.
func AddResponse(client *mongo.Client, response *Response) (interface{}, error) {
	collection := client.Database(database).Collection(collection)
	result, err := collection.InsertOne(context.TODO(), response)

	return result.InsertedID, err
}

func onGetResponses(ctx echo.Context, c *mongo.Client) error {
	// TODO: get body and do filtering from there
	r, err := GetResponses(c, bson.M{})

	ifOnError(ctx, err)

	return ctx.JSON(http.StatusOK, r)
}

func onPostResponse(ctx echo.Context, c *mongo.Client) error {
	r := new(Response)
	r.Response = "Golang"

	id, err := AddResponse(c, r)

	ifOnError(ctx, err)

	return ctx.JSON(http.StatusOK, id)
}

func ifOnError(ctx echo.Context, err error) error {
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return nil
}
