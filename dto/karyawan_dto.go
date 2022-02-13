package dto

type Kariawan struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Posisi string `json:"posisi"`
	Alamat string `json:"alamat"`
	Umur   int    `json:"umur"`
	Kontak string `json:"kontak"`
}
