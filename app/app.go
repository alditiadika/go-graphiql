package app

import (
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/alditiadika/go-graphiql/app/graphql/schema"
	typedefs "github.com/alditiadika/go-graphiql/app/graphql/type-defs"
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
)

const port = ":3000"

//App struct
type App struct{}

//Init App
func (app App) Init() {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphql", gqlHandler())
	http.Handle("/graphiql", graphiqlHandler)
}

//Run func
func (app App) Run() {
	fmt.Println("server run at", port)
	http.ListenAndServe(port, nil)
}

func gqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No query data", 400)
			return
		}

		var rBody typedefs.RootType
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		fmt.Fprintf(w, "%s", processQuery(rBody.Query))

	})
}

func processQuery(query string) (result string) {
	params := graphql.Params{Schema: schema.InitGraphQLSchema(), RequestString: query}
	r := graphql.Do(params)
	rJSON, _ := json.Marshal(r)
	return fmt.Sprintf("%s", rJSON)

}
