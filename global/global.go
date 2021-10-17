/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-10-17 12:03:05
 */
package global

import "os"

var Path string

func init() {
	Path, _ = os.Getwd()
}
