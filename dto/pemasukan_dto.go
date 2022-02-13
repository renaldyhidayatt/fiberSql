package dto

type Pemasukan struct {
	ID            int    `json:"id"`
	TGL_PEMASUKAN string `json:"tgl_pemasukan"`
	JUMLAH        int    `json:"jumlah"`
	ID_SUMBER     int    `json:"id_sumber"`
}
