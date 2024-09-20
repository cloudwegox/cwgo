package iflytek

import (
	"embed"

	"github.com/cloudwego/kitex/tool/internal_pkg/tpl"
)

//go:embed *
var files embed.FS

func Init() {
	tpl.ClientTpl = readAll("cwgo/client.tpl")
	tpl.ServiceTpl = readAll("cwgo/service.tpl")
}

func readAll(path string) string {
	bytes, err := files.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
