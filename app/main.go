package main

import (
	"fmt"
	you "github.com/q2rd/yougile_api_wrapp"
)

func main() {
	uClient := you.NewYouGileClient(
		"252650c8-b680-436d-93df-f9650fcce7f4",
		"GiYZoRkKR8DKEfZpnilK-8JJTDpJSYNpcOYzkNBAvHyowUCXHalhF+GSYNL+rWDL",
	)
	//task1 := you.Task{
	//	Id: "c5d76103-8f8e-4652-b124-9106448f8846",
	//}
	//task2 := you.Task{Id: "fd10acd2-d2fa-48cf-9280-58b255501e62"}
	//task3 := you.Task{Id: "e5407f92-e8f0-431a-901e-9722e57e3a3d"}
	c := you.Column{
		YouGileClient: uClient,
		Id:            "79df02e0-902a-4e30-a172-ee56348b033e",
	}

	list, err := c.GetTaskList()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range list.Content {
		fmt.Println(v.Title)
	}
	err = uClient.DeleteMultiTask(list.Content)
	if err != nil {
		return
	}

	//ts := you.Task{
	//	ColumnID: "79df02e0-902a-4e30-a172-ee56348b033e",
	//	Title:    "test",
	//}
	//err := uClient.CreateTask(&ts, you.Defaults())
	//if err != nil {
	//	fmt.Println(err)
	//}
}
