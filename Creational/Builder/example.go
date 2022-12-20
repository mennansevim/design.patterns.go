package main

import "fmt"

func main() {

	var builder = newNotificationBuilder()

	builder.SetTitle("Uyarı")
	builder.SetSubTitle("SubTitle")
	builder.SetIcon("icon.png")
	builder.SetImage("image.jpg")
	builder.SetPriority(1)
	builder.SetMessage("Bu bir test uyarı mesajıdır")
	builder.SetTitle("alert")

	notif, err := builder.Build()

	if err != nil {
		fmt.Println("Error creating the noty:", err)
	} else {
		fmt.Printf("Notification: %+v", notif)
	}
}
