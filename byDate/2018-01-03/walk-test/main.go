// To build the stopwatch app:
//   go get github.com/lxn/walk
//   go get github.com/akavel/rsrc
//   rsrc -manifest test.manifest -o rsrc.syso
//   go build -ldflags="-H windowsgui"
//   walk-test.exe
package main

import (
	"context"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var stopwatch *walk.TextEdit
	var start *walk.PushButton
	started := false
	var cancel context.CancelFunc

	MainWindow{
		Title:   "GoStopwatch",
		MinSize: Size{400, 140},
		Layout:  VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo:  &stopwatch,
				ReadOnly:  true,
				Alignment: AlignCenter,
				Font:      Font{PointSize: 34, Bold: true},
				Text:      "Press Start"},
			PushButton{
				AssignTo: &start,
				Text:     "Start",
				OnClicked: func() {
					if started {
						cancel()
						started = false
						start.SetText("Start")
					} else {
						started = true
						stopwatch.SetText("0s")
						start.SetText("Stop")

						var ctx context.Context
						ctx, cancel = context.WithCancel(context.Background())
						go timer(ctx, stopwatch)
					}
				},
			},
		},
	}.Run()
}

func timer(ctx context.Context, output *walk.TextEdit) {
	start := time.Now()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			output.SetText(time.Since(start).Round(time.Second).String())
		}
	}
}
