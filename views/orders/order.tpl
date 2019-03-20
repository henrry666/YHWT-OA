<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->
      <form class="searchform" action="/order/manage" method="get">
        <select name="status" class="form-control">
          <option value="">订单状态</option>
          <option value="1" {{if eq "1" .condArr.status}}selected{{end}}>挂起</option>
          <option value="2" {{if eq "2" .condArr.status}}selected{{end}}>延期</option>
          <option value="3" {{if eq "3" .condArr.status}}selected{{end}}>进行</option>
          <option value="4" {{if eq "4" .condArr.status}}selected{{end}}>结束</option>
        </select>
        <input type="text" class="form-control" name="keywords" placeholder="请输入名称" value="{{.condArr.keywords}}"/>
        <button type="submit" class="btn btn-primary">搜索</button>
      </form>
      <!--search end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 订单管理 {{template "orders/nav.tpl" .}}</h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/order/manage">订单管理</a> </li>
        <li class="active"> 订单 </li>
      </ul>
      <div class="pull-right"><a href="/order/add" class="btn btn-success">+新订单</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 订单 / 总数：{{.countOrder}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div id="orderpanel" class="panel-body">
              <section id="unseen">
                <form id="order-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>订单编号</th>
                        <th>订单来源</th>
						<th>任务承接单位</th>
                        <th>客户负责人</th>
                        <th>生产负责人</th>
                        <th>状态</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    
                    {{range $k,$v := .orders}}
                    <tr>
                      <td><a href="/order/{{$v.Id}}">{{$v.Orderno}}</a></td>
                      <td>{{$v.Source}}</td>
					  <td>{{$v.Task_unit}}</td>
                      <td>{{$v.Customer_responser}}</td>
                      <td>{{$v.Product_responser}}</td>
                      <td>{{if eq 1 $v.Status}}挂起{{else if eq 2 $v.Status}}延期{{else if eq 3 $v.Status}}进行{{else if eq 4 $v.Status}}结束{{end}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                          <ul class="dropdown-menu">
                            <li><a href="/order/edit/{{$v.Id}}">编辑</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-order-single" data-id="{{$v.Id}}" data-status="1">挂起</a></li>
                            <li><a href="javascript:;" class="js-order-single" data-id="{{$v.Id}}" data-status="2">延期</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-order-single" data-id="{{$v.Id}}" data-status="3">进行</a></li>
                            <li><a href="javascript:;" class="js-order-single" data-id="{{$v.Id}}" data-status="4">结束</a></li>
                          </ul>
                        </div></td>
                    </tr>
                    {{end}}
                    </tbody>
                    
                  </table>
                </form>
                {{template "inc/page.tpl" .}}
				 </section>
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
{{template "inc/foot.tpl" .}}
</body>
</html>
