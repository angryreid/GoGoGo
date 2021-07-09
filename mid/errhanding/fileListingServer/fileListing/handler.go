package fileListing

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type UserError string

func (e UserError) Error() string {
	//panic("implement me")
	return e.Message()
}

func (e UserError) Message() string {
	//panic("implement me")
	return string(e)
}

func HandlerFileList(
	writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return UserError("path must start as " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}