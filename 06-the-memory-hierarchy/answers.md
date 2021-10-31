# loop order

option 1:
```
==255== I   refs:      240,150,279
==255== I1  misses:            869
==255== LLi misses:            857
==255== I1  miss rate:        0.00%
==255== LLi miss rate:        0.00%
==255== 
==255== D   refs:      112,057,421  (96,042,496 rd   + 16,014,925 wr)
==255== D1  misses:     16,001,887  (     1,331 rd   + 16,000,556 wr)
==255== LLd misses:      1,001,703  (     1,170 rd   +  1,000,533 wr)
==255== D1  miss rate:        14.3% (       0.0%     +       99.9%  )
==255== LLd miss rate:         0.9% (       0.0%     +        6.2%  )
==255== 
==255== LL refs:        16,002,756  (     2,200 rd   + 16,000,556 wr)
==255== LL misses:       1,002,560  (     2,027 rd   +  1,000,533 wr)
==255== LL miss rate:          0.3% (       0.0%     +        6.2%  )
```

option 2:
```
==236== I   refs:      240,150,279
==236== I1  misses:            869
==236== LLi misses:            857
==236== I1  miss rate:        0.00%
==236== LLi miss rate:        0.00%
==236== 
==236== D   refs:      112,057,421  (96,042,496 rd   + 16,014,925 wr)
==236== D1  misses:      1,001,887  (     1,331 rd   +  1,000,556 wr)
==236== LLd misses:      1,001,703  (     1,170 rd   +  1,000,533 wr)
==236== D1  miss rate:         0.9% (       0.0%     +        6.2%  )
==236== LLd miss rate:         0.9% (       0.0%     +        6.2%  )
==236== 
==236== LL refs:         1,002,756  (     2,200 rd   +  1,000,556 wr)
==236== LL misses:       1,002,560  (     2,027 rd   +  1,000,533 wr)
==236== LL miss rate:          0.3% (       0.0%     +        6.2%  )
```

- option 2 takes 4-5x longer to run than option 1
- looks like the same number of instructions for each program
- option 1's cache utilization is 16x better than option 2. This makes sense! A cache line is 64 bytes, and so we can get 16 ints at a time without a cache miss if we're getting adjacent values. That's what option 1 does – one cache miss for every 16 ints. option 2 does not read sequentially, so every single int read is a cache miss.
