/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-17 11:12:30
 */
package main

import (
	files "github.com/gphper/ginadminctl/cli/file"
	_ "github.com/gphper/ginadminctl/global"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "ginadminctl"}
	rootCmd.AddCommand(files.CmdFile)
	rootCmd.Execute()

}
