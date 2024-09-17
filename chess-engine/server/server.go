package main

import (
	"chess-engine/engine"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func init() {
	engine.InitBitboards()
	engine.InitTables()
	engine.InitZobrist()
	engine.InitEvalBitboards()
	engine.InitSearchTables()
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/evaluate", handleEvaluate)

	fmt.Println("Starting server on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Chess Engine Server!")
}

func handleEvaluate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req engine.EvaluationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resultChan := make(chan engine.EvaluationResponse)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		evaluatePositionAsync(req.FEN, resultChan)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	select {
	case result, ok := <-resultChan:
		if !ok {
			http.Error(w, "Evaluation failed", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	case <-time.After(10 * time.Second):
		http.Error(w, "Evaluation timed out", http.StatusRequestTimeout)
	}
}

func evaluatePositionAsync(fen string, resultChan chan<- engine.EvaluationResponse) {

	search := engine.Search{}
	search.TT.Resize(engine.DefaultTTSize, engine.SearchEntrySize)

	search.Setup(fen)
	
	search.Timer.Setup(300000, 2, 1500, 15, getRandomMovesAhead(), 100000)

	bestMove := search.Search()
	score := engine.EvaluatePos(&search.Pos)

	evaluations := make([]engine.MoveEvaluation, 0, 1)
	evaluations = append(evaluations, engine.MoveEvaluation{
		Move:  bestMove.String(),
		Score: float64(score) / 100.0, // Convert centipawns to pawns
	})
	result := engine.EvaluationResponse{
		Evaluations: evaluations,
	}
	resultChan <- result
}

func getRandomMovesAhead() uint8 {

	random := rand.Intn(10)

	if(random >= 8) {
		return 1
	} else if (random >= 6) {
		return 2
	}

    return uint8(rand.Intn(4) + 3)
}
