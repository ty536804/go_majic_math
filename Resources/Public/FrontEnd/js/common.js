$(function () {
    $('.c-com').val(window.location.href);
    $('.bottom_submit').on('click',function () {
        if ($('#myform .footer_name').val() == "") {
            layer.tips('姓名不能为空', '.footer_name', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            return false;
        }
        if ($('#myform .footer_tel').val()=="") {
            layer.tips('电话号码不能为空┖', '.header_from .footer_tel', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            return false;
        }
        if ($('#myform .city').val()=="") {
            layer.tips('地区不能为空', '.header_from .footer_city', {
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
                layer.alert(result.msg);
                return false
            }
        })
        return false;
    })
})

$(function() {
    $("img.lazy").lazyload({
        effect : "fadeIn"
    });
});