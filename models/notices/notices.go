package notices

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"opms/models"
)
type Notices struct {
    Noticeid int64
    Title string
    Content string
    Created int
    Status int
}
func (this *Notices) TableName() string {
    	return models.TableName("notices")
}
func init() {
    orm.RegisterModel(new(Notices))
}
func GetNotices(id int64) (Notices, error) { 
	var notices Notices                       
	var err error 
 
	o := orm.NewOrm()
	notices = Notices{Noticeid: id}
	err = o.Read(&notices)
  
	return notices, err
} 
//修改notices信息 
func UpdateNotices(id int64, updNotices Notices) error {
	o := orm.NewOrm() 
	_, err := o.Update(&updNotices ,"noticeid" ,"title" ,"content" ,"created" ,"status")
	return err   
} 

//删除notices
func RemoveNotices(id int64, updNotices Notices) error {             
	o := orm.NewOrm()  
	p, err := o.Raw("Delete from " + models.TableName("notices") + " where id=?").Prepare()   
	res, err := p.Exec(id)             
	if res == nil {     
		return err      
	}    
	p.Close()           
         
	//_, err := o.Delete(&updNotices, "noticeid")    
	return err          
}        
         
//新增notices
func AddNotices(updNotices Notices) error {           
	o := orm.NewOrm()   
         
	_, err := o.Insert(&updNotices)       
	return err          
} 
  
//notices列表
func ListNotices(condArr map[string]string, page int, offset int) (num int64, err error, notices []Notices) {
	orm.Debug = true    
	o := orm.NewOrm()   
	o.Using("default")
	qs := o.QueryTable(models.TableName("notices"))
	cond := orm.NewCondition()         
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
         
	var noticess []Notices          
	qs = qs.OrderBy("-noticesid")   
	num, err1 := qs.Limit(offset, start).All(&noticess)             
	return num, err1, notices        
}        
         
//统计数量
func CountNotices(condArr map[string]string) int64 { 
	o := orm.NewOrm()   
	qs := o.QueryTable(models.TableName("notices"))
	qs = qs.RelatedSel()
	cond := orm.NewCondition()         
	if condArr["status"] != "" {   
		cond = cond.And("status", condArr["status"])             
	}    
	num, _ := qs.SetCond(cond).Count() 
	return num          
}        
         
//改变notices状态          
func ChangeNoticesStatus(id int64, status int) error {
	o := orm.NewOrm()   
         
	notices := Notices{Noticeid: id}
	err := o.Read(&notices, "id")
	if nil != err { 
		return err 
	} else { 
		notices.Status = status 
		_, err := o.Update(&notices) 
		return err 
	}    
} 
