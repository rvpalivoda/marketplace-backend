package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/zhashkevych/courses-backend/docs"
	v1 "github.com/zhashkevych/courses-backend/internal/delivery/http/v1"
	"github.com/zhashkevych/courses-backend/internal/service"
	"github.com/zhashkevych/courses-backend/pkg/auth"
	"net/http"

	_ "github.com/zhashkevych/courses-backend/docs"
)

type Handler struct {
	schoolsService    service.Schools
	studentsService   service.Students
	coursesService    service.Courses
	promoCodesService service.PromoCodes
	offersService     service.Offers
	modulesService    service.Modules
	ordersService     service.Orders
	paymentsService   service.Payments
	adminsService     service.Admins
	tokenManager      auth.TokenManager
}

func NewHandler(schoolsService service.Schools, studentsService service.Students, coursesService service.Courses, promoCodesService service.PromoCodes,
	offersService service.Offers, modulesService service.Modules, ordersService service.Orders,
	paymentsService service.Payments, adminsService service.Admins, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		schoolsService:    schoolsService,
		studentsService:   studentsService,
		coursesService:    coursesService,
		promoCodesService: promoCodesService,
		offersService:     offersService,
		modulesService:    modulesService,
		ordersService:     ordersService,
		paymentsService:   paymentsService,
		adminsService:     adminsService,
		tokenManager:      tokenManager,
	}
}

func (h *Handler) Init(host, port string) *gin.Engine {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.schoolsService, h.studentsService, h.coursesService, h.promoCodesService,
		h.offersService, h.modulesService, h.ordersService, h.paymentsService, h.adminsService, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
