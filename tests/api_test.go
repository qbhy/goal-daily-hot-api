package tests

import (
	"fmt"
	dailyhotapi "github.com/qbhy/goal-daily-hot-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTenAPI(t *testing.T) {
	api := dailyhotapi.New("http://localhost:6688")

	weiboRes, err := api.GetWeiboHotSearch()
	assert.NoError(t, err, err)
	fmt.Println(weiboRes)
}
