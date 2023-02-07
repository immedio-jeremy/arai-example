package task

type FrontendTasker interface {
	PerformFrontendTask()
}

type BackendTasker interface {
	PerformBackendTask()
}

type FullStackTasker interface { // independent of frontend and backend interfaces. this could live in a different package even.
	PerformFrontendTask()
	PerformBackendTask()
}

/* -------------------------------------------------------------------------- */

// notice that the consumer, not the implementation's package is where the interface is implemented.

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
