package daily_hot_api

// WeiboData defines the structure of individual hot list data items for Weibo
type WeiboData struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Author    string `json:"author"`
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// WeiboResponse defines the structure of the response for Weibo hot list methods
type WeiboResponse struct {
	Code        int         `json:"code"`
	Name        string      `json:"name"`
	Title       string      `json:"title"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Link        string      `json:"link"`
	Total       int         `json:"total"`
	UpdateTime  string      `json:"update_time"`
	FromCache   bool        `json:"from_cache"`
	Data        []WeiboData `json:"data"`
}

// GetWeiboHotSearch fetches the hot search list for Weibo
func (c *Client) GetWeiboHotSearch() (*WeiboResponse, error) {
	var response WeiboResponse
	err := c.fetchHots("weibo", &response)
	return &response, err
}
