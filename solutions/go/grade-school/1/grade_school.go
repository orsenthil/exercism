package school

// Define the Grade and School types here.

type Grade struct {
	number   int
	students []string
}

type School struct {
	grades []Grade
}

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) {
	panic("Please implement the Add function")
}

func (s *School) Grade(level int) []string {
	panic("Please implement the Grade function")
}

func (s *School) Enrollment() []Grade {
	panic("Please implement the Enrollment function")
}
