package gqlServices

type GqlServices struct {
	JobsService  *JobsService
	PostsService *PostsService
}

func (gs *GqlServices) Init() {
	gs.JobsService = &JobsService{}
	gs.PostsService = &PostsService{}
}
