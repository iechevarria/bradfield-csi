default rel

section .text
global volume
volume:
	mulss xmm0, xmm0
	mulss xmm0, xmm1
	mulss xmm0, [piover3]
	ret
section .rodata
piover3:	dd	1.04719		; >:)
