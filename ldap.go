// LDAP 検索

package main

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// Params は ldapSearch のパラメータを指定するための構造体型
type Params struct {
	Server     string
	BaseDN     string
	BindDN     string
	Password   string
	Pattern    string
	Attributes []string
}

// ldapSearch は LDAPサーバに接続し、検索する関数
func ldapSearch(p Params) (sr *ldap.SearchResult, err error) {
	var con *ldap.Conn

	// LDAPサーバへ接続
	fmt.Printf("\nConnecting %s...\n", p.Server)
	con, err = ldap.Dial("tcp", p.Server)
	if err != nil {
		return
	}

	fmt.Println("Connected.")
	defer func() {
		con.Close()
		fmt.Println("Disconnected.")
	}()

	// LDAPサーバ認証（LDAP BIND）
	err = con.Bind(p.BindDN, p.Password)
	if err != nil {
		return
	}
	fmt.Println("Logged in.")

	/*
		検索リクエスト作成

		コンストラクタ関数の引数
		func NewSearchRequest(
			BaseDN          string,
			Scope           int,
			DerefAliases    int,
			SizeLimit       int,
			TimeLimit       int,
			TypesOnly       bool,
			Filter          string,
			Attributes      []string,
			Controls        []Control,
		) *SearchRequest
	*/

	searchRequest := ldap.NewSearchRequest(
		p.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		p.Pattern,
		p.Attributes,
		nil,
	)

	//検索リクエストを基にディレクトリ内検索
	sr, err = con.Search(searchRequest)

	return
}
