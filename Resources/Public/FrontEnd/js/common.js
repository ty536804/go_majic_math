$(function () {
    $('.c-com').val(window.location.href);
    $('.bottom_submit').on('click',function () {
        var reg =/[^u4e00-u9fa5]/
        let name = $.trim($('#myform .footer_name').val())
        if (name== "" || !reg.test(name)) {
            layer.tips('姓名不能为空', '#myform .footer_name', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            return false;
        }
        // 验证手机号
        var pattern = /^1\d{10}$/;
        let phone = $.trim($('#myform .footer_tel').val())
        if (phone == "" ||  !pattern.test(phone)) {
            layer.tips('电话号码不能为空', '#myform .footer_tel', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            return false;
        }
        let cityName = $.trim($('#myform .footer_city').val())
        if (cityName ==""  || !reg.test(cityName)) {
            layer.tips('地区不能为空', '#myform .footer_city', {
                tips: [1, '#3595CC'],
                time: 4000
            });
            return false;
        }
        console.log(111)
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