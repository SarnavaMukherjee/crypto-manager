/**
 * File: server.go
 * Author: Kedar Kekan
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	log "github.com/SarnavaMukherjee/crypto-manager/pkg/logger"
)

var (
	port   int
	g      errgroup.Group
	server *http.Server
)

//Setup ...
func Setup(_port int, loglevel string) *gin.Engine {
	if _port == 0 {
		_port = 8080
	}

	if strings.ToLower(loglevel) != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DisableConsoleColor()

	r := gin.New()

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {

			// your custom format
			return fmt.Sprintf("[HTTP-API] %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
	}))

	r.Use(gin.Recovery())

	r.Use(cors.Default())
	port = _port
	return r
}

//Start ...
func Start(r *gin.Engine) {
	server = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: r,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal("", "", err.Error())
	}
}

//Stop ...
func Stop() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	log.Info("", "", "Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("", "", err.Error())
	}
	log.Info("", "", "Server exiting")
}

type AppResponse struct {
	Status  int         `json:"-"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// NewAppResponse generates an application response
func NewAppResponse(status int, data interface{}, msg string) *AppResponse {
	return &AppResponse{Status: status, Data: data, Message: msg}
}

//WriteResponse ...
func WriteResponse(
	c *gin.Context,
	resp *AppResponse,
) {
	c.JSON(resp.Status, resp)
}

// AppError is the default error struct containing detailed information about the error
type AppError struct {
	// HTTP Status code to be set in response
	Status int `json:"-"`
	// Message is the error message that may be displayed to end users
	Message string `json:"message,omitempty"`
}

// WriteResponseWithError writes an error response to client
func WriteResponseWithError(c *gin.Context, err error) {
	switch err.(type) {
	case *AppError:
		e := err.(*AppError)
		if e.Message == "" {
			c.AbortWithStatus(e.Status)
		} else {
			c.AbortWithStatusJSON(e.Status, e)
		}
		return
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}

var (
	// Generic error
	Generic = NewAppErrorStatus(http.StatusInternalServerError)
	// DB represents database related errors
	DB = NewAppErrorStatus(http.StatusInternalServerError)
	// Forbidden represents access to forbidden resource error
	Forbidden = NewAppErrorStatus(http.StatusForbidden)
	// BadRequest represents error for bad requests
	BadRequest = NewAppErrorStatus(http.StatusBadRequest)
	// NotFound represents errors for not found resources
	NotFound = NewAppErrorStatus(http.StatusNotFound)
	// Unauthorized represents errors for unauthorized requests
	Unauthorized = NewAppErrorStatus(http.StatusUnauthorized)
)

// NewAppErrorStatus generates new error containing only http status code
func NewAppErrorStatus(status int) *AppError {
	return &AppError{Status: status}
}

// NewAppError generates an application error
func NewAppError(status int, msg string) *AppError {
	return &AppError{Status: status, Message: msg}
}

// Error returns the error message.
func (e AppError) Error() string {
	return e.Message
}

var validationErrors = map[string]string{
	"required": " is required, but was not received",
	"min":      "'s value or length is less than allowed",
	"max":      "'s value or length is bigger than allowed",
}

func getVldErrorMsg(s string) string {
	if v, ok := validationErrors[s]; ok {
		return v
	}
	return " failed on " + s + " validation"
}
