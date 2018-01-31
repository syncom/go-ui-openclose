package main

import (
	"github.com/andlabs/ui"
        "os"
        "os/exec"
        "log"
        "path/filepath"
)


// Global variables
var appname = "数据隐身衣"
var dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
var cloak_script = filepath.Join(dir, "cloak.sh")
var uncloak_script = filepath.Join(dir, "uncloak.sh")
var current_state = 0 // Hidden area closed by default

// sta: 0 - close hidden area
//      1 - open hidden area
func set_state(sta int) (string, bool) {
    status := false
    var script string
    var outstr string
    if current_state == 0 {
        outstr = "Data cloaking failed"
        script = uncloak_script
    } else if current_state == 1 {
        outstr = "Data uncloaking failed"
        script = cloak_script
    } else {
        panic("Impossible state encountered. Parallel universe?")
    }
    outbytes, err := exec.Command(script).Output()
    if err != nil {
        log.Print(err)
    } else {
        outstr = string(outbytes[:])
        status = true
    }
    return outstr, status
}

func main() {
        state_text := map[int] string{
            0: "Uncloak",
            1: "Cloak",
        }
	err := ui.Main(func() {
		button := ui.NewButton(state_text[current_state])
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow(appname, 162, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
                        out, status := set_state(1-current_state)
			greeting.SetText(out)
                        if status == true {
                            current_state = 1 - current_state
                            button.SetText(state_text[current_state])
                        }
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
