package dto

type Pengeluaran struct {
	ID              int    `json:"id"`
	TGL_PENGELUARAN string `json:"tgl_pengeluaran"`
	JUMLAH          int    `json:"jumlah"`
	ID_SUMBER       int    `json:"id_sumber"`
}
