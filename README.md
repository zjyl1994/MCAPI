# MCAPI

A lightweight Minecraft Remote API written by Golang

一个用Golang编写的轻量级Minecraft远程控制API。

Sometimes I need to control my Minecraft server , but I'm on the bus without a computer or other places , so I need a remote control minecraft server machine usage and view the Web site , so that you can control on your phone , which that is what I did .

有些时候我需要控制我的Minecraft服务器，但是我在公交车上或者其他没有电脑的地方，所以我需要一个可以远程控制minecraft服务器并且查看机器使用情况的网站，这样就可以在手机上进行控制，这就是我所做的东西。

# Authcode

I set up a password-based authentication mechanism , ensure that the site will not be used illegally. I give it named called AuthCode.

我设置了一个基于密码的验证机制，确保网站不会被非法使用。我给它起名叫做AuthCode。

The current version of the authorization code is still static, so the project is not recommended for use in large-scale servers , to prevent password sniffing and endanger the security server .

目前版本的授权码仍然是静态的，所以本项目不建议在大型服务器中使用，以免密码被嗅探而危害服务器安全。

# Usage

/   Access root , you can see a simple panel.（These pages languages are Chinese , I will update the English version in the next time .）

/   访问根目录，你可以看到一个简单的面板。

/start POST to this URL with authcode，you can start the Minecraft server.

/start 对该路径POST授权码可以启动你的Minecraft服务器。

/stop POST to this URL with authcode，you can stop the Minecraft server.

/stop 对该路径POST授权码可以关闭你的Minecraft服务器。

/restart POST to this URL with authcode，you can restart the Minecraft server.

/restart 对该路径POST授权码可以重新启动你的Minecraft服务器。

More examples Please see pnnel.html code

更多调用例子，请查看pannel.html的代码
