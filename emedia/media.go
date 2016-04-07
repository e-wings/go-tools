package emedia

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Media struct {
	Id         int
	TypeId     int
	MemberId   int
	KeyId      int
	Url        string
	Status     int
	Weight     int
	Remark     string
	AddTime    string
	UpdateTime string
}

func (this *Media) TableName() string {
	return "tbl_media"
}

//增
func AddMedia(str_type_id, str_member_id, str_key_id, url, str_weight, remark string) error {
	var err error
	o := orm.NewOrm()
	model := new(Media)
	//赋值
	type_id, _ := strconv.Atoi(str_type_id)
	member_id, _ := strconv.Atoi(str_member_id)
	key_id, _ := strconv.Atoi(str_key_id)
	weight, _ := strconv.Atoi(str_weight)
	model.TypeId = type_id
	model.MemberId = member_id
	model.KeyId = key_id
	model.Url = url
	model.Remark = remark
	model.Weight = weight
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
func UpdateMedia(str_id, str_type_id, str_member_id, str_key_id, url, str_weight, remark, str_status string) error {
	//获取原值
	id, _ := strconv.Atoi(str_id)
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_media")
	model := &Media{}
	err := qs.Filter("id", id).One(model)
	if err != nil {
		return err
	}
	if str_status == "0" {
		//删除
		model.Status = 0
	} else {
		//更新
		type_id, _ := strconv.Atoi(str_type_id)
		member_id, _ := strconv.Atoi(str_member_id)
		key_id, _ := strconv.Atoi(str_key_id)
		weight, _ := strconv.Atoi(str_weight)
		status, _ := strconv.Atoi(str_status)
		model.Id = id
		model.Status = status
		model.TypeId = type_id
		model.MemberId = member_id
		model.KeyId = key_id
		model.Url = url
		model.Remark = remark
		model.Weight = weight
		model.UpdateTime = Substr(time.Now().String(), 0, 10)
	}
	_, err = o.Update(model)
	if err != nil {
		return err
	}
	return nil
}

//查
func ViewMedia(id string) (*Media, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_media")
	model := &Media{}
	err := qs.Filter("id", id).One(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func ListMedia(type_id, member_id, key_id, status, str_page_num, str_page_size string) ([]*Media, int, int, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tbl_media")
	var model []*Media
	//每页显示多少条
	page_size, _ := strconv.Atoi(str_page_size)
	page_num, _ := strconv.Atoi(str_page_num)
	//筛选
	if type_id != "" {
		qs = qs.Filter("type_id", type_id)
	}
	if member_id != "" {
		qs = qs.Filter("member_id", member_id)
	}
	if key_id != "" {
		qs = qs.Filter("key_id", key_id)
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
