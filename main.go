package main

import (
	"log"
	"fmt"
	"bytes"
	"time"
	"os"
	"github.com/nanu-c/qml-go"
)

const path="/home/phablet/Documents/be-remote.txt"

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
	if err := writeEntry("note", testvar.Message); err==nil {
		testvar.Message = ""
		testvar.Output = "Note Enqueued!"
		qml.Changed(testvar, &testvar.Message)
		qml.Changed(testvar, &testvar.Output)
	} else {
		log.Fatal(err)
		testvar.Output = err.Error();
		qml.Changed(testvar, &testvar.Output)
	}
}

func (testvar *TestStruct) DoTodo() {
	if err := writeEntry("todo", testvar.Message); err==nil {
		testvar.Message = ""
		testvar.Output = "Todo Enqueued!"
		qml.Changed(testvar, &testvar.Message)
		qml.Changed(testvar, &testvar.Output)
	} else {
		log.Fatal(err)
		testvar.Output = err.Error();
		qml.Changed(testvar, &testvar.Output)
	}
}

func writeEntry(key string, message string) error {
	var buf bytes.Buffer;
	fmt.Fprintf(&buf, "%x\t%s\t%s\n", time.Now().Unix(), "alpha", "beta");
	//ioutil.WriteFile(path, buf.Bytes(), os.ModeAppend|0777);

	//https://godoc.org/os#example-OpenFile--Append
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err;
	}
	if _, err := f.Write(buf.Bytes()); err != nil {
		return err;
	}
	if err := f.Close(); err != nil {
		return err;
	}
	return nil;
}

