package handlers

import (
	"github.com/radyatamaa/loyalti-go-echo/src/api/graphql"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"log"
	"net/http"
)
type Server struct {
	GqlSchema *graphql.Schema
}
type ReqBody struct {
	Query string `json:"query"`
}

func QuerySong(c echo.Context) error {

	var rBody ReqBody
	// Decode the request body into rBody
	err := json.NewDecoder(c.Request().Body).Decode(&rBody)
	if err != nil {
		//http.Error(w, "Error parsing JSON request body", 400)
		log.Printf("Failed processing  request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Execute graphql query
	result := graphQL.ExecuteQuery(rBody.Query)

	return c.JSON(http.StatusOK, result)

}


