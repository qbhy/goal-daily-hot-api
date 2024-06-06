package daily_hot_api

// TiebaData defines the structure of individual hot list data items for Tieba
type TiebaData struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Cover     string `json:"cover"`
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// TiebaResponse defines the structure of the response for Tieba hot list methods
type TiebaResponse struct {
	Code        int         `json:"code"`
	Name        string      `json:"name"`
	Title       string      `json:"title"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Link        string      `json:"link"`
	Total       int         `json:"total"`
	UpdateTime  string      `json:"update_time"`
	FromCache   bool        `json:"from_cache"`
	Data        []TiebaData `json:"data"`
}

// GetTiebaHotDiscussions fetches the hot discussion list for Tieba
func (c *Client) GetTiebaHotDiscussions() (*TiebaResponse, error) {
	var response TiebaResponse
	err := c.fetchHots("tieba", &response)
	return &response, err
}
