package model

type Matomo_log_action struct {
	Idaction  int    `gorm:"idaction" json:"idaction"`
	Name      string `gorm:"name" json:"name"`
	Hash      int    `gorm:"hash" json:"hash"`
	Type      int    `gorm:"type" json:"type"`
	UrlPrefix int    `gorm:"url_prefix" json:"url_prefix"`
}
