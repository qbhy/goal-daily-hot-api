package daily_hot_api

// ZhihuData defines the structure of individual hot list data items for Zhihu
type ZhihuData struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Cover     string `json:"cover"`
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// ZhihuResponse defines the structure of the response for Zhihu hot list methods
type ZhihuResponse struct {
	Code       int         `json:"code"`
	Name       string      `json:"name"`
	Title      string      `json:"title"`
	Type       string      `json:"type"`
	Link       string      `json:"link"`
	Total      int         `json:"total"`
	UpdateTime string      `json:"update_time"`
	FromCache  bool        `json:"from_cache"`
	Data       []ZhihuData `json:"data"`
}

// GetZhihuHotList fetches the hot list for Zhihu
func (c *Client) GetZhihuHotList() (*ZhihuResponse, error) {
	var response ZhihuResponse
	err := c.fetchHots("zhihu", &response)
	return &response, err
}
