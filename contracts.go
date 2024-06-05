package daily_hot_api

type DailyHotAPI interface {
	GetBilibiliHotList() (*BilibiliResponse, error)
	GetAcFunRanking() (*AcFunResponse, error)
	GetWeiboHotSearch() (*WeiboResponse, error)
	GetZhihuHotList() (*ZhihuResponse, error)
	GetZhihuDailyRecommendations() (*ZhihuDailyResponse, error)
	GetBaiduHotSearch() (*BaiduResponse, error)
}
