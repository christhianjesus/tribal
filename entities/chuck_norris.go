package entities

type ApiResponse struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Value string `json:"value"`
}

const (
	MaxTries       = 3
	JokesNum       = 25
	ChuckNorrisURL = "https://api.chucknorris.io/jokes/random"
)
