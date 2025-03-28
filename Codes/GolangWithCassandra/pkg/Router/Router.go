package router

import (
    "fmt"
    "log"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/username/projectname/pkg/Users"
)

func StartServer() *chi.Mux {
    router := chi.NewRouter()
    router.Mount("/api/users", Users.UsersRoutes())
    
    fmt.Println("Server is listening on port 8080.....")
    log.Fatal(http.ListenAndServe(":8080", router))
    
    return router
}
