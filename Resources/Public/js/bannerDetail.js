$('.subCon').on("click",function () {
    if ($("#bname").val() == "") {
        sweetAlert("操作失败","标题不能为空",'error');
        return false;
    }
    if ($("#base_url").val() == "") {
        sweetAlert("操作失败","链接不能为空",'error');
        return false;
    }
    if ($("#bposition").val() == "") {
        sweetAlert("操作失败","选择显示位置",'error');
        return false;
    }
    if ($("#imgurl").val() == "") {
        sweetAlert("操作失败","请上传图片",'error');
        return false;
    }
    $.ajax({
        type: "POST",
        dataType: "json",
        url:"/api/v1/AddBanner",
        data:$("#add_form").serialize(),
        success:function (result) {
            if (Number(result.code) == 200) {
                swal({title:result.msg,type: 'success'},
                    function () {
                        window.location.href="/api/v1/bannerList";
                    });
            } else {
                sweetAlert("操作失败","操作失败",'error');
            }
        }
    })
    return  false;
})

getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/api/v1/getBanner",
        data:{'id':$("#id").val()},
        success: function (result) {
            let _html= '<option value="">请选择</option>';
            if (Number(result.code) == 200) {
                $.each(result.data.list,function (k,v) {
                    if (Number(result.data.detail.bposition) == Number(v.id)) {
                        _html += '<option value="'+v.id+'" selected>'+v.name+'</option>'
                    } else {
                        _html += '<option value="'+v.id+'" '+(Number(result.data.bposition)==Number(v.id) ? "selected" : "")+'>'+v.name+'</option>'
                    }
                })
                $("#id").val(result.data.detail.id)
                $("#target_link").val(result.data.detail.target_link)
                $("#info").val(result.data.detail.info)
                $("#base_url").val(result.data.detail.base_url)
                $("#bname").val(result.data.detail.bname)
                if (result.data.detail.imgurl != "") {
                    let _imgURL = '/static/upload/'+result.data.detail.imgurl
                    $("#imgurl").val(_imgURL)
                    $("#demo1").show();
                    $('#demo1').attr('src', _imgURL);
                }
            }
            $("#bposition").empty().append(_html)
        }
});
}
layui.use('upload', function(){
    var $ = layui.jquery
        ,upload = layui.upload;
    //普通图片上传
    var uploadInst = upload.render({
        elem: '#test1'
        ,url: '/api/v1/upload' //改成您自己的上传接口
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
                $("#imgurl").val(res.data)
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