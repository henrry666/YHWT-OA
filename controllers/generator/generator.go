package main

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"opms/controllers"
	"opms/initial"
	"opms/utils"
	"os"
)

type GeneratorController struct {
	controllers.BaseController
}

type TableColumns struct {
	Column_name    string
	Data_type      string
	Column_comment string
	Column_key     string
}

func GenerateModelFile(tableName string, dbname string) string {
	var queryStr string
	var buffer bytes.Buffer
	initial.InitSql()
	buffer.WriteString("select column_name,data_type,column_comment,column_key from information_schema.columns ")
	buffer.WriteString("where table_name='")
	buffer.WriteString("pms_" + tableName)
	buffer.WriteString("' and table_schema='")
	buffer.WriteString(dbname)
	buffer.WriteString("'")
	queryStr = buffer.String()
	o := orm.NewOrm()

	var columns []TableColumns
	num, err := o.Raw(queryStr).QueryRows(&columns)
	if err != nil {
		return ""
	}
	if num == 0 {
		return ""
	}
	var prk string
	for i := 0; i < len(columns); i++ {
		if columns[i].Column_key == "PRI" {
			prk = utils.StringFirstToUpper(columns[i].Column_name)
		}
		fmt.Println(columns[i])
	}
	var typeName = utils.StringFirstToUpper(tableName)

	var fileContent string
	buffer.Reset()
	//生成package名和引用
	buffer.WriteString("package " + tableName + "\n")
	buffer.WriteString("import (\n")
	buffer.WriteString("    \"fmt\"\n")
	buffer.WriteString("    \"github.com/astaxie/beego\"\n")
	buffer.WriteString("    \"github.com/astaxie/beego/orm\"\n")
	buffer.WriteString("    \"opms/models\"\n")
	buffer.WriteString("    \"opms/utils\"\n")
	buffer.WriteString(")\n")

	//生成type结构
	buffer.WriteString("type " + typeName + " struct {\n")

	for i := 0; i < len(columns); i++ {
		buffer.WriteString("    " + utils.StringFirstToUpper(columns[i].Column_name) + " " + beego.AppConfig.String(columns[i].Data_type) + "\n")
		//fmt.Println(columns[i])
	}
	buffer.WriteString("}\n")
	//生成TableName函数
	buffer.WriteString("func (this *" + typeName + ") TableName() string {\n")
	buffer.WriteString("    	return models.TableName(\"" + tableName + "\")\n")
	buffer.WriteString("}\n")
	//生成init函数
	buffer.WriteString("func init() {\n")
	buffer.WriteString("    orm.RegisterModel(new(" + typeName + "))\n")
	buffer.WriteString("}\n")

	//生成Get函数
	buffer.WriteString("func Get" + typeName + "(id int64) (" + typeName + ", error) { \n")
	buffer.WriteString("	var " + tableName + " " + typeName + "                       \n")
	buffer.WriteString("	var err error \n")
	buffer.WriteString(" \n")
	buffer.WriteString("	o := orm.NewOrm()\n")
	buffer.WriteString("	" + tableName + " = " + typeName + "{" + prk + ": id} \n")
	buffer.WriteString("	err = o.Read(&" + tableName + ")\n ")
	buffer.WriteString(" \n")
	buffer.WriteString("	return " + tableName + ", err\n")
	buffer.WriteString("} \n")

	//////////////////////////////////////////////////
	buffer.WriteString("//修改"+tableName+"信息 \n")
	buffer.WriteString("func Update" + typeName + "(id int64, upd" + typeName + " " + typeName + ") error {\n")
	buffer.WriteString("	o := orm.NewOrm() \n")
	buffer.WriteString("	_, err := o.Update(&upd" + typeName )
	for i := 0; i < len(columns); i++ {
		buffer.WriteString( " ,\""+columns[i].Column_name +"\"")

	}
	buffer.WriteString(")\n")
	buffer.WriteString("	return err   \n")
	buffer.WriteString("} \n")
	buffer.WriteString("\n")

	buffer.WriteString("//删除"+tableName+"\n")
	buffer.WriteString("func Remove" + typeName + "(id int64, upd" + typeName + " " + typeName + ") error {             \n")
	buffer.WriteString("	o := orm.NewOrm()  \n")
	buffer.WriteString("	p, err := o.Raw(\"Delete from \" + models.TableName(\"" + tableName + "\") + \" where id=?\").Prepare()   \n")
	buffer.WriteString("	res, err := p.Exec(id)             \n")
	buffer.WriteString("	if res == nil {     \n")
	buffer.WriteString("		return err      \n")
	buffer.WriteString("	}    \n")
	buffer.WriteString("	p.Close()           \n")
	buffer.WriteString("         \n")
	buffer.WriteString("	//_, err := o.Delete(&upd" + typeName + ", \"" + prk + "\")    \n")
	buffer.WriteString("	return err          \n")
	buffer.WriteString("}        \n")
	buffer.WriteString("         \n")

	buffer.WriteString("//新增"+tableName+"\n")
	buffer.WriteString("func Add" + typeName + "(upd" + typeName + " " + typeName + ") error {           \n")
	buffer.WriteString("	o := orm.NewOrm()   \n")
	buffer.WriteString("         \n")
	buffer.WriteString("	_, err := o.Insert(&upd" + typeName + ")       \n")
	buffer.WriteString("	return err          \n")
	buffer.WriteString("} \n")
	buffer.WriteString("  \n")

	buffer.WriteString("//"+tableName+"列表\n")
	buffer.WriteString("func List" + typeName + "(condArr map[string]string, page int, offset int) (num int64, err error, " + tableName + " []" + typeName + ") {\n")
	buffer.WriteString("	orm.Debug = true    \n")
	buffer.WriteString("	o := orm.NewOrm()   \n")
	buffer.WriteString("	o.Using(\"default\")\n")
	buffer.WriteString("	qs := o.QueryTable(models.TableName(\"" + tableName + "\"))\n")
	buffer.WriteString("	cond := orm.NewCondition()         \n")
	buffer.WriteString("	if condArr[\"status\"] != \"\" {   \n")
	buffer.WriteString("		cond = cond.And(\"status\", condArr[\"status\"]) \n")
	buffer.WriteString("	} \n")
	buffer.WriteString("	qs = qs.SetCond(cond)\n")
	buffer.WriteString("	if page < 1 { \n")
	buffer.WriteString("		page = 1  \n")
	buffer.WriteString("	}    \n")
	buffer.WriteString("	if offset < 1 {\n")
	buffer.WriteString("		offset, _ = beego.AppConfig.Int(\"pageoffset\")\n")
	buffer.WriteString("	} \n")
	buffer.WriteString("	start := (page - 1) * offset       \n")
	buffer.WriteString("	qs = qs.RelatedSel()\n")
	buffer.WriteString("         \n")
	buffer.WriteString("	var " + tableName + "s []" + typeName + "          \n")
	buffer.WriteString("	qs = qs.OrderBy(\"-" + tableName + "id\")   \n")
	buffer.WriteString("	num, err1 := qs.Limit(offset, start).All(&" + tableName + "s)             \n")
	buffer.WriteString("	return num, err1, " + tableName + "        \n")
	buffer.WriteString("}        \n")
	buffer.WriteString("         \n")

	buffer.WriteString("//统计数量\n")
	buffer.WriteString("func Count" + typeName + "(condArr map[string]string) int64 { \n")
	buffer.WriteString("	o := orm.NewOrm()   \n")
	buffer.WriteString("	qs := o.QueryTable(models.TableName(\"" + tableName + "\"))\n")
	buffer.WriteString("	qs = qs.RelatedSel()\n")
	buffer.WriteString("	cond := orm.NewCondition()         \n")
	buffer.WriteString("	if condArr[\"status\"] != \"\" {   \n")
	buffer.WriteString("		cond = cond.And(\"status\", condArr[\"status\"])             \n")
	buffer.WriteString("	}    \n")
	buffer.WriteString("	num, _ := qs.SetCond(cond).Count() \n")
	buffer.WriteString("	return num          \n")
	buffer.WriteString("}        \n")
	buffer.WriteString("         \n")

	buffer.WriteString("//改变"+tableName+"状态          \n")
	buffer.WriteString("func Change" + typeName + "Status(id int64, status int) error {\n")
	buffer.WriteString("	o := orm.NewOrm()   \n")
	buffer.WriteString("         \n")
	buffer.WriteString("	" + tableName + " := " + typeName + "{" + prk + ": id} \n")
	buffer.WriteString("	err := o.Read(&" + tableName + ", \"id\")\n")
	buffer.WriteString("	if nil != err { \n")
	buffer.WriteString("		return err \n")
	buffer.WriteString("	} else { \n")
	buffer.WriteString("		" + tableName + ".Status = status \n")
	buffer.WriteString("		_, err := o.Update(&" + tableName + ") \n")
	buffer.WriteString("		return err \n")
	buffer.WriteString("	}    \n")
	buffer.WriteString("} \n")

	//////////////////////////////////////////////////
	fileContent = buffer.String()

	return fileContent
}

func GenerateControllerFile(tableName string, dbname string) string {
	var queryStr string
	var buffer bytes.Buffer
	//initial.InitSql()
	buffer.WriteString("select column_name,data_type,column_comment,column_key from information_schema.columns ")
	buffer.WriteString("where table_name='")
	buffer.WriteString("pms_" + tableName)
	buffer.WriteString("' and table_schema='")
	buffer.WriteString(dbname)
	buffer.WriteString("'")
	queryStr = buffer.String()
	o := orm.NewOrm()

	var columns []TableColumns
	num, err := o.Raw(queryStr).QueryRows(&columns)
	if err != nil {
		return ""
	}
	if num == 0 {
		return ""
	}
	var prk string
	for i := 0; i < len(columns); i++ {
		if columns[i].Column_key == "PRI" {
			prk = columns[i].Column_name
		}
		fmt.Println(columns[i])
	}
	var typeName = utils.StringFirstToUpper(tableName)

	var fileContent string
	buffer.Reset()
	//生成package名和引用
	buffer.WriteString("package " + tableName + "\n")
	buffer.WriteString("import (\n")
	buffer.WriteString("    \"fmt\"\n")
	buffer.WriteString("    \"github.com/astaxie/beego\"\n")
	buffer.WriteString("    \"github.com/astaxie/beego/utils/pagination\"\n")
	buffer.WriteString("    .\"opms/models/"+tableName+"\"\n")
	buffer.WriteString("    \"strings\"\n")
	buffer.WriteString(")\n")

	//生成type结构
	buffer.WriteString("type " + typeName + "Controller struct {\n")
	buffer.WriteString("    controllers.BaseController\n")
	buffer.WriteString("}\n")

	buffer.WriteString("// @router /"+tableName+"/manage [get]           \n")
	buffer.WriteString("func (this *"+typeName+"Controller) ListAll"+typeName+"s() {       \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-manage\") {\n")
	buffer.WriteString("		this.Redirect(\"/my/task\", 302)      \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("		//this.Abort(\"401\")                 \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	page, err := this.GetInt(\"p\")           \n")
	buffer.WriteString("	status := this.GetString(\"status\")      \n")
	buffer.WriteString("	keywords := this.GetString(\"keywords\")                 \n")
	buffer.WriteString("	if err != nil {          \n")
	buffer.WriteString("		page = 1             \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	offset, err1 := beego.AppConfig.Int(\"pageoffset\")      \n")
	buffer.WriteString("	if err1 != nil {         \n")
	buffer.WriteString("		offset = 15          \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	condArr := make(map[string]string)      \n")
	buffer.WriteString("	condArr[\"status\"] = status              \n")
	buffer.WriteString("	condArr[\"keywords\"] = keywords          \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	count"+typeName+" := Count"+typeName+"(condArr)                \n")
	buffer.WriteString("	paginator := pagination.SetPaginator(this.Ctx, offset, count"+typeName+") \n")
	buffer.WriteString("	_, _, "+tableName+" := List"+typeName+"(condArr, page, offset)                \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	this.Data[\"paginator\"] = paginator      \n")
	buffer.WriteString("	this.Data[\"condArr\"] = condArr          \n")
	buffer.WriteString("	this.Data[\""+tableName+"\"] = "+tableName+"      \n")
	buffer.WriteString("	this.Data[\"count"+typeName+"\"] = count"+typeName+"             \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	this.TplName = \"orders/"+tableName+".tpl\"    \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")
	buffer.WriteString("// @router /"+tableName+"/showOrder [get]        \n")
	buffer.WriteString("func (this *"+typeName+"Controller) ShowOrder() {              \n")
	buffer.WriteString("              \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")

	buffer.WriteString("// @router /"+tableName+"/edit/:id [get]         \n")
	buffer.WriteString("func (this *"+typeName+"Controller) Edit"+typeName+"Get() {        \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-edit\") {  \n")
	buffer.WriteString("		this.Abort(\"401\")    \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	idstr := this.Ctx.Input.Param(\":id\")    \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	id, err := strconv.Atoi(idstr)          \n")
	buffer.WriteString("	if err != nil {          \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	"+tableName+", err := Get"+typeName+"(int64(id))                \n")
	buffer.WriteString("	if err != nil {          \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")

	buffer.WriteString("	this.Data[\""+tableName+"\"] = "+tableName+"        \n")
	buffer.WriteString("	this.TplName = \"orders/"+tableName+"-edit.tpl\"              \n")
	buffer.WriteString("}             \n")

	buffer.WriteString("// @router /"+tableName+"/edit [post]            \n")
	buffer.WriteString("func (this *"+typeName+"Controller) Edit"+typeName+"Post() {       \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-edit\") {  \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"无权设置\"}     \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")

	for i := 0; i < len(columns); i++ {
		if beego.AppConfig.String(columns[i].Data_type)=="string"{
			buffer.WriteString("    "+columns[i].Column_name+":= this.GetString(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if \"\" == "+columns[i].Column_name+" {  \n")
			buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"请填写"+columns[i].Column_comment+"\"}               \n")
			buffer.WriteString("		this.ServeJSON()     \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}
		if beego.AppConfig.String(columns[i].Data_type)=="int64" {
			buffer.WriteString("    "+columns[i].Column_name+",error:= this.GetInt64(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if error != nil {  \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}
		if beego.AppConfig.String(columns[i].Data_type)=="int" {
			buffer.WriteString("    "+columns[i].Column_name+", error:= this.GetInt(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if error != nil {  \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}

	}

	buffer.WriteString("	var "+tableName+" "+typeName+"                  \n")

	for i := 0; i < len(columns); i++ {
		buffer.WriteString("	"+tableName+"." + utils.StringFirstToUpper(columns[i].Column_name) + " = "+columns[i].Column_name+"\n")
	}
	buffer.WriteString("              \n")
	buffer.WriteString("var	err = Update"+typeName+"("+prk+","+tableName+")              \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	if err == nil {          \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 1, \"message\": \""+tableName+"信息修改成功\"}             \n")
	buffer.WriteString("	} else {                 \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \""+tableName+"信息修改失败\"}             \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	this.ServeJSON()         \n")
	buffer.WriteString("	return                   \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")

	buffer.WriteString("// @router /"+tableName+"/remove/:id [get]       \n")
	buffer.WriteString("func (this *"+typeName+"Controller) RemoveOrder() {            \n")
	buffer.WriteString("              \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")

	buffer.WriteString("// @router /"+tableName+"/add [get]              \n")
	buffer.WriteString("func (this *"+typeName+"Controller) Add"+typeName+"Get() {         \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-add\") {   \n")
	buffer.WriteString("		this.Abort(\"401\")    \n")
	buffer.WriteString("	}         \n")

	buffer.WriteString("	var "+tableName+" "+typeName+"                  \n")

	buffer.WriteString("	this.Data[\""+tableName+"\"] = "+tableName+"        \n")
	buffer.WriteString("	this.TplName = \""+tableName+"/"+tableName+"-form.tpl\"              \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")

	buffer.WriteString("// @router /"+tableName+"/add [post]             \n")
	buffer.WriteString("func (this *"+typeName+"Controller) Add"+typeName+"Post() {        \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-add\") {   \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"无权设置\"}     \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	for i := 0; i < len(columns); i++ {
		if beego.AppConfig.String(columns[i].Data_type)=="string"{
			buffer.WriteString("    "+columns[i].Column_name+":= this.GetString(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if \"\" == "+columns[i].Column_name+" {  \n")
			buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"请填写"+columns[i].Column_comment+"\"}               \n")
			buffer.WriteString("		this.ServeJSON()     \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
	}
		if beego.AppConfig.String(columns[i].Data_type)=="int64" {
			buffer.WriteString("    "+columns[i].Column_name+", error:= this.GetInt64(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if error != nil {  \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}
		if beego.AppConfig.String(columns[i].Data_type)=="int" {
			buffer.WriteString("    "+columns[i].Column_name+", error:= this.GetInt(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if error != nil {  \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}


	}

	buffer.WriteString("	var "+tableName+" "+typeName+"                  \n")

	for i := 0; i < len(columns); i++ {
		buffer.WriteString("	"+tableName+"." + utils.StringFirstToUpper(columns[i].Column_name) + " = "+columns[i].Column_name+"\n")
	}

	buffer.WriteString("              \n")
	buffer.WriteString("var	err = Add"+typeName+"("+tableName+")             \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	if err == nil {          \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 1, \"message\":  \n")
	buffer.WriteString("			\""+tableName+"信息添加成功\", \"id\": fmt.Sprintf(\"%d\", "+prk+")}      \n")
	buffer.WriteString("	} else {                 \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \""+tableName+"信息添加失败\"}             \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	this.ServeJSON()         \n")
	buffer.WriteString("	return                   \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")


	buffer.WriteString("// @router /"+tableName+"/ajax/status [post]     \n")
	buffer.WriteString("func (this *"+typeName+"Controller) AjaxStatus() {             \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \"order-edit\") {     \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"无权设置\"}     \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	id, _ := this.GetInt64(\"id\")            \n")
	buffer.WriteString("	if id <= 0 {             \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"请选择"+tableName+"\"}   \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	status, _ := this.GetInt(\"status\")      \n")
	buffer.WriteString("	if status <= 0 || status >= 5 {         \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"请选择操作状态\"}               \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	err := Change"+typeName+"Status(id, status)    \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	if err == nil {          \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 1, \"message\": \""+tableName+"状态更改成功\"}             \n")
	buffer.WriteString("	} else {                 \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \""+tableName+"状态更改失败\"}             \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	this.ServeJSON()         \n")
	buffer.WriteString("}             \n")


	buffer.WriteString("\n")

	//////////////////////////////////////////////////
	fileContent = buffer.String()

	return fileContent
}

//写view页面
func GenerateViewFile(tableName string, dbname string) string {
	var queryStr string
	var buffer bytes.Buffer
	//initial.InitSql()
	buffer.WriteString("select column_name,data_type,column_comment,column_key from information_schema.columns ")
	buffer.WriteString("where table_name='")
	buffer.WriteString("pms_" + tableName)
	buffer.WriteString("' and table_schema='")
	buffer.WriteString(dbname)
	buffer.WriteString("'")
	queryStr = buffer.String()
	o := orm.NewOrm()

	var columns []TableColumns
	num, err := o.Raw(queryStr).QueryRows(&columns)
	if err != nil {
		return ""
	}
	if num == 0 {
		return ""
	}
	var prk string
	for i := 0; i < len(columns); i++ {
		if columns[i].Column_key == "PRI" {
			prk = columns[i].Column_name
		}
		fmt.Println(columns[i])
	}
	var typeName = utils.StringFirstToUpper(tableName)

	var fileContent string
	buffer.Reset()
	//生成package名和引用
	buffer.WriteString("package " + tableName + "\n")
	buffer.WriteString("import (\n")
	buffer.WriteString("    \"fmt\"\n")
	buffer.WriteString("    \"github.com/astaxie/beego\"\n")
	buffer.WriteString("    \"github.com/astaxie/beego/utils/pagination\"\n")
	buffer.WriteString("    .\"opms/models/"+tableName+"\"\n")
	buffer.WriteString("    \"strings\"\n")
	buffer.WriteString(")\n")

	//生成type结构
	buffer.WriteString("type " + typeName + "Controller struct {\n")
	buffer.WriteString("    controllers.BaseController\n")
	buffer.WriteString("}\n")

	buffer.WriteString("// @router /"+tableName+"/manage [get]           \n")
	buffer.WriteString("func (this *"+typeName+"Controller) ListAll"+typeName+"s() {       \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-manage\") {\n")
	buffer.WriteString("		this.Redirect(\"/my/task\", 302)      \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("		//this.Abort(\"401\")                 \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	page, err := this.GetInt(\"p\")           \n")
	buffer.WriteString("	status := this.GetString(\"status\")      \n")
	buffer.WriteString("	keywords := this.GetString(\"keywords\")                 \n")
	buffer.WriteString("	if err != nil {          \n")
	buffer.WriteString("		page = 1             \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	offset, err1 := beego.AppConfig.Int(\"pageoffset\")      \n")
	buffer.WriteString("	if err1 != nil {         \n")
	buffer.WriteString("		offset = 15          \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	condArr := make(map[string]string)      \n")
	buffer.WriteString("	condArr[\"status\"] = status              \n")
	buffer.WriteString("	condArr[\"keywords\"] = keywords          \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	count"+typeName+" := Count"+typeName+"(condArr)                \n")
	buffer.WriteString("	paginator := pagination.SetPaginator(this.Ctx, offset, count"+typeName+") \n")
	buffer.WriteString("	_, _, "+tableName+" := List"+typeName+"(condArr, page, offset)                \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	this.Data[\"paginator\"] = paginator      \n")
	buffer.WriteString("	this.Data[\"condArr\"] = condArr          \n")
	buffer.WriteString("	this.Data[\""+tableName+"\"] = "+tableName+"      \n")
	buffer.WriteString("	this.Data[\"count"+typeName+"\"] = count"+typeName+"             \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	this.TplName = \"orders/"+tableName+".tpl\"    \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")
	buffer.WriteString("// @router /"+tableName+"/showOrder [get]        \n")
	buffer.WriteString("func (this *"+typeName+"Controller) ShowOrder() {              \n")
	buffer.WriteString("              \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")

	buffer.WriteString("// @router /"+tableName+"/edit/:id [get]         \n")
	buffer.WriteString("func (this *"+typeName+"Controller) Edit"+typeName+"Get() {        \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-edit\") {  \n")
	buffer.WriteString("		this.Abort(\"401\")    \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	idstr := this.Ctx.Input.Param(\":id\")    \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	id, err := strconv.Atoi(idstr)          \n")
	buffer.WriteString("	if err != nil {          \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	"+tableName+", err := Get"+typeName+"(int64(id))                \n")
	buffer.WriteString("	if err != nil {          \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")

	buffer.WriteString("	this.Data[\""+tableName+"\"] = "+tableName+"        \n")
	buffer.WriteString("	this.TplName = \"orders/"+tableName+"-edit.tpl\"              \n")
	buffer.WriteString("}             \n")

	buffer.WriteString("// @router /"+tableName+"/edit [post]            \n")
	buffer.WriteString("func (this *"+typeName+"Controller) Edit"+typeName+"Post() {       \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-edit\") {  \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"无权设置\"}     \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")

	for i := 0; i < len(columns); i++ {
		if beego.AppConfig.String(columns[i].Data_type)=="string"{
			buffer.WriteString("    "+columns[i].Column_name+":= this.GetString(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if \"\" == "+columns[i].Column_name+" {  \n")
			buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"请填写"+columns[i].Column_comment+"\"}               \n")
			buffer.WriteString("		this.ServeJSON()     \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}
		if beego.AppConfig.String(columns[i].Data_type)=="int64" {
			buffer.WriteString("    "+columns[i].Column_name+",error:= this.GetInt64(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if error != nil {  \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}
		if beego.AppConfig.String(columns[i].Data_type)=="int" {
			buffer.WriteString("    "+columns[i].Column_name+", error:= this.GetInt(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if error != nil {  \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}

	}

	buffer.WriteString("	var "+tableName+" "+typeName+"                  \n")

	for i := 0; i < len(columns); i++ {
		buffer.WriteString("	"+tableName+"." + utils.StringFirstToUpper(columns[i].Column_name) + " = "+columns[i].Column_name+"\n")
	}
	buffer.WriteString("              \n")
	buffer.WriteString("var	err = Update"+typeName+"("+prk+","+tableName+")              \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	if err == nil {          \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 1, \"message\": \""+tableName+"信息修改成功\"}             \n")
	buffer.WriteString("	} else {                 \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \""+tableName+"信息修改失败\"}             \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	this.ServeJSON()         \n")
	buffer.WriteString("	return                   \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")

	buffer.WriteString("// @router /"+tableName+"/remove/:id [get]       \n")
	buffer.WriteString("func (this *"+typeName+"Controller) RemoveOrder() {            \n")
	buffer.WriteString("              \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")

	buffer.WriteString("// @router /"+tableName+"/add [get]              \n")
	buffer.WriteString("func (this *"+typeName+"Controller) Add"+typeName+"Get() {         \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-add\") {   \n")
	buffer.WriteString("		this.Abort(\"401\")    \n")
	buffer.WriteString("	}         \n")

	buffer.WriteString("	var "+tableName+" "+typeName+"                  \n")

	buffer.WriteString("	this.Data[\""+tableName+"\"] = "+tableName+"        \n")
	buffer.WriteString("	this.TplName = \""+tableName+"/"+tableName+"-form.tpl\"              \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")

	buffer.WriteString("// @router /"+tableName+"/add [post]             \n")
	buffer.WriteString("func (this *"+typeName+"Controller) Add"+typeName+"Post() {        \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \""+tableName+"-add\") {   \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"无权设置\"}     \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	for i := 0; i < len(columns); i++ {
		if beego.AppConfig.String(columns[i].Data_type)=="string"{
			buffer.WriteString("    "+columns[i].Column_name+":= this.GetString(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if \"\" == "+columns[i].Column_name+" {  \n")
			buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"请填写"+columns[i].Column_comment+"\"}               \n")
			buffer.WriteString("		this.ServeJSON()     \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}
		if beego.AppConfig.String(columns[i].Data_type)=="int64" {
			buffer.WriteString("    "+columns[i].Column_name+", error:= this.GetInt64(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if error != nil {  \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}
		if beego.AppConfig.String(columns[i].Data_type)=="int" {
			buffer.WriteString("    "+columns[i].Column_name+", error:= this.GetInt(\""+columns[i].Column_name+"\") \n")
			buffer.WriteString("	if error != nil {  \n")
			buffer.WriteString("		return               \n")
			buffer.WriteString("	}         \n")
		}


	}

	buffer.WriteString("	var "+tableName+" "+typeName+"                  \n")

	for i := 0; i < len(columns); i++ {
		buffer.WriteString("	"+tableName+"." + utils.StringFirstToUpper(columns[i].Column_name) + " = "+columns[i].Column_name+"\n")
	}

	buffer.WriteString("              \n")
	buffer.WriteString("var	err = Add"+typeName+"("+tableName+")             \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	if err == nil {          \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 1, \"message\":  \n")
	buffer.WriteString("			\""+tableName+"信息添加成功\", \"id\": fmt.Sprintf(\"%d\", "+prk+")}      \n")
	buffer.WriteString("	} else {                 \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \""+tableName+"信息添加失败\"}             \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	this.ServeJSON()         \n")
	buffer.WriteString("	return                   \n")
	buffer.WriteString("}             \n")
	buffer.WriteString("              \n")


	buffer.WriteString("// @router /"+tableName+"/ajax/status [post]     \n")
	buffer.WriteString("func (this *"+typeName+"Controller) AjaxStatus() {             \n")
	buffer.WriteString("	//权限检测               \n")
	buffer.WriteString("	if !strings.Contains(this.GetSession(\"userPermission\").(string), \"order-edit\") {     \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"无权设置\"}     \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	id, _ := this.GetInt64(\"id\")            \n")
	buffer.WriteString("	if id <= 0 {             \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"请选择"+tableName+"\"}   \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	status, _ := this.GetInt(\"status\")      \n")
	buffer.WriteString("	if status <= 0 || status >= 5 {         \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \"请选择操作状态\"}               \n")
	buffer.WriteString("		this.ServeJSON()     \n")
	buffer.WriteString("		return               \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	err := Change"+typeName+"Status(id, status)    \n")
	buffer.WriteString("              \n")
	buffer.WriteString("	if err == nil {          \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 1, \"message\": \""+tableName+"状态更改成功\"}             \n")
	buffer.WriteString("	} else {                 \n")
	buffer.WriteString("		this.Data[\"json\"] = map[string]interface{}{\"code\": 0, \"message\": \""+tableName+"状态更改失败\"}             \n")
	buffer.WriteString("	}         \n")
	buffer.WriteString("	this.ServeJSON()         \n")
	buffer.WriteString("}             \n")


	buffer.WriteString("\n")

	//////////////////////////////////////////////////
	fileContent = buffer.String()

	return fileContent
}

func WriteRouterFile(tableName string){
	fileName := "./routers/router.go"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	defer file.Close()

	fileString := "\n    beego.Include(&"+tableName+"."+utils.StringFirstToUpper(tableName)+"Controller{}) \n}"
	file.Seek(-2, 2)    // 最后增加
	file.WriteString(fileString)
}



func WriteWithIoutil(name, content string) {

	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入文件成功:", content)
	}
}

func main() {
	var modelfile = GenerateModelFile("notices", "aiopms")
	os.Mkdir("d:\\temp\\opms\\models\\notices",0777)
	WriteWithIoutil("d:\\temp\\opms\\models\\notices\\notices.go", modelfile)

	var controllerfile = GenerateControllerFile("notices","aiopms")
	os.Mkdir("d:\\temp\\opms\\controller\\notices",0777)
	WriteWithIoutil("d:\\temp\\opms\\controller\\notices\\notices.go", controllerfile)

	WriteRouterFile("notices")

}
