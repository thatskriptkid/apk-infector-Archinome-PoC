# Apk infector Archinome PoC

This program infects APK with malicious code using DEX/Manifest patching. It also can be used to inject frida gadget for using frida on non-rooted device

**Full description about What is it and How it works:**

https://www.orderofsixangles.com/en/2020/04/07/android-infection-the-new-way.html (EN)

https://www.orderofsixangles.com/ru/2020/07/04/Infecting-android-app-the-new-way.html (RU)

**Please read article berfore use it!**

# Prerequisite

Install and add to PATH

1. Android SDK
2. zipalign
3. apksigner 

# Usage

```
Usage:
main input.apk output.apk -o [option]
options:
        1 - custom payload
        2 - frida inject
```

zipalign and siging with test key:

```
./build.sh
```

To inject your malicious code, you should place file named payload_custom.dex with malicious code that follow rules:

1. Class name within payload.dex - `aaaaaaaaaaaa.payload`

2. Method `public void executePayload()`

After you infect apk please align and sign it. You can use `build.sh` script for it.

If there are problems make sure that:
   1. The original application works
   2. All file paths in PoC are correct
   3. There's nothing unusual in apkinfector.log.
   4. The name of the original Application class in the patched InjectedApp.dex is really in its place. 
   5. The target application uses its Application class. Otherwise, PoC inoperability is predictable.

If nothing helped, try to play with the `-min-api` parameter when compiling payload classes.
If nothing worked, then create an issue on github.

PoC includes files from https://github.com/avast/apkparser.

I am not a Go developer so forgive me for the quality of code

# TODO

1. Add signing in golang
2. Create anti debug frida gadget according to https://docs.google.com/presentation/d/1BktWJ91ill5iI_-ENzh2Uq14BGIHxxpONzNYybYJIC4/edit#slide=id.p. Add this to project.
3. Get rid of Manifest unpacking
