package http

import (
	productApi "ivixlabs.com/goweb/internal/controller/http/handlers/api/product"
	userApi "ivixlabs.com/goweb/internal/controller/http/handlers/api/user"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web/dashboard"
	productWeb "ivixlabs.com/goweb/internal/controller/http/handlers/web/product"
	"ivixlabs.com/goweb/internal/controller/http/handlers/web/static"
	userWeb "ivixlabs.com/goweb/internal/controller/http/handlers/web/user"
	"ivixlabs.com/goweb/internal/controller/http/middleware"
	"ivixlabs.com/goweb/internal/model/product/usecase"
	"ivixlabs.com/goweb/internal/model/user"
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
	userRepository user.Repository,
) http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.GetContextMiddleware(sessionStore))

	router.Handle("/",
		middleware.GretPreloadMiddleware(
			web.GetHomeHandler(),
		),
	)

	router.Handle("/signup",
		middleware.GretPreloadMiddleware(
			userWeb.GetSignupHandler(userService, formValidator),
		),
	)

	router.Handle("/auth",
		middleware.GretPreloadMiddleware(
			userWeb.GetAuthHandler(userService, formValidator),
		),
	)

	router.Handle("/logout", userWeb.GetLogoutHandler())

	router.Handle("/products",
		middleware.GretPreloadMiddleware(
			middleware.GretAuthMiddleware(
				productWeb.GetListHandler(productService),
			),
		),
	)

	router.Handle("/products/form",
		middleware.GretPreloadMiddleware(
			middleware.GretAuthMiddleware(
				productWeb.GetSaveHandler(formValidator, productRepository, productUpdating, productCreation),
			),
		),
	)

	router.Handle("/products/delete",
		middleware.GretPreloadMiddleware(
			middleware.GretAuthMiddleware(
				productWeb.GetDeleteHandler(productRepository),
			),
		),
	)

	//router.Handle("/video",
	//	middleware.GretPreloadMiddleware(
	//		video.GetIndexHandler(),
	//	),
	//)

	//router.Handle("/video/room", video.GetRoomHandler())
	//router.Handle("/video/room/ws", video.GetSignalHandler())
	//router.Handle("/video/room/enter", video.GetEnterInRoomHandler())

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.GetBasicAuthMiddleware())
	apiRouter.Use(middleware.GetCorsMiddleware())

	userRouter := apiRouter.PathPrefix("/users").Subrouter()

	userRouter.Handle("/", userApi.GetListHandler(userService)).Methods("GET")
	userRouter.Handle("/{id}", userApi.GetGetHandler(userRepository)).Methods("GET")
	userRouter.Handle("/{id}", userApi.GetUpdateHandler(userRepository, formValidator)).Methods("PUT")
	userRouter.Handle("/", userApi.GetCreateHandler(userRepository, formValidator)).Methods("POST")
	userRouter.Handle("/{id}", userApi.GetDeleteHandler(userRepository)).Methods("DELETE")

	productRouter := apiRouter.PathPrefix("/products").Subrouter()
	productRouter.Handle("/", productApi.GetListHandler(productService)).Methods("GET")
	productRouter.Handle("/{id}", productApi.GetGetHandler(productRepository)).Methods("GET")
	productRouter.Handle("/{id}", productApi.GetDeleteHandler(productRepository)).Methods("DELETE")
	productRouter.Handle("/", productApi.GetCreateHandler(productRepository, formValidator)).Methods("POST")
	productRouter.Handle("/{id}", productApi.GetUpdateHandler(productRepository, formValidator)).Methods("PUT")

	dashboardHandler := middleware.GetBasicAuthHandler(
		dashboard.GetDashboardHandler(),
	)

	router.Handle("/dashboard", dashboardHandler)
	router.Handle("/dashboard/", dashboardHandler)
	router.Handle("/dashboard/{webapp}", dashboardHandler)

	//router.Handle("/front1", front.GetFrontHandler())
	//router.Handle("/api/property/create", middleware.GetCorsHandler(property.GetUpdateHandler(propertyRepository, formValidator)))
	//router.Handle("/api/property/delete", middleware.GetCorsHandler(property.GetDeleteHandler(propertyRepository)))
	//router.Handle("/api/property/list", middleware.GetCorsHandler(property.GetListHandler(propertyRepository)))

	router.
		PathPrefix("/static/").
		Handler(
			http.StripPrefix("/static/", static.GetFileHandler(developmentMode, staticDir)),
		)

	return router
}
