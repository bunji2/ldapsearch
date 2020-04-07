package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

const (
	winTitle = "LDAP Client"
)

var conf Config

// フォームの入力エントリ群
var emailEntry *widget.Entry
var passwdEntry *widget.Entry
var patternEntry *widget.Entry

var ap fyne.App
var w fyne.Window

func main() {
	ap = app.New()
	w = ap.NewWindow(winTitle)

	var err error

	// 設定ファイルの読み込み
	conf, err = LoadConfig()
	if err != nil {
		dialog.ShowError(err, w)
	}

	// メールアドレスの入力エントリ
	emailEntry = &widget.Entry{
		PlaceHolder: "メールアドレスを入力して下さい",
		Text:        conf.Email,
	}

	// パスワードの入力エントリ
	passwdEntry = &widget.Entry{
		PlaceHolder: "パスワードを入力して下さい",
		Password:    true,
	}

	// 検索パターンの入力エントリ
	patternEntry = &widget.Entry{
		PlaceHolder: "検索パターンを入力して下さい",
		Text:        conf.Filter,
	}

	w.SetContent(makeForm())
	w.ShowAndRun()
}

// checkFormEntry はフォームの入力エントリに値が設定されているかを確認する関数
func checkFormEntry() bool {
	return emailEntry.Text != "" &&
		passwdEntry.Text != "" &&
		passwdEntry.Text != ""
}

// makeForm はメールアドレス・パスワード・検索パターンの
// 入力エントリからなるフォームを作る関数
func makeForm() fyne.CanvasObject {
	form := &widget.Form{
		OnSubmit: submit,
	}
	form.Append("email", emailEntry)
	form.Append("password", passwdEntry)
	form.Append("pattern", patternEntry)
	return form
}

// submit は LDAP サーバから検索結果を取得する関数
func submit() {
	// サーバ設定が空のときはエラーダイアログを表示して終了する
	if conf.Server == "" {
		dialog.ShowError(fmt.Errorf("server is empty"), w)
		return
	}

	// フォームの値をチェックする関数
	if !checkFormEntry() {
		return
	}

	// 検索結果のテキストを格納する変数
	text := ""

	if !conf.Debug {
		// デバッグモードではないときには LDAP サーバに接続し、検索をお粉う
		sr, err := ldapSearch(Params{
			Server:     conf.Server,       // 接続先
			BaseDN:     conf.BaseDN,       //
			BindDN:     conf.BindDN,       //
			Password:   passwdEntry.Text,  // パスワード
			Pattern:    patternEntry.Text, // 検索パターン
			Attributes: conf.Attributes,   // 取得する属性
		})
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		if len(sr.Entries) > 0 {
			//検索結果のリスト出力
			for _, entry := range sr.Entries {
				text += entry.DN + "\n"
				for _, attrName := range conf.Attributes {
					text += fmt.Sprintf("%s=%v\n", attrName, entry.GetAttributeValue(attrName))
				}
				text += "\n"
			}
		}
	} else {
		// デバッグモードのときには、LDAP サーバに接続しない
		text = "email=" + emailEntry.Text + "\npattern=" + patternEntry.Text
	}

	// 結果を表示
	TextViewer(ap, text)
}
