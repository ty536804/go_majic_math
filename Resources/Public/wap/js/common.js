$(function() {
    $("img.lazy").lazyload({
        effect : "fadeIn"
    });
});
$('.home_icon').on('click',function () {
    let flag = $(".home_nav").is(":hidden");
    if(flag){
        $(".home_nav").show();
    }else{
        $(".home_nav").hide();
    }
})

$('.c-com').val(window.location.href);

$('.f_btn').on('click',function () {
    if ($('.footer_con_right .c-area').val()=="") {
        layer.tips('姓名不能为空', '.footer_con_right .c-area', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.footer_con_right .c-tel').val()=="") {
        layer.tips('电话不能为空', '.footer_con_right .c-tel', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.footer_con_right .c-city').val()=="") {
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
        data:$('#myform').serialize(),
        success: function (result) {
            if (result.code == 200) {
                layer.alert("留言成功");
                return false
            }
            layer.alert("留言失败");
            return false
        }
    })
    return false;
})
