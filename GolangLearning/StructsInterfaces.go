package golanglearning

import (
	"fmt"
)


type Student interface {
	GetName() string
	GetAge() int
	GetGrade() float64
	SetGrade(grade float64)
	IsPassing() bool
	GetTotalCredits() int
	IncrementCredits(credits int)
	GetGPA() float64

}



type HighSchoolStudent struct {
	name  string
	age   int
	grade float64
	totalCredits int 
	gpa float64

}


func NewHighSchoolStudent(name string, age int) *HighSchoolStudent {
	return &HighSchoolStudent{
		name:  name,
		age:   age,
		grade: 0.0,
	}
}


func (s *HighSchoolStudent) GetName() string {
	return s.name
}

func (s *HighSchoolStudent) GetAge() int {
	return s.age
}

func (s *HighSchoolStudent) GetGrade() float64 {
	return s.grade
}

func (s *HighSchoolStudent) SetGrade(grade float64) {
	s.grade = grade
}

func (s *HighSchoolStudent) IsPassing() bool {
	return s.grade >= 60.0
}
func (s * HighSchoolStudent) GetTotalCredits() int {
	return s.totalCredits
}
func (s * HighSchoolStudent) IncrementCredits(credits int) {
	s.totalCredits+=credits
}
func (s * HighSchoolStudent) GetGPA() float64 {
	return s.gpa
}
func (s * HighSchoolStudent) SetGPA(gpa float64) {
	s.gpa=gpa
}




type School struct {
	students []Student
}

func NewSchool() *School {
	return &School{
		students: make([]Student, 0),
	}
}

func (s *School) AddStudent(student Student) {
	s.students = append(s.students, student)
}

func (s *School) GetPassingStudents() []Student {
	passing := make([]Student, 0)
	for _, student := range s.students {
		if student.IsPassing() {
			passing = append(passing, student)
		}
	}
	return passing
}

func SimulateSchool() {
	// Create a new school
	school := NewSchool()

	// Create some students
	student1 := NewHighSchoolStudent("John", 16)
	student2 := NewHighSchoolStudent("Alice", 15)
	student3 := NewHighSchoolStudent("Bob", 17)

	// Set grades
	student1.SetGrade(85.5)
	student2.SetGrade(45.0)
	student3.SetGrade(92.0)

	// Add students to school
	school.AddStudent(student1)
	school.AddStudent(student2)
	school.AddStudent(student3)

	// Get passing students
	passingStudents := school.GetPassingStudents()

	// Print results
	fmt.Println("School Management System")
	fmt.Println("----------------------")
	fmt.Printf("Total Students: %d\n", len(school.students))
	fmt.Printf("Passing Students: %d\n", len(passingStudents))

	fmt.Println("\nStudent Details:")
	for _, student := range school.students {
		fmt.Printf("Name: %s, Age: %d, Grade: %.1f, Passing: %v\n",
			student.GetName(),
			student.GetAge(),
			student.GetGrade(),
			student.IsPassing())
	}
}
