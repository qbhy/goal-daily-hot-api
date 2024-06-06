package daily_hot_api

// ThePaperData defines the structure of individual hot list data items for The Paper
type ThePaperData struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// ThePaperResponse defines the structure of the response for The Paper hot list methods
type ThePaperResponse struct {
	Code       int            `json:"code"`
	Name       string         `json:"name"`
	Title      string         `json:"title"`
	Type       string         `json:"type"`
	Link       string         `json:"link"`
	Total      int            `json:"total"`
	UpdateTime string         `json:"update_time"`
	FromCache  bool           `json:"from_cache"`
	Data       []ThePaperData `json:"data"`
}

// GetThePaperHotList fetches the hot list for The Paper
func (c *Client) GetThePaperHotList() (*ThePaperResponse, error) {
	var response ThePaperResponse
	err := c.fetchHots("thepaper", &response)
	return &response, err
}
