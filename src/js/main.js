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
