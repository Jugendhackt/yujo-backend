package jsonmodels

type PlayerNames struct {
	CreatorName  string `json:"creatorName"`
	TeamMateName string `json:"teammateName"`
}

type HealthPoints struct {
	Creator  int `json:"0"`
	TeamMate int `json:"1"`
	Enemy    int `json:"2"`
}

type GameInfo struct {
	Names         PlayerNames `json:"names"`
	HealthPoints  HealthPoints
	CorrectAnswer bool `json:"correctAnswer"`
	NextRoundID   int  `json:"nextRoundID"`
}
