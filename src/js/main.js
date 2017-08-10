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
