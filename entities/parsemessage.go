package entities

// ParseMessage is a struct
type ParseMessage struct {
	Distances map[string]float32
	Messages  [][]string
}

// SateliteDistance is a struct
type SateliteDistance struct {
	Name     string
	Distance float32
}
