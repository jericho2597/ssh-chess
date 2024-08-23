package engine

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
