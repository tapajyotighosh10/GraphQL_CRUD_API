package main

// import (
// 	"fmt"
// 	"strconv"
// 	"github.com/360EntSecGroup-Skylar/excelize"
// )

// func Datastore(input []*model.User) {
// 	f := excelize.NewFile()
// 	f.SetCellValue("Sheet1","A1","userId")
// 	f.SetCellValue("Sheet1","B1","FirstName")
// 	f.SetCellValue("Sheet1","C1","LastName")
// 	f.SetCellValue("Sheet1","D1","Dob")
	
// 	for i, value := range input {
	
// 		a := strconv.Itoa(i + 2)
// 		f.SetCellValue("Sheet1", ("A" + a), value.ID)
// 		f.SetCellValue("Sheet1", ("B" + a), value.FirstName)
// 		f.SetCellValue("Sheet1", ("C" + a), value.LastName)
// 		f.SetCellValue("Sheet1", ("D" + a), value.Dob)
		
// 		i++

// 	}

// err:=f.SaveAs("MyData.xlsx")
// if err!=nil {
// 	fmt.Println(err)
// } else {
// 	fmt.Println("File created and data inserted")
// }

// // 	i:=0
// // for {

// // 	a := strconv.Itoa(i + 2)
// // 	c1 := f.GetCellValue("Sheet1", ("A" + a))
// // 	c2 := f.GetCellValue("Sheet1", ("B" + a))
// // 	c3 := f.GetCellValue("Sheet1", ("C" + a))
// // 	c4 := f.GetCellValue("Sheet1", ("D" + a))
// // 	i++
// // }


// }