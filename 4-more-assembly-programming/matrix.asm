section .text
global index
index:
	; rdi: matrix
	; rsi: rows
	; rdx: cols
	; rcx: rindex
	; r8: cindex
	imul rcx, rdx
	add rcx, r8
	mov rax, [rdi + 4 * rcx]
	ret
