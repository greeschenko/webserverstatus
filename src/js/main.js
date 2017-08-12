let LoginPage = new Hata(
    '#res',
    'login/form', {},
    function() {
        $('#loginbutton').bind('click', function() {
            MainPage.render();
        });
    }
);

let MainPage = new Hata(
    '#res',
    'main/layout', {},
    function() {
        let Sites = new Hata(
            '.mainpage_sites',
            'main/siteitem', {},
            function() {
                let sl = this.el;
                sl.find('.siteitem').each(function(index) {
                    let domen = $(this).data('domen');
                    new Hata(
                        $(this).find('.siteitem_stats'),
                        'main/item', {},
                        function() {
                            let self = this;
                            self.el.find('.statitem_graph').each(function(index) {
                                let el = $(this);
                                let data = self.data[index].Graph;
                                el.html('');
                                for (var i in data) {
                                    let h = data[i];
                                    if (h > 42) {
                                        h = 42;
                                    }
                                    el.append('<div class="statitem_graph_item" style="height: ' + h + 'px"></div>');
                                }
                            });
                        },
                        '/api/sites-stats?index=' + index,
                        3000
                    ).render();
                    new Hata(
                        $(this).find('.mainpage_header_status'),
                        'main/site_status', {},
                        undefined,
                        '/api/sites-status?index=' + index,
                        3000
                    ).render();
                });
                //let ri = sl.children();
                //let interval = 10000;
                //ri.each(function(index) {
                //if (index != ri.length - 1) {
                //setInterval(function() {
                //ri.eq(0).css('marginLeft', '-' + (100 * (index + 1)) + '%');
                //}, interval * (index + 1));
                //} else {
                //setInterval(function() {
                //ri.eq(0).css('marginLeft', '0%');
                //}, interval * (index + 1));
                //}
                //});
            },
            '/api/sites'
        ).render();

        let StatusInd = new Hata(
            '.mainpage_header_status',
            'main/status', {},
            undefined,
            '/api/status',
            '3000',
            function() {
                console.log('testfail');
                StatusInd.data = {
                    "Status": "FAIL"
                };
            }
        ).render();

        let StatList = new Hata(
            '.mainpage_list',
            'main/item', {},
            function() {
                StatList.el.find('.statitem_graph').each(function(index) {
                    let el = $(this);
                    let data = StatList.data[index].Graph;
                    el.html('');
                    for (var i in data) {
                        let h = data[i];
                        if (h > 42) {
                            h = 42;
                        }
                        el.append('<div class="statitem_graph_item" style="height: ' + h + 'px"></div>');
                    }
                });
            },
            '/api/stat',
            '3000'
        ).render();
    },
    '/api/main'
).render();

$('.mainhelp').bind('click', function() {
    let ww = $(window).width();
    let wh = $(window).height();
    $('#helpwrap').attr('width', ww);
    $('#helpwrap').attr('height', wh);
    var c = document.getElementById("helpwrap");
    var ctx = c.getContext("2d");
    ctx.font = "20px Arial";
    ctx.fillStyle = 'tomato';
    ctx.fillText("site or server online status", ww / 2, 20);
    ctx.fillText("stat name", 20, wh - 20);
    ctx.fillText("stat graph", ww / 4, wh - 40);
    ctx.fillText("stat status text", ww / 2, wh - 60);
    ctx.fillText("stat status indicator", 3 * ww / 4, wh - 80);

    let control = $('.mainpage_list').children().eq(0);
    let p1 = control.children().eq(0).position();
    let p2 = control.children().eq(1).position();
    let p3 = control.children().eq(2).position();
    let p4 = control.children().eq(3).position();

    let s1 = $('#mainserverstatus').position();
    let s2 = $('.mainpage_sites').children().eq(0).find('.mainpage_header_status').position();

    ctx.moveTo(s1.left + 20, s1.top + 20);
    ctx.lineTo(ww / 2 - 5, 15);
    ctx.moveTo(s2.left + 20, s2.top + 20);
    ctx.lineTo(ww / 2 - 5, 15);

    ctx.moveTo(p1.left + 20, p1.top + 20);
    ctx.lineTo(20, wh - 20 - 15);
    ctx.moveTo(p2.left + 20, p2.top + 20);
    ctx.lineTo(ww / 4, wh - 40 - 15);
    ctx.moveTo(p3.left + 20, p3.top + 20);
    ctx.lineTo(ww / 2, wh - 60 - 15);
    ctx.moveTo(p4.left + 20, p4.top + 20);
    ctx.lineTo(3 * ww / 4, wh - 80 - 15);
    ctx.strokeStyle = 'tomato';
    ctx.stroke();
    $('#helpwrap').toggle();
});
