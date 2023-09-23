package structs

type Urls struct {
	URL        string `json:"url" binding:"required"`
	Shorted    string `json:"shorted" binding:"required"`
	Created_at int32  `json:"created_at" binding:"required"`
}
