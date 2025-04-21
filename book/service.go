package book

type Service interface {
	FindAll() ([]Book, error)              //method untuk mendapatkan semua buku
	FindById(id int) (Book, error)         //method untuk mendapatkan buku berdasarkan id
	Create(book BookRequest) (Book, error) //method untuk membuat buku baru
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll() //mencari semua buku di repository
	return books, err                    //mengembalikan semua buku dari repository
}

func (s *service) FindById(id int) (Book, error) {
	book, err := s.repository.FindById(id) //mencari buku berdasarkan id di repository
	return book, err                       //mengembalikan buku dari repository
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64() //mengubah harga buku dari json.Number ke int64
	book := Book{                         //membuat buku baru
		Title:       bookRequest.Title,       //mengisi judul buku
		Description: bookRequest.Description, //mengisi deskripsi buku
		Price:       int(price),              //mengisi harga buku
	}
	book, err := s.repository.Create(book) //membuat buku baru di repository
	return book, err                       //mengembalikan buku dari repository
}