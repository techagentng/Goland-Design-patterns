package main

import "fmt"

func main()  {
	// TODO: Create a notification builder and use it to set property
	var bldr = newNotificationBuilder()
	bldr.SetMessage("This is a notification message")
	bldr.SetIcon("icon.png")
	bldr.SetPriority(10)
	bldr.SetType("alert")
	bldr.SetTitle("notification cart")
	bldr.SetSubTitle("cart")

	// TODO:Use the build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("error occurred while building notification")
	} else {
		fmt.Printf("succefully built %+v", notif)
	}
}


