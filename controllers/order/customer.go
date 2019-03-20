package orders

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"opms/controllers"
	. "opms/models/orders"
	"opms/models/users"
	"opms/utils"
	"strconv"
	"strings"
)

type CustomerController struct {
	controllers.BaseController
}

// @router /customer/manage [get]
func (this *CustomerController) ListAllCustomers() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-manage") {
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

	countCustomer := CountCustomer(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countCustomer)
	_, _, customers := ListCustomer(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["customers"] = customers
	this.Data["countCustomer"] = countCustomer

	this.TplName = "orders/customer.tpl"
}

// @router /customer/showOrder [get]
func (this *CustomerController) ShowOrder() {

}

// @router /customer/edit/:id [get]
func (this *CustomerController) EditCustomerGet() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return
	}
	customer, err := GetCustomer(int64(id))
	if err != nil {
		return
	}
	conArr := make(map[string] string)
	_, _, teams := users.ListUser(conArr, 1, 100)


	this.Data["teams"] = teams
	this.Data["customer"] = customer
	this.TplName = "orders/customer-edit.tpl"
}
// @router /customer/edit [post]
func (this *CustomerController) EditCustomerPost() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-edit") {
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
	customerid,err := this.GetInt64("customerid")

	if err !=nil{
		return
	}

	var customer Customers
	customer.Customerid = customerid
	customer.Userid = userid
	customer.Name = name
	customer.Code = code
	customer.Detail = detail
	customer.Status = status

	err = UpdateCustomer(customerid,customer)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "客户信息修改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "客户信息修改失败"}
	}
	this.ServeJSON()
	return
}

// @router /customer/remove/:id [get]
func (this *CustomerController) RemoveOrder() {

}

// @router /customer/add [get]
func (this *CustomerController) AddCustomerGet() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "customer-add") {
		this.Abort("401")
	}
	conArr := make(map[string] string)
	_, _, teams := users.ListUser(conArr, 1, 100)
	var customer Customers
	customer.Status = 1
	customer.Userid = this.UserUserId
	this.Data["teams"] = teams
	this.Data["customer"] = customer
	this.TplName = "orders/customer-form.tpl"
}

// @router /customer/add [post]
func (this *CustomerController) AddCustomerPost() {
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
		this.Data["json"] = map[string]interface{}{"code": 1, "message":
			"客户信息添加成功", "id": fmt.Sprintf("%d", customerid)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "客户信息添加失败"}
	}
	this.ServeJSON()
	return
}

// @router /customer/ajax/status [post]
func (this *CustomerController) AjaxStatus() {
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


