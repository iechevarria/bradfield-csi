nasm -fmacho64 hello_mac.asm \
&& ld -lSystem hello_mac.o -L/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/lib \
&& ./a.out
