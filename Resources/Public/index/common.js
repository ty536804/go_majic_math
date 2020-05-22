$('.btn').on('click',function () {

    if ($('.c-area').val()=="") {
        layer.tips('姓名不能为空', '.c-area', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        // layer.alert("姓名不能为空");
        return false;
    }
    if ($('.c-tel').val()=="") {
        layer.tips('电话不能为空', '.c-tel', {
            tips: [1, '#3595CC'],
            time: 4000
        });
        return false;
    }
    if ($('.c-city').val()=="") {
        layer.tips('地区不能为空', '.c-city', {
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