package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

// var message string

func Login(id string, name string) string {
	// jika input id dan name kosong
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}
	// jika imput id < 5 atau > 5 karakter
	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	for _, student := range Students {
		informasi := strings.Split(student, "_")
		if informasi[0] == id && informasi[1] == name {
			//  informasi diatas merupakan sebuah slice sehingga ambil bagiannya, informasi [2] => berisi major
			// inf := fmt.Sprintf("Login berhasil: %v (%v)", name, informasi[2]) ==> ini diganti kaya line 42
			inf := fmt.Sprintf("Login berhasil: %v", name)
			return inf
		}
	}

	// pindah return gagal kebawah
	return "Login gagal: data mahasiswa tidak ditemukan"
}

func Register(id string, name string, major string) string {
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}
	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}
	for _, student := range Students {
		data := strings.Split(student, "_")
		if data[0] == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	Students = append(Students, fmt.Sprintf("%s_%s_%s", id, name, major))
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

func GetStudyProgram(code string) string {
	program, cari := StudentStudyPrograms[code]
	if !cari {
		return "Kode program studi tidak ditemukan"
	}
	return program
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	var cari bool
	for i, student := range Students {
		data := strings.Split(student, "_")
		if data[1] == nama {
			fn(programStudi, &Students[i])
			cari = true
			break
		}
	}
	if cari {
		return "Program studi mahasiswa berhasil diubah."
	}
	return "Mahasiswa tidak ditemukan."

}

func UpdateStudyProgram(programStudi string, students *string) {
	data := strings.Split(*students, "_")
	data[2] = programStudi
	*students = strings.Join(data, "_")
}

// func Login(id string, name string) string {
// 	logDetail := id + "_" + name
// 	logDetailUp := strings.ToUpper(logDetail)

// 	// Jika input id dan name kosong, program menampilkan pesan "ID or Name is undefined!"
// 	if id == "" || name == "" {
// 		message = "ID or Name is undefined!"
// 		return message
// 	}

// 	// Jika ditemukan, fungsi akan mengembalikan pesan "Login berhasil: name"
// 	for _, student := range Students {
// 		student = strings.ToUpper(student[:len(student)-3])
// 		if student == logDetailUp {
// 			message = "Login berhasil: " + name
// 			return message
// 		}
// 	}

// 	// Jika tidak ditemukan, fungsi akan mengembalikan pesan "Login gagal: data mahasiswa tidak ditemukan".
// 	message = "Login gagal: data mahasiswa tidak ditemukan"
// 	return message
// }

// func Register(id string, name string, major string) string {
// 	newId := strings.ToUpper(id)

// 	// Jika input id, name dan major kosong, program menampilkan pesan "ID, Name or Major is undefined!"
// 	if id == "" || name == "" || major == "" {
// 		message = "ID, Name or Major is undefined!"
// 		return message
// 	}

// 	// Jika id telah digunakan, fungsi akan mengembalikan pesan "Registrasi gagal: id sudah digunakan".
// 	for _, student := range Students {
// 		studentId := strings.ToUpper(student[:5])
// 		if newId == studentId {
// 			message = "Registrasi gagal: id sudah digunakan"
// 			return message
// 		}
// 	}

// 	// Jika id belum pernah digunakan sebelumnya, fungsi akan menambahkan data mahasiswa baru ke dalam slice Students dan mengembalikan pesan "Registrasi berhasil: name (major)".
// 	newStudent := strings.ToUpper(id) + "_" + name + "_" + strings.ToUpper(major)
// 	Students = append(Students, newStudent)
// 	message = "Registrasi berhasil: " + name + " (" + strings.ToUpper(major) + ")"
// 	return message
// }

// func GetStudyProgram(code string) string {
// 	newCode := strings.ToUpper(code)

// 	// Jika ditemukan, fungsi akan mengembalikan nama program studi.
// 	for key, value := range StudentStudyPrograms {
// 		if newCode == key {
// 			message = value
// 			return message
// 		}
// 	}

// 	// Jika tidak ditemukan, fungsi akan mengembalikan pesan "Kode program studi tidak ditemukan".
// 	message = "Kode program studi tidak ditemukan"
// 	return message
// }

// func ModifyStudent(programStudi, nama string, fn studentModifier) string {
// 	newNama := strings.ToUpper(nama)
// 	// mencari data mahasiswa dengan nama yang sesuai pada slice Students, jika ada dan berhasil diubah  mengembalikan pesan "Program studi mahasiswa berhasil diubah."
// 	for i, student := range Students {
// 		studentName := strings.ToUpper(student[6 : len(student)-3])
// 		if newNama == studentName {
// 			UpdateStudyProgram(programStudi, &Students[i])
// 			message = "Program studi mahasiswa berhasil diubah."
// 			return message
// 		}
// 	}

// 	// Jika modifikasi gagal, maka program ini akan mengembalikan pesan "Mahasiswa tidak ditemukan.".
// 	message = "Mahasiswa tidak ditemukan."
// 	return message
// }

// func UpdateStudyProgram(programStudi string, students *string) {
// 	parts := strings.Split(*students, "_")
// 	parts[2] = strings.ToUpper(programStudi)
// 	*students = strings.Join(parts, "_")
// }

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

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
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
