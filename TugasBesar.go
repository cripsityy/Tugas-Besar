package main

import "fmt"

// struktur dan variabel global
type akun struct {
	username string
	password string
	id       string
}
type teman struct {
	arrchat      [10]chat
	rc           [10]int
	nChat        int
	nteman       int
	nama1, nama2 string
}
type chat struct {
	pengirim string
	isichat  isiChat
	nchat    int
}
type isiChat struct {
	pengirim string
	IsiChat  string
}
type grup struct {
	nama    string
	ngrup   int
	anggota [10]anggotagrup
	arrChat [10]chat
	nChat   int
}
type anggotagrup struct {
	name1 string
	name2 string
	nName int
}
type akunBaru [1000]akun
type regis struct {
	username string
	password string
	id       string
}
type acc [1000]regis
type Teman [10]teman
type akungrup [10]grup
type akunTeman [10]teman

var inde int
var Friend teman
var nT int
var friend akunTeman
var akunLogin akun
var T akungrup
var ng int
var newUser akunBaru
var nData int
var terima acc
var nAcc int

func main() {
	var opsi int
	fmt.Println("------------------------")
	fmt.Println("     SELAMAT DATANG     ")
	fmt.Println("------------------------")
	fmt.Println("1. User")
	fmt.Println("2. Admin")
	fmt.Println("3. Keluar Aplikasi")
	fmt.Println("------------------------")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		User()
	case 2:
		Admin()
	case 3:
		fmt.Println("Terimakasih sudah menggunakan aplikasi kami.")
		break
	}
}

func User() {
	var opsi int
	fmt.Println("-----------------------")
	fmt.Println("        U S E R        ")
	fmt.Println("-----------------------")
	fmt.Println("1.Registrasi")
	fmt.Println("2.Login")
	fmt.Println("3.Exit")
	fmt.Println("-----------------------")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		register(&newUser, &nData)
	case 2:
		login()
	case 3:
		main()
	}

}

func register(A *akunBaru, n *int) {
	fmt.Print("Buat Id: ")
	fmt.Scan(&A[*n].id)
	for i := 0; i < *n; i++ {
		if A[i].id == A[*n].id {
			fmt.Print("Buat Id: ")
			fmt.Scan(&A[*n].id)
		}
	}
	fmt.Print("Buat username: ")
	fmt.Scan(&A[*n].username)
	fmt.Print("Buat Password: ")
	fmt.Scan(&A[*n].password)
	*n++
	User()
}

func login() {
	var x, y, z string
	fmt.Print("Masukkan id: ")
	fmt.Scan(&z)
	fmt.Print("Masukkan username: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan Passwod: ")
	fmt.Scan(&y)
	if mencariAkuni(terima, nAcc, x, y, z) {
		fmt.Println("Berhasil login")
		akunLogin.id = z
		akunLogin.username = x
		Login()

	} else {
		fmt.Println("Akun anda tidak ditemukan.")
		User()
	}

}

func mencariAkuni(B acc, nAcc int, x, y, z string) bool {
	var i int
	var idx bool
	idx = false
	for i = 0; i < nAcc; i++ {
		if B[i].id == z && B[i].username == x && B[i].password == y {
			idx = true
		}
	}
	return idx
}

func Admin() {
	var opsi int
	fmt.Println("-----------------------")
	fmt.Println("       A D M I N       ")
	fmt.Println("-----------------------")
	fmt.Println("1.Daftar Akun Terbuat")
	fmt.Println("2.Penerimaan Akun")
	fmt.Println("3.Daftar Akun Diterima")
	fmt.Println("4.Exit")
	fmt.Println("-----------------------")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		cetakData(newUser, nData)
	case 2:
		admin(&newUser, &nData, &terima, &nAcc)
	case 3:
		cetakDataTerima(terima, nAcc)
	case 4:
		main()
	}
}

func cetakData(A akunBaru, n int) {
	fmt.Println("Daftar Akun Terbuat: ")
	for i := 0; i < n; i++ {
		fmt.Println(A[i].id, A[i].username, A[i].password)
	}
	Admin()
}

func mencariId(A akunBaru, n int, x string) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if A[i].id == x {
			idx = i
		}
	}
	return idx
}

func admin(A *akunBaru, n *int, B *acc, nAcc *int) {
	var x string
	var idx int
	fmt.Print("Pilih akun yang akan diterima: ")
	fmt.Scan(&x)
	idx = mencariId(*A, *n, x)
	if idx == -1 {
		fmt.Println("Id tidak ditemukan")
	} else {

		B[*nAcc].username, B[*nAcc].password, B[*nAcc].id = A[idx].username, A[idx].password, A[idx].id
		*nAcc++
	}
	fmt.Println("Data teregistrasi")
	Admin()
}

func cetakDataTerima(B acc, n int) {
	fmt.Println("Daftar Akun Diterima: ")
	for i := 0; i < n; i++ {
		fmt.Println(B[i].id, B[i].username, B[i].password)
	}
	Admin()
}

func Login() {
	var opsi int
	fmt.Println("------------------------")
	fmt.Println("    H O M E  P A G E    ")
	fmt.Println("------------------------")
	fmt.Println("1. Chat Pribadi")
	fmt.Println("2. Chat Group")
	fmt.Println("3. Logout")
	fmt.Println("4. Keluar Aplikasi")
	fmt.Println("------------------------")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		ChatPribadi()
	case 2:
		ChatGrup()
	case 3:
		User()
	case 4:
		fmt.Println("Terimakasih sudah menggunakan aplikasi kami.")
	}
}

func ChatPribadi() {
	var opsi int
	fmt.Println("------------------------------")
	fmt.Println("    C H A T  P R I B A D I    ")
	fmt.Println("------------------------------")
	fmt.Println("1. Tambah Teman")
	fmt.Println("2. Daftar Roomchat Teman")
	fmt.Println("3. Chat Teman")
	fmt.Println("4. Exit")
	fmt.Println("------------------------------")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		tambahTeman(&terima, &nAcc, &friend, &nT)
	case 2:
		dataTeman(friend, nT)
	case 3:
		chatTeman(&friend, nT)
	case 4:
		Login()
	}

}

func tambahTeman(B *acc, nAcc *int, C *akunTeman, nTeman *int) {
	var id string
	var idx int
	fmt.Print("Masukkan id yang akan ditambahkan: ")
	fmt.Scan(&id)
	idx = MencariId(*B, *nAcc, id)
	if idx == -1 {
		fmt.Println("Id tidak ditemukan")
	} else {
		C[*nTeman].nama1 = akunLogin.username
		C[*nTeman].nama2 = B[idx].username
		*nTeman++

		fmt.Println("Teman ditambahkan")

	}
	ChatPribadi()
}

func dataTeman(C akunTeman, nTeman int) {
	fmt.Println("Daftar Roomchat")
	for i := 0; i < nTeman; i++ {
		fmt.Println(i+1, C[i].nama1, "dan", C[i].nama2)
	}
	ChatPribadi()
}

func chatTeman(C *akunTeman, nTeman int) {
	var x akunTeman
	var y int
	var idx int
	fmt.Println("Daftar Roomchat")
	for i := 0; i < nTeman; i++ {
		fmt.Println(i+1, C[i].nama1, "dan", C[i].nama2)
	}
	fmt.Print("Masukkan pasangan roomchat secara berurutan: ")
	fmt.Scan(&x[nTeman].nama1, &x[nTeman].nama2)
	idx = mencariTemani(*C, nTeman, x[nTeman].nama1, x[nTeman].nama2)
	if mencariTeman(friend, nT, x[nTeman].nama1, x[nTeman].nama2) {
		for {
			if idx == -1 {
				fmt.Println("Roomchat tidak ditemukan.")
				ChatPribadi()
			} else {
				nTeman = idx
				fmt.Println("------------------------")
				fmt.Println("1. Kirim Pesan")
				fmt.Println("2. Tampilkan Pesan")
				fmt.Println("3. Exit")
				fmt.Println("------------------------")
				fmt.Print("Pilih Menu: ")
				fmt.Scan(&y)
				switch y {
				case 1:

					fmt.Scan(&C[nTeman].arrchat[C[nTeman].nChat].isichat.IsiChat)

					C[nTeman].arrchat[C[nTeman].nChat].isichat.pengirim = akunLogin.username
					C[nTeman].nChat++

					for i := 0; i < C[nTeman].nChat; i++ {
						fmt.Printf("%s: %s\n", C[nTeman].arrchat[i].isichat.pengirim, C[nTeman].arrchat[i].isichat.IsiChat)
					}
				case 2:
					for i := 0; i < C[nTeman].nChat; i++ {
						fmt.Printf("%s: %s\n", C[nTeman].arrchat[i].isichat.pengirim, C[nTeman].arrchat[i].isichat.IsiChat)
					}
				case 3:
					ChatPribadi()
				}
			}
		}
	} else {
		fmt.Println("Roomchat tidak ditemukan.")
		ChatPribadi()
	}
}

func mencariTeman(C akunTeman, n int, x, y string) bool {
	var i int
	var idx bool
	idx = false
	for i = 0; i < n; i++ {
		if C[i].nama2 == y && C[i].nama1 == x {
			idx = true
		}
	}
	return idx
}
func mencariTemani(C akunTeman, n int, x, y string) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if C[i].nama2 == y && C[i].nama1 == x {
			idx = i
		}
	}
	return idx
}

func MencariId(B acc, nAcc int, x string) int {
	var i, idx int
	idx = -1
	for i = 0; i < nAcc; i++ {
		if B[i].id == x {
			idx = i
		}
	}
	return idx
}

func ChatGrup() {
	var opsi int
	fmt.Println("--------------------------")
	fmt.Println("    C H A T  G R O U P    ")
	fmt.Println("--------------------------")
	fmt.Println("1. Buat Group")
	fmt.Println("2. Daftar Group")
	fmt.Println("3. Chat Group")
	fmt.Println("4. Exit")
	fmt.Println("--------------------------")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&opsi)
	switch opsi {
	case 1:
		buatGrup(&T, &ng)
	case 2:
		daftarGrup(T, ng)
	case 3:
		mauchatgrup(&T, ng)
	case 4:
		Login()
	}
}

//membuat grup
func buatGrup(G *akungrup, nGrup *int) {
	fmt.Println("Masukkan Nama Group: ")
	fmt.Scan(&G[*nGrup].nama)
	ng++
	ttg(&terima, &nAcc, &T, ng-1)
}

//memasukkan anggota grup
func ttg(z *acc, nAcc *int, G *akungrup, nGrup int) {
	var id string
	var x int
	var idx int
	for {
		fmt.Println("Menambahkan teman? 1/2?")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		fmt.Scan(&x)
		switch x {
		case 1:
			fmt.Print("Masukkan id yang akan ditambahkan: ")
			fmt.Scan(&id)
			idx = MencariId(*z, *nAcc, id)
			if idx == -1 {
				fmt.Println("Id tidak ditemukan")
			} else {
				G[nGrup].anggota[G[nGrup].ngrup].name1 = akunLogin.username
				G[nGrup].anggota[G[nGrup].ngrup].name2 = (*z)[idx].username
				G[nGrup].ngrup++
				fmt.Println("Anggota Group Ditambahkan")
			}
		case 2:
			ChatGrup()
		}
	}
}

func daftarGrup(G akungrup, nGrup int) {
	fmt.Println("Daftar Group")
	for i := 0; i < nGrup; i++ {
		fmt.Printf("%d. Grup %s:\n", i+1, G[i].nama)
		fmt.Println("anggota: ", G[i].anggota[G[nGrup].ngrup].name1)
		for j := 0; j < G[i].ngrup; j++ {
			fmt.Println("anggota: ", G[i].anggota[j].name2)
		}
	}
	ChatGrup()
}

func mauchatgrup(G *akungrup, nGrup int) {
	var x akungrup
	var y int
	var idx int
	fmt.Println("Daftar Group")
	for i := 0; i < nGrup; i++ {
		fmt.Printf("%d. Grup %s:\n", i+1, G[i].nama)
	}
	fmt.Scan(&x[nGrup].nama)
	idx = mencariGrupi(*G, nGrup, x[nGrup].nama)
	if mencarigrup(T, ng, x[nGrup].nama) {
		for {
			if idx == -1 {
				fmt.Println("Group tidak ditemukan.")
				ChatGrup()
			} else {
				nGrup = idx
				fmt.Println("------------------------")
				fmt.Println("1. Kirim Pesan")
				fmt.Println("2. Tampilkan Pesan")
				fmt.Println("3. Exit")
				fmt.Println("------------------------")
				fmt.Print("Pilih Menu: ")
				fmt.Scan(&y)
				switch y {
				case 1:
					fmt.Scan(&G[nGrup].arrChat[G[nGrup].nChat].isichat.IsiChat)
					G[nGrup].arrChat[G[nGrup].nChat].isichat.pengirim = akunLogin.username
					G[nGrup].nChat++
					for i := 0; i < G[nGrup].nChat; i++ {
						fmt.Printf("%s: %s\n", G[nGrup].arrChat[i].isichat.pengirim, G[nGrup].arrChat[i].isichat.IsiChat)

					}
				case 2:
					for i := 0; i < G[nGrup].nChat; i++ {
						fmt.Printf("%s: %s\n", G[nGrup].arrChat[i].isichat.pengirim, G[nGrup].arrChat[i].isichat.IsiChat)

					}
				case 3:
					ChatGrup()
				}
			}
		}
	} else {
		fmt.Println("Group tidak ditemukan.")
		ChatGrup()
	}
}

func mencarigrup(G akungrup, n int, x string) bool {
	var i int
	var idx bool
	idx = false
	for i = 0; i < n; i++ {
		if G[i].nama == x {
			idx = true
		}
	}
	return idx
}

func mencariGrupi(G akungrup, n int, x string) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if G[i].nama == x {
			idx = i
		}
	}
	return idx
}
