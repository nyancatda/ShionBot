[中文](ReverseProxyAPI.md) | English | [日本語](ReverseProxyAPI-ja-JP.md)
# Set up a LineBotAPI reverse proxy server
*This tutorial will show you how to use CloudFlare Workers to build your own LineBotAPI reverse proxy server*  
*If you don’t want to build, you can use the service I built, URL：https://linebotapi.h123hh.workers.dev/*
1. Go to [CloudFlare Workers](https://workers.cloudflare.com/), register an account and create a new worker
2. Enter the following code in the script  
```
const whitelist = ["/v2/bot"];
const tg_host = "api.line.me";

addEventListener('fetch', event => {
    event.respondWith(handleRequest(event.request))
})

function validate(path) {
    for (var i = 0; i < whitelist.length; i++) {
        if (path.startsWith(whitelist[i]))
            return true;
    }
    return false;
}

async function handleRequest(request) {
    var u = new URL(request.url);
    u.host = tg_host;
    if (!validate(u.pathname))
        return new Response('Unauthorized', {
            status: 403
        });
    var req = new Request(u, {
        method: request.method,
        headers: request.headers,
        body: request.body
    });
    const result = await fetch(req);
    return result;
}
```
3. Save and deploy
4. Modify the `BotAPILink` in the configuration to this Worker‘s address
