package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func useSubTempl() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	return (substrs[0] + " at " + substrs[1])
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "sub" {
		useSubTempl()
		return
	}

	f1 := Friend{Fname: "Friend1"}
	f2 := Friend{Fname: "Friend2"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(
		`hello {{.UserName}}!
		{{range .Emails}}
		an email {{.|emailDeal}}
		{{end}}
		{{with .Friends}}
		{{range .}}
		my friend name is {{.Fname}}
		{{end}}
		{{end}}`)
	p := Person{UserName: ",Ana",
		Emails:  []string{"ana@tmp.me", "anme@hoge.jp"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

	fmt.Println()
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("からの pipeline if demo : {{if ``}} 出力できません{{end}}\n"))
	tEmpty.Execute(os.Stdout, p)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("からでない pipeline if demo : {{if `anything`}} 出力します{{end}}\n"))
	tWithValue.Execute(os.Stdout, p)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo : {{if `hoge`}} if {{else}} else {{end}}\n"))
	tIfElse.Execute(os.Stdout, p)
}
