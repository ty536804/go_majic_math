layui.use('upload', function(){
    var $ = layui.jquery
        ,upload = layui.upload;
    //普通图片上传
    var uploadInst = upload.render({
        elem: '#test1'
        ,url: '/api/v1/upload'
        ,before: function(obj){
            //预读本地文件示例，不支持ie8
            obj.preview(function(index, file, result){
                $("#demo1").show();
                $('#demo1').attr('src', result); //图片链接（base64）
            });
        }
        ,done: function(res){
            //如果上传失败
            if(Number(res.code) ==200){
                $("#school_img").val(res.data)
            } else {
                return layer.msg('上传失败');
            }
            //上传成功
        }
        ,error: function(){
            //演示失败状态，并实现重传
            var demoText = $('#demoText');
            demoText.html('<span style="color: #FF5722;">上传失败</span> <a class="layui-btn layui-btn-xs demo-reload">重试</a>');
            demoText.find('.demo-reload').on('click', function(){
                uploadInst.upload();
            });
        }
    });
});
getAjax(1)
//请求数据
function getAjax(page)
{
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/campusList",
        data: {"page":page},
        success: function (result) {
            let _html= "";
            let _province = "";
            if (Number(result.code) == 200) {
                $.each(result.data.list,function (k,v) {
                    _html += "<tr><td class='w5'>"+v.id+"</td><td>"+v.province_name+"</td>"
                    _html += "<td class='w5'>"+v.school_name+"</td><td>"+v.school_tel+"</td>"
                    _html += "<td class='w5'>"+v.worker_time+"</td><td>"+v.address+"</td>"
                    _html += "<td class='w5'>"+v.school_img+"</td>"
                    _html += '<td class="td-manage">' +
                        '<a title="编辑" onclick="member_restore('+v.id+')" href="javascript:;"><i class="fa fa-pencil-square-o" aria-hidden="true"></i>编辑</a>\n' +
                        '</td></tr>';
                })
                $.each(result.data.areacode,function (k,v) {
                    if (Number(result.data.aid) == Number(v.id)) {
                        _province += '<option value="'+v.gaode_id+'" selected>'+v.aname+'</option>'
                    } else {
                        _province += '<option value="'+v.gaode_id+'" '+(Number(result.data.aid)==Number(v.id) ? "selected" : "")+'>'+v.aname+'</option>'
                    }
                })
            } else {
                _html = "<tr><td colspan='8' class='text-center'>暂无内容</td></tr>";
            }
            pageList(page,result.data.count)
            $(".tbody").empty().append(_html)
            $("#province").empty().append(_province)
            $("#province_name").val($("#province option:eq(0)").html());
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
        url: "/api/v1/addCampuses",
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

//查看校区信息
function member_restore(id) {
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/campusDetail",
        data: {"id":id},
        success: function (result) {
            if (Number(result.code) == 200) {
                $("#addform #id").val(result.data.detail.id)
                $("#addform #school_name").val(result.data.detail.school_name)
                $("#addform #school_tel").val(result.data.detail.school_tel)
                $("#addform #school_img").val(result.data.detail.school_img)
                $("#addform #province_name").val(result.data.detail.province_name)
                $("#addform #province").val(result.data.detail.province)
                $("#addform #address").val(result.data.detail.address)
                $("#addform #worker_time").val(result.data.detail.worker_time)
                console.log(result.data.detail.province)
                $("#addform #province").val(result.data.detail.province)
                if (result.data.detail.school_img != "") {
                    let _imgURL = '/static/upload/'+result.data.detail.school_img
                    $("#imgurl").val(_imgURL)
                    $("#demo1").show();
                    $('#demo1').attr('src', _imgURL);
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


$("#province").change(function () {
    $("#province_name").val($("#province option:selected").html());
});

$('.btn-white').on('click',function () {
    initInput()
    $('#demo1').hide()
})

$('.close').on('click',function () {
    initInput()
    $('#demo1').hide()
})