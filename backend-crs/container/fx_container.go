package container

import (
	"backend-crs/config"
	"backend-crs/controller"
	"backend-crs/repository"
	"backend-crs/route"
	"backend-crs/service"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewDatabase(lc fx.Lifecycle) (*gorm.DB, error) {
	
	db, err := config.ConnectDatabase();
	if err != nil {
		panic("Failed to commect to the Database "+err.Error())
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			postgresDB, _ := db.DB()
			postgresDB.Close()
			return err 
		},
	})
	return db, nil
}

func StartServer(r *gin.Engine) {
	fmt.Println("Server successfully started in the PORT : 8080")
	r.Run(":8080")
}

var Module = fx.Module("Application Entry", fx.Provide(
	NewDatabase,

	// repositories
	repository.NewStudentRepository,

	// services
	service.NewStudentService,

	//controllers
	controller.NewAuthController,

	//routers
	route.SetupAuthRouter,


), fx.Invoke(StartServer))
