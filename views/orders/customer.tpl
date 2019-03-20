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
          <option value="">客户状态</option>
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
        <li> <a href="/order/manage">客户管理</a> </li>
        <li class="active"> 客户 </li>
      </ul>
      <div class="pull-right"><a href="/customer/add" class="btn btn-success">+新客户</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 客户 / 总数：{{.countOrder}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div id="orderpanel" class="panel-body">
              <section id="unseen">
                <form id="order-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>客户代码</th>
                        <th>客户名称</th>
						<th>开票信息</th>
                        <th>客户负责人</th>
                        <th>状态</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    
                    {{range $k,$v := .customers}}
                    <tr>
                      <td><a href="/customer/{{$v.Customerid}}">{{$v.Code}}</a></td>
                      <td>{{$v.Name}}</td>
					  <td><div>{{$v.Detail}}</div></td>
                      <td><a href="/user/show/{{$v.Userid}}">{{getRealname $v.Userid}}</a></td>

                      <td>{{if eq 1 $v.Status}}正常{{else if eq 2 $v.Status}}屏蔽{{end}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                          <ul class="dropdown-menu">
                            <li><a href="/customer/edit/{{$v.Customerid}}">编辑</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-order-single" data-id="{{$v.Customerid}}" data-status="1">正常</a></li>
                            <li><a href="javascript:;" class="js-order-single" data-id="{{$v.Customerid}}" data-status="2">屏蔽</a></li>

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
