package file

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/gphper/ginadminctl/comment"
	"github.com/gphper/ginadminctl/global"

	"github.com/spf13/cobra"
)

type CParams struct {
	ClassName   string
	Pagename    string
	PackageName string
	UpClassName string
	Methods     []string
}

var controllerStr string = `package {{.PackageName}}

import (
	"github/gphper/ginadmin/internal/controllers/admin"
	"github.com/gin-gonic/gin"
)

type {{.ClassName}} struct {
	admin.BaseController
}

var {{.UpClassName}} = {{.ClassName}}{}

{{range .Methods}}

func (con *{{$.ClassName}}) {{.}}(c *gin.Context) {

}


{{end}}
`

var cmdController = &cobra.Command{
	Use:   "controller [-p pagename -c controllerName -l methods]",
	Short: "create controller file",
	Run:   controllerFunc,
}

var (
	pagename       string
	controllerName string
	methods        string
	packagename    string
	upName         string
)

func init() {
	cmdController.Flags().StringVarP(&pagename, "pagename", "p", "", "input pagename eg: setting")
	cmdController.Flags().StringVarP(&controllerName, "controllerName", "c", "", "input controller name eg: AdminController")
	cmdController.Flags().StringVarP(&methods, "methods", "m", "", "input methods eg: index,add,del")
}

func controllerFunc(cmd *cobra.Command, args []string) {
	if len(pagename) == 0 || len(controllerName) == 0 {
		cmd.Help()
		return
	}

	pageSlice := strings.Split(pagename, "\\")
	packagename = pageSlice[len(pageSlice)-1]

	upName, _ = comment.StrFirstToUpper(packagename)

	err := writeController()
	if err != nil {
		fmt.Printf("[error] %s", err.Error())
		return
	}
}

func writeController() error {
	parms := CParams{
		ClassName:   controllerName,
		Pagename:    pagename,
		PackageName: packagename,
		UpClassName: upName,
	}

	mSlice := strings.Split(methods, ",")
	if len(mSlice) != 0 {
		tempMethod := make([]string, len(mSlice))
		for k, v := range mSlice {
			temp, _ := comment.StrFirstToUpper(v)
			tempMethod[k] = temp
		}
		parms.Methods = tempMethod
	}

	basePath := global.Path + "\\internal\\controllers\\" + pagename
	_, err := os.Lstat(basePath)
	if err != nil {
		os.Mkdir(basePath, os.ModeDir)
	}

	newPath := basePath + "\\" + controllerName + ".go"
	_, err = os.Lstat(newPath)
	if err == nil {
		return fmt.Errorf("%s file already exist", controllerName+".go")
	}

	file, err := os.Create(newPath)
	if err != nil {
		cobra.CompError(err.Error())
		return err
	}
	defer file.Close()

	tem, _ := template.New("controller_file").Parse(controllerStr)
	tem.ExecuteTemplate(file, "controller_file", parms)
	return nil
}
