<!DOCTYPE html>
<html>
<head>
    <title></title>
    <link rel="stylesheet" type="text/css" href="../css/css.css"/>
    <link rel="stylesheet" type="text/css" href="/themes/bootstrap/easyui.css">
    <link rel="stylesheet" type="text/css" href="/themes/icon.css">
    <script type="text/javascript" src="/js/jquery.js"></script>
    <script type="text/javascript" src="/js/jquery.easyui.min.js"></script>
    <script type="text/javascript" src="/js/easyui-lang-zh_CN.js"></script>
    <script type="text/javascript" src="/js/bll.js"></script>
    <script type="text/javascript">
        function addSet(arg) {
            if(arg.length > 0)
            {
                top.$('#pid').val(arg[0])
            }else{
                top.$('#pid').val(0)
            }
        }

        function editSet(data, arg) {
            top.$("#username").val(data["username"]);
            top.$("#address").val(data["address"]);
            top.$("#phone").val(data["phone"]);
            top.$("#isspend").val(data["isspend"]);
            top.$("#pid").val(data["pid"]);
        }

        function addSubmitSet(arg) {
        }

        function editSubmitSet(arg){
        }

        function initGrid() {          
            $('#dg').treegrid({
                singleSelect: true,
                checkOnSelect: false,
                selectOnCheck: false,
                idField:'id',    
                treeField:'username',   
                url: "?action=getAll&" + $('#f_form [name]').serialize(),
                columns:[[    
                    {field: 'id', title: 'ID',align: 'center', width: 150, checkbox: true},
                    {field: 'username', title: '姓名',align: 'left', width: 350},
                    {field: 'address', title: '超市地址',align: 'center', width: 200},
                    {field: 'phone', title: '联系方式',align: 'center', width: 100},
                    {field: 'isspend', title: '佣金是否发放',align: 'center', width: 100, formatter: function (value, row, index){
                        return row.isspend == 1 ? "是" : "否";
                    }},
                    {field: 'code', title: '操作', align: 'center', width: 200, formatter: function (value, row, index) {
                        return '<a href="javascript:void(0);" onclick="add(' + row.id + ')">添加</a> | <a href="javascript:void(0);" onclick="edit(' + row.id + ')">编辑</a> | <a href="javascript:void(0);" onclick="del(' + row.id + ')">删除</a> | <a href="javascript:void(0)" onclick="getChild('+row.id+')">下级加盟商</a>';
                    }}
                ]]
            });
        }

        function getChild(pid){
            $.post('?action=getAll&pid='+pid, function(data, textStatus, xhr) {
                childs = $('#dg').treegrid('getChildren', pid);
                for (var i = 0; i < childs.length; i++) {
                    $('#dg').treegrid('remove', childs[i].id);
                };
                $('#dg').treegrid('append',{
                    parent: pid,
                    data: data
                });
            }, 'json');
        }
    </script>
    <style type="text/css">
        .datagrid-header-check input, .datagrid-cell-check input {
            height: 12px;
            width: 12px;
            vertical-align: text-bottom;
        }
    </style>
</head>
<body class="easyui-layout">
<div data-options="region:'north'" style="height:45px; padding-top: 10px; border: none;">
    <table width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
            <td>
                <div class="head_title">加盟商管理</div>
            </td>
            <td align="right" style="padding-right:12px;" id="f_form">
                姓名：
                <input type="text" id="f_username" name="f_username" class="textbox" />
                <a id="btn3" href="javascript:void(0);" onclick="init()" class="easyui-linkbutton"
                   data-options="iconCls:'icon-search'">查询</a>
                <a id="btn2" href="javascript:void(0);" onclick="add();" class="easyui-linkbutton"
                   data-options="iconCls:'icon-add'">添加</a>
                <a id="btn2" href="javascript:void(0);" onclick="batchDel();" class="easyui-linkbutton"
                   data-options="iconCls:'icon-add'">删除</a>
            </td>
        </tr>
    </table>
</div>
<div data-options="region:'center'" style="border-left: none;">
    <table id="dg" class="easyui-layout" data-options="fit:true"></table>
</div>

<div style="display: none; width: 400px;" id="form">
    <form method="post" id="dataRow">
        <table border="0" cellpadding="3" cellspacing="5" width="100%">
            <tr>
                <td>姓名：</td>
                <td><input name="username" id="username" type="text"  style="width: 250px; height: 22px;" value=""/></td>
            </tr>
            <tr>
                <td>超市地址：</td>
                <td><input name="address" id="address" type="text"  style="width: 250px; height: 22px;" value=""/></td>
            </tr>
            <tr>
                <td>联系方式：</td>
                <td><input name="phone" id="phone" type="text"  style="width: 250px; height: 22px;" value=""/></td>
            </tr>
            <tr>
                <td>佣金是否发放：</td>
                <td>
                    <select id="isspend" name="isspend">
                        <option value="0">否</option>
                        <option value="1">是</option>
                    </select>
                </td>
            </tr>
        </table>
        <input type="hidden" id="pid" name="pid" />
    </form>
</div>
</body>
</html>