package api

import (
	"database/sql"
	"net/http"
	"github.com/labstack/echo/v4"
)

// Response struct
type Response struct {
	id int
	body string
}

const (
	root string = "/responses"
)

// InitRoutes adds all routes to an Echo Group
func InitRoutes(g *echo.Group, db *sql.DB) {
	responsesGroup := g.Group(root)

	responsesGroup.GET("", func (ctx echo.Context) error {
		return onGetResponses(ctx, db)
	})

	// responsesGroup.POST("", func (ctx echo.Context) error {
	// 	return onPostResponse(ctx, db)
	// })
}

// GetResponses returns list of all Responses in db
func GetResponses(db *sql.DB) ([]*Response, error) {
	var responses []*Response

	rows, err := db.Query("SELECT * from Responses")

	for rows.Next() {
		var r Response
		err = rows.Scan(&r.id, &r.body)

		responses = append(responses, &r)
	}

	return responses, err
}

// AddResponse adds response document to mongo client.
// func AddResponse(client *mongo.Client, response *Response) (interface{}, error) {
// 	collection := client.Database(database).Collection(collection)
// 	result, err := collection.InsertOne(context.TODO(), response)

// 	return result.InsertedID, err
// }

func onGetResponses(ctx echo.Context, db *sql.DB) error {
	// TODO: get body and do filtering from there
	r, err := GetResponses(db)

	ifOnError(ctx, err)

	return ctx.JSON(http.StatusOK, r)
}

// func onPostResponse(ctx echo.Context, c *mongo.Client) error {
// 	r := new(Response)
// 	r.Response = "Golang"

// 	id, err := AddResponse(c, r)

// 	ifOnError(ctx, err)

// 	return ctx.JSON(http.StatusOK, id)
// }

func ifOnError(ctx echo.Context, err error) error {
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return nil
}
