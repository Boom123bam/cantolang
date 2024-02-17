import Token

class Lexer:
    def __init__(self, input):
        self._input = input
        self._pos = -1
        self._char = ""
        self._peek_char = ""
        self.advance()

    def advance(self):
        self._pos += 1
        self._char = self._input[self._pos] if self._pos < len(self._input) else ""
        self._peek_char = self._input[self._pos + 1] if self._pos + 1 < len(self._input) else ""

    def read_identifier(self):
        result = ""
        while self._char.isalpha() or self._char == "_" :
            result += self._char
            self.advance()
        return result

    def read_number(self):
        result = ""
        while self._char.isdigit():
            result += self._char
            self.advance()
        return result

    def read_token(self):
        while self._char in [" ", "\n", "\r", "\t"]:
            self.advance()
        if not self._char:
            return 0

        # single char tokens
        if self._char in [
            Token.OPEN_PAREN,
            Token.CLOSE_PAREN,
            Token.OPEN_BRACE,
            Token.CLOSE_BRACE,
            Token.FULLSTOP,
            Token.COMMA,
            Token.ASSIGN,
            Token.TO,
            Token.EQUAL_TO,
            Token.THEN,
            Token.WHILE,
            Token.SI,
            Token.ADD,
            Token.MINUS,
            Token.MULTIPLY,
            Token.DIVIDE,
        ]:
            c = self._char
            self.advance()
            return c

        # double char tokens
        elif self._char in [
            Token.INITIALIZE[0],
            Token.LESS_THAN[0],
            Token.GREATER_THAN[0],
            Token.AND[0],
            Token.OR[0],
            Token.NOT[0],
            Token.IF[0],
            Token.GEWA[0],
        ] and self._char + self._peek_char in [
            Token.INITIALIZE,
            Token.LESS_THAN,
            Token.GREATER_THAN,
            Token.AND,
            Token.OR,
            Token.NOT,
            Token.IF,
            Token.GEWA,
        ]:
            tok = self._char + self._peek_char
            self.advance()
            self.advance()
            return tok

        elif self._char.isdigit():
            n = self.read_number()
            return Token.NUMBER

        elif self._char + self._peek_char == "//":
            # commemt, skip until new line
            while not self._char == "\n":
                self.advance()
            return Token.COMMENT

        # else identifier
        i = self.read_identifier()
        if not i:
            self.advance()
            return Token.INVALID
        return Token.IDENTIFIER
