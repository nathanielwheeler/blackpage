This is a program that simply serves a black webpage.  It is intended to be used as a default new tab page in your browser.


# Installation
### Prerequisites
- linux
- go 1.16+

Install with:
```sh
$ go install github.com/nathanielwheeler/blackpage@latest
```
This makes one neat binary in your $GOBIN.  Start the server by typing `blackpage` into your terminal.

# Setup

Once installed, go to your browser settings.  In the tab settings section, find the new tab behavior and place this address into 'custom website':
```
http://localhost:9999
```

# Autostart
If you don't want to start the binary every time you restart your computer, you'll want to use some sort of autostart service.

If you use systemd, then you can make a service file `/etc/systemd/system/blackpage.service`:
```
[Unit]
Description=blackpage local website

[Service]
ExecStart=/home/{username}/go/bin/blackpage
Restart=always
RestartSec=30

[Install]
WantedBy=multi-user.target
```

Then, you reload systemctl and enable the new service:
```sh
$ systemctl daemon-reload
$ systemctl enable blackpage.service
$ sudo service blackpage start
```