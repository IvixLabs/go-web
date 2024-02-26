package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"ivixlabs.com/goweb/api/handlers"
	"ivixlabs.com/goweb/api/handlers/product"
	"ivixlabs.com/goweb/api/handlers/user"
	"ivixlabs.com/goweb/api/middleware"
	productUseCase "ivixlabs.com/goweb/internal/product"
	userUseCase "ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/validation/form"
)

func NewRouter(sessionStore sessions.Store, userService userUseCase.Service,
	formValidator *form.Validator, productService productUseCase.Service, staticDir string) http.Handler {
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

	return router
}
