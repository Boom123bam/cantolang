# Cantolang

Cantonese programming language

Cantolang is currently an interpreted language using a tree walking interpreter. The code is based on Thorsten Ball's great book "Writing an Interpreter in Go".

## How to run

To run an external file:

```
go run main.go input.txt
```

To run REPL:

```
go run main.go
```

## Syntax

##### Assignment

```
塞 2 入 i。
講（i）。 // 2
```

##### Increment, decrement

```
a 大D。
a 細D。
```

##### Boolean

```
啱 // true
錯 // false
```

##### Array

```
塞【1，“二”， 啱】入 i
i【2】// 啱
```

##### Comparison

```
（a 係 b）
（a 細過 b）
（a 大過 b）
```

##### Logic

```
a 同埋 b
a 或者 b
唔係 a
```

##### If conditional

```
如果 （唔係（a 係 b）） 嘅話，就「
    講（“OK”）。
」唔係就「
    講（“NO”）。
」
```

##### While loop

```
當 （i 細過 8） 時，就「
    講（i）。
  	塞 i 加 1 入 i。
」
```

##### Function

```
聽到 add（x，y） 嘅話，就「
    俾我 x + y。 // return
」
講（add（2，3））// 5
```

##### Builtin funcitons

```
// length
有幾長（“hello”）// 5
// print
講（“OK”）// prints OK
```

## Features

- dynamic types

```
塞 2 入 i。
講（i）。 // 2
塞 “hi” 入 i。
講（i）。 // hi
```

- interchangable with standard symbols

```
（） ->  ()
【】 ->  []
「」 ->  {}
“”  ->  ""
。 -> ;
， -> ,
加減乘除 -> +-*/
```

- cantonese swag
