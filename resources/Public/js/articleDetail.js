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
                $("#thumb_img").val(res.data)
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

//提交内容
$(".subCon").on("click",function () {
    if ($("#title").val() == "") {
        sweetAlert("登录失败","标题不能为空",'error');
        return false
    }
    if ($("#com").val() == "") {
        sweetAlert("登录失败","来源不能为空",'error');
        return false
    }
    if ($("#summary").val() == "") {
        sweetAlert("登录失败","摘要不能为空",'error');
        return false
    }
    if ($("#thumb_img").val() == "") {
        sweetAlert("登录失败","缩率图不能为空",'error');
        return false
    }
    if ($("#admin").val() == "") {
        sweetAlert("登录失败","发布者不能为空",'error');
        return false
    }
    if ($("#content").val() == "") {
        sweetAlert("登录失败","内容不能为空",'error');
        return false
    }
    if ($("#sort").val() == "") {
        sweetAlert("登录失败","排序不能为空",'error');
        return false
    }

    $.ajax({
        url: "/api/v1/addArticle",
        type:"POST",
        dataType: "json",
        data:$("#add_form").serialize(),
        success:function (result) {
            if (Number(result.code) == 200) {
                swal({title:result.msg,type: 'success'},
                    function () {
                        window.location.href="/api/v1/article";
                    });
            } else {
                sweetAlert("操作失败","操作失败",'error');
            }
        }
    })
    return false
})
getAjax()

//请求数据
function getAjax()
{
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/getArticle",
        data: {"id":$("#id").val()},
        success: function (result) {
            if (Number(result.code) == 200) {
                $("#title").val(result.data.title)
                $("#com").val(result.data.com)
                $("#summary").val(result.data.summary)
                $("#thumb_img").val(result.data.thumb_img)
                $("#admin").val(result.data.admin)
                $(".panel-body").html(result.data.content)
                $("#sort").val(result.data.sort)
                if (Number(result.data.is_show) ==1) {
                    $('input[name="is_show"]:eq(0)').prop("checked",true);
                } else {
                    $('input[type="is_show"]:eq(1)').prop("checked",true);
                }
                if (Number(result.data.hot) ==1) {
                    $('input[name="hot"]:eq(0)').prop("checked",true);
                } else {
                    $('input[name="hot"]:eq(1)').prop("checked",true);
                }
                if (result.data.thumb_img != "") {
                    let _imgURL = '/static/upload/'+result.data.thumb_img
                    $("#thumb_img").val(result.data.thumb_img)
                    $("#demo1").show();
                    $('#demo1').attr('src', _imgURL);
                }
            }
        }
    });
}