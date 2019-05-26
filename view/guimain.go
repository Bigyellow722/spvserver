package view

import (
	"fmt"
	"github.com/andlabs/ui"
	"github.com/kongyixueyuan.com/kongyixueyuan/service"
	"github.com/kongyixueyuan.com/kongyixueyuan/view/spvwindow"
)

func Spvservice(serviceSetup service.ServiceSetup) {

	msg, err := serviceSetup.SetInfo("wqy", "12345")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	msg, err = serviceSetup.GetInfo("wqy")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	err = ui.Main(spvwindow.App)
	if err != nil {
		panic(err)
	}
}
