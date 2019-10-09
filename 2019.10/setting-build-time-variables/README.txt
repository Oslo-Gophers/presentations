
What:
Set variables from the "outside" when compiling (go buidl, install, run etc..)

Why:
Some information is too dynamic to hardcode in the source code, but, at the same time, not something we can figure out at runtime.

The version of a program, for example, is a pain to have in in the source code because it's very easy to forget to update it every time you compile a new version.

SLIDES:
    https://docs.google.com/presentation/d/1YGkPflHXeo4P3t4Hhb3HvASoTxqqzzwEbaj-8wMnEPA/edit?usp=sharing

REPO:
    https://github.com/Oslo-Gophers/presentations

LINKS:
    https://stackoverflow.com/questions/47509272/how-to-set-package-variable-using-ldflags-x-in-golang-build
    https://blog.cloudflare.com/setting-go-variables-at-compile-time/
    https://github.com/golang/go/wiki/GcToolchainTricks (at the end)
