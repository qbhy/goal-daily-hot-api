package daily_hot_api

// ToutiaoData defines the structure of individual hot list data items for Toutiao
type ToutiaoData struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// ToutiaoResponse defines the structure of the response for Toutiao hot list methods
type ToutiaoResponse struct {
	Code       int           `json:"code"`
	Name       string        `json:"name"`
	Title      string        `json:"title"`
	Type       string        `json:"type"`
	Link       string        `json:"link"`
	Total      int           `json:"total"`
	UpdateTime string        `json:"update_time"`
	FromCache  bool          `json:"from_cache"`
	Data       []ToutiaoData `json:"data"`
}

// GetToutiaoHotList fetches the hot list for Toutiao
func (c *Client) GetToutiaoHotList() (*ToutiaoResponse, error) {
	var response ToutiaoResponse
	err := c.fetchHots("toutiao", &response)
	return &response, err
}
