// テキスト表示用のウィンドウ
// ファイル保存ダイアログ付き

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	osdialog "github.com/sqweek/dialog"
)

const (
	textViewerTitle = "textviewer"
	width           = 640
	height          = 480
)

// TextViewer はテキストを表示するウィンドウを作成する関数
func TextViewer(a fyne.App, text string) {

	w := a.NewWindow(textViewerTitle)
	w.Resize(fyne.NewSize(width, height))

	// テキスト用エントリ
	le := &widget.Entry{
		Text:      text,
		MultiLine: true,
	}

	// ファイル保存ダイアログを呼び出すボタン
	bt := widget.NewButton("Save as...", func() {
		// ファイル保存ダイアログを呼び出し
		err := saveAs(le.Text)
		if err != nil && !strings.HasPrefix(err.Error(), "Cancelled") {
			// キャンセル以外のエラー時にはダイアログで表示
			fmt.Println(err)
			dialog.ShowError(err, w)
			return
		}
	})

	// ボーダーレイアウト付きコンテナ
	box := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, bt, nil, nil),
		widget.NewScrollContainer(le), // スクロール付きテキストエントリ
		bt)

	w.SetContent(box)

	w.Show()
	//w.ShowAndRun()
}

// saveAs はファイル保存ダイアログを表示し、選択されたファイルにテキストを保存する関数
func saveAs(text string) (err error) {
	var filename string
	// ファイル保存ダイアログの表示
	filename, err = osdialog.File().Filter("Text files", "txt").Title("Save as a text file").Save()
	if err == nil {
		// テキストの保存
		err = ioutil.WriteFile(filename, []byte(text), os.ModePerm)
	}
	return
}
