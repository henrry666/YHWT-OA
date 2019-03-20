package notices
import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/utils/pagination"
	"opms/controllers"
	."opms/models/notices"
	"strconv"
	"strings"
)
type NoticesController struct {
    controllers.BaseController
}
// @router /notices/manage [get]           
func (this *NoticesController) ListAllNoticess() {       
	//权限检测               
	if !strings.Contains(this.GetSession("userPermission").(string), "notices-manage") {
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
              
	countNotices := CountNotices(condArr)                
	paginator := pagination.SetPaginator(this.Ctx, offset, countNotices) 
	_, _, notices := ListNotices(condArr, page, offset)                
              
	this.Data["paginator"] = paginator      
	this.Data["condArr"] = condArr          
	this.Data["notices"] = notices      
	this.Data["countNotices"] = countNotices             
              
	this.TplName = "orders/notices.tpl"    
}             
              
// @router /notices/showOrder [get]        
func (this *NoticesController) ShowOrder() {              
              
}             
              
// @router /notices/edit/:id [get]         
func (this *NoticesController) EditNoticesGet() {        
	//权限检测               
	if !strings.Contains(this.GetSession("userPermission").(string), "notices-edit") {  
		this.Abort("401")    
	}         
	idstr := this.Ctx.Input.Param(":id")    
              
	id, err := strconv.Atoi(idstr)          
	if err != nil {          
		return               
	}         
	notices, err := GetNotices(int64(id))                
	if err != nil {          
		return               
	}         
	this.Data["notices"] = notices        
	this.TplName = "orders/notices-edit.tpl"              
}             
// @router /notices/edit [post]            
func (this *NoticesController) EditNoticesPost() {       
	//权限检测               
	if !strings.Contains(this.GetSession("userPermission").(string), "notices-edit") {  
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}     
		this.ServeJSON()     
		return               
	}         
    noticeid,error:= this.GetInt64("noticeid") 
	if error != nil {  
		return               
	}         
    title:= this.GetString("title") 
	if "" == title {  
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写"}               
		this.ServeJSON()     
		return               
	}         
    content:= this.GetString("content") 
	if "" == content {  
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写"}               
		this.ServeJSON()     
		return               
	}         
    created, error:= this.GetInt("created") 
	if error != nil {  
		return               
	}         
    status, error:= this.GetInt("status") 
	if error != nil {  
		return               
	}         
	var notices Notices                  
	notices.Noticeid = noticeid
	notices.Title = title
	notices.Content = content
	notices.Created = created
	notices.Status = status
              
var	err = UpdateNotices(noticeid,notices)              
              
	if err == nil {          
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "notices信息修改成功"}             
	} else {                 
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "notices信息修改失败"}             
	}         
	this.ServeJSON()         
	return                   
}             
              
// @router /notices/remove/:id [get]       
func (this *NoticesController) RemoveOrder() {            
              
}             
              
// @router /notices/add [get]              
func (this *NoticesController) AddNoticesGet() {         
	//权限检测               
	if !strings.Contains(this.GetSession("userPermission").(string), "notices-add") {   
		this.Abort("401")    
	}         
	var notices Notices                  
	this.Data["notices"] = notices        
	this.TplName = "notices/notices-form.tpl"              
}             
              
// @router /notices/add [post]             
func (this *NoticesController) AddNoticesPost() {        
	//权限检测               
	if !strings.Contains(this.GetSession("userPermission").(string), "notices-add") {   
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}     
		this.ServeJSON()     
		return               
	}         
    noticeid, error:= this.GetInt64("noticeid") 
	if error != nil {  
		return               
	}         
    title:= this.GetString("title") 
	if "" == title {  
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写"}               
		this.ServeJSON()     
		return               
	}         
    content:= this.GetString("content") 
	if "" == content {  
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写"}               
		this.ServeJSON()     
		return               
	}         
    created, error:= this.GetInt("created") 
	if error != nil {  
		return               
	}         
    status, error:= this.GetInt("status") 
	if error != nil {  
		return               
	}         
	var notices Notices                  
	notices.Noticeid = noticeid
	notices.Title = title
	notices.Content = content
	notices.Created = created
	notices.Status = status
              
var	err = AddNotices(notices)             
              
	if err == nil {          
		this.Data["json"] = map[string]interface{}{"code": 1, "message":  
			"notices信息添加成功", "id": fmt.Sprintf("%d", noticeid)}      
	} else {                 
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "notices信息添加失败"}             
	}         
	this.ServeJSON()         
	return                   
}             
              
// @router /notices/ajax/status [post]     
func (this *NoticesController) AjaxStatus() {             
	//权限检测               
	if !strings.Contains(this.GetSession("userPermission").(string), "order-edit") {     
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}     
		this.ServeJSON()     
		return               
	}         
	id, _ := this.GetInt64("id")            
	if id <= 0 {             
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择notices"}   
		this.ServeJSON()     
		return               
	}         
	status, _ := this.GetInt("status")      
	if status <= 0 || status >= 5 {         
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}               
		this.ServeJSON()     
		return               
	}         
              
	err := ChangeNoticesStatus(id, status)    
              
	if err == nil {          
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "notices状态更改成功"}             
	} else {                 
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "notices状态更改失败"}             
	}         
	this.ServeJSON()         
}             

