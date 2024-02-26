package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
	"ivixlabs.com/goweb/api/handlers"
	"ivixlabs.com/goweb/api/handlers/product"
	"ivixlabs.com/goweb/api/handlers/user"
	"ivixlabs.com/goweb/api/middleware"
	product2 "ivixlabs.com/goweb/internal/product"
	user2 "ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/validation/form"
	user3 "ivixlabs.com/goweb/internal/validation/user"
)

func StartServer(addr string, staticDir string, db *gorm.DB, sessionsDir string) {
	var sessionStore = sessions.NewFilesystemStore(sessionsDir, []byte("abc123"))

	userService := user2.NewService(db)
	productService := product2.NewService(db)

	formValidator := form.NewValidator()
	user3.InitEmailValidation(formValidator, userService)

	router := mux.NewRouter()

	router.Use(middleware.GetContextMiddleware(sessionStore))

	router.Handle("/", handlers.GetHomeHandler())

	router.HandleFunc("/signup", user.GetSignupHandlerFunc(userService, formValidator))
	router.HandleFunc("/users", middleware.GretAuthMiddleware(user.GetListHandlerFunc(userService), sessionStore))
	router.HandleFunc("/auth", user.GetAuthHandlerFunc(userService, formValidator, sessionStore))
	router.HandleFunc("/logout", user.GetLogoutHandlerFunc())
	router.Handle("/products", middleware.GretAuthMiddleware(product.GetListHandler(productService).ServeHTTP, sessionStore))
	router.Handle("/products/form", middleware.GretAuthMiddleware(product.GetSaveHandler(formValidator, productService).ServeHTTP, sessionStore))
	router.Handle("/products/delete", middleware.GretAuthMiddleware(product.GetDeleteHandler(productService).ServeHTTP, sessionStore))

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	server := http.Server{Addr: addr, Handler: router}

	err := server.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}
