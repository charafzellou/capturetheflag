package main

type (
	Request struct {
		User string `json:"User"`
	}
	AuthRequest struct {
		User   string `json:"User"`
		Secret string `json:"Secret"`
	}
	FullRequest struct {
		User    string `json:"User"`
		Secret  string `json:"Secret"`
		Content struct {
			Level     uint `json:"Level"`
			Challenge struct {
				Username string `json:"Username"`
				Secret   string `json:"Secret"`
				Points   uint   `json:"Points"`
			} `json:"Challenge"`
			Protocol  string `json:"Protocol"`
			SecretKey string `json:"SecretKey"`
		} `json:"Content"`
	}
)
