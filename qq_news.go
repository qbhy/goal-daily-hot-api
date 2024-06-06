package daily_hot_api

// QQNewsData defines the structure of individual hot list data items for QQ News
type QQNewsData struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Author    string `json:"author"`
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// QQNewsResponse defines the structure of the response for QQ News hot list methods
type QQNewsResponse struct {
	Code       int          `json:"code"`
	Name       string       `json:"name"`
	Title      string       `json:"title"`
	Type       string       `json:"type"`
	Link       string       `json:"link"`
	Total      int          `json:"total"`
	UpdateTime string       `json:"update_time"`
	FromCache  bool         `json:"from_cache"`
	Data       []QQNewsData `json:"data"`
}

// GetQQNewsHotList fetches the hot list for QQ News
func (c *Client) GetQQNewsHotList() (*QQNewsResponse, error) {
	var response QQNewsResponse
	err := c.fetchHots("qq-news", &response)
	return &response, err
}
