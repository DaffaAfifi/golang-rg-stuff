package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

// var message string

func Login(id string, name string) string {

	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	// student := make(map[string]string)
	pisahinputan := strings.Split(Students, ", ")
	// fmt.Println(pisahinputan)

	for _, data := range pisahinputan {
		pisahdata := strings.Split(data, "_")
		if pisahdata[0] == id && pisahdata[1] == name {
			return fmt.Sprintf("Login berhasil: %v (%v)", name, pisahdata[2])
		}
	}

	return "Login gagal: data mahasiswa tidak ditemukan"

}

func Register(id string, name string, major string) string {

	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}
	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	pisahinputan := strings.Split(Students, ",")
	for _, data2 := range pisahinputan {
		pisahdata2 := strings.Split(data2, "_")
		studentid := pisahdata2[0]
		// studentname := pisahdata2[1]
		// Major := pisahdata2[2]
		if studentid == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}
	Students += fmt.Sprintf(",%v_%v_%v", id, name, major)
	return fmt.Sprintf("Registrasi berhasil: %v (%v)", name, major)
}

func GetStudyProgram(code string) string {
	if code == "" {
		return "Code is undefined!"
	}

	pisahketerangan := strings.Split(StudentStudyPrograms, ",")
	for _, keterangan1 := range pisahketerangan {
		pisahketerangan2 := strings.Split(keterangan1, "_")
		ket1 := pisahketerangan2[0]
		ket2 := pisahketerangan2[1]
		if strings.Contains(ket1, code) {
			return ket2
		}
	}
	return "Code is undefined!"
}

// func Login(id string, name string) string {
// 	var student string
// 	var studentsWoSpace string
// 	logDetail := id + "_" + name
// 	logDetailLow := strings.ToLower(logDetail)

// 	// Jika input id dan name kosong, program menampilkan pesan "ID or Name is undefined!"
// 	if id == "" || name == "" {
// 		message = "ID or Name is undefined!"
// 		return message
// 	}

// 	// Jika input id kurang atau lebih dari 5 karakter, program akan menampilkan pesan "ID must be 5 characters long!"
// 	if len(id) != 5 {
// 		message = "ID must be 5 characters long!"
// 		return message
// 	}

// 	// Delete space from Students
// 	for i := 0; i < len(Students); i++ {
// 		if string(Students[i]) != " " {
// 			studentsWoSpace += string(Students[i])
// 		}
// 	}

// 	// Jika data terdaftar, maka program akan menampilkan pesan login berhasil beserta nama dan jurusan mahasiswa. Format "Login berhasil: <nama> (<jurusan mahasiswa>)" contoh: "Login berhasil: Aditira (TI)"
// 	for i := 0; i < len(studentsWoSpace); i++ {
// 		if string(studentsWoSpace[i]) == "," {
// 			newStudent := strings.ToLower(student[:len(student)-3])
// 			major := student[len(student)-2:]
// 			if logDetailLow == newStudent {
// 				message = "Login berhasil: " + name + " (" + major + ")"
// 				return message
// 			}
// 			student = ""
// 		} else {
// 			student += string(studentsWoSpace[i])
// 		}
// 	}

// 	if student != "" {
// 		newStudent := strings.ToLower(student[:len(student)-3])
// 		major := student[len(student)-2:]
// 		if logDetailLow == newStudent {
// 			message = "Login berhasil: " + name + " (" + major + ")"
// 			return message
// 		}
// 	}

// 	// Jika data tidak terdaftar, program akan menampilkan pesan "Login gagal: data mahasiswa tidak ditemukan".
// 	message = "Login gagal: data mahasiswa tidak ditemukan"
// 	return message
// }

// func Register(id string, name string, major string) string {
// 	var student string
// 	var studentsWoSpace string
// 	newId := strings.ToUpper(id)
// 	var varId string

// 	// Jika input id, name dan major kosong, program menampilkan pesan "ID, Name or Major is undefined!"
// 	if id == "" || name == "" || major == "" {
// 		message = "ID, Name or Major is undefined!"
// 		return message
// 	}

// 	// Jika input id kurang atau lebih dari 5 karakter, program akan menampilkan pesan "ID must be 5 characters long!"
// 	if len(id) != 5 {
// 		message = "ID must be 5 characters long!"
// 		return message
// 	}

// 	// Delete space from Students
// 	for i := 0; i < len(Students); i++ {
// 		if string(Students[i]) != " " {
// 			studentsWoSpace += string(Students[i])
// 		}
// 	}

// 	// Jika ID sudah terdaftar, maka program akan menampilkan pesan "Registrasi gagal: id sudah digunakan".
// 	for i := 0; i < len(studentsWoSpace); i++ {
// 		if string(studentsWoSpace[i]) == "," {
// 			varId = student[:5]
// 			studentId := strings.ToUpper(varId)
// 			if newId == studentId {
// 				message = "Registrasi gagal: id sudah digunakan"
// 				return message
// 			}
// 			student = ""
// 		} else {
// 			student += string(studentsWoSpace[i])
// 		}
// 	}

// 	if student != "" {
// 		varId = student[:5]
// 		studentId := strings.ToUpper(varId)
// 		if newId == studentId {
// 			message = "Registrasi gagal: id sudah digunakan"
// 			return message
// 		}
// 	}

// 	// Jika ID belum terdaftar, maka program akan menambahkan data mahasiswa baru ke dalam variabel Students dan menampilkan pesan registrasi berhasil. Format "Registrasi berhasil: <nama> (<jurusan mahasiswa>)" contoh: "Registrasi berhasil: Aditira (TI)"
// 	Students += ", " + strings.ToUpper(id) + "_" + name + "_" + strings.ToUpper(major)
// 	message = "Registrasi berhasil: " + name + " (" + strings.ToUpper(major) + ")"
// 	return message
// }

// func GetStudyProgram(code string) string {
// 	if code != "" {
// 		programs := strings.Split(StudentStudyPrograms, ", ")
// 		for _, program := range programs {
// 			if strings.Contains(program, code) {
// 				programDetail := strings.Split(program, "_")
// 				return programDetail[1]
// 			}
// 		}
// 	}
// 	return "Code is undefined!"
// }

// func GetStudyProgram(code string) string {
// 	var studyProgram string
// 	var studyProgramWoSpace string
// 	newCode := strings.ToLower(code)

// 	// Jika input code kosong, program menampilkan pesan "Code is undefined!"
// 	if code == "" {
// 		message = "Code is undefined!"
// 		return message
// 	}

// 	// Delete space from StudentStudyPrograms
// 	for i := 0; i < len(StudentStudyPrograms); i++ {
// 		if string(StudentStudyPrograms[i]) != " " {
// 			studyProgramWoSpace += string(StudentStudyPrograms[i])
// 		}
// 	}

// 	// Jika kode dari program studi ditemukan, maka program akan menampilkan nama program studi. Contoh input "TI" maka keluar output "Teknik Informatika"
// 	for i := 0; i < len(studyProgramWoSpace); i++ {
// 		if string(studyProgramWoSpace[i]) == "," {
// 			var finalStudyProgram string
// 			studyProgramCode := strings.ToLower(studyProgram[:2])
// 			if newCode == studyProgramCode {
// 				newStudyProgram := studyProgram[3:]
// 				for i := 0; i < len(newStudyProgram); i++ {
// 					if i > 0 && unicode.IsUpper(rune(newStudyProgram[i])) {
// 						finalStudyProgram += " " + string(newStudyProgram[i])
// 					} else {
// 						finalStudyProgram += string(newStudyProgram[i])
// 					}
// 				}
// 				message = finalStudyProgram
// 				return message
// 			}
// 			studyProgram = ""
// 		} else {
// 			studyProgram += string(studyProgramWoSpace[i])
// 		}
// 	}

// 	if studyProgram != "" {
// 		var finalStudyProgram string
// 		studyProgramCode := strings.ToLower(studyProgram[:2])
// 		if newCode == studyProgramCode {
// 			newStudyProgram := studyProgram[3:]
// 			for i := 0; i < len(newStudyProgram); i++ {
// 				if i > 0 && unicode.IsUpper(rune(newStudyProgram[i])) {
// 					finalStudyProgram += " " + string(newStudyProgram[i])
// 				} else {
// 					finalStudyProgram += string(newStudyProgram[i])
// 				}
// 			}
// 			message = finalStudyProgram
// 			return message
// 		}
// 	}

// 	// Jika tidak ada kode yang dicari
// 	message = "Pencarian gagal: kode program studi tidak ditemukan "
// 	return message
// }

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
