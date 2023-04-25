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

I wrote an article about how I use this, along with `pbcopy` and `pbpaste` over
SSH. You can read it
[here](https://carlosbecker.com/posts/pbcopy-pbpaste-open-ssh/). It gives a few
tips on how to make the experience better than using `nc` 😄

## Alternatives

As pointed out by @pbnj on #1, you can also use `ncat`, which comes with the
`nmap` package. There are a few subtle differences, like the lack of logging and
listening to `0.0.0.0` instead of `localhost`, but they should work more or less
the same apart from that.

---

I hope you find this useful :)
