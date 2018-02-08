package main

import (
	"github.com/andlabs/ui"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Global variables
const appname = "数据隐身衣"
const button_text0 = "显形 (Uncloak)"
const button_text1 = "隐身 (Cloak)"

var dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
var cloak_script = filepath.Join(dir, "cloak.sh")
var uncloak_script = filepath.Join(dir, "uncloak.sh")

const c2s_msg = "Go Boilers"

var current_state = 0 // Hidden area closed by default

// Button text while in state
var state_text = map[int]string{
	0: button_text0,
	1: button_text1,
}

// sta: 0 - close hidden area
//      1 - open hidden area
func set_state(sta int) (string, bool) {
	status := false
	var script string
	var outstr string

	if sta == 1 {
		outstr = "Data uncloaking failed"
		script = uncloak_script
	} else if sta == 0 {
		outstr = "Data cloaking failed"
		script = cloak_script
	} else {
		panic("Impossible state encountered. Parallel universe?")
	}
	outbytes, err := exec.Command(script).Output()
	//log.Print("exec.Command")
	if err != nil {
		log.Print(err)
	} else {
		outstr = string(outbytes[:])
		status = true
	}
	return outstr, status
}

// For offline test only. Replace set_state with set_state1 in
// state_changer_server to test without a device
func set_state1(sta int) (string, bool) {
	status := false
	var outstr string
	if sta == 0 {
		outstr = "Data cloaked"
	} else if sta == 1 {
		outstr = "Data uncloaked"
	} else {
		panic("Impossible state encountered. Parallel universe?")
	}
	log.Print("server dummy operations: Sleep 10 secs")
	time.Sleep(time.Second * 10)
	status = true
	return outstr, status
}

func state_changer_server(c2s chan string, s2c chan result) {
	for {
		var r result
		msg := <-c2s
		log.Print("server received message from c2s channel:" + msg)
		if msg == c2s_msg {
			// Uncomment the following line for execution with a Rogue SSD
			// device
			out, status := set_state(1 - current_state)
			// Uncomment the following line for test without a Rogue SSD
			// device
			//            out, status := set_state1(1 - current_state)

			// Change state only upon success
			if status == true {
				current_state = 1 - current_state
			}
			r.out = out
			r.state = current_state
			s2c <- r

		} else {
			panic("Bad ping message. Traffic hijacked?")
		}

		time.Sleep(time.Second * 1)
	}

}

type result struct {
	out   string
	state int // 0: cloaked/hidden/closed; 1: uncloaked/visible/open
}

func main() {

	err := ui.Main(func() {
		// trigger from client to server
		var c2s chan string = make(chan string)
		// returned result from server to client
		var s2c chan result = make(chan result)
		button := ui.NewButton(state_text[current_state])
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow(appname, 180, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Processing...")
			c2s <- c2s_msg

			go func() {
				time.Sleep(time.Second * 2)
				msg := <-s2c
				button.SetText(state_text[msg.state])
				greeting.SetText(msg.out)
			}()
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
		go state_changer_server(c2s, s2c)
	})
	if err != nil {
		panic(err)
	}
}
