build:
	go build -o ./.sim/simgen

demo:
	cd .sim && ./simgen -template=./form.tpl.html -data=./form.tpl.json -output=./generated_form.html

generate:
	@go run simgen.go -template=./.sim/form.tpl.html -data=./.sim/form.tpl.json -output=./generated_form.html
