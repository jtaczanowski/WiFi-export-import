//go:generate goversioninfo -icon=assets/WiFi-export-import.ico -manifest=assets/goversioninfo.exe.manifest
package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comco"
	"github.com/rodrigocfd/windigo/win/com/shell"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
)

var Version = "dev"

func main() {
	runtime.LockOSThread()

	myWindow := NewMyWindow() // instantiate
	myWindow.wnd.RunAsMain()  // ...and run
}

// This struct represents our main window.
type MyWindow struct {
	wnd               ui.WindowMain
	btnExport         ui.Button
	btnImportAllUsers ui.Button
	lblInfo           ui.Static
}

func (me *MyWindow) exportProfilesFunction() {
	fod := shell.NewIFileOpenDialog(
		com.CoCreateInstance(
			shellco.CLSID_FileOpenDialog, nil,
			comco.CLSCTX_INPROC_SERVER,
			shellco.IID_IFileOpenDialog),
	)

	fod.SetOptions(fod.GetOptions() | shellco.FOS_PICKFOLDERS)
	isDirectorySelected := fod.Show(me.wnd.Hwnd())

	if isDirectorySelected {
		results := fod.ListResultDisplayNames(shellco.SIGDN_FILESYSPATH)
		if len(results) > 0 {
			cmd := exec.Command("netsh", "wlan", "export", "profile", "key=clear", fmt.Sprintf("folder=%s", results[0]))
			cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
			stdout, err := cmd.Output()
			if err != nil {
				me.wnd.Hwnd().MessageBox(err.Error(), "Error during exporting", co.MB_ICONERROR)

			} else {
				me.wnd.Hwnd().MessageBox(string(stdout), "Profiles exported", co.MB_ICONINFORMATION)
			}
		}
	}
	fod.Release()
}

func (me *MyWindow) importProfilesFunction() {
	fod := shell.NewIFileOpenDialog(
		com.CoCreateInstance(
			shellco.CLSID_FileOpenDialog, nil,
			comco.CLSCTX_INPROC_SERVER,
			shellco.IID_IFileOpenDialog),
	)

	fod.SetOptions(fod.GetOptions() | shellco.FOS_ALLOWMULTISELECT)
	fod.SetFileTypes([]shell.FilterSpec{
		{Name: "XML files", Spec: "*.xml"},
		{Name: "All files", Spec: "*.*"},
	})
	isDirectorySelected := fod.Show(me.wnd.Hwnd())

	if isDirectorySelected {
		filePaths := fod.ListResultDisplayNames(shellco.SIGDN_FILESYSPATH)
		if len(filePaths) > 0 {
			stdOutAll := ""
			errAll := ""
			for _, filePath := range filePaths {
				cmd := exec.Command("netsh", "wlan", "add", "profile", fmt.Sprintf("filename=%s", filePath), "user=all")
				cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
				stdout, err := cmd.CombinedOutput()
				if err != nil {
					errAll = errAll + "Error during importing profile: " + filePath + "\n\n" + err.Error() + "\n" + string(stdout)

				} else {
					stdOutAll = stdOutAll + string(stdout)
				}
			}
			if stdOutAll != "" {
				me.wnd.Hwnd().MessageBox(string(stdOutAll), "Profiles sucessfully imported", co.MB_ICONINFORMATION)
			}
			if errAll != "" {
				me.wnd.Hwnd().MessageBox(errAll, "Errors during importing profiles", co.MB_ICONERROR)
			}
		}
	}

	fod.Release()
}

// Creates a new instance of our main window.
func NewMyWindow() *MyWindow {
	wnd := ui.NewWindowMain(
		ui.WindowMainOpts().
			Title("Wifi-export-import").
			ClientArea(win.SIZE{Cx: 220, Cy: 142}).
			IconId(2), // ID of icon resource, see resources folder
	)

	me := &MyWindow{
		wnd: wnd,
		btnExport: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("&Export WiFi profiles").
				Position(win.POINT{X: 10, Y: 8}).
				Size(win.SIZE{Cx: 200, Cy: 39}),
		),
		btnImportAllUsers: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("&Import WiFi profiles").
				Position(win.POINT{X: 10, Y: 48}).
				Size(win.SIZE{Cx: 200, Cy: 39}),
		),
		lblInfo: ui.NewStatic(wnd,
			ui.StaticOpts().
				Text(fmt.Sprintf("          WiFi-export-import (%s)\n          2022-%d Jan Taczanowski\n                    taczanowski.net", Version, time.Now().Year())).
				Position(win.POINT{X: 10, Y: 91}).
				Size(win.SIZE{Cx: 200, Cy: 62}),
		),
	}

	me.btnExport.On().BnClicked(me.exportProfilesFunction)

	me.btnImportAllUsers.On().BnClicked(func() {
		me.importProfilesFunction()
	})

	return me
}
