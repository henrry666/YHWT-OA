package orders

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"opms/models"
	"opms/utils"
	"time"
)

type Orders struct {
	Id                 int64 `orm:"pk;column(id);"`
	Orderno            string
	Source             int64
	Task_unit          string
	Customer_responser string
	Product_responser  string
	Task_name          string
	Task_priviledge    string
	Is_lailiaojiagong  int
	Has_paper          int
	Paper_count        int
	Requied_date       time.Time
	Actual_date        time.Time
	Has_outer_kesu     int
	Has_inner_kesu     int
	Task_description   string
	Mcenter_opinion    string
	Approve_sign       string
	Order_arrival_time time.Time
	Special_illustrate string
	Creator            int64
	Status             int
}

func (this *Orders) TableName() string {
	return models.TableName("orders")
}
func init() {
	orm.RegisterModel(new(Orders))
}

func GetOrder(id int64) (Orders, error) {
	var order Orders
	var err error
	o := orm.NewOrm()
	order = Orders{Id: id}
	err = o.Read(&order)
	return order, err
}

func GetOrderNo(id int64) string {
	var err error
	var orderNo string

	err = utils.GetCache("GetOrderNo.id."+fmt.Sprintf("%d", id), &orderNo)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var order Orders
		o := orm.NewOrm()
		o.QueryTable(models.TableName("orders")).Filter("orderid", id).One(&order, "orderNo")
		orderNo = order.Orderno
		utils.SetCache("GetOrderNo.id."+fmt.Sprintf("%d", id), orderNo, cache_expire)
	}
	return orderNo
}

func UpdateOrder(id int64, updOrder Orders) error {
	var order Orders
	o := orm.NewOrm()
	order = Orders{Id: id}
	order.Id = updOrder.Id
	order.Orderno = updOrder.Orderno
	order.Source = updOrder.Source
	order.Task_unit = updOrder.Task_unit
	order.Customer_responser = updOrder.Customer_responser
	order.Product_responser = updOrder.Product_responser
	order.Task_name = updOrder.Task_name
	order.Task_priviledge = updOrder.Task_priviledge
	order.Is_lailiaojiagong = updOrder.Is_lailiaojiagong
	order.Has_paper = updOrder.Has_paper
	order.Paper_count = updOrder.Paper_count
	order.Requied_date = updOrder.Requied_date
	order.Actual_date = updOrder.Actual_date
	order.Has_outer_kesu = updOrder.Has_outer_kesu
	order.Has_inner_kesu = updOrder.Has_inner_kesu
	order.Task_description = updOrder.Task_description
	order.Mcenter_opinion = updOrder.Mcenter_opinion
	order.Approve_sign = updOrder.Approve_sign
	order.Order_arrival_time = updOrder.Order_arrival_time
	order.Special_illustrate = updOrder.Special_illustrate
	order.Creator = updOrder.Creator
	order.Status = updOrder.Status
	//pro.Status = updPro.Status
	_, err := o.Update(&order, "orderno", "source", "task_unit", "customer_responser", "product_responser")
	return err
}

func AddOrder(updOrder Orders) error {
	o := orm.NewOrm()
	order := new(Orders)

	order.Id = updOrder.Id
	order.Orderno = updOrder.Orderno
	order.Source = updOrder.Source
	order.Task_unit = updOrder.Task_unit
	order.Customer_responser = updOrder.Customer_responser
	order.Product_responser = updOrder.Product_responser
	order.Task_name = updOrder.Task_name
	order.Task_priviledge = updOrder.Task_priviledge
	order.Is_lailiaojiagong = updOrder.Is_lailiaojiagong
	order.Has_paper = updOrder.Has_paper
	order.Paper_count = updOrder.Paper_count
	order.Requied_date = updOrder.Requied_date
	order.Actual_date = updOrder.Actual_date
	order.Has_outer_kesu = updOrder.Has_outer_kesu
	order.Has_inner_kesu = updOrder.Has_inner_kesu
	order.Task_description = updOrder.Task_description
	order.Mcenter_opinion = updOrder.Mcenter_opinion
	order.Approve_sign = updOrder.Approve_sign
	order.Order_arrival_time = updOrder.Order_arrival_time
	order.Special_illustrate = updOrder.Special_illustrate
	order.Creator = updOrder.Creator
	order.Status = updOrder.Status
	_, err := o.Insert(order)
	return err
}

//项目列表
func ListOrder(condArr map[string]string, page int, offset int) (num int64, err error, user []Orders) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("orders"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("orderno__icontains", condArr["keywords"]).Or("source__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()

	var orders []Orders
	qs = qs.OrderBy("-id")
	num, err1 := qs.Limit(offset, start).All(&orders)
	return num, err1, orders
}

//统计数量
func CountOrder(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("orders"))
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("orderno__icontains", condArr["keywords"]).Or("source__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}

func ListMyOrder(userId int64, page int, offset int) (num int64, err error, ops []Orders) {
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	var orders []Orders

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("p.id", "p.orderno", "p.source", "p.task_unit", "p.customer_responser", "p.product_responser", "p.status").From("pms_orders AS t").
		Where("t.id=?").
		Limit(offset).Offset(start)
	sql := qb.String()
	o := orm.NewOrm()
	nums, err := o.Raw(sql, userId).QueryRows(&orders)
	return nums, err, orders
}

func ChangeOrderStatus(id int64, status int) error {
	o := orm.NewOrm()

	order := Orders{Id: id}
	err := o.Read(&order, "id")
	if nil != err {
		return err
	} else {
		order.Status = status
		_, err := o.Update(&order)
		return err
	}
}
