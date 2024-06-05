package daily_hot_api

// BaiduData defines the structure of individual hot list data items for Baidu
type BaiduData struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Cover     string `json:"cover"`
	Author    string `json:"author"`
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// BaiduParameType defines the structure of parameter types in the Baidu response
type BaiduParameType struct {
	Name string            `json:"name"`
	Type map[string]string `json:"type"`
}

// BaiduParameData defines the structure of parameter data in the Baidu response
type BaiduParameData struct {
	Type BaiduParameType `json:"type"`
}

// BaiduResponse defines the structure of the response for Baidu hot list methods
type BaiduResponse struct {
	Code       int             `json:"code"`
	Name       string          `json:"name"`
	Title      string          `json:"title"`
	Type       string          `json:"type"`
	ParameData BaiduParameData `json:"parameData"`
	Link       string          `json:"link"`
	Total      int             `json:"total"`
	UpdateTime string          `json:"update_time"`
	FromCache  bool            `json:"from_cache"`
	Data       []BaiduData     `json:"data"`
}

// GetBaiduHotSearch fetches the hot search list for Baidu
func (c *Client) GetBaiduHotSearch() (*BaiduResponse, error) {
	var response BaiduResponse
	err := c.fetchHots("baidu", &response)
	return &response, err
}
