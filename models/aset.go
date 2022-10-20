package models

type Aset struct {
	Id     int64   `gorm:"primaryKey" json:"id"`
	Nama   string  `gorm:"type:varchar(300)" json:"nama_aset"`
	Jumlah int32   `json:"jumlah"`
	Tahun  int32   `json:"tahun_pembelian"`
	Harga  float64 `json:"harga"`
	Lokasi string  `gorm:"type:varchar(300)" json:"lokasi_aset"`
}
