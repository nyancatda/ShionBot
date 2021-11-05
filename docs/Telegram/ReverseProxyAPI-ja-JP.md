[中文](ReverseProxyAPI.md) | [English](ReverseProxyAPI-en-US.md) | 日本語
# TelegramBotAPI反向代理服务器搭建
*本篇教程将告诉你如何使用CloudFlare Workers搭建一个自己的TelegramBotAPI反向代理服务器*
1. 前往[CloudFlare Workers](https://workers.cloudflare.com/)官网，注册一个账号，并新建一个Worker
2. 在脚本中写入以下代码  
`<Bot API Token>`:你的机器人token
```
const whitelist = ["/bot<Bot API Token>"];
const tg_host = "api.telegram.org";

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
3. 点击保存并部署
4. 在将配置文件中的`BotAPILink`修改为这个Worker的地址即可