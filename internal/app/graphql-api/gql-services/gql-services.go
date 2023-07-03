package gqlServices

type GqlServices struct {
	AlertsService    *AlertsService
	JobsService      *JobsService
	TradingService   *TradingService
	PositionsService *PositionsService
}

func (gs *GqlServices) Init() {
	gs.AlertsService = &AlertsService{}
	gs.JobsService = &JobsService{}
	gs.TradingService = &TradingService{}
	gs.PositionsService = &PositionsService{}
}
