package responses

// MessageInfo is the message's struct
type MessageInfo struct {
	Position SatelitePosition `json:"position"`
	Message  string           `json:"message"`
}

// SatelitePosition are the satelite's coords
type SatelitePosition struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
