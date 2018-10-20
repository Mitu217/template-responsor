package model

import "github.com/Mitu217/template-responsor/repository/model/enum"

type Response struct {
	Uuid          string
	Name          string
	RequestMethod enum.RequestMethod
	ContentType   enum.ContentType
	Route         string
	Data          string
}
