package gqlServices

type GqlServices struct {
	AlertsService  *AlertsService
	JobsService    *JobsService
	TradingService *TradingService
}

func (gs *GqlServices) Init() {
	gs.AlertsService = &AlertsService{}
	gs.JobsService = &JobsService{}
	gs.TradingService = &TradingService{}
}
