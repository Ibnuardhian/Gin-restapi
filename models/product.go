package models

type Product struct {
	Id              int64  `gorm:"primaryKey" json:"id"`
	NamaProduk      string `gorm:"type:varchar(300)" json:"nama_produk"`
	DeskripsiProduk string `gorm:"type:text" json:"deskripsi"`
}
