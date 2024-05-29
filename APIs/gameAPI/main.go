package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: func (r *http.Request) bool{
        return true
    },
}

var clients = make(map[*websocket.Conn]bool)
var brodcast = make(chan Task)

type Task struct {
    Usernaem string `json:"username"`
    TaskName string `json:"task_name"`
    TaskDone bool `json:"task_done"`
}

func handlerConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }

    defer  conn.Close()

    clients[conn] = true

    for {
        var task Task
        
        err := conn.ReadJSON(&task)
        if err != nil {
            log.Println(err)
            delete(clients, conn)
            return
        }

        brodcast <- task
    }
}

func handlerGames() {
    for {
        taks := <- brodcast

        for client := range clients {
            err := client.WriteJSON(taks)
            if err != nil {
                log.Println(err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main() {
	router := chi.NewRouter()

    portString := "3007"
    
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

    v1Router := chi.NewRouter()
    v1Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("connected")
    })
    v1Router.Get("/game", handlerConnections)

    go handlerGames()

    router.Mount("/v1", v1Router)

    srv := &http.Server{
        Handler: router,
        Addr: fmt.Sprintf(":%s", portString),

    }

    log.Printf("Server running on port %s\n", portString)
    log.Fatal(srv.ListenAndServe())
}
