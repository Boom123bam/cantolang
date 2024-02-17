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
    (Token.COMMENT,None),
    (Token.INITIALIZE,"叫佢"),
    (Token.IDENTIFIER,"i"),
    (Token.FULLSTOP,"。"),
    (Token.WHILE,"當"),
    (Token.OPEN_PAREN,"（"),
    (Token.IDENTIFIER,"i"),
    (Token.LESS_THAN,"細過"),
    (Token.NUMBER,"8"),
    (Token.CLOSE_PAREN,"）"),
    (Token.SI,"時"),
    (Token.COMMA,"，"),
    (Token.THEN, "就"),
    (Token.OPEN_BRACE, "「"),
    (Token.IDENTIFIER, "講"),
    (Token.OPEN_PAREN, "（"),
    (Token.IDENTIFIER, "i"),
    (Token.CLOSE_PAREN, "）"),
    (Token.FULLSTOP, "。"),
    (Token.ASSIGN, "塞"),
    (Token.IDENTIFIER, "i"),
    (Token.ADD, "+"),
    (Token.NUMBER, "1"),
    (Token.TO, "入"),
    (Token.IDENTIFIER, "i"),
    (Token.FULLSTOP, "。"),
    (Token.CLOSE_BRACE, "」"),
    (Token.IDENTIFIER, "a"),
    (Token.INCREMENT, "大D"),
    (Token.FULLSTOP, "。"),
    (Token.EOF, None),
]


lexer = Lexer.Lexer(input)
i=0
for [expected_type, expected_literal] in expected_tokens:
    got = lexer.read_token()
    if got.type != expected_type:
        print(f'Tokens [{i}] Expected Type "{expected_type}" got "{got.type}"')
    if got.literal != expected_literal:
        print(f'Tokens [{i}] Expected Literal "{expected_literal}" got "{got.literal}"')
    i+= 1
