package main

import (
	"log"

	"github.com/nanu-c/qml-go"
)

func main() {
	err := qml.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	engine := qml.NewEngine()
	component, err := engine.LoadFile("qml/Main.qml")
	if err != nil {
		return err
	}

	testvar := TestStruct{Message: ""}
	context := engine.Context()
	context.SetVar("testvar", &testvar)
	//testvar.GetMessage()

	win := component.CreateWindow(nil)
	testvar.Root = win.Root()
	win.Show()
	win.Wait()

	return nil
}

type TestStruct struct {
	Root    qml.Object
	Message string
	Output  string
	Number  int
}

func (testvar *TestStruct) DoCancel() {
	testvar.Message = ""
	testvar.Output = "Reset!"
	qml.Changed(testvar, &testvar.Message)
	qml.Changed(testvar, &testvar.Output)
}

func (testvar *TestStruct) DoNote() {
	testvar.Message = ""
	testvar.Output = "Note Sent!"
	qml.Changed(testvar, &testvar.Message)
	qml.Changed(testvar, &testvar.Output)
}

func (testvar *TestStruct) DoTodo() {
	testvar.Message = ""
	testvar.Output = "Todo Sent!"
	qml.Changed(testvar, &testvar.Message)
	qml.Changed(testvar, &testvar.Output)
}


