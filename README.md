# ginadminctl
ginadmin 命令控制

## 依赖
* golang > 1.16
* github.com/spf13/cobra

## 功能清单

:white_check_mark:创建文件

:black_square_button:安装模块

## 使用文档
- [创建文件](#创建文件)


### :small_blue_diamond:<a name="创建文件">创建文件</a>

```
ginadminctl.exe file -h 
File Operation

Usage:
  ginadminctl file [File Operation] [flags]
  ginadminctl file [command]

Available Commands:
  controller  create controller file
  model       create model

Flags:
  -h, --help   help for file
```

2. 创建model文件

   ```shell
   ginadminctl.exe file model -m article
   ```

3. 创建controller文件

   ```
   ginadminctl.exe file controller -p admin\article -c articleController -m add,list
   ```
   
   * -p 路径
   * -c 控制器名称
   * -m 控制器中包含的方法