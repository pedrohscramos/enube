package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/pedrohscramos/enube/configs"
	_ "github.com/pedrohscramos/enube/docs"
	"github.com/pedrohscramos/enube/internal/entity"
	"github.com/pedrohscramos/enube/internal/infra/database"
	"github.com/pedrohscramos/enube/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           API Enube
// @version         1.0
// @description     Api para gerenciamento de fornecedores
// @termsOfService  http://swagger.io/terms/

// @contact.name   Pedro ramos
// @contact.email  pedrohscramos@gmail.com

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("enube.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{})

	supplierDB := database.NewSupplier(db)
	supplierHandler := handlers.NewSupplierHandler(supplierDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JwtExpiresIn))

	r.Route("/suppliers", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", supplierHandler.GetSuppliers)
		r.Get("/{id}", supplierHandler.GetSupplier)

	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":"+configs.WebServerPort, r)

}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
