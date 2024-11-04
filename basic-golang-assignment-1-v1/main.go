package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	splitStudents := strings.Split(Students, ", ")

	if id == "" || name == "" {
		return "ID or Name is undefined!"
	} else if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	for _, student := range splitStudents {
		studentID := strings.Split(student, "_")[0]
		studentName := strings.Split(student, "_")[1]
		if studentID == id && studentName == name {
			studentMajor := strings.Split(student, "_")[2]
			return "Login berhasil: " + name + " (" + studentMajor + ")"
		}
	}

	return "Login gagal: data mahasiswa tidak ditemukan"
}

func Register(id string, name string, major string) string {
	splitStudents := strings.Split(Students, ", ")

	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	} else if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	for _, student := range splitStudents {
		studentID := strings.Split(student, "_")[0]
		if studentID == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	Students += ", " + id + "_" + name + "_" + major
	return "Registrasi berhasil: " + name + " " + "(" + major + ")"
}

func GetStudyProgram(code string) string {
	splitStudyPrograms := strings.Split(StudentStudyPrograms, ", ")
	if code == "" {
		return "Code is undefined!"
	}

	for _, studyProgram := range splitStudyPrograms {
		programCode := strings.Split(studyProgram, "_")[0]
		programName := strings.Split(studyProgram, "_")[1]
		if programCode == code {
			return programName
		}
	}
	return ""
}

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
