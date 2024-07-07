package main

import (
	"booking-system/config"
	"booking-system/controller"
	"booking-system/repository"
	"booking-system/service"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

type Server struct {
	pS      service.BookingService
	mS      service.MidtransService
	engine  *gin.Engine
	portApp string
}

func (s *Server) initiateRoute() {
	routeGroup := s.engine.Group("/api/v1")
	controller.NewBookingController(s.pS, routeGroup).Router()
	controller.NewMidtransController(s.mS, routeGroup).Router()
}

func (s *Server) Start() {
	s.initiateRoute()
	s.engine.Run(s.portApp)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open(cfg.Driver, urlConnection)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	validate := validator.New()
	paymentRepo := repository.NewBookingRepository(db)
	midtransService := service.NewMidtransServiceImpl(validate)
	paymentService := service.NewBookingService(paymentRepo)

	return &Server{
		pS:      paymentService,
		mS:      midtransService,
		engine:  gin.Default(),
		portApp: cfg.AppPort,
	}
}
