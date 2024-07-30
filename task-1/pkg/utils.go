package pkg

import (
	"fmt"
)

func PrintWelcomeMessage() {
	fmt.Println()
	fmt.Println()
	fmt.Println("|------------------------------------------------------------------|")
	fmt.Println("|                                                                  |")
	fmt.Println("|           Welcome to Student Grade Calculator                    |")
	fmt.Println("|                                                                  |")
	fmt.Println("|------------------------------------------------------------------|")
	fmt.Println()
	fmt.Println()
}

func PrintClosingMessage(name string, subjects *map[string]int) {
	fmt.Println()
	fmt.Println()
	fmt.Println("|------------------------------------------------------------------|")
	fmt.Println("|                                                                  |")
	fmt.Println("|                              Results                             |")
	fmt.Printf("|                    Name: %-40v|\n", name)
	fmt.Println("|                                                                  |")
	fmt.Println("|------------------------------------------------------------------|")
	fmt.Println()
	fmt.Println()
	fmt.Println("Subjects:")

	i := 1
	for subject, grade := range *subjects {
		fmt.Printf("%3v) %-10v ................................. %v \n", i, subject, grade)
		i++
	}

	fmt.Println()
	fmt.Println()
	fmt.Printf("Average: %v", CalculateAverage(subjects))
	fmt.Println()
	fmt.Println()
}

func GetNameFromUser() string {
	fmt.Print("Enter your name: ")
	input := stringInput()

	fmt.Println()

	return input
}

func GetSubjects() map[string]int {
	totalSubjects := getNumberOfSubjects()
	subjects := getDetailsOfSubjects(totalSubjects)

	return subjects
}

func getNumberOfSubjects() int {
	fmt.Print("Enter number of subjects: ")

	totalSubjects := getIntInput()

	fmt.Println()

	return totalSubjects
}

func getDetailsOfSubjects(totalSubjects int) map[string]int {
	fmt.Println("Insert the details of each subject based on the questions asked.")

	subjects := make(map[string]int)

	for i := range totalSubjects {
		subject, grade := getSubjectAndGrade(i+1, &subjects)
		subjects[subject] = grade
		fmt.Println()
	}

	fmt.Println()

	return subjects
}

func getSubjectAndGrade(subjectNumber int, subjects *map[string]int) (string, int) {
	fmt.Printf("Enter name of subject %v: ", subjectNumber)

	subjectName := stringInput()

	for _, ok := (*subjects)[subjectName]; ok; {
		fmt.Println()
		fmt.Println("🛑 Make sure that the subject is unique. Try again.")
		fmt.Println()

		fmt.Printf("Re-enter name of subject %v: ", subjectNumber)
		subjectName = stringInput()
	}
	fmt.Println()

	fmt.Printf("Enter grade of %v: ", subjectName)

	grade := getIntInput()

	for grade < 0 || grade > 100 {
		fmt.Println()
		fmt.Println("🛑 Make sure that the value of the grade is between 0 and 100. Try again.")
		fmt.Println()

		fmt.Printf("Re-enter grade of %v: ", subjectName)
		grade = getIntInput()
		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("✅ Finished inserting subject no. %v with name %v \n", subjectNumber, subjectName)
	fmt.Println()

	return subjectName, grade
}

func CalculateAverage(subjects *map[string]int) float32 {
	if len(*subjects) == 0 {
		return 0
	}

	var sum float32
	for _, grade := range *subjects {
		sum += float32(grade)
	}

	return sum / float32(len(*subjects))
}
