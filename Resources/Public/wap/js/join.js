$('.c_com').val(window.location.href);

$('.s_btn').on('click',function () {
    if ($('.banner_con_right .c-area').val()=="") {
        layer.tips('姓名不能为空', '.footer_con_right .c-area', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.banner_con_right .c-tel').val()=="") {
        layer.tips('电话不能为空', '.footer_con_right .c-tel', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.banner_con_right .c-city').val()=="") {
        layer.tips('地区不能为空', '.footer_con_right .c-city', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/AddMessage",
        data:$('#bannerForm').serialize(),
        success: function (result) {
            if (result.code == 200) {
                $('input').val('');
                layer.alert("留言成功");
                return false
            }
            layer.alert("留言失败");
            return false
        }
    })
    return false;
})