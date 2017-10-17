# 服务计算第二次作业

#### 要求：实现selpg命令，使用Go语言

以下为该命令参数形式
> -s start_page -e end_page [ -f | -l lines_per_page ] [ -d dest ] [ in_filename ]

各参数意义如下
>-s start page 表示打印开始的页数</br>
-e end page 表示打印结束页数</br>
-f 可选，表示文件是否以\f作为分页标志，与-l不能一起使用</br>
-l 可选，表示文件几行作为一页，与-f不能一起使用</br>
-d 可选，目标打印机</br>
in_filename 可选，读入的文件，若不选则从标准流输入</br>

#### 说明
类似实验报告的文档请看我博客 http://blog.csdn.net/caijhBlog/article/details/78265811
