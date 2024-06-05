package daily_hot_api

// ZhihuDailyData defines the structure of individual hot list data items for Zhihu Daily
type ZhihuDailyData struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Author    string `json:"author"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// ZhihuDailyResponse defines the structure of the response for Zhihu Daily hot list methods
type ZhihuDailyResponse struct {
	Code        int              `json:"code"`
	Name        string           `json:"name"`
	Title       string           `json:"title"`
	Type        string           `json:"type"`
	Description string           `json:"description"`
	Link        string           `json:"link"`
	Total       int              `json:"total"`
	UpdateTime  string           `json:"update_time"`
	FromCache   bool             `json:"from_cache"`
	Data        []ZhihuDailyData `json:"data"`
}

// GetZhihuDailyRecommendations fetches the recommendations for Zhihu Daily
func (c *Client) GetZhihuDailyRecommendations() (*ZhihuDailyResponse, error) {
	var response ZhihuDailyResponse
	err := c.fetchHots("zhihu-daily", &response)
	return &response, err
}
