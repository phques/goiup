package main

import (
	"fmt"
	"github.com/gonutz/goiup/iup"
	"github.com/gonutz/goiup/iuputil"
	"runtime"
	"time"
)

// Will hold Handles to widgets,
// filled by iuputil.FetchWidgets
type MyWidgets struct {
	MainDialog *iup.Handle `IUP:"mainDialog"`
	LocalRoot  *iup.Handle `IUP:"localRoot"`
	Files      *iup.Handle
	Push       *iup.Handle `IUP:"pushButton"`
}

var cmdChan chan string
var myWidgets MyWidgets

//---------

// Idle callback
// Called from goroutine to execute commands that change the GUI,
// since the GUI stuff must run in the original thread
func idleFunc1() int {

	select {
	case cmd := <-cmdChan:
		fmt.Println("got something to do in idle: ", cmd)

		if cmd == "addtofiles" {
			myWidgets.Files.SetAttribute("APPENDITEM", "some filename, push pressed")
		} else {
			myWidgets.LocalRoot.SetAttribute("VALUE", cmd)
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
		myWidgets.Files.SetAttribute("APPENDITEM", "push pressed live")

		// but this should work
		cmdChan <- "addtofiles"
	}()

	return iup.DEFAULT
}

//----------

func createDialog() {
	// load GUI definitions from file
	if errStr := iup.Load("androidGUI.led"); errStr != "" {
		panic(errStr)
	}

	// get widgets handles into myWidgets
	if err := iuputil.FetchWidgets(&myWidgets); err != nil {
		panic("FetchWidgets failed : ", err)
		return
	}

}

//----------

func main() {
	runtime.LockOSThread()

	iup.Open()
	defer iup.Close()

	createDialog()
	myWidgets.Push.SetCallback("ACTION", pushBtnCB)

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
	myWidgets.MainDialog.Show()
	iup.MainLoop()

}
