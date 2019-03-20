package orders

import (
	"fmt"
	"opms/controllers"
	. "opms/models/orders"
	"opms/models/users"
	"opms/models/orders"
	"opms/utils"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type OrderController struct {
	controllers.BaseController
}

// @router /order/manage [get]
func (this *OrderController) ListAllOrders() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "order-manage") {
		this.Redirect("/my/task", 302)
		return
		//this.Abort("401")
	}
	page, err := this.GetInt("p")
	status := this.GetString("status")
	keywords := this.GetString("keywords")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	condArr := make(map[string]string)
	condArr["status"] = status
	condArr["keywords"] = keywords

	countOrder := CountOrder(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countOrder)
	_, _, orders := ListOrder(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["orders"] = orders
	this.Data["countOrder"] = countOrder

	this.TplName = "orders/order.tpl"
}

// @router /order/showOrder [get]
func (this *OrderController) ShowOrder() {

}

// @router /order/edit/:id [get]
func (this *OrderController) EditOrder() {

}

// @router /order/remove/:id [get]
func (this *OrderController) RemoveOrder() {

}

// @router /order/add [get]
func (this *OrderController) AddOrder() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-add") {
		this.Abort("401")
	}
	conArr := make(map[string] string)
	_, _, userList := users.ListUser(conArr, 1, 100)
	var order Orders
	order.Status = 1
	_,_, customers := orders.ListCustomer(conArr,1,100 )
	_,_, departs := users.ListDeparts(conArr,1,100)
	this.Data["customers"] = customers
	this.Data["departs"] = departs
	this.Data["userList"] = userList
	this.Data["order"] = order
	this.TplName = "orders/order-form.tpl"
}

// @router /order/add [post]
func (this *OrderController) AddOrderPost() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写客户名称"}
		this.ServeJSON()
		return
	}
	code := this.GetString("code")
	if "" == code {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写客户代码"}
		this.ServeJSON()
		return
	}
	detail := this.GetString("detail")
	if "" == detail {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写客户开票信息"}
		this.ServeJSON()
		return
	}
	useridStr := this.GetString("userid")
	if "" == detail {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写客户负责人"}
		this.ServeJSON()
		return
	}

	var err error
	status,err := strconv.Atoi(this.GetString("status"))
	if err !=nil{
		return
	}


	userid,err := strconv.ParseInt(useridStr,10,64)

	if err !=nil{
		return
	}

	//雪花算法ID生成
	customerid := utils.SnowFlakeId()

	var customer Customers
	customer.Customerid = customerid
	customer.Userid = userid
	customer.Name = name
	customer.Code = code
	customer.Detail = detail
	customer.Status = status

	err = AddCustomer(customer)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "客户信息添加成功", "id": fmt.Sprintf("%d", customerid)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "客户信息添加失败"}
	}
	this.ServeJSON()
	return
}
// @router /order/ajax/status [post]
func (this *OrderController) AjaxStatus() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "order-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择项目"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status >= 5 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeOrderStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "项目状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "项目状态更改失败"}
	}
	this.ServeJSON()
}

