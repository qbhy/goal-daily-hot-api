package daily_hot_api

// AcFunData defines the structure of individual hot list data items for AcFun
type AcFunData struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Cover     string `json:"cover"`
	Author    string `json:"author"`
	Hot       int    `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobile_url"`
}

// AcFunParameType defines the structure of parameter types in the AcFun response
type AcFunParameType struct {
	Name string            `json:"name"`
	Type map[string]string `json:"type"`
}

// AcFunParameData defines the structure of parameter data in the AcFun response
type AcFunParameData struct {
	Type  AcFunParameType `json:"type"`
	Range AcFunParameType `json:"range"`
}

// AcFunResponse defines the structure of the response for AcFun hot list methods
type AcFunResponse struct {
	Code        int             `json:"code"`
	Name        string          `json:"name"`
	Title       string          `json:"title"`
	Type        string          `json:"type"`
	Description string          `json:"description"`
	ParameData  AcFunParameData `json:"parameData"`
	Link        string          `json:"link"`
	Total       int             `json:"total"`
	UpdateTime  string          `json:"update_time"`
	FromCache   bool            `json:"from_cache"`
	Data        []AcFunData     `json:"data"`
}

// GetAcFunRanking fetches the ranking list for AcFun
func (c *Client) GetAcFunRanking() (*AcFunResponse, error) {
	var response AcFunResponse
	err := c.fetchHots("acfun", &response)
	return &response, err
}
