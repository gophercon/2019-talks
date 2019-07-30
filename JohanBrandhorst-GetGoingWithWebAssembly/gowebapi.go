package main

func main() {
	doc := webapiall.GetWindow().Document()
	d := doc.GetElementById("demo")
	d.SetInnerHTML("Hello <b>Go!</b>")
}
