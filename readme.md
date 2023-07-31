# coverTool 
golang cover.html 转换工具
* 将go tool cover 工具生成的html文件以树形展示,方便查看

# start
```shell
#得到cover.html文件
go tool cover -html cover.out -o cover.html

# 用工具转换html文件 并在本地开启服务
coverTool --cover=cover.html --server=:8080 

# 或者直接得到输出文件
coverTool --cover=cover.html --out=coverOutDir 
```
