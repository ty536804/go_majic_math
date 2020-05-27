getAjax(1)

//请求数据
function getAjax(page)
{
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1/singleList",
        data: {"page":page},
        success: function (result) {
            let _html= "";
            if (Number(result.code) == 200) {
                $.each(result.data.list,function (k,v) {
                    _html += "<tr><td class='w5'>"+v.id+"</td><td>"+v.nav.name+"</td><td class='title'>"+v.name+"</td><td class='summary w20'>"+v.summary+"</td>"
                    if (v.thumb_img !="") {
                        _html += "<td class='thumb_img'><img src=/static/upload/"+v.thumb_img+"></td>"
                    } else {
                        _html += "<td class='thumb_img'></td>"
                    }
                    _html += "<td class='content w20'>"+v.content+"</td>"
                    _html += '<td class="td-manage">' +
                        '<a title="编辑" href="/api/v1/singleDetail?id='+v.id+'"><i class="fa fa-pencil-square-o" aria-hidden="true"></i>编辑</a>\n' +
                        '</td></tr>';
                })
            } else {
                _html = "<tr><td colspan='9' class='text-center'>暂无内容</td></tr>"
            }
            $(".tbody").empty().append(_html)
            pageList(page,result.data.count)
        }
    });
}


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