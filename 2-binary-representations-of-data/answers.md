## 1.1
```
9 -> 0x9
136 -> 0x88
247 -> 0xE7
```

## 1.2
16^6 = 16777216

## 1.3 
34 characters

```
6 0110
8 1000
6 0110
5 0101
6 0110
c 1101
6 0110
c 1101
6 0110
f 1111

->

01101000 01100101 01101101 01101101 01101111
```

## 2.1
```
4 -> 0b100
65 -> 0b1000001
105 -> 0b1101001
255 -> 0b111111111
```
```
10 -> 2
11 -> 3
1101100 -> 108
1010101 -> 85
```

## 2.2
```
11111111  carry
 11111111
 00001101
100001100
```

Sum is 0b100001100 with 9 bits. Sum is 0b00001100 with 8 bits and overflow.

## 2.3

```
127 -> 01111111
-128 -> 10000000
-1 -> 11111111
1 -> 00000001
-14 -> 11110010
```

```
10000011 -> -125
11000100 -> -60
```

## 2.4

```
01111111
10000000+
----------------
11111111
```

- The answer is -1 which is correct. This does match my expectations.
- We can negate a number in two's complement by flipping all the bits and adding one. We can compute subtraction by adding the negation of the number to subtract.
- Most significant bit is -128 in 8-bit two's complement. Most significant bit is -2^31 in 32-bit two's complement.


## 2.5

Can detect overflow by checking if the summands have the same most significant bit if it then differs from the sum's most significant bit.

## 3.1

## 3.2

## 3.3

## 4.1

## 4.2

## 5.1

## 5.2

## 5.3

```
echo -e \\x07\\x07\\x07
```