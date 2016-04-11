# media

###	安装方法
go get github.com/e-wings/go-tools/emedia

请确保已经安装如下模块：
```bash
	"errors"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	"github.com/disintegration/gift"
	_ "github.com/mattn/go-sqlite3"
	"image"
	"image/jpeg"
	"os"
	"path"
	"strconv"
	"testing"
```

### 文件结构	
* tbl_key.sql		是对应的mysql数据结构--关键词表
* tbl_media.sql		是对应的mysql数据结构--图片表
* tbl_media_type.sql 	是对应的mysql数据结构--图片类型表
* key.go	是对应的数据库操作文件--关键词增删改
* media.go	是对应的数据库操作文件--图片增删改
* mediaType.go	是对应的数据库操作文件--图片类型增删改
* emedia_test.go 测试文件

###	使用

* emedia_test.go中有对应的增删改操作和图片处理

* 图片处理更多详情查看 https://github.com/disintegration/gift

###	beego中使用mysql数据库：
1.在对应的mysql数据库中创建对应的三张表

2.在model中引入该emedia包

3.在model中init方法中添加	
orm.RegisterModel(new(emedia.Key), new(emedia.MediaType), new(emedia.Media),...)注：如有更多表则在后面添加

4.直接调用包中的增删改方法

###	测试
* 安装之后在emedia目录中运行go test命令，测试文件emedia_test.go会被之行

测试文件会生成数据库文件：github.com/e-wings/go-tools/edit/master/emedia/dataablog.db.db
并进行响应测试。如果测试出错请与我联系。

