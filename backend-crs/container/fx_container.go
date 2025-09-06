package container

import (
	"backend-crs/config"
	"backend-crs/controller"
	"backend-crs/middleware"
	"backend-crs/model"
	"backend-crs/repository"
	"backend-crs/route"
	"backend-crs/service"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.AddCorsMapping())
	return router
}

func NewDatabase(lc fx.Lifecycle) (*gorm.DB, error) {

	db, err := config.ConnectDatabase()
	if err != nil {
		panic("Failed to commect to the Database " + err.Error())
	}
	db.AutoMigrate(&model.Staff{}, &model.Student{}, &model.Admin{}, &model.Course{}, &model.CourseOffering{}, &model.Department{})

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

	// router
	NewRouter,

	// repositories
	repository.NewStudentRepository,
	repository.NewStaffRepository,
	repository.NewCourseRepository,

	// services
	service.NewStudentService,
	service.NewStaffService,
	service.NewCourseService,

	//controllers
	controller.NewAuthController,
	controller.NewTempController,
	controller.NewCourseController,
), fx.Invoke(
	route.SetupAuthRouter,
	route.SetupTempRouter,
	route.SetupCourseRouter,
	StartServer))
