# webServerStatus

golang based, highly configured web app for check server and sites condition

![help](https://raw.githubusercontent.com/greeschenko/webserverstatus/master/help.png)

## Installation

`git clone https://github.com/greeschenko/webserverstatus.git`
`mv -r scripts_sample/ scripts/`
`mv package.cfg_sample package.cfg`

## Configuration

edit package.cfg

first add server config

```
{
    "server":{
        "ip":"88.88.88.88"                                  //add your server IP
    },
    "stats":[
        ...
        {
            "name":"CPU",                                   // stat item name
            "condition":"bash scripts/cpu.sh",              // script output can be string
            "status":"bash scripts/cpu.sh status",          // script output must be in 'OK' 'WARN' 'FAIL'
            "graphs":[
                {
                    "command":"bash scripts/cpu.sh graph",  // graph present value, script ouput must be number
                    "max":"bash scripts/cpu.sh graphmax"    // graph max posible value, script ouput must be number
                }
            ]
        },
        ...
    ],
    ...
```

next add sites config if it need

```
    "sites":[                                               // You can add many sites that are hosted on your server
        ...
        {
            "domen": "site.com",                            // add site domen
            "checklinks": [                                 // add one or many links for checking it 200 status
                "https://site.com",
                "https://site.com/main/"
            ],
            "stats":[                                       // add array of site stats similarly as for the server
                ...
                {
                    "name":"NGINX",
                    "condition":"bash scripts/webserv.sh /var/log/nginx/access.log 3",
                    "status":"bash scripts/webserv.sh /var/log/nginx/access.log 3 status",
                    "graphs":[
                        {
                            "command":"bash scripts/webserv.sh /var/log/nginx/access.log 3 graph",
                            "max":"bash scripts/webserv.sh /var/log/nginx/access.log 3 graphmax"
                        }
                    ]
                }
                ...
            ]
        }
        ...
    ]
```

## Usage

configure nginx

```
server {
    listen 80;
    listen [::]:80;
    server_name yourdomenforstats.com;
    root /var/www/html/web;
    index index.html;
    sendfile  off;
    location / {
        try_files $uri $uri/ =404;
        access_log        off;
        expires           0;
        add_header        Cache-Control private;
    }
    location /api {
        include             fastcgi_params;
        proxy_pass  http://localhost:1888;
    }
}
```

start backend

`./webserverstatus &`

or add it in your server autostart script

go to yourdomenforstats.com and watch by your server and sites condition online


