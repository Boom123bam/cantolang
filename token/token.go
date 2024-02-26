package token

const (
	OPEN_PAREN  = "OPEN_PAREN"
	CLOSE_PAREN = "CLOSE_PAREN"
	OPEN_BRACE  = "OPEN_BRACE"
	CLOSE_BRACE = "CLOSE_BRACE"
	EOL         = "EOL"
	COMMA       = "COMMA"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	ADD      = "ADD"
	MINUS    = "MINUS"
	MULTIPLY = "MULTIPLY"
	DIVIDE   = "DIVIDE"

	INCREMENT = "INCREMENT"
	DECREMENT = "DECREMENT"

	INITIALIZE = "INITIALIZE"

	ASSIGN = "ASSIGN"
	TO     = "TO"

	EQUAL_TO     = "EQUAL_TO"
	LESS_THAN    = "LESS_THAN"
	GREATER_THAN = "GREATER_THAN"

	AND = "AND"
	OR  = "OR"
	NOT = "NOT"

	IF   = "IF"
	ELSE = "ELSE"
	GEWA = "GEWA"

	THEN = "THEN"

	WHILE = "WHILE"
	SI    = "SI"

	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"

	IDENTIFIER      = "IDENTIFIER"
	INVALID         = "INVALID"
	COMMENT         = "COMMENT"
	NUMBER          = "NUMBER"
	EOF             = "EOF"
	TEMP_NOT_SYMBOL = "TEMP_NOT_SYMBOL"
)

var symbols = map[rune]string{
	'（': OPEN_PAREN,
	'）': CLOSE_PAREN,
	'「': OPEN_BRACE,
	'」': CLOSE_BRACE,
	'。': EOL,
	'，': COMMA,

	'(': OPEN_PAREN,
	')': CLOSE_PAREN,
	'{': OPEN_BRACE,
	'}': CLOSE_BRACE,
	';': EOL,
	',': COMMA,

	'加': ADD,
	'減': MINUS,
	'乘': MULTIPLY,
	'除': DIVIDE,

	'+': ADD,
	'-': MINUS,
	'*': MULTIPLY,
	'/': DIVIDE,
}

var keywords = map[string]string{
	"係":   EQUAL_TO,
	"細過":  LESS_THAN,
	"大過":  GREATER_THAN,
	"同埋":  AND,
	"或者":  OR,
	"唔係":  NOT,
	"如果":  IF,
	"唔係就": ELSE,
	"嘅話":  GEWA,
	"大D":  INCREMENT,
	"細D":  DECREMENT,
	"叫佢":  INITIALIZE,
	"就":   THEN,
	"當":   WHILE,
	"時":   SI,
	"塞":   ASSIGN,
	"入":   TO,
	"聽到":  FUNCTION,
	"俾我":  RETURN,
	"啱":   TRUE,
	"錯":   FALSE,
}

func LookUpSymbol(symbol rune) string {
	ident, ok := symbols[symbol]
	if !ok {
		return TEMP_NOT_SYMBOL
	}
	return ident
}

func LookUpIdent(keyword string) string {
	ident, ok := keywords[keyword]
	if !ok {
		return IDENTIFIER
	}
	return ident
}

type Token struct {
	TokenType    string
	TokenLiteral string
}
