section .text
global binary_convert
binary_convert:
	xor rax, rax
loop:
	cmp byte [rdi], 0
	je exit
	shl rax, 1
	cmp byte [rdi], 48
	je next
	add rax, 1
next:
	inc rdi
	jmp loop
exit:
	ret
