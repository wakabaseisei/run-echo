package controllers

// POSTリクエストで受け取ったJSONをパースする際の構造体
type UserJsonRequest struct {
	Sex          int    `json:"sex" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
}
