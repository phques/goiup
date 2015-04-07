package main

//extern void led_load(void);
import "C"

import (
	"fmt"
	"github.com/gonutz/goiup/iup"
	"github.com/gonutz/goiup/iuputil"
	"runtime"
	"time"
)

// Will hold Handles to controls,
type MyControls struct {
	MainDialog *iup.Handle `IUP:"mainDialog"`
	LocalRoot  *iup.Handle `IUP:"localRoot"`
	Files      *iup.Handle
	Push       *iup.Handle `IUP:"pushButton"`
}

var cmdChan chan string
var myControls MyControls

//---------

// Idle callback
// Called from goroutine to execute commands that change the GUI,
// since the GUI stuff must run in the original thread
func idleFunc1() int {

	select {
	case cmd := <-cmdChan:
		fmt.Println("got something to do in idle: ", cmd)

		if cmd == "addtofiles" {
			myControls.Files.SetAttribute("APPENDITEM", "some filename, push pressed")
		} else {
			myControls.LocalRoot.SetAttribute("VALUE", cmd)
		}

	case <-time.After(time.Duration(100 * time.Millisecond)):
	}

	return iup.DEFAULT
}

//----------

// 'Push' button callback
func pushBtnCB() int {
	// do some work (in diff goroutine)
	go func() {
		fmt.Println("pushBtnAction")

		// fake work
		time.Sleep(time.Duration(1) * time.Second)

		// try to update GUI .. might not work if different thread !
		myControls.Files.SetAttribute("APPENDITEM", "push pressed live")

		// but this should work
		cmdChan <- "addtofiles"
	}()

	return iup.DEFAULT
}

func createDialog() {
	// call generated C function (gen from LED file with ledc)
	C.led_load()

	// get controls handles into myControls
	if err := iuputil.FetchControls(&myControls); err != nil {
		fmt.Println("FetchControls failed : ", err)
		return
	}

}

//----------

func main() {

	runtime.LockOSThread()

	iup.Open()
	defer iup.Close()

	createDialog()

	myControls.Push.SetCallback("ACTION", pushBtnCB)

	// prepare a channel for the idle callback msgs,
	// start a goroutine to send a msg on the channel after some time
	cmdChan = make(chan string)

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		cmdChan <- "command to process by UI thread"
	}()

	// hook our idle func
	iup.SetIdleFunc(idleFunc1)

	// show dialog and loop until last window closed
	myControls.MainDialog.Show()
	iup.MainLoop()

}
