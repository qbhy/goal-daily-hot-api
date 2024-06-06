package daily_hot_api

// NeteaseNewsData defines the structure of individual hot list data items for Netease News
type NeteaseNewsData struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Author    string `json:"author"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// NeteaseNewsResponse defines the structure of the response for Netease News hot list methods
type NeteaseNewsResponse struct {
	Code       int               `json:"code"`
	Name       string            `json:"name"`
	Title      string            `json:"title"`
	Type       string            `json:"type"`
	Link       string            `json:"link"`
	Total      int               `json:"total"`
	UpdateTime string            `json:"update_time"`
	FromCache  bool              `json:"from_cache"`
	Data       []NeteaseNewsData `json:"data"`
}

// GetNeteaseNewsHotList fetches the hot list for Netease News
func (c *Client) GetNeteaseNewsHotList() (*NeteaseNewsResponse, error) {
	var response NeteaseNewsResponse
	err := c.fetchHots("netease-news", &response)
	return &response, err
}
