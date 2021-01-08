//请求数据
function getAjax(page)
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/video",
        data: {"page":page},
        success: function (result) {
            let _html= "";
            if (Number(result.code) == 200) {
                $.each(result.data.list,function (k,v) {
                    let timeStr = v.created_at;
                    _html += "<a href='/de?id="+v.id+"'><dl><dt class='DR'><img src='"+v.thumb_img+"'></dt><dd class='DL'>"
                    _html += '<section class="new_right"><h3>'+v.title+'</h3><span class="read">'+timeStr.substring(0,10)+'</span>'
                    _html += '</section></dd></dl></a>'
                })
                pageList(page)
            }
            $(".new_ul").empty().append(_html)
        }
    });
}
//分页
pageList(1)
function pageList(page) {
    $('#pageLimit').bootstrapPaginator({
        currentPage: page,//当前页。
        totalPages: $('#countNum').val(),//总页数。
        size:"normal",//应该是页眉的大小。
        bootstrapMajorVersion: 3,//bootstrap的版本要求。
        alignment:"right",
        numberOfPages:3,//显示的页数
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