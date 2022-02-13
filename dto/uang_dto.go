package dto

type Uang struct {
	ID             int    `json:"id"`
	TGLUANG        string `json:"tgl_uang"`
	ID_PENGELUARAN int    `json:"id_pengeluaran"`
	ID_PENDAPATAN  int    `json:"id_pendapatan"`
	JUMLAH         int    `json:"jumlah"`
}
