package server

import (
	"fmt"
	handler2 "github.com/bootcamp-go/desafio-go-web/cmd/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/ticket"
	"github.com/gin-gonic/gin"
)

type router struct {
	engine *gin.Engine
	bd     *[]domain.Ticket
}

func NewRouter(engine *gin.Engine, tickets *[]domain.Ticket) *router {
	return &router{
		engine: engine,
		bd:     tickets,
	}
}

func (r *router) MapRoutes() {
	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())

	r.TicketRoutes()
}

func (r *router) TicketRoutes() {

	fmt.Printf("Dirección en memoria desde TicketRoutes INICIADA %p\n", r.bd)
	fmt.Printf("Dirección en memoria desde TicketRoutes TERMINADA \n")

	repo := ticket.NewRepository(r.bd)
	service := ticket.NewServiceImpl(repo)
	handler := handler2.NewHandler(service)

	tck := r.engine.Group("/ticket")

	tck.GET("", handler.GetAll())
}
