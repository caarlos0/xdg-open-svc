# xdg-open-svc

A simple service that you can run in your host, forward the socket over SSH, and
use it to open URLs and etc from your target host.

## How can I use this?

Basically, forward the socket to the target host:

```
Host myhost
  RemoteForward [localhost]:2226 [127.0.0.1]:2226
```

Then, run this service. You can run it as a native service, or just execute the
binary.

Finally, SSH into the host:

```bash
ssh myhost
```

And you should be able to "remote open" using `nc`:

```bash
echo "https://carlosbecker.com" | nc localhost 2226
```

## Making it better

Tools like `gh` and many others use `open` or `xdg-open` to open links, and they
obviously won't use `nc` automatically.

A way to work around that is to have your own `open` binary in a higher place in
the path. An example of how that binary could look like:

```bash
#!/bin/bash
set -eo pipefail
if test -n "$SSH_TTY"; then
	echo "$@" | nc -q1 localhost 2226
else
	/usr/bin/open $*
fi
```

---

I hope you find this useful :)
