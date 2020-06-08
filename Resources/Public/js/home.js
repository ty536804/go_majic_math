$('.update_info').on("click", function () {
    if ($("#commentForm #nick_name").val() == "") {
        layer.msg("用户不能为空");
        return false;
    }

    if ($("#commentForm #email").val() == "") {
        layer.msg("邮箱不能为空");
        return false;
    }

    if ($("#commentForm #tel").val() == "") {
        layer.msg("电话不能为空");
        return false;
    }
    PostData($("#commentForm").serialize())
    return false;
})

$('.head_sub').on("click", function () {
    if ($("#update_pwd #pwd").val() == "") {
        layer.msg("原始密码不能为空");
        return false;
    }

    if ($("#update_pwd #newpwd").val() == "") {
        layer.msg("新密码不能为空");
        return false;
    }

    if ($("#update_pwd #newpwd").val().length < 6) {
        layer.msg("新密码长度不能小于6位");
        return false;
    }
    PostData($("#update_pwd").serialize())
    return false;
})


GetUser()
function GetUser() {
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/api/v1/detailsUser",
        success: function (result) {
            if (Number(result.code) == 200) {
                $('.nickName').empty().html(result.data.nick_name);
                $('#commentForm #id').val(result.data.id);
                $('#update_pwd #id').val(result.data.id);
                $("#commentForm #login_name").val(result.data.login_name)
                $("#email").val(result.data.email)
                $("#tel").val(result.data.tel)
                $("#update_pwd #login_name").val(result.data.login_name)
            }
        },
    });
}
function PostData(data) {
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/editUser",
        data: data,
        success: function (result) {
            if (Number(result.code) == 200) {
                swal({title:"操作成功",type: 'success'},
                    function () {
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
}