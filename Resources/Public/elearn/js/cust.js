$(document).ready(function() {
	var _index5=0;
	$('.certificate_l').on('click',function() {
		let picLen = $('.offline_ul dl').length-1
		if(_index5<picLen){
			$(".offline_ul dl").eq(_index5).hide()
			_index5++;
			console.log(_index5);
		}
	})
	
	$('.certificate_r').on('click',function() {
		let picLen = $('.offline_ul dl').length
		if (_index5==0) {
			return false;
		}
		_index5--;
		if(_index5<picLen){
			$(".offline_ul dl").eq(_index5).show()
		}		
	})
	//品牌学术背景
	var picList = ['./static/img/zs1.png','./static/img/zs2.png','./static/img/zs1.jpg','./static/img/zs3.jpg']
	var _picLen = picList.length-1;
	var _start = 0;
	timer = setInterval(function(){
		$.each(picList,function(v) {
			$('.ppry dl.first dt img').attr('src',picList[_start])
			let _next = _start+1;
			let _last = _start+2;
			
			if (_start==(_picLen-1)) {
				_last = 0 ;
			}
			if (_start==_picLen) {
				_next = 0;
				_last = 1 ;
			}
			$('.ppry dl.active dt img').attr('src',picList[_next])
			$('.ppry dl.three dt img').attr('src',picList[_last])
			_start++
			if (_start==_picLen) {
				_start = 0;
			}
		});
	},3000)

	//品牌荣誉
	var ppList = ['./static/img/皮亚杰.png','./static/img/哈佛大学.png','./static/img/zml.png','./static/img/zxl.png']
	var textList = ['儿童心理发育是有迹可循的','MI（多元智力理论）包括语言智能、逻辑——数学智能、空间智能、肢体——运作智能，是全新人力思维智力结构理论。',
	'中科院专家 / 著名心理学家 / 儿童数理教育专家','北师大教授 / 脑与数学认知专家']
	var _ppLen = ppList.length-1;
	var _pStart = 0;
	timers = setInterval(function(){
		$.each(picList,function(v) {
			$('.xsbj dl.first dt img').attr('src',ppList[_pStart])
			let _n = _pStart+1;
			let _l = _pStart+2;
			
			if (_pStart==(_picLen-1)) {
				_l = 0 ;
			}
			if (_pStart==_picLen) {
				_n = 0;
				_l = 1 ;
			}
			
			$('.xsbj dl.xsbj_ul_mid dt img').attr('src',ppList[_n])
			$('.xsbj dl.xsbj_ul_mid dd').empty().html(textList[_n]);
			$('.xsbj dl.three dt img').attr('src',ppList[_l])
			_pStart++
			if (_pStart==_ppLen) {
				_pStart = 0;
			}
		});
	},3000);
});