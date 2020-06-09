var mt = '-'+$('.school_map_bg img').height();
$("#chart-panel").css("height",$('.school_map_bg img').height()+200+"px");
$("#chart-panel").css("margin",mt+'px auto 0')
var myChart = echarts.init(document.getElementById('chart-panel'));
var uploadedDataURL = "/static/index/data-1528971808162-BkOXf61WX.json";
myChart.showLoading();
var symbolSize =0;
var data = [
	{
		name: "北京",
		value: 0
	},
	{
		name: "天津",
		value: 0
	},
	{
		name: "河北",
		value: 0
	},
	{
		name: "山西",
		value: 0
	},
	{
		name: "内蒙古自治区",
		value: 0
	},
	{
		name: "辽宁",
		value: 0
	},
	{
		name: "吉林",
		value: 0
	},
	{
		name: "黑龙江",
		value: 0
	},
	{
		name: "上海",
		value: 0
	},
	{
		name: "江苏",
		value: 0
	},
	{
		name: "浙江",
		value: 0
	},
	{
		name: "安徽",
		value: 0
	},
	{
		name: "福建",
		value: 0
	},
	{
		name: "江西",
		value: 0
	},
	{
		name: "山东",
		value: 0
	},
	{
		name: "河南",
		value: 0
	},
	{
		name: "湖北",
		value: 0
	},
	{
		name: "湖南",
		value: 0
	},
	{
		name: "重庆",
		value: 0
	},
	{
		name: "四川",
		value: 0
	},
	{
		name: "贵州",
		value: 0
	},
	{
		name: "云南",
		value: 0
	},
	{
		name: "西藏自治区",
		value: 0
	},
	{
		name: "陕西",
		value: 0
	},
	{
		name: "甘肃",
		value: 0
	},
	{
		name: "青海",
		value: 0
	},
	{
		name: "宁夏回族自治区",
		value: 0
	},
	{
		name: "新疆维吾尔自治区",
		value: 0
	},
	{
		name: "广东",
		value: 0
	},
	{
		name: "广西壮族自治区",
		value: 0
	},
	{
		name: "海南",
		value: 0
	},
	{
		name: "台湾",
		value: 0
	},
	{
		name: "香港",
		value: 0
	},
	{
		name: "澳门",
		value: 0
	},
];
getCampus()
function getCampus() {
	$.ajax({
		type: "POST",
		dataType: "json",
		url: "/groupCampuses",
		success: function (result) {
			console.log(result)
			if (Number(result.code) == 200) {
				$.each(result.data.detail,function (rk,rv) {
					$.each(data,function (datak,dataV) {
						if (rv.name == dataV.name) {
							data[datak]["value"] = rv.c_province
						}
					})
				})
				symbolSize = result.data.detail.length
			}
		}
	});
}


$.getJSON(uploadedDataURL, function(geoJson) {
	echarts.registerMap('china', geoJson);
	myChart.hideLoading();
	var geoCoordMap = {
		'台湾': [121.5135, 25.0308],
		'黑龙江': [127.9688, 45.368],
		'内蒙古自治区': [110.3467, 41.4899],
		"吉林": [125.8154, 44.2584],
		'北京': [116.4551, 40.2539],
		"辽宁": [123.1238, 42.1216],
		"河北": [114.4995, 38.1006],
		"天津": [117.4219, 39.4189],
		"山西": [112.3352, 37.9413],
		"陕西": [109.1162, 34.2004],
		"甘肃": [103.5901, 36.3043],
		"宁夏": [106.3586, 38.1775],
		"青海": [101.4038, 36.8207],
		"新疆维吾尔自治区": [87.9236, 43.5883],
		"西藏自治区": [91.11, 29.97],
		"四川": [103.9526, 30.7617],
		"重庆": [108.384366, 30.439702],
		"山东": [117.1582, 36.8701],
		"河南": [113.4668, 34.6234],
		"江苏": [118.8062, 31.9208],
		"安徽": [117.29, 32.0581],
		"湖北": [114.3896, 30.6628],
		"浙江": [119.5313, 29.8773],
		"福建": [119.4543, 25.9222],
		"江西": [116.0046, 28.6633],
		"湖南": [113.0823, 28.2568],
		"贵州": [106.6992, 26.7682],
		"云南": [102.9199, 25.4663],
		"广东": [113.12244, 23.009505],
		"广西壮族自治区": [108.479, 23.1152],
		"海南": [110.3893, 19.8516],
		'上海': [121.4648, 31.2891],
		'香港': [114.2578, 22.3242],
		'澳门': [113.5547, 22.1484],
	};

	var max = 480,
		min = 9; // todo
	var maxSize4Pin = 100,
		minSize4Pin = 20;

	var convertData = function(data) {
		var res = [];
		for (var i = 0; i < data.length; i++) {
			var geoCoord = geoCoordMap[data[i].name];
			if (geoCoord) {
				res.push({
					name: data[i].name,
					value: geoCoord.concat(data[i].value)
				});
			}
		}
		return res;
	};



	option = {
		itemStyle:{
			normal:{
				label:{show:true}
				,
				areaStyle:{color:'green'}//设置地图背景色的颜色设置
				,color:'rgba(255,0,255,0.8)' //刚才说的图例颜色设置
			},
			emphasis:{label:{show:true}}
		},
		backgroundColor: {
			type: 'linear',
			x: 0,
			y: 0,
			x2: 0,
			y2: 0,
			areaStyle:{color:'green'},
			colorStops: [{
				offset: 0,
				color: '#F9FAFC' // 0% 处的颜色
			}, {
				offset: 1,
				color: '#F9FAFC' // 100% 处的颜色
			}],

			globalCoord: false // 缺省为 false
		},

		title: {
			top: 20,
			subtext: '',
			x: 'center',
			textStyle: {
				color: '#ccc'
			}
		},

		tooltip: {
			trigger: 'item',
			formatter: function(params) {
				if (typeof(params.value)[2] == "undefined") {
					return params.name + ' : ' + params.value;
				} else {
					return params.name + ' : ' + params.value[2];
				}
			}
		},
		legend: {
			orient: 'vertical',
			y: 'bottom',
			x: 'right',
			data: ['学校分布'],
			textStyle: {
				color: '#fff'
			}
		},
		visualMap: {
			show: false,
			min: 0,
			max: 500,
			left: 'left',
			top: 'bottom',
			text: ['高', '低'], // 文本，默认为数值文本
			calculable: true,
			seriesIndex: [1],
			inRange: {

			}
		},
		geo: {
			map: 'china',
			show: true,
			roam: false,
			label: {
				normal: {
					show: false
				},
				emphasis: {
					show: false,
				}
			},
			itemStyle: {
				normal: {
					areaColor: '#3a7fd5',
					borderColor: '#0a53e9', //线
					shadowColor: '#092f8f', //外发光
					shadowBlur: 20
				},
				emphasis: {
					areaColor: '#0a2dae', //悬浮区背景
				}
			}
		},
		series: [{

			symbolSize: 5,
			label: {
				normal: {
					formatter: '{b}',
					position: 'right',
					show: true
				},
				emphasis: {
					show: true
				}
			},
			itemStyle: {
				normal: {
					color: '#fff'
				}
			},
			name: 'light',
			type: 'scatter',
			coordinateSystem: 'geo',
			data: convertData(data),

		},
			{
				type: 'map',
				map: 'china',
				geoIndex: 0,
				aspectScale: 0.75, //长宽比
				showLegendSymbol: false, // 存在legend时显示
				label: {
					normal: {
						show: false
					},
					emphasis: {
						show: false,
						textStyle: {
							color: '#fff'
						}
					}
				},
				roam: false,
				itemStyle: {
					normal: {
						areaColor: '#031525',
						borderColor: '#FFFFFF',
					},
					emphasis: {
						areaColor: '#2B91B7'
					}
				},
				animation: false,
				data: data
			},
			{
				name: 'Top 5',
				type: 'scatter',
				coordinateSystem: 'geo',
				symbol: 'pin',
				symbolSize: [50, 50],
				label: {
					normal: {
						show: true,
						textStyle: {
							color: '#fff',
							fontSize: 9,
						},
						formatter(value) {
							return value.data.value[2]
						}
					}
				},
				itemStyle: {
					normal: {
						color: '#D8BC37', //标志颜色
					}
				},
				data: convertData(data.sort(function(a, b) {
					return b.value - a.value;
				}).slice(0, symbolSize)),
				showEffectOn: 'render',
				rippleEffect: {
					brushType: 'stroke'
				},
				hoverAnimation: true,
				zlevel: 1
			},

		]
	};
	myChart.setOption(option);
	myChart.on('click',function (params) {
		getCityRes(params.data.name);
	})
});
getCityRes("北京")
function getCityRes(tit) {
	$.ajax({
		type: "POST",
		dataType: "json",
		url: "/campusData",
		data: {"province":tit},
		success: function (result) {
			let _html = "";
			if (Number(result.code) == 200) {
				$.each(result.data,function (k,v) {
					_html +='<dl><dt><h3>'+v.school_name+'</h3><span class="shcool_line"></span>'
					_html +='<p class="tel">'+v.school_tel+'</p>'
					_html +='<p class="work">'+v.worker_time+'</p>'
					_html +='<p class="address">'+v.address+'</p>'
					_html +='</dt><dd><img src="/static/upload/'+v.school_img+'"></dd></dl>'
				})
			}
			$('.school_ul').empty().append(_html);
		}
	});
}