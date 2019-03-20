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
      <h3> 订单管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/order/manage">订单管理</a> </li>
        <li class="active"> 订单 </li>
      </ul>
      <div class="pull-right"><a href="/order/add" class="btn btn-success">添加新订单</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="order-form">
                <div class="form-group">
                    <label class="col-sm-2 control-label">订单编号：</label>
                    <div class="col-sm-4">
                        <input id="orderno" name="orderno" class="form-control" type="text">
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label">订单来源：</label>
                    <div class="col-sm-4">
                        <select id="source" name="source" class="form-control">
                          <option>请选择订单来源</option>
                        {{range .customers}}
                          <option value="{{.Customerid}}" {{if eq .Customerid $.order.Source}}selected{{end}}>{{getCustomername .Customerid}}</option>
                        {{end}}
                        </select>
                    </div>

                    <label class="col-sm-2 col-sm-2 control-label">任务承接单位：</label>
                    <div class="col-sm-4">
                        <select id="task_unit" name="task_unit" class="form-control">
                          <option>请选择任务承接单位</option>
                        {{range .departs}}
                          <option value="{{.Id}}" {{if eq .Id $.order.Source}}selected{{end}}>{{getDepartName .Id}}</option>
                        {{end}}
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">客户负责人：</label>
                    <div class="col-sm-4">
                        <select id="task_unit" name="task_unit" class="form-control">
                          <option>请选择任务承接单位</option>
                        {{range .departs}}
                          <option value="{{.Id}}" {{if eq .Id $.order.Source}}selected{{end}}>{{getDepartName .Id}}</option>
                        {{end}}
                        </select>
                    </div>

                    <label class="col-sm-2 control-label">生产负责人：</label>
                    <div class="col-sm-4">
                        <select id="task_unit" name="task_unit" class="form-control">
                          <option>请选择任务承接单位</option>
                        {{range .userList}}
                          <option value="{{.Id}}" {{if eq .Id $.order.Source}}selected{{end}}>{{getRealname .Id}}</option>
                        {{end}}
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">任务名称：</label>
                    <div class="col-sm-4">
                        <input id="taskName" name="taskName" class="form-control" type="text">
                    </div>

                    <label class="col-sm-2 control-label">任务性质：</label>
                    <div class="col-sm-4">
                        <select id="taskPriviledge" name="taskPriviledge" class="selectpicker">
                            <option value="1" selected>请选择任务性质</option>
                            <option value="1">一般</option>
                            <option value="2">加急</option>
                            <option value="3">特急</option>
                        </select>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">是否来料加工：</label>
                    <div class="col-sm-4">
                        <label class="radio-inline"> <input type="radio" name="isLailiaojiagong" value="1" /> 是</label>
                        <label class="radio-inline"> <input type="radio" name="isLailiaojiagong" value="0" /> 否</label>
                    </div>
                    <label class="col-sm-2 control-label">是否有图纸：</label>
                    <div class="col-sm-1">
                        <label class="radio-inline"> <input type="radio" name="hasPaper" value="1" /> 是</label>
                        <label class="radio-inline"> <input type="radio" name="hasPaper" value="0" /> 否</label>
                    </div>
                    <label class="col-sm-2 control-label">图纸数量：</label>
                    <div class="col-sm-1">
                        <input id="paperCount" name="paperCount" class="form-control" type="text">
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">要求完成时间：</label>
                    <div class="col-sm-4">
                        <div class="input-group date" id='date1'>
                            <input id="requiedDate" name="requiedDate" placeholder="2019-01-01" type="text" class="form-control">
                            <span class="input-group-addon">
                                <i class="glyphicon glyphicon-calendar"></i>
                            </span>
                        </div>
                    </div>

                    <label class="col-sm-2 control-label">实际完成时间：</label>
                    <div class="col-sm-4">
                        <div class="input-group date" id='date2'>
                            <input id="actualDate" name="actualDate" placeholder="2019-01-01"  type="text" class="form-control">
                            <span class="input-group-addon">
                                <i class="glyphicon glyphicon-calendar"></i>
                            </span>
                        </div>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">是否有外部客诉：</label>
                    <div class="col-sm-4">
                        <label class="radio-inline"> <input type="radio" name="hasOuterKesu" value="1" /> 是</label>
                        <label class="radio-inline"> <input type="radio" name="hasOuterKesu" value="0" /> 否</label>
                    </div>
                    <label class="col-sm-2 control-label">是否有内部客诉：</label>
                    <div class="col-sm-4">
                        <label class="radio-inline"> <input type="radio" name="hasInnerKesu" value="1" /> 是</label>
                        <label class="radio-inline"> <input type="radio" name="hasInnerKesu" value="0" /> 否</label>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">任务内容描述：</label>
                    <div class="col-sm-10">
                        <textarea id="taskDescription" style="width:100%;" name="taskDescription" rows="8"></textarea>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">制造中心审批意见：</label>
                    <div class="col-sm-10">
                        <textarea id="mcenterOpinion" style="width:100%;" name="mcenterOpinion" rows="4" ></textarea>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">审批签字：</label>
                    <div class="col-sm-4">
                        <input id="approveSign" name="approveSign" class="form-control" type="text">
                    </div>

                    <label class="col-sm-2 control-label"> 派工单下达时间：</label>
                    <div class="col-sm-4">
                        <div class="input-group date" id='date3'>
                            <input id="orderArrivalTime" name="orderArrivalTime" placeholder="2019-01-01 13:15:00"  type="text" class="form-control">
                            <span class="input-group-addon">
                                <i class="glyphicon glyphicon-calendar"></i>
                            </span>
                        </div>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label">特别说明：</label>
                    <div class="col-sm-10">
                        <textarea id="specialIllustrate" style="width:100%;" name="specialIllustrate" rows="4" ></textarea>
                    </div>
                </div>
                <div class="ibox-content">
                    <label class="col-sm-2 control-label">请选择图纸文件上传：</label>
                    <div class="file-manager">
                        <button type="button" class="layui-btn" id="test1" >
                            <i class="fa fa-cloud"></i>上传图纸文件
                        </button>
                        <div class="hr-line-dashed"></div>
                        <ul class="folder-list" style="padding: 0">

                        </ul>
                    </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 control-label"></label>
                    <div class="col-sm-12">
                        <div class="file-box" >
                            <div class="file">

                            </div>
                        </div>
                        <div id="incomeNum"></div>
                    </div>
                </div>

                <div class="form-group">
                    <div class="col-sm-8 col-sm-offset-3">
                        <button type="button" onclick="save()" class="btn btn-primary">提交</button>
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
<div aria-hidden="true" aria-labelledby="orderModalLabel" role="dialog" tabindex="-1" id="orderModal" class="modal fade">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
            <h4 class="modal-title">新建订单成功，请先按订单流程设置</h4>
          </div>
          <div class="modal-body">
            
            
            
          </div>
          <div class="modal-footer">
            <a href="/order/manage" class="btn btn-primary">去设置管理</a>
          </div>
        </div>
      </div>
    </div>
{{template "inc/foot.tpl" .}}
<script src="/static/js/bootstrap-datepicker/js/bootstrap-datepicker.js"></script>
<script src="/static/keditor/kindeditor-min.js"></script>
<script>
$(function(){
	var editor = KindEditor.create('textarea[name="desc"]', {
	    uploadJson: "/kindeditor/upload",
	    allowFileManager: true,
	    filterMode : false,
	    afterBlur: function(){this.sync();}
	});
	
	var nowTemp = new Date();
    var now = new Date(nowTemp.getFullYear(), nowTemp.getMonth(), nowTemp.getDate(), 0, 0, 0, 0);

    var checkin = $('.dpd1').datepicker({
		 format: 'yyyy-mm-dd',
        onRender: function(date) {
            return date.valueOf() < now.valueOf() ? 'disabled' : '';
        }
    }).on('changeDate', function(ev) {
            if (ev.date.valueOf() > checkout.date.valueOf()) {
                var newDate = new Date(ev.date)
                newDate.setDate(newDate.getDate() + 1);
                checkout.setValue(newDate);
            }
            checkin.hide();
            $('.dpd2')[0].focus();
        }).data('datepicker');
    var checkout = $('.dpd2').datepicker({
		 format: 'yyyy-mm-dd',
        onRender: function(date) {
            return date.valueOf() <= checkin.date.valueOf() ? 'disabled' : '';
        }
    }).on('changeDate', function(ev) {
            checkout.hide();
        }).data('datepicker');
})
</script>
</body>
</html>
