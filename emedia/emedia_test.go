package emedia

import (
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
)

const (
	_DB_Name        = "/dataablog.db.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//新建测试
func TestCreate(t *testing.T) {
	err := CreateDB()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("CreateDB Success")
	}

	err = CreateMediaType()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("CreateMediaType Success")
	}

	err = CreateKey()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("CreateKey Success")
	}

	err = CreateMedia()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("CreateMedia Success")
	}
}

//修改-删除 测试
func TestUpdate(t *testing.T) {
	err := UpdateMediaType("1", "风景改", "1024", "800", "", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("UpdateMediaType Success")
	}

	err = UpdateKey("1", "风景图改", "60", "", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("UpdateKey Success")
	}

	err = UpdateMedia("1", "", "", "", "", "", "", "0")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("DeleteMedia Success")
	}
}

//创建数据库
func CreateDB() error {
	if !com.IsExist(_DB_Name) {
		os.MkdirAll(path.Dir(_DB_Name), os.ModePerm)
		os.Create(_DB_Name)
	}

	orm.RegisterModel(new(Key), new(Media), new(MediaType))
	err := orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_Name, 10)
	if err != nil {
		return err
	} else {
		orm.RunSyncdb("default", true, true) //数据库链接、是否强制覆盖、是否更新
		return nil
	}
}

//新建图片类型
func CreateMediaType() error {
	err := AddMediaType("风景", "1024", "1024", "")
	if err != nil {
		return err
	} else {
		return nil
	}
}

//新建关键字
func CreateKey() error {
	err := AddKey("风景图", "50", "")
	if err != nil {
		return err
	} else {
		return nil
	}
}

//图片处理并入库
func CreateMedia() error {
	file, err := os.Open("test.jpg")
	if err != nil {
		return errors.New("加载原图片失败")
	}
	defer file.Close()

	test_file, err := os.Create("test_img.jpg")
	if err != nil {
		return errors.New("创建新图片失败")
	}
	defer test_file.Close()

	img, err := jpeg.Decode(file) //解码
	if err != nil {
		return errors.New("原图片解码失败")
	}
	//获取图片类型
	media_type, err := ViewMediaType("1")
	if err != nil {
		return errors.New("没有图片类型数据")
	}
	//改变图片
	g := gift.New(
		//图片宽高赋值
		gift.Resize(media_type.Width, media_type.Height, gift.LanczosResampling),
		gift.UnsharpMask(1.0, 1.0, 0.0),
		gift.Convolution(
			[]float32{
				-1, -1, 0,
				-1, 1, 1,
				0, 1, 1,
			},
			false, false, false, 0.0,
		),
	)
	dst := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(dst, img)
	//图片入库
	err = AddMedia(strconv.Itoa(media_type.Id), "0", "1", "test_img.jpg", "50", "")
	if err != nil {
		return errors.New("图片入库失败")
	}
	//图片写入文件
	err = jpeg.Encode(test_file, dst, &jpeg.Options{100}) //编码
	if err != nil {
		return err
	} else {
		return nil
	}
}
