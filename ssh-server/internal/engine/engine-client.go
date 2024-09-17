package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type EvaluationRequest struct {
	FEN string `json:"fen"`
}

type MoveEvaluation struct {
	Move  string  `json:"move"`
	Score float64 `json:"score"`
}

type EvaluationResponse struct {
	Evaluations []MoveEvaluation `json:"evaluations"`
}

func GetBestMoveFromChessEngine(fen string) (string, error) {
	// Prepare request payload
	evalReq := EvaluationRequest{FEN: fen}
	jsonReq, err := json.Marshal(evalReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	// Make POST request to the chess engine's evaluate endpoint
	url := "http://chess-engine:80/evaluate"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonReq))

	if err != nil {
		return "", fmt.Errorf("failed to call chess engine: %v", err)
	}
	defer resp.Body.Close()

	log.Println("RESPONSE")
	log.Println(resp)
	log.Println(resp.Body)

	// Parse response
	var evalResp EvaluationResponse
	if err := json.NewDecoder(resp.Body).Decode(&evalResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if len(evalResp.Evaluations) == 0 {
		return "", fmt.Errorf("no move evaluations returned")
	}

	// Return the best move from the response
	bestMove := evalResp.Evaluations[0].Move
	return bestMove, nil
}
