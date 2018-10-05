package main

import (
	"log"
	"fmt"
	"bytes"
	"time"
	"os"
	"strings"
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

func (testvar *TestStruct) DoNote(raw string) {
	testvar.ProcessRawEntry(raw, "note", "Note Enqueued!")
}

func (testvar *TestStruct) DoTodo(raw string) {
	testvar.ProcessRawEntry(raw, "todo", "Todo Enqueued!")
}

func (testvar *TestStruct) ProcessRawEntry(raw string, key string, happyMessage string) {
	//DNW: message := testvar.Message
	if err := writeEntry(key, raw); err==nil {
		//BUG: this only seems to clear the message the first time!?!
		testvar.Message = ""
		testvar.Output = happyMessage
		qml.Changed(testvar, &testvar.Message)
		qml.Changed(testvar, &testvar.Output)
	} else {
		//dies: log.Fatal(err)
		testvar.Output = err.Error();
		qml.Changed(testvar, &testvar.Output)
	}
}

func writeEntry(key string, raw string) error {
	var buf bytes.Buffer;
	//ioutil.WriteFile(path, buf.Bytes(), os.ModeAppend|0777);

	//https://godoc.org/os#example-OpenFile--Append
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err;
	}

	bits:=strings.Split(raw, "\n")
	
	for _, message := range bits {
		log.Printf("%s: %s\n", key, message);
		if len(message) > 0 {
			fmt.Fprintf(&buf, "%x\t%s\t%s\n", time.Now().Unix(), key, message);
		}
	}

	if _, err := f.Write(buf.Bytes()); err != nil {
		return err;
	}
	if err := f.Close(); err != nil {
		return err;
	}
	return nil;
}

