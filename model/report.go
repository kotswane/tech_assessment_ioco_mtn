package model

type Reports struct{
	ID 				 int64 `json:"id"`
	SurID 			 string `json:"sur_id" bson:"sur_id" binding:"required"`
}