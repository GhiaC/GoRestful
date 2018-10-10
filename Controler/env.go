package Controler

import "../Models"

var Configuration = Models.Configuration{}

func GetAdminPhonenumber()  string {
	return Configuration.AdminPhonenumber
}

func GetLineNumber() string{
	return Configuration.LineNumber
}