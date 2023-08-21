package api

import (
	"TaskAPP/database"
	"TaskAPP/models"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(tasks) == 0 {
		response := map[string]string{
			"message": "Maaf, tidak ada data task ya mungkin kamu belum menambahkan",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Initialize subtasks and attachments as empty arrays
	task.Subtasks = []models.Subtask{}
	task.Attachments = []models.Attachment{}

	insertedID, err := database.InsertTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task.ID = insertedID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskID := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	updatedTask.ID = objectID

	err = database.UpdateTask(updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteTask(objectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Berhasil menghapus Task ",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreateSubTask(w http.ResponseWriter, r *http.Request) {
	var subtask models.Subtask
	err := json.NewDecoder(r.Body).Decode(&subtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskID := r.URL.Query().Get("task_id")
	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	// Tambahkan subtask ke task
	err = database.AddSubtaskToTask(taskObjectID, subtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ambil ulang task setelah menambahkan subtask
	updatedTask, err := database.GetTaskByID(taskObjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedTask)
}

func UpdateSubtask(w http.ResponseWriter, r *http.Request) {
	var updatedSubtask models.Subtask
	err := json.NewDecoder(r.Body).Decode(&updatedSubtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskID := r.URL.Query().Get("task_id")
	subtaskID := r.URL.Query().Get("subtask_id")

	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	subtaskObjectID, err := primitive.ObjectIDFromHex(subtaskID)
	if err != nil {
		http.Error(w, "Invalid subtask ID", http.StatusBadRequest)
		return
	}

	// Update subtask menggunakan subtask ID
	err = database.UpdateSubtaskInTask(taskObjectID, subtaskObjectID, updatedSubtask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update task dengan subtask yang telah diubah
	updatedTask := models.Task{
		ID: taskObjectID,
	}
	err = database.UpdateTask(updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedSubtask)
}

func DeleteSubtask(w http.ResponseWriter, r *http.Request) {
	subtaskID := r.URL.Query().Get("subtask_id")
	subtaskObjectID, err := primitive.ObjectIDFromHex(subtaskID)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return

	}
	err = database.DeleteSubtask(subtaskObjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"message": "Subtask deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func CreateAttachment(w http.ResponseWriter, r *http.Request) {
	var attachment models.Attachment
	err := json.NewDecoder(r.Body).Decode(&attachment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskID := r.URL.Query().Get("task_id")
	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = database.AddAttachmentToTask(taskObjectID, attachment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attachment)
}

func UpdateAttachment(w http.ResponseWriter, r *http.Request) {
	var updatedAttachment models.Attachment
	err := json.NewDecoder(r.Body).Decode(&updatedAttachment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskID := r.URL.Query().Get("task_id")
	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	attachmentID := r.URL.Query().Get("attachment_id")
	attachmentObjectID, err := primitive.ObjectIDFromHex(attachmentID)
	if err != nil {
		http.Error(w, "Invalid attachment ID", http.StatusBadRequest)
		return
	}

	err = database.UpdateAttachmentInTask(taskObjectID, attachmentObjectID, updatedAttachment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedAttachment)
}

func DeleteAttachment(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Query().Get("task_id")
	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	attachmentID := r.URL.Query().Get("attachment_id")
	attachmentObjectID, err := primitive.ObjectIDFromHex(attachmentID)
	if err != nil {
		http.Error(w, "Invalid attachment ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteAttachmentFromTask(taskObjectID, attachmentObjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Attachment deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
