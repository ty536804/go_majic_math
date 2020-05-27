getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/api/v1/getNavs",
        success: function (result) {
            let _html= "";
            if (Number(result.code) == 200) {
                $.each(result.data,function (k,v) {
                    _html += "<tr><td>"+v.id+"</td><td>"+v.name+"</td><td>"+v.base_url+"</td>"
                    _html += "<td>"+(Number(v.is_show)==1 ? '启用' : '禁用')+"</td>"
                    _html += '<td class="td-manage">' +
                        '<a title="编辑" onclick="member_restore('+v.id+')" href="javascript:;"><i class="fa fa-pencil-square-o" aria-hidden="true"></i>编辑</a>\n' +
                        '</td></tr>';
                })
            } else {
                _html = "<tr><td colspan='5' class='text-center'>暂无内容</td></tr>"
            }
            $(".tbody").empty().append(_html)
        }
    });
}
$("#addpower").on("click", function () {
    if ($("#addform #name").val() == "") {
        sweetAlert("操作失败","名称不能为空",'error');
        return false;
    }
    if ($("#addform #base_url").val() == "") {
        sweetAlert("操作失败","跳转地址不能为空",'error');
        return false;
    }
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/addNav",
        data: $("#addform").serialize(),
        success: function (result) {
            if (Number(result.code) == 200) {
                swal({title:"操作成功",type: 'success'},
                    function () {
                        // $("#myModal2").modal('hide');
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

//查看当个用户信息
function member_restore(id) {
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/getNav",
        data: {"id":id},
        success: function (result) {
            if (Number(result.code) == 200) {
                $("#addform #id").val(result.data.id)
                $("#addform #name").val(result.data.name)
                $("#addform #base_url").val(result.data.base_url)
                if (Number(result.data.is_show) ==1) {
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