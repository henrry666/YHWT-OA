<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/js/bootstrap-datepicker/css/datepicker-custom.css" rel="stylesheet" />
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 客户管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/customer/manage">客户管理</a> </li>
        <li class="active"> 客户 </li>
      </ul>
      <div class="pull-right"><a href="/customer/add" class="btn btn-success">添加新客户</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="customer-form" action="/customer/add" method="post">
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>客户名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="name" value="{{.customer.Name}}" class="form-control" placeholder="请填写名称">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>客户代码</label>
                  <div class="col-sm-10">
                    <input type="text" name="code" value="{{.customer.Code}}" class="form-control" placeholder="请填写客户代码">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">开票信息</label>
                  <div class="col-sm-10">
                    <textarea name="detail" placeholder="请填写开票信息" style="height:300px;" class="form-control">{{.customer.Detail}}</textarea>
                  </div>
                </div>
				<div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户负责人</label>
                  <div class="col-sm-10">
                    <select name="userid" class="form-control">
                      <option>请选择产品负责人</option>
					{{range .teams}}
                      <option value="{{.Id}}" {{if eq .Id $.customer.Userid}}selected{{end}}>{{getRealname .Id}}</option>
                    {{end}}
                    </select>
                  </div>
                </div>

                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="customerid" value="{{.customer.Customerid}}">
                    <input type="hidden" name="status" value="{{.customer.Status}}">
                    <button type="submit" class="btn btn-primary">提 交</button>
                  </div>
                </div>
              </form>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
<div aria-hidden="true" aria-labelledby="customerModalLabel" role="dialog" tabindex="-1" id="customerModal" class="modal fade">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
            <h4 class="modal-title">新建客户成功</h4>
          </div>
          <div class="modal-footer">
            <a href="/customer/manage" class="btn btn-primary">去设置管理</a>
          </div>
        </div>
      </div>
    </div>
{{template "inc/foot.tpl" .}}

<script src="/static/keditor/kindeditor-min.js"></script>
<script>
$(function(){
	var editor = KindEditor.create('textarea[name="detail"]', {
	    uploadJson: "/kindeditor/upload",
	    allowFileManager: true,
	    filterMode : false,
	    afterBlur: function(){this.sync();}
	});
	

})
</script>
</body>
</html>
