package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	var bldr = newNotificationBuilder()
	// TODO: Use the builder to set some properties
	bldr.SetTitle("New notification")
	bldr.SetSubTitle("some sub title")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(1)
	bldr.SetMessage("this is basic notification")
	bldr.SetType("alert")
	// TODO: Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating notification", err)
	} else {
		fmt.Printf("Created notification: %+v", notif)
	}
}
