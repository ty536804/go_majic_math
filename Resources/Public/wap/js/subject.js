var setupBottom = new Array(
    '移多补少，趣味乘法除法，整数拆分，等量代换<br>\n' +
    '            间隔问题，长度单位换算，趣味应用题，解应用题初步，智巧趣题，排队的学问，解应用题<br>\n' +
    '            生活中的可能性，重叠问题，简单的经济问题，生活中的统筹规划，相遇问题等',
    '容斥原理，计算中的去重问题，乘法原理进阶<br/>立体几何综合，几何图形拼接，角度计算，奇偶分析与应用 水管问题，变速行程，概率初步，往返初步 模式问题，洛书幻方'
);
$('.zy li').each(function () {
    $(this).on('click',function () {
        let _index = $(this).index();
        $('.setup_last_r').empty().html(setupBottom[_index])
        $(this).addClass('active_setup').siblings().removeClass('active_setup')
    })
})

var qm_active = new Array(
    '数概念初步建立<br/>\n' +
    '            基本图形认知，图形特征应用<br/>\n' +
    '            视觉训练，运笔训练，基础对应<br/>\n' +
    '            方位认知，坐标认知，多视角成像<br/>\n' +
    '            简单关系比较，序数认知<br/>\n' +
    '            简单规律辨识，事物内在联系，分析判断',
    '大数目认知，简单运算，时间概念 立体图形认知，平面与立体的转换，初步抽象思考<br/>实物空间稳定性建立，巧数立方体，空间位置关系<br/>推理顺序，面积认知及比较',
    '数的概念以及其高级应用，复杂运算，实际应用 图形转化，空间几何体构建<br/>多视角成像应用，事物空间稳定性作用<br/>体积认知与比较，量差关系 排列组合，枚举应用，数形策列'
)
$('.qm li').each(function () {
    $(this).on('click',function () {
        let _index = $(this).index();
        $('.qm_desc').empty().html(qm_active[_index])
        $(this).addClass('qm_active').siblings().removeClass('qm_active')
    })
})