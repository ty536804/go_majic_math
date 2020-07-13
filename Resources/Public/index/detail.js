$(function () {
    getAjax()
//请求数据
    function getAjax()
    {
        $.ajax({
            type: "GET",
            dataType: "json",
            url: "/newDetail",
            data:{"id":$('.detail_id').val()},
            success: function (result) {
                console.log(result)
                $(".detail_tit").empty().html(result.data.detail.title)
                let detail_date = result.data.detail.created_at
                $(".detail_date").empty().html(detail_date.substring(0,10))
                $(".detail_con").empty().html(result.data.detail.content)
            }
        });
    }
})