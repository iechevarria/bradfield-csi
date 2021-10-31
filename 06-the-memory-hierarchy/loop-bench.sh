gcc -O0 main.c -o loop_order.out && valgrind --tool=cachegrind ./loop_order.out
