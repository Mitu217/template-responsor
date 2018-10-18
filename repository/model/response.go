package model

import "github.com/Mitu217/template-responsor/repository/model/enum"

type Response struct {
	Uuid         string
	Name         string
	AccessType   enum.AccessType
	EncodingType enum.EncodingType
	Route        string
	Data         string
}
