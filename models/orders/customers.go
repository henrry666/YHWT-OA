package orders

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"opms/models"
	"opms/utils"
)

type Customers struct {
	Customerid int64 `orm:"pk;column(customerid);"`
	Code       string
	Name       string
	Detail     string
	Userid     int64
	Status     int
}

func (this *Customers) TableName() string {
	return models.TableName("customers")
}
func init() {
	orm.RegisterModel(new(Customers))
}

func GetCustomer(id int64) (Customers, error) {
	var customer Customers
	var err error

	o := orm.NewOrm()
	customer = Customers{Customerid: id}
	err = o.Read(&customer)

	return customer, err
}

func GetCode(id int64) string {
	var err error
	var code string

	err = utils.GetCache("GetOrderNo.id."+fmt.Sprintf("%d", id), &code)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var order Orders
		o := orm.NewOrm()
		o.QueryTable(models.TableName("orders")).Filter("orderid", id).One(&order, "orderNo")
		code = order.Orderno
		utils.SetCache("GetOrderNo.id."+fmt.Sprintf("%d", id), code, cache_expire)
	}
	return code
}

//修改客户信息
func UpdateCustomer(id int64, updCstm Customers) error {
	o := orm.NewOrm()
	_, err := o.Update(&updCstm, "customerid", "code", "name", "detail", "userid", "status")
	return err
}

//删除客户
func RemoveCustomer(id int64, updCstm Customers) error {
	o := orm.NewOrm()
	p, err := o.Raw("Delete from " + models.TableName("orders") + " where id=?").Prepare()
	res, err := p.Exec(id)
	if res == nil {
		return err
	}
	p.Close()

	//_, err := o.Delete(&updCstm, "customerid")
	return err
}

//获取客户名称
func GetCustomerName(id int64) string {
	var err error
	var name string
	err = utils.GetCache("GetCustomerName.id."+fmt.Sprintf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var customer Customers
		o := orm.NewOrm()
		o.QueryTable(models.TableName("customers")).Filter("customerid", id).One(&customer, "name")
		name = customer.Name
		utils.SetCache("GetCustomerName.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}

//添加客户
func AddCustomer(updCstm Customers) error {
	o := orm.NewOrm()

	_, err := o.Insert(&updCstm)
	return err
}

//客户列表
func ListCustomer(condArr map[string]string, page int, offset int) (num int64, err error, customer []Customers) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("customers"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("name__icontains", condArr["keywords"]).Or("code__icontains", condArr["keywords"]))
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

	var customers []Customers
	qs = qs.OrderBy("-customerid")
	num, err1 := qs.Limit(offset, start).All(&customers)
	return num, err1, customers
}

//统计数量
func CountCustomer(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("customers"))
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("name__icontains", condArr["keywords"]).Or("code__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}

//改变客户状态
func ChangeCustomerStatus(id int64, status int) error {
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
