package book

import "gorm.io/gorm"

type Repository interface { //interface adalah kontrak yang harus diimplementasikan oleh struct
	FindAll() ([]Book, error)       //method untuk mendapatkan semua buku
	FindById(id int) (Book, error)  //method untuk mendapatkan buku berdasarkan id
	Create(book Book) (Book, error) //method untuk membuat buku baru
	//Update(book Book) (Book, error) //method untuk memperbarui buku
}

type repository struct { //struct untuk implementasi repository
	db *gorm.DB //db adalah pointer ke struct DB
	//DB adalah struct yang berisi koneksi ke database
}

func NewRepository(db *gorm.DB) *repository { //fungsi untuk membuat repository baru
	return &repository{db} //mengembalikan pointer ke struct repository
}	

func (r *repository) FindAll() ([]Book, error) { //method untuk mendapatkan semua buku
	var books []Book //books adalah slice dari struct book
	//slice adalah array yang dapat berubah ukuran

	err := r.db.Find(&books).Error //mencari semua buku di database
	return books, err //mengembalikan slice buku dan error
}

func (r *repository) FindById(id int) (Book,error) { //method untuk mendapatkan buku berdasarkan id
	var book Book //book adalah struct book
	err := r.db.Find(&book,id).Error //mencari buku berdasarkan id di database
	return book, err //mengembalikan buku dan error

}

func (r *repository) Create(book Book) (Book, error) { //method untuk membuat buku baru
	err := r.db.Create(&book).Error //membuat buku baru di database	
	return book, err //mengembalikan buku dan error
}







		//CRUD
		//==========================
		// Simpan / Hapus Data
		//==========================
		// var book book.Book //untuk menampung data dari tabel book
		// err = db.Debug().Where("id = ?",4).First(&book).Error
		// if err != nil {
		// 	log.Println("=======================================")
		// 	log.Println("Failed to find book record")
		// 	log.Println("=======================================")
		// }

		// book.Title = "Belajar Golang Revisi"
		// err = db.Save(&book).Error//untuk mengupdate data book dengan id 4
		// //err = db.Delete(&book).Error//untuk menghapus data book dengan id 4
		// if err != nil {
		// 	fmt.Println("=======================================")
		// 	fmt.Println("Failed to update book record")
		// 	fmt.Println("=======================================")
		// }
		//==========================
		// Create Data
		//==========================
		// book := book.Book{}
		// book.Title=       "Atomic Habits"
		// book.Description= "Belajar habit atau kebiasaan menggunakan buku ini"
		// book.Price=       120000
		// book.Discount= 	25
		// book.Rating =   4

		// err = db.Create(&book).Error
		// if err != nil {
		// 	fmt.Println("=======================================")
		// 	fmt.Println("Failed to create book record")
		// 	fmt.Println("=======================================")
		// } 


		//==========================
		// Read Data
		//==========================
		// var books []book.Book //untuk menampung data dari tabel book

		// err = db.Debug().Where("title = ?","Belajar Golang").Find(&books).Error //untuk mengisi variabel books dengan data dari tabel book
		// if err != nil {
		// 	fmt.Println("=======================================")
		// 	fmt.Println("Failed to find book record")
		// 	fmt.Println("=======================================")	
		// }

		// for _, book := range books {//untuk menampilkan data book
		// 	fmt.Println("=======================================")
		// 	fmt.Println("Title: ", book.Title)
		// 	fmt.Println("Description: ", book.Description)
		// 	fmt.Printf("Book Object %v\n", book) //untuk menampilkan data book
		// }

		//var book book.Book //untuk menampung data dari tabel book

		// err= db.Debug().First(&book,4).Error //untuk mengisi variabel book dengan data pertama dari tabel book
		// if err != nil {
		// 	fmt.Println("=======================================")
		// 	fmt.Println("Failed to find book record")
		// 	fmt.Println("=======================================")
		// }

		// fmt.Println("Title: ", book.Title)
		// fmt.Println("Description: ", book.Description)
		// fmt.Printf("Book Object %v ", book.Price) //untuk menampilkan data book

		
		