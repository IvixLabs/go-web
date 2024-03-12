package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/sessions"
	httpController "ivixlabs.com/goweb/internal/controller/http"
	"ivixlabs.com/goweb/internal/gorm"
	internalHttp "ivixlabs.com/goweb/internal/http"
	"ivixlabs.com/goweb/internal/model"
	"ivixlabs.com/goweb/internal/product"
	"ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/validation/form"
	userValidation "ivixlabs.com/goweb/internal/validation/user"
)

func Run(addr string, staticDir string, dbUrl string, sessionsDir string, developmentMode bool) {
	sessionStore := sessions.NewFilesystemStore(sessionsDir, []byte("abc123"))
	sessionStore.MaxAge(3600)

	gormDb := gorm.NewGormDb(dbUrl)

	model.GormInitModels(gormDb)

	userRepository := model.NewGormUserRepository(gormDb)
	userService := user.NewService(userRepository)

	productRepository := model.NewGormProductRepository(gormDb)
	productService := product.NewService(productRepository)

	formValidator := form.NewValidator()
	userValidation.InitEmailValidation(formValidator, userService)

	router := httpController.NewRouter(sessionStore,
		userService, formValidator, productService, staticDir, developmentMode)

	httpServer := internalHttp.NewServer(addr, router)
	httpServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("Os.Signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Println("HttpServer.Notify: %w", err)
	}

	err := httpServer.Stop()
	if err != nil {
		log.Println(err)
	}

}
