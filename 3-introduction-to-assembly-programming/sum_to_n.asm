section .text
global sum_to_n
sum_to_n:
	xor rax, rax
.L1:
	add rax, rdi
	dec rdi
	cmp rdi, 0		; redundant but just checking if behavior changes
	jg .L1			; jnz doesn't work here - infinite loop
	ret
