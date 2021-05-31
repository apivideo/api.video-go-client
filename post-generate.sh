cp -R ../../templates/go/statics/.github ./
cp -R ../../templates/go/statics/* ./
gofmt -s -w *.go