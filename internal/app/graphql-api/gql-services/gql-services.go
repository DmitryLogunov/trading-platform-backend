package gqlServices

type GqlServices struct {
	JobsService    *JobsService
	TradingService *TradingService
}

func (gs *GqlServices) Init() {
	gs.JobsService = &JobsService{}
	gs.TradingService = &TradingService{}
}
