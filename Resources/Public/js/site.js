$(".subCon").on("click", function () {
    if ($("#addform #site_title").val() == "") {
        sweetAlert("操作失败","网站标题不能为空",'error');
        return false;
    }
    if ($("#addform #site_desc").val() == "") {
        sweetAlert("操作失败","网站描述不能为空",'error');
        return false;
    }
    if ($("#addform #site_keyboard").val() == "") {
        sweetAlert("操作失败","网站关键字不能为空",'error');
        return false;
    }
    if ($("#addform #site_keyboard").val() == "") {
        sweetAlert("操作失败","网站关键字不能为空",'error');
        return false;
    }
    if ($("#addform #land_line").val() == "" && $("#addform #client_tel").val() == "" && $("#addform #client_tel").val() == "") {
        sweetAlert("操作失败","电话联系方式，必须填写一项",'error');
        return false;
    }
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/addSite",
        data: $("#addform").serialize(),
        success: function (result) {
            if (Number(result.code) == 200) {
                swal({title:"操作成功",type: 'success'},function () {
                    getSiteInfo()
                })
            } else {
                sweetAlert("操作失败",result.msg,'error');
                return false;
            }
        }
    });
    return false
})
getSiteInfo()
function getSiteInfo() {

    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/api/v1/getSite",
        success: function (result) {
            if (Number(result.code) == 200) {
                $("#addform #id").val(result.data.id)
                $("#addform #site_title").val(result.data.site_title)
                $("#addform #site_desc").val(result.data.site_desc)
                $("#addform #site_keyboard").val(result.data.site_keyboard)
                $("#addform #site_copyright").val(result.data.site_copyright)
                $("#addform #site_tel").val(result.data.site_tel)
                $("#addform #land_line").val(result.data.land_line)
                $("#addform #client_tel").val(result.data.client_tel)
                $("#addform #site_email").val(result.data.site_email)
                $("#addform #site_address").val(result.data.site_address)
                $("#addform #record_number").val(result.data.record_number)
            }
        }
    });
}