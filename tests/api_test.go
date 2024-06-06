package tests

import (
	"fmt"
	dailyhotapi "github.com/qbhy/goal-daily-hot-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTenAPI(t *testing.T) {
	api := dailyhotapi.New("http://localhost:6688")

	bilibiliRes, weiboErr := api.GetWeiboHotSearch()
	assert.NoError(t, weiboErr, weiboErr)
	fmt.Println(bilibiliRes)

	acfun, acfunErr := api.GetAcFunRanking()
	assert.NoError(t, acfunErr, acfunErr)
	fmt.Println(acfun)

	weiboRes, err := api.GetWeiboHotSearch()
	assert.NoError(t, err, err)
	fmt.Println(weiboRes)

	zhihuRes, zhihuErr := api.GetZhihuHotList()
	assert.NoError(t, zhihuErr, zhihuErr)
	fmt.Println(zhihuRes)

	zhihuRecommendsRes, err := api.GetZhihuDailyRecommendations()
	assert.NoError(t, err, err)
	fmt.Println(zhihuRecommendsRes)

	baiduRes, err := api.GetBaiduHotSearch()
	assert.NoError(t, err, err)
	fmt.Println(baiduRes)

	toutiaoRes, err := api.GetToutiaoHotList()
	assert.NoError(t, err, err)
	fmt.Println(toutiaoRes)

	qqRes, err := api.GetQQNewsHotList()
	assert.NoError(t, err, err)
	fmt.Println(qqRes)

	neteaseRes, err := api.GetNeteaseNewsHotList()
	assert.NoError(t, err, err)
	fmt.Println(neteaseRes)

	thepaperRes, err := api.GetThePaperHotList()
	assert.NoError(t, err, err)
	fmt.Println(thepaperRes)

}
