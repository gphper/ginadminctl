package file

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/gphper/ginadminctl/comment"
	"github.com/gphper/ginadminctl/global"

	"github.com/spf13/cobra"
)

type Params struct {
	Name  string
	Table string
	Short string
}

var modelStr string = `package models

type {{.Name}} struct {
	BaseModle
}

func ({{.Short}} *{{.Name}}) TableName() string {
	return "{{.Table}}"
}

func ({{.Short}} *{{.Name}}) FillData() {
	
}
`

var cmdModel = &cobra.Command{
	Use:   "model [-m modelName]",
	Short: "create model",
	Run:   modelFunc,
}

var modelName string

func init() {
	cmdModel.Flags().StringVarP(&modelName, "model", "m", "", "input model name eg: shop_items")
}

func modelFunc(cmd *cobra.Command, args []string) {
	if len(modelName) == 0 {
		cmd.Help()
		return
	}
	fileName, firstName := comment.StrFirstToUpper(modelName)
	err := writeModel(fileName, firstName)

	if err != nil {
		fmt.Printf("[error] %s", err.Error())
		return
	}
	modifyDefault(fileName)
}

func writeModel(fileName string, firstName string) error {

	parms := Params{
		Name:  fileName,
		Table: modelName,
		Short: firstName,
	}

	newPath := global.Path + "\\internal\\models\\" + fileName + ".go"
	_, err := os.Lstat(newPath)
	if err == nil {
		return errors.New("file already exist")
	}

	file, err := os.Create(newPath)
	if err != nil {
		cobra.CompError(err.Error())
		return err
	}
	defer file.Close()

	tem, _ := template.New("models_file").Parse(modelStr)
	tem.ExecuteTemplate(file, "models_file", parms)
	return nil
}

func modifyDefault(fileName string) {
	oldPath := global.Path + "\\internal\\models\\default.go"
	newPath := global.Path + "\\internal\\models\\default_tmp.go"
	file, err := os.Open(oldPath)
	if err != nil {
		cobra.CompError(err.Error())
	}

	reader := bufio.NewReader(file)

	file_tmp, _ := os.Create(newPath)
	var flagTag int
	for {
		bytes, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if strings.Contains(string(bytes), "GetModels") {
			flagTag++
		}

		if flagTag > 0 {
			flagTag++
		}

		if flagTag == 4 {
			file_tmp.Write([]byte("\t\t&" + fileName + "{},\n"))
		}

		file_tmp.Write(bytes)
	}
	file.Close()
	file_tmp.Close()

	err = os.Remove(oldPath)
	if err != nil {
		cobra.CompError(err.Error())
	}
	err = os.Rename(newPath, oldPath)
	if err != nil {
		cobra.CompError(err.Error())
	}
}
