certifySwiper = new Swiper('#certify .swiper-container', {
    slidesPerView: 'auto',
    centeredSlides: true,
    loop: true,
    loopedSlides: 3,
    autoplay: 3000,
    prevButton: '.swiper-button-prev',
    nextButton: '.swiper-button-next',
    //paginationClickable :true,
    onProgress: function(swiper, progress) {
        for (i = 0; i < swiper.slides.length; i++) {
            var slide = swiper.slides.eq(i);
            var slideProgress = swiper.slides[i].progress;
            modify = 1;
            if (Math.abs(slideProgress) > 1) {
                modify = (Math.abs(slideProgress) - 1) * 0.3 + 1;
            }
            translate = slideProgress * modify * 260 + 'px';
            scale = 1 - Math.abs(slideProgress) / 5;
            zIndex = 999 - Math.abs(Math.round(10 * slideProgress));
            slide.transform('translateX(' + translate + ') scale(' + scale + ')');
            slide.css('zIndex', zIndex);
            slide.css('opacity', 1);
            if (Math.abs(slideProgress) > 3) {
                slide.css('opacity', 0);
            }
        }
    },
    onSetTransition: function(swiper, transition) {
        for (var i = 0; i < swiper.slides.length; i++) {
            var slide = swiper.slides.eq(i)
            slide.transition(transition);
        }
    }
});
var mobileAgent = new Array("iphone", "ipod", "ipad", "android", "mobile", "blackberry", "webos", "incognito", "webmate", "bada", "nokia", "lg", "ucweb", "skyfire");
var browser = navigator.userAgent.toLowerCase();
var isMobile = false;
for (var i=0; i<mobileAgent.length; i++) {
    if (browser.indexOf(mobileAgent[i])!=-1){
        isMobile = true;
        location.href = '/wap';
        break;
    }
}