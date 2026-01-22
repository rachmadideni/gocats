package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Categories = []Category{
	{ID: 1, Name: "Elektronik", Description: "Perangkat elektronik seperti ponsel, laptop, dan televisi."},
}

func getCategories(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Categories)
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the path
	path := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid category ID"})
		return
	}

	for _, category := range Categories {
		if category.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}
	}

	// Category not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a category
	w.Header().Set("Content-Type", "application/json")
	path := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid category ID"})
		return
	}
	var updatedCategory Category
	if err := json.NewDecoder(r.Body).Decode(&updatedCategory); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	for i, category := range Categories {
		if category.ID == id {
			Categories[i].Name = updatedCategory.Name
			Categories[i].Description = updatedCategory.Description
			json.NewEncoder(w).Encode(Categories[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid category ID"})
		return
	}

	for i, p := range Categories {
		if p.ID == id {
			Categories = append(Categories[:i], Categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "success deleting a category",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a new category
	w.Header().Set("Content-Type", "application/json")
	var newCategory Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}
	newCategory.ID = len(Categories) + 1
	Categories = append(Categories, newCategory)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)

}

func main() {

	// Health Check Endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Category service is running",
		})
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCategories(w)
		} else if r.Method == http.MethodPost {
			createCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCategoryByID(w, r)
		} else if r.Method == http.MethodPut {
			updateCategory(w, r)
		} else if r.Method == http.MethodDelete {
			deleteCategory(w, r)
		}
	})

	fmt.Println("Server starting on port 6000...")
	if err := http.ListenAndServe(":6000", nil); err != nil {
		fmt.Printf("HTTP server failed: %v\n", err)
	}
}
