$("#addpower").on("click", function () {
    if ($("#addform #nick_name").val() == "") {
        sweetAlert("操作失败","用户不能为空",'error');
        return false;
    }
    if ($("#addform #login_name").val() == "") {
        sweetAlert("操作失败","登陆账号不能为空",'error');
        return false;
    }
    if ($("#addform #email").val() == "") {
        sweetAlert("操作失败","邮箱不能为空",'error');
        return false;
    }

    if ($("#addform #id").val()=="" && $("#addform #pwd").val() == "") {
        sweetAlert("操作失败","密码不能为空",'error');
        return false;
    }

    if ($("#addform #tel").val() == "") {
        sweetAlert("操作失败","电话不能为空",'error');
        return false;
    }
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/AddUser",
        data: $("#addform").serialize(),
        success: function (result) {
            if (Number(result.code) == 200) {
                swal({title:"操作成功",type: 'success'},
                    function () {
                        initInput()
                        pageList(1,result.data.count)
                        $("#myModal2").modal('hide');
                        window.location.reload();
                    });
            } else {
                sweetAlert("操作失败",result.msg,'error');
            }
        },
        error: function (result) {
            $.each(result.responseJSON.errors, function (k, val) {
                sweetAlert("操作失败",val[0],'error');
                return false;
            });
        }
    });
    return false;
})
getResult()
function getResult() {
getAjax(1)
initInput()
}
//请求数据
function getAjax(page)
{
$.ajax({
    type: "GET",
    dataType: "json",
    url: "/api/v1/userData",
    data: {"page":page},
    success: function (result) {
        let _html= "";
        if (Number(result.code) == 200) {
            $.each(result.data.list,function (k,v) {
                _html += "<tr><td>"+v.id+"</td><td>"+v.login_name+"</td><td>"+v.nick_name+"</td><td>"+v.tel+"</td><td>"+v.email+"</td>"
                _html += "<td>"+(Number(v.statues)==1 ? '启用' : '禁用')+"</td>"
                _html += '<td class="td-manage">' +
                    '<a title="编辑" onclick="member_restore('+v.id+')" href="javascript:;"><i class="fa fa-pencil-square-o" aria-hidden="true"></i>编辑</a>\n' +
                    '</td></tr>';
            })
            pageList(page,result.data.count)
        } else {
            _html = "<tr><td colspan='9' class='text-center'>暂无内容</td></tr>"
        }
        $(".tbody").empty().append(_html)
    }
});
}

//查看当个用户信息
function member_restore(id) {
$.ajax({
    type: "POST",
    dataType: "json",
    url: "/api/v1/GetUser",
    data: {"id":id},
    success: function (result) {
        if (Number(result.code) == 200) {
            console.log(result.code)
            $("#addform #id").val(result.data.id)
            $("#addform #nick_name").val(result.data.nick_name)
            $("#addform #login_name").val(result.data.login_name)
            $("#addform #email").val(result.data.email)
            $("#addform #pwd").val("")
            $("#addform #tel").val(result.data.tel)
            if (Number(result.data.statues) ==1) {
                $('input[type="radio"]:eq(0)').prop("checked",true);
            } else {
                $('input[type="radio"]:eq(1)').prop("checked",true);
            }
            $("#myModal2").modal("show");
        } else {
            initInput()
        }
    }
});
}

//初始化提交表单
function initInput() {
    $("#addform input[type='text']").val("")
    $("#addform #city_id").val(10000)
    $("#addform #position_id").val(1)
    $("#addform #department_id").val(1)
    $('input[type="radio"]:eq(0)').prop("checked",true);
}
//分页
function pageList(page,PageCount) {
    $('#pageLimit').bootstrapPaginator({
        currentPage: page,//当前页。
        totalPages: PageCount,//总页数。
        size:"normal",//应该是页眉的大小。
        bootstrapMajorVersion: 3,//bootstrap的版本要求。
        alignment:"right",
        numberOfPages:5,//显示的页数
        itemTexts: function (type, page, current) {//如下的代码是将页眉显示的中文显示我们自定义的中文。
            switch (type) {
                case "first": return "首页";
                case "prev": return "上一页";
                case "next": return "下一页";
                case "last": return "末页";
                case "page": return page;
            }
        },
        onPageClicked: function (event, originalEvent, type, page) {//给每个页眉绑定一个事件，其实就是ajax请求，其中page变量为当前点击的页上的数字。
            getAjax(page)
        }
    });
}