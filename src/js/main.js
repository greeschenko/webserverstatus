let LoginPage = new Hata(
    '#res',
    'login/form', {},
    function() {
        $('#loginbutton').bind('click', function() {
            MainPage.render();
        });
    }
).render();

let MainPage = new Hata(
    '#res',
    'main/layout', {},
    function() {
        let StatusInd = new Hata(
            '.mainpage_header_status',
            'main/status', {},
            undefined,
            '/data/status.json',
            '3000',
        ).render();

        let StatList = new Hata(
            '.mainpage_list',
            'main/item', {},
            undefined,
            '/data/stat.json',
            '3000',
        ).render();
    },
    '/data/main.json'
);
