package main

import (
	"fmt"
	"github.com/gonutz/goiup/iup"
	//	"github.com/gonutz/goiup/iuputil"
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

	// Local Root frame + text
	localRoot := iup.Text("").SetAttributes(`EXPAND=HORIZONTAL, MARGIN=5`)
	localRootVbox := iup.Vbox(localRoot).SetAttributes(`NMARGIN=5x5, NGAP=5x5, EXPAND=YES`)
	localRootFrame := iup.Frame(localRootVbox).SetAttributes(`TITLE="Local Root"`)

	// Destination frame + vbox
	// Destination - Root list
	destRoot := iup.List("").
		SetAttributes(`1="aaa",2="bbb",3="ccc", DROPDOWN="YES"`).
		SetAttributes(`EXPAND=HORIZONTAL`)
	destVbox1_1 := iup.Vbox(iup.Label("Root"), destRoot)

	// Destination - Files list
	destFiles := iup.List("").SetAttributes(`EXPAND=YES`)
	destVbox1_2 := iup.Vbox(iup.Label("Files"), destFiles)

	// Destination - frame & vbox
	destVbox1 := iup.Vbox(destVbox1_1, destVbox1_2).
		SetAttributes(`NMARGIN=5x5, NGAP=5x5, EXPAND=YES`)
	destFrame := iup.Frame(destVbox1).SetAttributes(`TITLE="Destination", SIZE=x150`)

	// buttons - Push
	pushButt := iup.Button("Push", "")
	buttsHbox := iup.Hbox(iup.Fill(), pushButt).
		// forces dialog to be wider
		SetAttributes(`EXPAND=HORIZONTAL, SIZE=200`)

	// main dialog + vbox
	mainvbox := iup.Vbox(localRootFrame, destFrame, buttsHbox).
		SetAttributes("NMARGIN=5x5, NGAP=5x5, EXPAND=YES")

	mainDialog := iup.Dialog(mainvbox).
		SetAttributes(`TITLE="Android Push", MARGINS=5x5`)

	// save handles
	myControls.LocalRoot = localRoot
	myControls.MainDialog = mainDialog
	myControls.Push = pushButt
	myControls.Files = destFiles
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
