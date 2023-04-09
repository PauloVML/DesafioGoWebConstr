package handler

import (
	"errors"
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/ticket"
	"github.com/gin-gonic/gin"
)

type HandlerTicket struct {
	service ticket.Service
}

func NewHandler(s ticket.Service) *HandlerTicket {
	return &HandlerTicket{
		service: s,
	}
}

func (handler *HandlerTicket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := handler.service.GetTicketByDestination(destination)
		if err != nil {
			if errors.Is(ticket.ErrEmptySlice, err) {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (handler *HandlerTicket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := handler.service.AverageDestination(destination)
		if err != nil {
			if errors.Is(ticket.ErrEmptySlice, err) {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}

func (handler *HandlerTicket) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		tickets, err := handler.service.GetAll()

		if err != nil {
			if errors.Is(ticket.ErrEmptySlice, err) {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(http.StatusOK, tickets)

	}
}
