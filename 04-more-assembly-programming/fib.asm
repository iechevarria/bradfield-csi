section .text
global fib
fib:
	push rbx
	push rbp

	mov rbx, rdi	; save rdi

	cmp rdi, 0		; leq 0, return 0
	jle case0

	cmp rdi, 2 		; leq 2, return 1
	jle case1

	sub rdi, 1		; call for n - 1
	call fib

	mov rbp, rax 	; save return value
	mov rdi, rbx	; get original rdi value back

	sub rdi, 2		; call for n - 2
	call fib

	add rax, rbp	; add fib(n-1) to fib(n-2)

	jmp exit

case1:				; fib(1) and fib(2)
	mov rax, 1
	jmp exit

case0:				; fib(0)
	mov rax, 0
	jmp exit

exit:
	pop rbp
	pop rbx
	ret
