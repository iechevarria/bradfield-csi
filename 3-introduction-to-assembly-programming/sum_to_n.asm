section .text
global sum_to_n
sum_to_n:
	add rax, rdi
	dec rdi
	jg  sum_to_n
	ret
