// 設定ファイルの読み出し

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	confFileName = "conf.json"
	uidPos       = "%UID%"
	emailPos     = "%EMAIL%"
)

// Config は設定情報の型
type Config struct {
	Server     string   `json:"server"`
	Attributes []string `json:"attributes"`
	Email      string   `json:"email"`
	BaseDN     string   `json:"base_dn"`
	BindDN     string   `json:"bind_dn"`
	UIDFilter  string   `json:"uid_filter"`
	Filter     string   `json:"filter"`
	Debug      bool     `json:"debug"`
}

// LoadConfig はファイルに保存された JSON オブジェクトを読み出す関数
func LoadConfig() (conf Config, err error) {

	// バイト列読み出し
	var bytes []byte
	bytes, err = ioutil.ReadFile(resolvConfFile())
	if err != nil {
		return
	}

	// json 形式のデコード
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return
	}

	// LDAP サーバの値チェック
	if conf.Server == "" {
		err = fmt.Errorf("server is empty")
		return
	}

	// メールアドレスの値チェック
	if conf.Email == "" {
		err = fmt.Errorf("email is empty")
		return
	}

	// ベースDNの値チェック
	if conf.BaseDN == "" {
		err = fmt.Errorf("base_dn is empty")
		return
	}

	// バインドDNの値チェック
	if conf.BindDN == "" {
		err = fmt.Errorf("bind_dn is empty")
		return
	}

	// 取得する属性リストのチェック
	if len(conf.Attributes) < 1 {
		err = fmt.Errorf("attributes is empty")
		return
	}

	// BindDN の中の "%EMAIL%" を conf.Email で置換する
	conf.BindDN = strings.Replace(conf.BindDN, emailPos, conf.Email, -1)

	return
}

// resolvConfFile は設定ファイルのパスを特定する関数。
// 実行ファイルと同じディレクトリ配下の設定ファイルのパスとする。
func resolvConfFile() string {
	// 実行ファイルのパスを特定
	exe, err := os.Executable()
	if err == nil {
		// 実行ファイルのあるディレクトリ配下の設定ファイルのパス
		return filepath.Dir(exe) + "/" + confFileName
	}

	// 実行カレントディレクトリ配下の設定ファイルのパス
	return confFileName
}
