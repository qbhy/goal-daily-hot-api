package daily_hot_api

import "github.com/goal-web/contracts"

func NewService() contracts.ServiceProvider {
	return Service{}
}

type Service struct {
}

func (s Service) Register(application contracts.Application) {
	application.Singleton("dailyHotAPI", func(config contracts.Config) DailyHotAPI {
		conf, _ := config.Get("daily-hot-api").(Config)
		return New(conf.BaseUrl)
	})
}

func (s Service) Start() error {
	return nil
}

func (s Service) Stop() {
}
