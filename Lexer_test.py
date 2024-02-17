import Lexer

lexer = Lexer.Lexer(
    '''
    // initialize
    叫佢 i。

    // assignment
    塞 0 入 i。

    // compare
    （a 係 b）
    （a 細過 b）
    （a 大過 b）

    // logic
    a 同埋 b
    a 或者 b
    唔係 a

    // if conditional
    如果 （唔係（a 係 b）） 嘅話，就「
        講（“OK”）。
    」

    // while loop
    當 （i 細過 12） 時，就「
        講（i）。
        塞 i + 1 入 i。
    」
''')


tok = lexer.read_token()
while tok:
    print(tok)
    tok = lexer.read_token()
