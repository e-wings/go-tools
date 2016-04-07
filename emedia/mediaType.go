package emedia

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type MediaType struct {
	Id         int
	Name       string
	Width      int
	Height     int
	Remark     string
	Status     int
	AddTime    string
	UpdateTime string
}

func (this *MediaType) TableName() string {
	return "tbl_media_type"
}

//增
func AddMediaType(name, str_width, str_height, remark string) error {
	var err error
	o := orm.NewOrm()
	model := new(MediaType)
	//赋值
	width, _ := strconv.Atoi(str_width)
	height, _ := strconv.Atoi(str_height)
	model.Name = name
	model.Width = width
	model.Height = height
	model.Remark = remark
	model.Status = 1
	model.AddTime = Substr(time.Now().String(), 0, 10)
	model.UpdateTime = Substr(time.Now().String(), 0, 10)
	if err != nil {
		return err
	}
	//保存
	_, err = o.Insert(model)
	return err
}

//改
func UpdateMediaType(str_id, name, str_width, str_height, remark, status string) error {
	//获取原值
	id, _ := strconv.Atoi(str_id)
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_media_type")
	model := &MediaType{}
	err := qs.Filter("id", id).One(model)
	if err != nil {
		return err
	}
	if status == "0" {
		//删除
		model.Status = 0
	} else {
		//更新
		width, _ := strconv.Atoi(str_width)
		height, _ := strconv.Atoi(str_height)
		model.Id = id
		model.Name = name
		model.Width = width
		model.Height = height
		model.Remark = remark
		model.UpdateTime = Substr(time.Now().String(), 0, 10)
	}
	_, err = o.Update(model)
	if err != nil {
		return err
	}
	return nil
}

//查
func ViewMediaType(id string) (*MediaType, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_media_type")
	model := &MediaType{}
	err := qs.Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func ListMediaType(name, status, str_page_num, str_page_size string) ([]*MediaType, int, int, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_media_type")
	var model []*MediaType
	//每页显示多少条
	page_size, _ := strconv.Atoi(str_page_size)
	page_num, _ := strconv.Atoi(str_page_num)
	//筛选
	if name != "" {
		qs = qs.Filter("name__icontains", name)
	}
	if status != "" {
		qs = qs.Filter("status", status)
	}
	//排序
	qs = qs.OrderBy("-status", "-id")
	//查询共有多少数据
	cnt, _ := qs.Count()
	//共有多少页
	page_sum := (int(cnt) + page_size - 1) / page_size
	if page_num > page_sum {
		return nil, 0, 0, errors.New("404 page_num too big")
	}
	//分页
	qs = qs.Limit(page_size, (page_num-1)*page_size)

	_, err := qs.All(&model)
	if err != nil {
		return nil, 0, 0, err
	}
	return model, page_sum, int(cnt), nil
}
