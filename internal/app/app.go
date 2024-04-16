package app

import (
	"ivixlabs.com/goweb/internal/clickhouse"
	product2 "ivixlabs.com/goweb/internal/gorm/repository/product"
	"ivixlabs.com/goweb/internal/gorm/repository/property"
	user2 "ivixlabs.com/goweb/internal/gorm/repository/user"
	"ivixlabs.com/goweb/internal/model/product/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/sessions"
	httpController "ivixlabs.com/goweb/internal/controller/http"
	"ivixlabs.com/goweb/internal/gorm"
	internalHttp "ivixlabs.com/goweb/internal/http"
	"ivixlabs.com/goweb/internal/model"
	"ivixlabs.com/goweb/internal/model/product"
	"ivixlabs.com/goweb/internal/user"
	"ivixlabs.com/goweb/internal/validation/form"
	userValidation "ivixlabs.com/goweb/internal/validation/user"
)

func Run(
	addr string,
	staticDir string,
	dbUrl string,
	sessionsDir string,
	developmentMode bool,
	clickhouseAddr []string,
) {
	sessionStore := sessions.NewFilesystemStore(sessionsDir, []byte("abc123"))
	sessionStore.MaxAge(3600)

	gormDb := gorm.NewGormDb(dbUrl)

	//model.GormInitModels(gormDb)
	gorm.InitModels(gormDb)

	userRepository := user2.New(gormDb)
	userService := user.NewService(userRepository)

	productRepository := product2.New(gormDb)
	productCreation := usecase.NewProductCreation(productRepository)
	productUpdating := usecase.NewProductUpdating(productRepository)
	productService := product.NewService(productRepository)

	propertyRepository := property.New(gormDb)

	formValidator := form.NewValidator()
	userValidation.InitEmailValidation(formValidator, userService)

	clickhouseConn := clickhouse.NewConn(clickhouseAddr)
	entityPropertyRepo := model.NewClickHouseEntityPropertyRepository(clickhouseConn)
	entityRepo := model.NewClickHouseEntityRepository(clickhouseConn, entityPropertyRepo)

	router := httpController.NewRouter(
		sessionStore,
		userService,
		formValidator,
		productCreation,
		productUpdating,
		productRepository,
		productService,
		staticDir,
		developmentMode,
		propertyRepository,
		entityRepo,
		entityPropertyRepo,
	)

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
