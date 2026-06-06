package main

import "fmt"

const maxPatients = 100
const maxServices = 100 
const maxVisits = 500

type Patient struct {
	ID    int
	Name  string
	Age   int
	Phone string
}

type Service struct {
	ID    int
	Name  string
	Price int
}

type Visit struct {
	PatientID int
	ServiceID int
	Date      int 
	Total     int
	Note      string
}

var patients [maxPatients]Patient
var patientCount int

var services [maxServices]Service
var serviceCount int

var visits [maxVisits]Visit
var visitCount int

func addPatient() { // prosedur untuk menanmbah atau mendaftarkan pasien baru
	if patientCount >= maxPatients {
		fmt.Println("Kapasitas pasien penuh")
		return
	}
	var name string
	var age int
	var phone string
	var id int
	fmt.Println("Masukkan nama (tanpa spasi):")
	fmt.Scan(&name)
	fmt.Println("Masukkan umur:")
	fmt.Scan(&age)
	fmt.Println("Masukkan nomor handPhone:")
	fmt.Scan(&phone)
	id = patientCount + 1
	patients[patientCount] = Patient{ID: id, Name: name, Age: age, Phone: phone}
	patientCount = patientCount + 1
	fmt.Println("Pasien ditambahkan dengan ID:", id)
}

func findPatientIndexByID(id int) int { // ini adalah fungsi untuk melakukan searching melalui metode binary search ketika data nya ascending 
	left := 0
	right := patientCount - 1
	for left <= right {
		mid := (left + right) / 2
		if patients[mid].ID == id {
			return mid
		}
		if patients[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func updatePatient() { // ini adalah prosedur untuk memperbarui data pasiean berdasarkan id yang di cari melalui fungsi binary search sebelum nya 
	var id int
	fmt.Println("Masukkan ID pasien yang akan diubah:")
	fmt.Scan(&id)
	idx := findPatientIndexByID(id)
	if idx == -1 {
		fmt.Println("Pasien tidak ditemukan")
		return
	}
	var name string
	var age int
	var phone string
	fmt.Println("Masukkan nama baru (tanpa spasi):")
	fmt.Scan(&name)
	fmt.Println("Masukkan umur baru:")
	fmt.Scan(&age)
	fmt.Println("Masukkan nomor telepon baru:")
	fmt.Scan(&phone)
	patients[idx].Name = name
	patients[idx].Age = age
	patients[idx].Phone = phone
	fmt.Println("Data pasien diperbarui")
}

func deletePatient() { // lalu diprosedur ini di gunakan untuk menghapus data pasien melalui id nya dengan cara menimpa data yang ingin di hapus hingga akhir data dan data di kurangi satu
	var id int
	fmt.Println("Masukkan ID pasien yang akan dihapus:")
	fmt.Scan(&id)
	idx := findPatientIndexByID(id)
	if idx == -1 {
		fmt.Println("Pasien tidak ditemukan")
		return
	}
	i := idx
	for i < patientCount-1 {
		patients[i] = patients[i+1]
		i = i + 1
	}
	patientCount = patientCount - 1
	fmt.Println("Pasien dihapus")
}

func listPatients() { // prosedur ini di gunakan untuk menampilkan pasien yang sudah terdaftar 
	if patientCount == 0 {
		fmt.Println("Belum ada pasien")
		return
	}
	fmt.Println("Daftar pasien:")
	i := 0
	for i < patientCount {
		p := patients[i]
		fmt.Println(p.ID, p.Name, p.Age, p.Phone)
		i = i + 1
	}
}

func addService() { // untuk prosedur ini hampir sama dengan addppasien hanya saja ini menambahkan service nya sesuai dengan urutan id nya jadi mulai dari id 1 hingga seterusnya 
	if serviceCount >= maxServices {
		fmt.Println("Kapasitas layanan penuh")
		return
	}
	var name string
	var price int
	fmt.Println("Masukkan nama layanan (Peeling/Facial/Hair_Treatment/Massage/Botox):")
	fmt.Scan(&name)
	fmt.Println("Masukkan harga layanan:")
	fmt.Scan(&price)
	id := serviceCount + 1
	services[serviceCount] = Service{ID: id, Name: name, Price: price}
	serviceCount = serviceCount + 1
	fmt.Println("Layanan ditambahkan dengan ID:", id)
}

func findServiceIndexByID(id int) int { // fungsi ini di gunakan untuk searching service dari id tetapi dengan metode sequential search dan fungsi mengembalikan berupa id int 
	i := 0
	for i < serviceCount {
		if services[i].ID == id {
			return i
		}
		i = i + 1
	}
	return -1
}

func updateService() { //prosedur ini pada dasar nya sama saja fungsi nya dengan update pasien tapi yang di update service atau layanannya dan berdasarkan id juga 
	var id int
	fmt.Println("Masukkan ID layanan yang akan diubah:")
	fmt.Scan(&id)
	idx := findServiceIndexByID(id)
	if idx == -1 {
		fmt.Println("Layanan tidak ditemukan")
		return
	}
	var name string
	var price int
	fmt.Println("Masukkan nama layanan baru (tanpa spasi):")
	fmt.Scan(&name)
	fmt.Println("Masukkan harga baru:")
	fmt.Scan(&price)
	services[idx].Name = name
	services[idx].Price = price
	fmt.Println("Data layanan diperbarui")
}

func deleteService() { //pada dasar nya sama saja dengan prosedur delete pasien 
	var id int
	fmt.Println("Masukkan ID layanan yang akan dihapus:")
	fmt.Scan(&id)
	idx := findServiceIndexByID(id)
	if idx == -1 {
		fmt.Println("Layanan tidak ditemukan")
		return
	}
	i := idx
	for i < serviceCount-1 {
		services[i] = services[i+1]
		i = i + 1
	}
	serviceCount = serviceCount - 1
	fmt.Println("Layanan dihapus")
}

func listServices() { //pada dasar nya sama dengan list pasien 
	if serviceCount == 0 {
		fmt.Println("Belum ada layanan")
		return
	}
	fmt.Println("Daftar layanan:")
	i := 0
	for i < serviceCount {
		s := services[i]
		fmt.Println(s.ID, s.Name, s.Price)
		i = i + 1
	}
}

func recordVisit() { //prosedur ini di gunakan untuk mendata kapan tanggal kunjungan dan juga catatan singkat mengenai kunjungan 
	if visitCount >= maxVisits {
		fmt.Println("Kapasitas riwayat kunjungan penuh")
		return
	}
	var pid int
	var sid int
	var date int
	var note string
	fmt.Println("Masukkan ID pasien:")
	fmt.Scan(&pid)
	if findPatientIndexByID(pid) == -1 {
		fmt.Println("Pasien tidak ditemukan")
		return
	}
	fmt.Println("Masukkan ID layanan:")
	fmt.Scan(&sid)
	sidx := findServiceIndexByID(sid)
	if sidx == -1 {
		fmt.Println("Layanan tidak ditemukan")
		return
	}
	fmt.Println("Masukkan tanggal kunjungan (YYYYMMDD):")
	fmt.Scan(&date)
	fmt.Println("Masukkan catatan singkat (tanpa spasi):")
	fmt.Scan(&note)
	total := services[sidx].Price
	visits[visitCount] = Visit{PatientID: pid, ServiceID: sid, Date: date, Total: total, Note: note}
	visitCount = visitCount + 1
	fmt.Println("Kunjungan tercatat")
}

func listVisits() { //untuk mencatat daftar kunjungan dari pasien mirip dengan list sebelum nya 
	if visitCount == 0 {
		fmt.Println("Belum ada riwayat kunjungan")
		return
	}
	fmt.Println("Daftar kunjungan:")
	i := 0
	for i < visitCount {
		v := visits[i]
		fmt.Println("PasienID:", v.PatientID, "LayananID:", v.ServiceID, "Tanggal:", v.Date, "Total:", v.Total, "Catatan:", v.Note)
		i = i + 1
	}
}

func searchPatientByNameSequential() { //prosedur ini mencari searching pasien dari nama dan menggunakan sequential search 
	var name string
	fmt.Println("Masukkan nama pasien untuk pencarian (tanpa spasi):")
	fmt.Scan(&name)
	found := false
	i := 0
	for i < patientCount {
		if patients[i].Name == name {
			p := patients[i]
			fmt.Println("Ditemukan:", p.ID, p.Name, p.Age, p.Phone)
			found = true
		}
		i = i + 1
	}
	if !found {
		fmt.Println("Tidak ditemukan")
	}
}

func binarySearchPatientByID() { //prosedur ini untuk meng output hasil dari fungsi binarysearch yang pertama
	var id int
	fmt.Println("Masukkan ID pasien untuk binary search:")
	fmt.Scan(&id)
	idx := findPatientIndexByID(id)
	if idx == -1 {
		fmt.Println("Tidak ditemukan")
	} else {
		p := patients[idx]
		fmt.Println("Ditemukan:", p.ID, p.Name, p.Age, p.Phone)
	}
}

func selectionSortVisitsByDate() { //sorting pasien berdasarkan tanggal kunjungan dengan selection sort secara ascending 
	i := 0
	for i < visitCount-1 {
		minIdx := i
		j := i + 1
		for j < visitCount {
			if visits[j].Date < visits[minIdx].Date {
				minIdx = j
			}
			j = j + 1
		}
		if minIdx != i {
			tmp := visits[i]
			visits[i] = visits[minIdx]
			visits[minIdx] = tmp
		}
		i = i + 1
	}
	fmt.Println("Kunjungan diurutkan berdasarkan tanggal (Selection Sort)")
	for k := 0; k < visitCount; k++ {
		fmt.Printf("Tanggal: %d, Total: %d\n", visits[k].Date, visits[k].Total)
	}
}

func insertionSortVisitsByTotal() { //sorting pasien berdasarkan tanggal kunjungan dengan insertion sort secara ascending
	i := 1
	for i < visitCount {
		key := visits[i]
		j := i - 1
		for j >= 0 && visits[j].Total > key.Total {
			visits[j+1] = visits[j]
			j = j - 1
		}
		visits[j+1] = key
		i = i + 1
	}
	fmt.Println("Kunjungan diurutkan berdasarkan total (Insertion Sort)")
	for k := 0; k < visitCount; k++ {
		fmt.Printf("Tanggal: %d, Total: %d\n", visits[k].Date, visits[k].Total)
	}
}

func statsDailyVisits() { //prosedur ini di gunakan untuk melihat jumlah kunjungan dari hari nya dengan mengcount jumlah date yang sama dan memberikan jumlah
	var date int
	fmt.Println("Masukkan tanggal untuk statistik (YYYYMMDD):")
	fmt.Scan(&date)
	cnt := 0
	i := 0
	for i < visitCount {
		if visits[i].Date == date {
			cnt = cnt + 1
		}
		i = i + 1
	}
	fmt.Println("Jumlah kunjungan pada", date, ":", cnt)
}

func statsMostPopularService() { //prosedur ini di gunakan untuk melihat layanan yang paling banyak di gunakan oleh pasien dan berapa jumlah nya
	if serviceCount == 0 {
		fmt.Println("Belum ada layanan")
		return
	}
	var counts [maxServices]int
	i := 0
	for i < visitCount {
		sid := visits[i].ServiceID
		if sid >= 1 && sid <= serviceCount {
			counts[sid-1] = counts[sid-1] + 1
		}
		i = i + 1
	}
	maxCount := 0
	maxIdx := 0
	k := 0
	for k < serviceCount {
		if counts[k] > maxCount {
			maxCount = counts[k]
			maxIdx = k
		}
		k = k + 1
	}
	if maxCount == 0 {
		fmt.Println("Belum ada kunjungan untuk menghitung layanan populer")
		return
	}
	s := services[maxIdx]
	fmt.Println("Layanan paling populer:", s.ID, s.Name, "dengan jumlah:", maxCount)
}

func main() {
	var choice int
    choice = -1
    for choice != 0 {
        fmt.Println("SIM-KLIK: Laman Menu")
        fmt.Println("1 Tambah pasien")
        fmt.Println("2 Ubah pasien")
        fmt.Println("3 Hapus pasien")
		fmt.Println("4 Daftar list pasien")
		fmt.Println("5 Tambah layanan")
		fmt.Println("6 Ubah layanan")
		fmt.Println("7 Hapus layanan")
		fmt.Println("8 Daftar list layanan")
		fmt.Println("9 Catat kunjungan")
		fmt.Println("10 Daftar kunjungan")
		fmt.Println("11 Pencarian pasien (Sequential by name)")
		fmt.Println("12 Pencarian pasien (Binary by ID)")
		fmt.Println("13 Urutkan kunjungan (Tanggal - Selection)")
		fmt.Println("14 Urutkan kunjungan (Total - Insertion)")
		fmt.Println("15 Statistik: jumlah kunjungan harian")
		fmt.Println("16 Statistik: layanan paling populer")
        fmt.Println("0 Keluar")
        fmt.Scan(&choice)
        
        switch choice {
        case 1:
            addPatient()
        case 2:
            updatePatient()
        case 3:
			deletePatient()
		case 4:
			listPatients()
		case 5:
			addService()
		case 6:
			updateService()
		case 7:
			deleteService()
		case 8:
			listServices()
		case 9:
			recordVisit()
		case 10:
			listVisits()
		case 11:
			searchPatientByNameSequential()
		case 12:
			binarySearchPatientByID()
		case 13:
			selectionSortVisitsByDate()
		case 14:
			insertionSortVisitsByTotal()
		case 15:
			statsDailyVisits()
		case 16:
			statsMostPopularService()
        case 0:
            fmt.Println("Keluar")
        default:
            fmt.Println("Pilihan tidak sesuai menu, mohon pilih nomor yang sesuai")
        }
    }
}
