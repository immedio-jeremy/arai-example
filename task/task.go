package task

type FrontendTasker interface {
	PerformFrontendTask()
}

type BackendTasker interface {
	PerformBackendTask()
}

type FullStackTasker interface {
	PerformFrontendTask()
	PerformBackendTask()
}

/* -------------------------------------------------------------------------- */

func PerformFrontendTask(ft FrontendTasker) {
	ft.PerformFrontendTask()
}

func PerformBackendTask(bt BackendTasker) {
	bt.PerformBackendTask()
}

func PerformFullStackTask(fst FullStackTasker) {
	fst.PerformBackendTask()
	fst.PerformFrontendTask()
}
