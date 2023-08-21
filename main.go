package main

import (
	"TaskAPP/api"
	"TaskAPP/database"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDatabase()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", api.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", api.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", api.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks", api.DeleteTask).Methods("DELETE")

	// Definisikan handler untuk subtask
	router.HandleFunc("/tasks/subtask", api.CreateSubTask).Methods("POST")
	router.HandleFunc("/tasks/subtask", api.UpdateSubtask).Methods("PUT")
	router.HandleFunc("/tasks/subtask", api.DeleteSubtask).Methods("DELETE")

	// Definisikan handler untuk attachment
	router.HandleFunc("/tasks/attachment", api.CreateAttachment).Methods("POST")
	router.HandleFunc("/tasks/attachment", api.UpdateAttachment).Methods("PUT")
	router.HandleFunc("/tasks/attachment", api.DeleteAttachment).Methods("DELETE")
	http.Handle("/", router)

	serverAddr := ":8080"
	go func() {
		fmt.Printf("Server sedang berjalan di http://localhost%s\n", serverAddr)
		if err := http.ListenAndServe(serverAddr, nil); err != nil {
			panic(err)
		}
	}()

	select {}
}
