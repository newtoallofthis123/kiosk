package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/ollama/ollama/api"
)

type Server struct {
	Db   *DbInstance
	Addr string
	subs map[*websocket.Conn]bool
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewServer(env *Env) (*Server, error) {
	db, err := NewDb(env)
	if err != nil {
		return nil, err
	}

	return &Server{
		Db:   db,
		Addr: env.ListenAddr,
		subs: make(map[*websocket.Conn]bool),
	}, nil
}

func (s *Server) Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": "1.0",
	})
}

func (s *Server) handleSubscribe(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading websocket: %v", err)
		return
	}

	waiting := false

	for {
		if waiting {
			continue
		} else {
			waiting = true
		}
		_, prompt, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		client, err := api.ClientFromEnvironment()
		if err != nil {
			log.Fatal(err)
		}

		// By default, GenerateRequest is streaming.
		req := &api.GenerateRequest{
			Model:  "orca-mini",
			Prompt: string(prompt),
		}

		ctx := context.Background()

		respFunc := func(resp api.GenerateResponse) error {
			if resp.Done {
				waiting = false
			}
			conn.WriteMessage(websocket.TextMessage, []byte(resp.Response))
			return nil
		}

		err = client.Generate(ctx, req, respFunc)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Server) handleCreateRecord(c *gin.Context) {
	var createRecordRequest ConstructRequest
	err := c.BindJSON(&createRecordRequest)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	id, err := s.Db.CreateRecord(&createRecordRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating record"})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func (s *Server) handleGetRecord(c *gin.Context) {
	id := c.Param("id")

	record, err := s.Db.GetRecord(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error getting record"})
		return
	}

	c.JSON(200, record)
}

func (s *Server) handleListRecords(c *gin.Context) {
	records, err := s.Db.ListRecords()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error listing records"})
		return
	}

	c.JSON(200, records)
}

type ContructRequest struct {
	Symtoms  []string `json:"symtoms,omitempty"`
	Diseases []string `json:"diseases,omitempty"`
}

func (s *Server) handlerConstructPrompt(c *gin.Context) {
	var req ContructRequest
	err := c.BindJSON(&req)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	prompt := `
	Imagine you are doctor
	`

	final := fmt.Sprintf(prompt, strings.Join(req.Symtoms, ","), strings.Join(req.Diseases, ","))

	c.JSON(200, map[string]string{"prompt": final})
}

func (s *Server) Start() error {
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	})

	r.GET("/v", s.Version)
	r.POST("/records", s.handleCreateRecord)
	r.GET("/records/:id", s.handleGetRecord)
	r.GET("/records", s.handleListRecords)
	r.GET("/generate", s.handleSubscribe)

	err := r.Run(s.Addr)
	if err != nil {
		return err
	}

	return nil
}
