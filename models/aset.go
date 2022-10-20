package models

type Aset struct {
	Id              int64   `gorm:"primaryKey" json:"id"`
	Nama_aset       string  `gorm:"varchar(300)" json:"nama_aset"`
	Jumlah          int32   `json:"jumlah"`
	Tahun_pembelian int32   `json:"tahun_pembelian"`
	Harga_satuan    float64 `json:"harga"`
	Lokasi_aset     string  `gorm:"varchar(300)" json:"lokasi_aset"`
}
