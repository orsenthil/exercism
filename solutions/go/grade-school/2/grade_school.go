package school

import "sort"

// Define the Grade and School types here.

type Grade struct {
	number   int
	students []string
}

type School struct {
	grades []Grade
}

func New() *School {
	return &School{grades: make([]Grade, 0)}
}

func (s *School) Add(student string, grade int) {
	for i := 0; i < len(s.grades); i++ {
		if s.grades[i].number == grade {
			s.grades[i].students = append(s.grades[i].students, student)
			return
		}
	}
	s.grades = append(s.grades, Grade{number: grade, students: []string{student}})
}

func (s *School) Grade(level int) []string {
	for i := 0; i < len(s.grades); i++ {
		if s.grades[i].number == level {
			sort.Strings(s.grades[i].students)
			return s.grades[i].students
		}
	}

	return nil
}

func (s *School) Enrollment() []Grade {
	sort.Slice(s.grades, func(i, j int) bool {
		return s.grades[i].number < s.grades[j].number
	})

	for i := 0; i < len(s.grades); i++ {
		sort.Strings(s.grades[i].students)
	}

	return s.grades
}
