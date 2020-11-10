package entities

// Satelite is a representation of a group of satelites
type Satelite struct {
	Satelites []Satelites `json:"satellites"`
}

// Satelites is a representation of satelite's parameters
type Satelites struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
