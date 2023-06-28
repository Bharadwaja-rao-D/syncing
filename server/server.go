package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func collaborator(w http.ResponseWriter, r *http.Request){
    upgrader := new(websocket.Upgrader)
    conn, err := upgrader.Upgrade(w, r, nil);
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    for {
        mtype, msg, _ := conn.ReadMessage();
        if mtype == websocket.CloseMessage || len(msg) == 0{
            log.Printf("Closing connection: %s\n", conn.RemoteAddr() );
            break;
        }
        log.Printf("From client %s: %s\n", conn.RemoteAddr(), msg );
        conn.WriteMessage(mtype, msg);
    }
}

func Server(){
    http.HandleFunc("/collaborate", collaborator)
    http.ListenAndServe(":8080", nil)
}
