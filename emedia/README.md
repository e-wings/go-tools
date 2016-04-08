# media


首先要有如下模块：
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

* tbl_key.sql  tbl_media.sql  tbl_media_type.sql 是对应的mysql数据结构、

* key.go media.go mediaType.go 是对应的数据库操作文件

###	使用

* emedia_test.go中有对应的增删改操作和图片处理

* 图片处理更多详情查看 https://github.com/disintegration/gift

###	beego中使用mysql数据库：
* 0.在对应的mysql数据库中创建对应的三张表
* 1.在model中引入该emedia包
* 2.在model中init方法中添加	
orm.RegisterModel(new(emedia.Key), new(emedia.MediaType), new(emedia.Media),...)
* 3.直接调用包中的增删改方法

###	测试
* 下载之后在emedia目录中运行go test命令 对应的是运行测试文件emedia_test.go

会在对应的盘符下生成数据库文件：dataablog.db.db

