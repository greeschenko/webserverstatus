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
                        undefined,
                        '/api/sites-stats?index=' + index,
                        3000
                    ).render();
                    new Hata(
                        $(this).find('.mainpage_header_status'),
                        'main/status', {},
                        undefined,
                        '/api/sites-status?index=' + index,
                        3000
                    ).render();
                });
                let ri = sl.children();
                let interval = 10000;
                ri.each(function(index) {
                    if (index != ri.length - 1) {
                        setInterval(function() {
                            ri.eq(0).css('marginLeft', '-' + (100 * (index + 1)) + '%');
                        }, interval * (index + 1));
                    } else {
                        setInterval(function() {
                            ri.eq(0).css('marginLeft', '0%');
                        }, interval * (index + 1));
                    }
                });
            },
            '/api/sites'
        ).render();

        let StatusInd = new Hata(
            '.mainpage_header_status',
            'main/status', {},
            undefined,
            '/api/status'
            //'3000'
        ).render();

        let StatList = new Hata(
            '.mainpage_list',
            'main/item', {},
            undefined,
            '/api/stat',
            '3000'
        ).render();
    },
    '/api/main'
).render();
