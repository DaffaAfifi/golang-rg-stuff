package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	students             []model.Student
	studentStudyPrograms map[string]string
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
	}
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	data := sm.students
	return data
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	if id == "" || name == "" {
		return "", errors.New("ID or Name is undefined!")
	}

	for _, student := range sm.students {
		if id == student.ID && name == student.Name {
			message := "Login berhasil: " + name
			return message, nil
		}
	}

	return "", errors.New("Login gagal: data mahasiswa tidak ditemukan")
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	if id == "" || name == "" || studyProgram == "" {
		return "", errors.New("ID, Name or StudyProgram is undefined!")
	}

	checkExistSp := false
	for key := range sm.studentStudyPrograms {
		if key == studyProgram {
			checkExistSp = true
		}
	}

	if !checkExistSp {
		return "", errors.New("Study program " + studyProgram + " is not found")
	}

	for _, student := range sm.students {
		if id == student.ID {
			return "", errors.New("Registrasi gagal: id sudah digunakan")
		}
	}

	sm.students = append(sm.students,
		model.Student{
			ID:           id,
			Name:         name,
			StudyProgram: studyProgram,
		},
	)
	message := "Registrasi berhasil: " + name + " (" + studyProgram + ")"
	return message, nil
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	if code == "" {
		return "", errors.New("Code is undefined!")
	}

	for key, value := range sm.studentStudyPrograms {
		if key == code {
			return value, nil
		}
	}

	return "", errors.New("Kode program studi tidak ditemukan")
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	for i, student := range sm.students {
		if name == student.Name {
			if err := fn(&sm.students[i]); err != nil {
				return "", err
			}
			return "Program studi mahasiswa berhasil diubah.", nil
		}
	}

	return "", errors.New("Mahasiswa tidak ditemukan.")
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		for key := range sm.studentStudyPrograms {
			if key == programStudi {
				s.StudyProgram = programStudi
				return nil
			}
		}
		return errors.New("Kode program studi tidak ditemukan")
	}
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "5":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
