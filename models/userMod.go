package models

import "time"

type User struct {
  ID uint `json:"id"` 
  Name string `json:"name"`
  Address string `json:"address"`
  BornDate time.Time `json:"born_date" gorm:"default:current_timestamp"` 
}

type UserReq struct {
  Name string `json:"name"`
  Address string `json:"address"`
}
