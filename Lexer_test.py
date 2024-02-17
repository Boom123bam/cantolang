import Lexer
import Token

input = '''
    // initialize
    叫佢 i。

    當 （i 細過 8） 時，就「
        講（i）。
        塞 i + 1 入 i。
    」

    a 大D。
'''

expected_tokens = [
   Token.COMMENT,
   Token.INITIALIZE,
   Token.IDENTIFIER,
   Token.FULLSTOP,
   Token.WHILE,
   Token.OPEN_PAREN,
   Token.IDENTIFIER,
   Token.LESS_THAN,
   Token.NUMBER,
   Token.CLOSE_PAREN,
   Token.SI,
   Token.COMMA,
   Token.THEN,
   Token.OPEN_BRACE,
   Token.IDENTIFIER,
   Token.OPEN_PAREN,
   Token.IDENTIFIER,
   Token.CLOSE_PAREN,
   Token.FULLSTOP,
   Token.ASSIGN,
   Token.IDENTIFIER,
   Token.ADD,
   Token.NUMBER,
   Token.TO,
   Token.IDENTIFIER,
   Token.FULLSTOP,
   Token.CLOSE_BRACE,
   Token.IDENTIFIER,
   Token.INCREMENT,
   Token.FULLSTOP,
]


lexer = Lexer.Lexer(input)
for expected in expected_tokens:
    got = lexer.read_token()
    if got != expected:
        raise ValueError(f'Expected "{expected}" got "{got}"')

print("OK!")
