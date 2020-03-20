package main

import (
	"encoding/gob"
	"os"
	"log"
)

type Address struct {
	Type             string
	City             string
	Country          string
}
type VCard struct {
	FirstName	string
	LastName	string
	Addresses	[]Address
	Remark		string
}

func main()  {
	path :="./vcard.gob"

	file,err:=os.OpenFile(path,os.O_RDONLY,0777)
	defer file.Close()
	if err!=nil {
		log.Fatal(err)
	}
	vcar:=new(VCard)
	dec:=gob.NewDecoder(file)
	err=dec.Decode(vcar)
	if err!=nil {
		log.Fatal(err)
	}

	log.Println(vcar.Addresses,vcar.FirstName,vcar.LastName)

}