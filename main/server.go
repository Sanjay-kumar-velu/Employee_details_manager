package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Employee struct
type Employee struct {
	ID       int     `json:"id" bson:"id"`
	Name     string  `json:"name" bson:"name"`
	Role     string  `json:"role" bson:"role"`
	Language string  `json:"language" bson:"language"`
	Salary   float64 `json:"salary" bson:"salary"`
}

var (
	client      *mongo.Client
	emp_table   *mongo.Collection
	dept_table  *mongo.Collection
	dev_table   *mongo.Collection
	test_table  *mongo.Collection
	idCounter   = 1
	idCounterMu sync.Mutex
)

// employeesHandler handles API requests
func employeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		idCounterMu.Lock()
		emp.ID = idCounter
		idCounter++
		idCounterMu.Unlock()

		_, err := emp_table.InsertOne(context.TODO(), emp)
		if err != nil {
			http.Error(w, "Failed to save employee", http.StatusInternalServerError)
			return
		}

		dept_table.InsertOne(context.TODO(), bson.M{"id": emp.ID, "role": emp.Role})
		desg := strings.ToLower(emp.Role)
		switch desg {
		case "developer":
			dev_table.InsertOne(context.TODO(), emp)
		case "tester":
			test_table.InsertOne(context.TODO(), emp)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(emp)

	case http.MethodGet:
		cur, err := emp_table.Find(context.TODO(), bson.M{})
		if err != nil {
			http.Error(w, "Failed to fetch employees", http.StatusInternalServerError)
			return
		}
		defer cur.Close(context.TODO())

		var employees []Employee
		for cur.Next(context.TODO()) {
			var emp Employee
			if err := cur.Decode(&emp); err != nil {
				http.Error(w, "Error decoding data", http.StatusInternalServerError)
				return
			}
			employees = append(employees, emp)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// initIDCounter initializes the counter from the highest existing ID
func initIDCounter() {
	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	var emp Employee
	err := emp_table.FindOne(context.TODO(), bson.M{}, opts).Decode(&emp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			idCounter = 1
			log.Println("No employees found. Starting IDs from 1.")
		} else {
			log.Fatalf("Failed to initialize ID counter: %v", err)
		}
	} else {
		idCounter = emp.ID + 1
		log.Printf("Initialized ID counter. Starting from %d\n", idCounter)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://martinmystery885:Martin123@cluster0.lbvhff5.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0",
	))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("vuejs&go")
	emp_table = db.Collection("employees")
	dept_table = db.Collection("departments")
	dev_table = db.Collection("developers")
	test_table = db.Collection("testers")

	initIDCounter()

	// Serve API with CORS support for development
	http.HandleFunc("/api/employees", func(w http.ResponseWriter, r *http.Request) { // Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			return
		} // Call existing handler
		employeesHandler(w, r)
	})

	// Serve static assets
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/assets/", fs)

	// Catch-all: serve index.html for SPA routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If the request is for an actual file, let it pass
		if _, err := os.Stat("./frontend/dist" + r.URL.Path); err == nil {
			fs.ServeHTTP(w, r)
			return
		}

		// Otherwise serve index.html
		http.ServeFile(w, r, "./frontend/dist/index.html")
	})

	log.Println("Server running at http://localhost:3030")
	log.Fatal(http.ListenAndServe(":3030", nil))
}
