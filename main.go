package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-webkit/webkit"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

var addr string
var filename string

var dark_theme bool

const port = ":5926"

const blank_page = `
<html class="blank_page">

<head>
<style>
%s
</style>
</head>
<body class="blank_page">
</body>
</html>
`

const page_template = `
<html>

<head>
<style>
%s
</style>
</head>
<body>
%s
</body>
</html>
`

var webview *webkit.WebView
var menu_reload *gtk.MenuItem
var menu_exit *gtk.MenuItem

func getContent() string {
	if filename != "" {
		input, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output := blackfriday.MarkdownCommon(input)

		var html string
		if dark_theme {
			html = fmt.Sprintf(page_template, css+css_dark, string(output))
		} else {
			html = fmt.Sprintf(page_template, css, string(output))
		}

		return html

	}
	return fmt.Sprintf(blank_page, css)
}

func reload() {
	webview.LoadString(getContent(), "text/html", "utf-8", ".")
}

func main() {
	if len(os.Args) < 2 {
		filename = ""
	} else {
		filename = os.Args[1]
	}

	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle(filename)
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(false, 1)
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	webview = webkit.NewWebView()
	vbox.Add(webview)

	filemenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(filemenu)
	submenu := gtk.NewMenu()
	filemenu.SetSubmenu(submenu)

	menu_open := gtk.NewMenuItemWithMnemonic("_Open")
	menu_open.Connect("activate", func() {
		filechooserdialog := gtk.NewFileChooserDialog(
			"Choose file",
			window,
			gtk.FILE_CHOOSER_ACTION_OPEN,
			gtk.STOCK_OK,
			gtk.RESPONSE_ACCEPT)
		filter := gtk.NewFileFilter()
		filter.AddPattern("*.md")
		filter.AddPattern("*.markdown")
		filter.SetName("Markdown files")
		filechooserdialog.AddFilter(filter)
		filechooserdialog.Response(func() {
			filename = (filechooserdialog.GetFilename())
			filechooserdialog.Destroy()
		})
		filechooserdialog.Run()
		reload()
	})
	submenu.Append(menu_open)

	menu_reload = gtk.NewMenuItemWithMnemonic("_Reload")
	menu_reload.Connect("activate", func() {
		reload()
	})
	submenu.Append(menu_reload)

	menu_exit = gtk.NewMenuItemWithMnemonic("E_xit")
	menu_exit.Connect("activate", func() {
		gtk.MainQuit()
	})
	submenu.Append(menu_exit)

	viewmenu := gtk.NewMenuItemWithMnemonic("_View")
	menubar.Append(viewmenu)
	viewsubmenu := gtk.NewMenu()
	viewmenu.SetSubmenu(viewsubmenu)

	menu_dark := gtk.NewCheckMenuItemWithMnemonic("D_ark Theme")
	menu_dark.Connect("activate", func() {
		if menu_dark.GetActive() {
			dark_theme = true
		} else {
			dark_theme = false
		}
		reload()
	})
	viewsubmenu.Append(menu_dark)

	window.Add(vbox)
	window.SetSizeRequest(480, 600)
	window.SetIconFromFile("./icons/icon-small.png")
	window.ShowAll()
	webview.LoadString(getContent(), "text/html", "utf-8", ".")

	gtk.Main()
}
