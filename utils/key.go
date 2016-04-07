package utils

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Key struct {
	Id         int
	Name       string
	Status     int
	Weight     int
	Remark     string
	AddTime    string
	UpdateTime string
}

func (this *Key) TableName() string {
	return "tbl_key"
}

//增
func AddKey(name, str_weight, remark string) error {
	var err error
	o := orm.NewOrm()
	model := new(Key)
	//赋值
	weight, _ := strconv.Atoi(str_weight)
	model.Name = name
	model.Weight = weight
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
func UpdateKey(str_id, name, str_weight, remark, status string) error {
	//获取原值
	id, _ := strconv.Atoi(str_id)
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_key")
	model := &Key{}
	err := qs.Filter("id", id).One(model)
	if err != nil {
		return err
	}
	if status == "0" {
		//删除
		model.Status = 0
	} else {
		//更新
		weight, _ := strconv.Atoi(str_weight)
		model.Id = id
		model.Name = name
		model.Weight = weight
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
func ViewKey(id string) (*Key, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_key")
	model := &Key{}
	err := qs.Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func ListKey(name, status, str_page_num, str_page_size string) ([]*Key, int, int, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_key")
	var model []*Key
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
