package daily_hot_api

// BilibiliData defines the structure of individual hot list data items for Bilibili
type BilibiliData struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"` // 新增字段
	Cover     string `json:"cover"`
	Author    string `json:"author"` // 新增字段
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// BilibiliParameType defines the structure of parameter types in the Bilibili response
type BilibiliParameType struct {
	Name string            `json:"name"`
	Type map[string]string `json:"type"`
}

// BilibiliParameData defines the structure of parameter data in the Bilibili response
type BilibiliParameData struct {
	Type BilibiliParameType `json:"type"`
}

// BilibiliResponse defines the structure of the response for Bilibili hot list methods
type BilibiliResponse struct {
	Code        int                `json:"code"`
	Name        string             `json:"name"`
	Title       string             `json:"title"`
	Type        string             `json:"type"`
	Description string             `json:"description"` // 新增字段
	ParameData  BilibiliParameData `json:"parameData"`  // 新增字段
	Link        string             `json:"link"`
	Total       int                `json:"total"`
	UpdateTime  string             `json:"update_time"`
	FromCache   bool               `json:"from_cache"`
	Data        []BilibiliData     `json:"data"`
}

// GetBilibiliHotList fetches the hot list for Bilibili
func (c *Client) GetBilibiliHotList() (*BilibiliResponse, error) {
	var response BilibiliResponse
	err := c.fetchHots("bilibili", &response)
	return &response, err
}
