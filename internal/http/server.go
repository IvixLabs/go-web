package http

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	httpController "ivixlabs.com/goweb/internal/controller/http"
	"ivixlabs.com/goweb/internal/gorm"
	"ivixlabs.com/goweb/internal/model"
	product2 "ivixlabs.com/goweb/internal/product"
	user2 "ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/validation/form"
	user3 "ivixlabs.com/goweb/internal/validation/user"
)

func StartServer(addr string, staticDir string, dbUrl string, sessionsDir string) {
	sessionStore := sessions.NewFilesystemStore(sessionsDir, []byte("abc123"))

	gormDb := gorm.NewGormDb(dbUrl)

	model.GormInitModels(gormDb)

	userRepository := model.NewGormUserRepository(gormDb)
	userService := user2.NewService(userRepository)

	productRepository := model.NewGormProductRepository(gormDb)
	productService := product2.NewService(productRepository)

	formValidator := form.NewValidator()
	user3.InitEmailValidation(formValidator, userService)

	router := httpController.NewRouter(sessionStore,
		userService, formValidator, productService, staticDir)

	server := http.Server{Addr: addr, Handler: router}

	err := server.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}
