# path-formatter

This is a small helper tool, which you can pass a path and it formats it in a certain way.

It is best explained by an example.

```shell
path-formatter /home/user/projects/important/path-formatter
> ~/p/i/path-formatter
```

So in a nutshell, `path-formatter` will replace your home path with `~`, use just the first
letter of the individual path segments and adds the last path segment at full again.

This tool exists for the sole reason of me needing a way of formatting paths used to set the
title of a tmux session.
