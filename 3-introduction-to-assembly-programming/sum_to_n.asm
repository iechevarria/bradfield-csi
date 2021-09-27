section .text
global sum_to_n
sum_to_n:
	xor rax, rax
loop:
	add rax, rdi
	dec rdi
	jg  loop
	ret
