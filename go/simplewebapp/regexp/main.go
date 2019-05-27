package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func usage() {
	fmt.Printf("USAGE:\n%s <http|find>", os.Args[0])
}

func regHTTP() {
	url := "http://google.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get err")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read err")
		return
	}

	src := string(body)

	// HTMLタグを全て小文字化
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// <style>タグを除去
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// <script>タグを除去
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// <>内のHTMLコードを削除,改行に置き換える
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	// 連続した改行を除去
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}

func regFind() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// 正規表現にマッチした最初のものを探し出す
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	//正規表現にマッチするすべてのsliceを探し出す.nが0よりも小さかった場合はすべてのマッチする文字列を返す
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	//条件にマッチするindexの位置を探し出す.開始位置と終了位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	//Submatchを探し出し配列を返す.
	//以下の出力でははじめの要素は"am learning Go language"です。
	//２つ目の要素は" learning Go "です。空白を含んで出力することに注意してください。
	//３つ目の要素は"uage"です。
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	//定義と上のFindIndexは同じ
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println("FindSubmatchIndex", submatchindex)

	//FindAllSubmatchは条件マッチする全てのサブマッチを探し出す
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println("FindAllSubmatch", submatchall)

	//FindAllSubmatchIndexは全てのサブマッチのindexを探し出す
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "http":
			regHTTP()
		case "find":
			regFind()
		default:
			usage()
		}
	}
}
