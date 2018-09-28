package assertions

type Asserts struct {
	JsonAsserts *[]JsonAssert
}

type JsonAssert struct {
	path   string
	Expect string
}
