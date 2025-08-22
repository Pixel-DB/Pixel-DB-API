package dto

type APIResponseData struct {
	TotalRequests    int64
	TotalUsers       int64
	TotalPixelArts   int64
	TotalGithubStars int64
}

type APIResponse struct {
	Status  string
	Message string
	Data    APIResponseData
}
