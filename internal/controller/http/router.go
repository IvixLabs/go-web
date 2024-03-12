package http

import (
	userApi "ivixlabs.com/goweb/internal/controller/http/handlers/api/user"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web"
	productWeb "ivixlabs.com/goweb/internal/controller/http/handlers/web/product"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web/static"
	userWeb "ivixlabs.com/goweb/internal/controller/http/handlers/web/user"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web/video"
	"ivixlabs.com/goweb/internal/controller/http/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	productUseCase "ivixlabs.com/goweb/internal/product"
	userUseCase "ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/validation/form"
)

func NewRouter(sessionStore sessions.Store, userService userUseCase.Service,
	formValidator *form.Validator, productService productUseCase.Service, staticDir string, developmentMode bool) http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.GetContextMiddleware(sessionStore))

	router.Handle("/", middleware.GretPreloadMiddleware(web.GetHomeHandler()))

	router.Handle("/signup", middleware.GretPreloadMiddleware(
		userWeb.GetSignupHandler(userService, formValidator)))
	router.Handle("/users", middleware.GretPreloadMiddleware(
		middleware.GretAuthMiddleware(userWeb.GetListHandler(userService))))
	router.Handle("/auth", middleware.GretPreloadMiddleware(
		userWeb.GetAuthHandler(userService, formValidator)))
	router.Handle("/logout", userWeb.GetLogoutHandler())
	router.Handle("/products", middleware.GretPreloadMiddleware(
		middleware.GretAuthMiddleware(productWeb.GetListHandler(productService))))
	router.Handle("/products/form", middleware.GretPreloadMiddleware(
		middleware.GretAuthMiddleware(productWeb.GetSaveHandler(formValidator, productService))))
	router.Handle("/products/delete", middleware.GretPreloadMiddleware(
		middleware.GretAuthMiddleware(productWeb.GetDeleteHandler(productService))))
	router.Handle("/video", middleware.GretPreloadMiddleware(
		video.GetIndexHandler()))

	router.Handle("/video/room", video.GetRoomHandler())
	router.Handle("/video/room/ws", video.GetSignalHandler())
	router.Handle("/video/room/enter", video.GetEnterInRoomHandler())

	router.Handle("/api/user/list", userApi.GetListHandler(userService))

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", static.GetFileHandler(developmentMode, staticDir)))

	return router
}
