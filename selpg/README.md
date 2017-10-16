# 服务计算第二次作业

#### 要求：实现selpg命令，使用Go语言

以下为该命令参数形式
> -s start_page -e end_page [ -f | -l lines_per_page ] [ -d dest ] [ in_filename ]

各参数意义如下
>-s start page 表示打印开始的页数
-e end page 表示打印结束页数
-f 可选，表示文件是否以\f作为分页标志，与-l不能一起使用
-l 可选，表示文件几行作为一页，与-f不能一起使用
-d 可选，目标打印机
in_filename 可选，读入的文件，若不选则从标准流输入

#### 说明
本次代码中使用cat命令代替测试，如果有打印机可把对应cat命令那行代码注释掉，把其上面一行代码从注释变成代码
