package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"school/internal/config"
	"school/internal/handlers"
	"school/internal/middleware"
	"school/internal/services"
	"school/pkg/jwt"
)

func Register(r *gin.Engine, db *pgxpool.Pool, cfg *config.Config) {
	jm := jwt.New(cfg.JWTSecret, cfg.JWTIssuer, cfg.JWTExpireH)
	authSvc := services.NewAuthService(db)
	authH := handlers.NewAuthHandler(authSvc, jm)

	studentSvc := services.NewStudentService(db)
	studentH := handlers.NewStudentHandler(studentSvc)

	classSvc := services.NewClassroomService(db)
	classH := handlers.NewClassroomHandler(classSvc)

	regSvc := services.NewRegistrarService(db)
	regH := handlers.NewRegistrarHandler(regSvc)

	
	// OpenAPI docs & Swagger UI
	r.StaticFile("/openapi.yaml", "./api/openapi.yaml")
	r.GET("/docs", func(c *gin.Context) {
		html := `<!DOCTYPE html><html><head><meta charset="utf-8"/>
		<title>API Docs</title>
		<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css" />
		</head><body>
		<div id="swagger-ui"></div>
		<script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
		<script>
		window.onload = () => { window.ui = SwaggerUIBundle({ url: '/openapi.yaml', dom_id: '#swagger-ui' }); };
		</script></body></html>`
		c.Data(200, "text/html; charset=utf-8", []byte(html))
	})

	api := r.Group("/api")
	{
		api.POST("/auth/login", authH.Login)

		protected := api.Group("")
		protected.Use(middleware.AuthRequired(jm))
		{
			protected.GET("/me", func(c *gin.Context) { claims, _ := c.Get("claims"); c.JSON(200, claims) })

			stu := protected.Group("/students")
			{
				stu.GET("", middleware.RequirePermission("STUDENT_READ"), studentH.List)
				stu.GET("/:id", middleware.RequirePermission("STUDENT_READ"), studentH.Get)
				stu.POST("", middleware.RequirePermission("STUDENT_WRITE"), studentH.Create)
				stu.PUT("/:id", middleware.RequirePermission("STUDENT_WRITE"), studentH.Update)
				stu.DELETE("/:id", middleware.RequirePermission("STUDENT_WRITE"), studentH.Delete)
			}

			cl := protected.Group("/classrooms")
			{
				cl.GET("", middleware.RequirePermission("CLASSROOM_READ"), classH.List)
				cl.GET("/:id", middleware.RequirePermission("CLASSROOM_READ"), classH.Get)
				cl.POST("", middleware.RequirePermission("CLASSROOM_WRITE"), classH.Create)
				cl.PUT("/:id", middleware.RequirePermission("CLASSROOM_WRITE"), classH.Update)
				cl.DELETE("/:id", middleware.RequirePermission("CLASSROOM_WRITE"), classH.Delete)
			}

			en := protected.Group("/enrollments")
			{
				en.POST("", middleware.RequirePermission("ENROLLMENT_WRITE"), regH.Enroll)
				en.PUT(":/id/status", middleware.RequirePermission("ENROLLMENT_WRITE"), regH.UpdateStatus)
				en.GET("/by-classroom/:classroom_id", middleware.RequirePermission("ENROLLMENT_READ"), regH.ListByClassroom)
				en.GET("/by-student/:student_id", middleware.RequirePermission("ENROLLMENT_READ"), regH.ListByStudent)
			}
		}
	}
}
