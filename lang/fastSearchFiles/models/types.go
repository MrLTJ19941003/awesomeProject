package typesF

type WorkerResult struct {
	FilesName []Workers
	FindsName []string
}

type WorderFuncs func (workers Workers) (WorkerResult,error)

type Workers struct {
	Fliepath string
	WorderFunc WorderFuncs
}
