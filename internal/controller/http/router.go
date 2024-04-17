package http

import (
	productApi "ivixlabs.com/goweb/internal/controller/http/handlers/api/product"
	"ivixlabs.com/goweb/internal/controller/http/handlers/api/property"
	userApi "ivixlabs.com/goweb/internal/controller/http/handlers/api/user"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web/dashboard"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web/front"
	productWeb "ivixlabs.com/goweb/internal/controller/http/handlers/web/product"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web/static"
	userWeb "ivixlabs.com/goweb/internal/controller/http/handlers/web/user"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web/video"
	"ivixlabs.com/goweb/internal/controller/http/middleware"
	"ivixlabs.com/goweb/internal/model/product/usecase"
	property2 "ivixlabs.com/goweb/internal/model/property"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	productModel "ivixlabs.com/goweb/internal/model/product"
	userUseCase "ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/validation/form"
)

func NewRouter(sessionStore sessions.Store, userService userUseCase.Service,
	formValidator *form.Validator,
	productCreation usecase.ProductCreation,
	productUpdating usecase.ProductUpdating,
	productRepository productModel.Repository,
	productService productModel.Service,
	staticDir string,
	developmentMode bool,
	propertyRepository property2.Repository,
) http.Handler {
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
		middleware.GretAuthMiddleware(productWeb.GetSaveHandler(formValidator, productRepository, productUpdating, productCreation))))
	router.Handle("/products/delete", middleware.GretPreloadMiddleware(
		middleware.GretAuthMiddleware(productWeb.GetDeleteHandler(productRepository))))
	router.Handle("/video", middleware.GretPreloadMiddleware(
		video.GetIndexHandler()))

	router.Handle("/video/room", video.GetRoomHandler())
	router.Handle("/video/room/ws", video.GetSignalHandler())
	router.Handle("/video/room/enter", video.GetEnterInRoomHandler())

	router.Handle("/api/user/list", middleware.GetCorsMiddleware(userApi.GetListHandler(userService)))

	router.Handle("/api/product/list", middleware.GetCorsMiddleware(productApi.GetListHandler(productService)))

	router.Handle("/dashboard", dashboard.GetDashboardHandler())
	router.Handle("/front1", front.GetFrontHandler())

	router.Handle("/api/property/create", middleware.GetCorsMiddleware(property.GetCreateHandler(propertyRepository, formValidator)))
	router.Handle("/api/property/delete", middleware.GetCorsMiddleware(property.GetDeleteHandler(propertyRepository)))
	router.Handle("/api/property/list", middleware.GetCorsMiddleware(property.GetListHandler(propertyRepository)))

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", static.GetFileHandler(developmentMode, staticDir)))

	return router
}
