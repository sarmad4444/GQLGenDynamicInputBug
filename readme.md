### todo app

This is the simplest example of a graphql server.

to run this server
```bash
go run ./server/server.go
```

and open http://localhost:8081 in your browser


### Problem
Running `go run github.com/99designs/gqlgen` skips generating slice for `Map` types.

Check out `TodoCreateInput` in `schema.graphql` and the specific resolver in `resolvers/schema.resolvers.go` -> `CreateTodoMany()` 

The expected resolver should be
```go
func (r *mutationsResolver) CreateTodoMany(ctx context.Context, inputs []map[string]interface{}) ([]*todo.Todo, error)
```


But currently GQLGen is generating incorrect type for `inputs` parameter:
```go
func (r *mutationsResolver) CreateTodoMany(ctx context.Context, inputs map[string]interface{}) ([]*todo.Todo, error)
```