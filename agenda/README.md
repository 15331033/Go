#### CLI 命令行实用程序开发实战 - Agenda
##### 说明：目前数据持久化由于遇到json不能转小写成员，在利用反射转的过程bug还未解决故部分代码注释了

#### 1.项目结构
-agenda //根目录
main.go
</br>
-- cmd
---  query.go
--- register.go
--- root.go
</br>
-- data
---  agenda.log
---  data.json
</br>
-- entity
---  logger.go
---  meeting.go
---  storage.go
---  user.go
</br>
-- test
---  meeting_test.go
</br>
#### 2.各文件功能
#####1.基本的文件
model层：meeting.go和user.go分别是实现user和meeting
存储：storage.go
log工具：logger.go
controller层：query.go和register.go
UI层：root.go
#####2.数据文件
data.json用于存数据，agenda.log用于存日志记录
####3.测试文件
meeting_test.go用于测试meeting中的一些函数，由于此项目中其他函数大多数为get,set故没写什么测试