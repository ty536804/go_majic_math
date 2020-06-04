var $slider = $('.slider ul');
var $slider_child_l = $('.slider ul li').length;
var $slider_width = $('.slider ul li').width();
var speed =4;
$slider.width($slider_child_l * $slider_width);

var slider_count = 0;

if ($slider_child_l < 4) {
    $('#btn-right').css({cursor: 'auto'});
}

$('#btn-right').click(function() {
    if ($slider_child_l < 4 || slider_count >= $slider_child_l - 4) {
        return false;
    }
    console.log(slider_count);
    slider_count++;
    $slider.animate({left: '-=' + $slider_width + 'px'}, 'slow');
    slider_pic();
});

$('#btn-left').click(function() {
    if (slider_count <= 0) {
        return false;
    }
    slider_count--;
    $slider.animate({left: '+=' + $slider_width + 'px'}, 'slow');
    slider_pic();
});

function slider_pic() {
    if (slider_count >= $slider_child_l - 4) {
        $('#btn-right').css({cursor: 'auto'});
    }
    else if (slider_count > 0 && slider_count <= $slider_child_l - 4) {
        $('#btn-left').css({cursor: 'pointer'});
        $('#btn-right').css({cursor: 'pointer'});
    }
    else if (slider_count <= 0) {
        $('#btn-left').css({cursor: 'auto'});
    }
}

getAjax()
//请求数据
function getAjax()
{
    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/aboutData",
        success: function (result) {
            let _html = "";
            if (Number(result.code) == 200) {
               $.each(result.data.magic,function (k,magicVal) {
                   $.each(magicVal,function (k1,v1) {
                       if (k=="魔数简介") {
                           $('.about_con').empty().html(v1.content);
                       }
                       if (k=="魔数起源") {
                           $('.com_con .tit h3').empty().html(k)
                           if (k1 < 1) {
                               $('.Jean .humanity').empty().html(v1.summary)
                               $('.Jean p.all').empty().html(v1.content)
                               // $('.Jean dd').empty().html('<img src="" data-original="/static/upload/'+v1.thumb_img+'" class="lazy">')
                           } else {
                               $('.Harvard h3.Harvard_tit').empty().html(v1.summary)
                               $('.Harvard dt').append(v1.content)
                               // $('.Harvard dd').empty().html('<img src="" data-original="/static/upload/'+v1.thumb_img+'" class="lazy">')
                           }
                       }
                       if (k=="发展历程") {
                           _html += '<li><p class="date">'+v1.summary+'<span class="bar"></span></p>'
                           _html += '<span class="developing_c">'+v1.content+'</span></li>'
                       }
                   })
               })
            }
            $('.developing_ul_right').empty().html(_html);
        }
    });
}