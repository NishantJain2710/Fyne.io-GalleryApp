package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type image struct {
	name      string
	path      string
	size      int64
	extension string
}

func main() {
	app := app.New()

	w := app.NewWindow("Gallery App")

	root_src := "C:\\Users\\nisha\\OneDrive\\Pictures\\Camera Roll"

	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}

	images := []image{}
	for _, file := range files {
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpg" || extension == "jpeg" {
				img := image{name: strings.Split(file.Name(), ".")[0], path: root_src + "\\" + file.Name(), size: file.Size(), extension: extension}
				images = append(images, img)
			}
		}
	}

	name := widget.NewEntry()
	name.TextStyle = fyne.TextStyle{Bold: true}
	name.SetPlaceHolder("null")
	name.Disable()
	name.SetText(images[0].name)

	Extension := widget.NewEntry()
	Extension.TextStyle = fyne.TextStyle{Bold: true}
	Extension.SetPlaceHolder("null")
	Extension.Disable()
	Extension.SetText(images[0].extension)

	Size := widget.NewEntry()
	Size.TextStyle = fyne.TextStyle{Bold: true}
	Size.SetPlaceHolder("null")
	Size.Disable()
	Size.SetText(strconv.FormatInt(images[0].size, 10))

	Path := widget.NewEntry()
	Path.TextStyle = fyne.TextStyle{Bold: true}
	Path.SetPlaceHolder("null")
	Path.Disable()
	Path.SetText(images[0].path)

	tabs := container.NewAppTabs(container.NewTabItem(images[0].name,
		container.NewHSplit(
			container.New(layout.NewGridWrapLayout(fyne.NewSize(750, 650)), canvas.NewImageFromFile(images[0].path)),
			container.NewVBox(
				container.New(
					layout.NewGridLayout(2),
					canvas.NewText("Name: ", color.White),
					name,
				), container.New(
					layout.NewGridLayout(2),
					canvas.NewText("Extension: ", color.White),
					Extension,
				), container.New(
					layout.NewGridLayout(2),
					canvas.NewText("Size: ", color.White),
					Size,
				), container.New(
					layout.NewGridLayout(2),
					canvas.NewText("Path: ", color.White),
					Path,
				),
			),
		),
	))
	for i := 1; i < len(images); i++ {
		name1 := widget.NewEntry()
		name1.TextStyle = fyne.TextStyle{Bold: true}
		name1.SetPlaceHolder("null")
		name1.Disable()
		name1.SetText(images[i].name)

		Extension1 := widget.NewEntry()
		Extension1.TextStyle = fyne.TextStyle{Bold: true}
		Extension1.SetPlaceHolder("null")
		Extension1.Disable()
		Extension1.SetText(images[i].extension)

		Size1 := widget.NewEntry()
		Size1.TextStyle = fyne.TextStyle{Bold: true}
		Size1.SetPlaceHolder("null")
		Size1.Disable()
		Size1.SetText(strconv.FormatInt(images[i].size, 10))

		Path1 := widget.NewEntry()
		Path1.TextStyle = fyne.TextStyle{Bold: true}
		Path1.SetPlaceHolder("null")
		Path1.Disable()
		Path1.SetText(images[i].path)

		tabs.Append(container.NewTabItem(images[i].name,
			container.NewHSplit(
				container.New(layout.NewGridWrapLayout(fyne.NewSize(750, 650)), canvas.NewImageFromFile(images[i].path)),
				container.NewVBox(
					container.New(
						layout.NewGridLayout(2),
						canvas.NewText("Name: ", color.White),
						name1,
					), container.New(
						layout.NewGridLayout(2),
						canvas.NewText("Extension: ", color.White),
						Extension1,
					), container.New(
						layout.NewGridLayout(2),
						canvas.NewText("Size: ", color.White),
						Size1,
					), container.New(
						layout.NewGridLayout(2),
						canvas.NewText("Path: ", color.White),
						Path1,
					),
				),
			),
		))
	}

	tabs.SetTabLocation(container.TabLocationBottom)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(900, 650))
	w.SetContent(tabs)
	w.ShowAndRun()
}
