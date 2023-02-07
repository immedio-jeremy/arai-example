package developer

type FrontendDeveloper struct {
}

func (fd FrontendDeveloper) PerformFrontendTask() {

}

/* -------------------------------------------------------------------------- */

type BackendDeveloper struct {
}

func (bd BackendDeveloper) PerformBackendTask() {

}

/* -------------------------------------------------------------------------- */

type FullStackDeveloper struct {
	FrontendDeveloper // mix and match embedded structs to provide a custom set of functionality
	BackendDeveloper
}

/* -------------------------------------------------------------------------- */

type FrontendClone struct {
	FrontendDeveloper
}

func (fd FrontendClone) PerformFrontendTask() {
	fd.FrontendDeveloper.PerformFrontendTask() //can invoke developer implementation similar to inheritance
	// if you dont want to call it this implementation will hide embedded "base" implementation
	// do something extra because you are a clone
}

/* -------------------------------------------------------------------------- */

type FrontendRobot struct {
	// no embedded struct, completely independent implementation
}

func (fr FrontendRobot) PerformFrontendTask() {
	// completely independent implementation
}
