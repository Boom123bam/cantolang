import Lexer

lexer = Lexer.Lexer(
    '''
    （a 多過 b）//com
    叫佢 i。
''')


tok = lexer.read_token()
while tok:
    print(tok)
    tok = lexer.read_token()
