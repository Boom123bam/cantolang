package token

const (
	OPEN_PAREN  = "（"
	CLOSE_PAREN = "）"
	OPEN_BRACE  = "「"
	CLOSE_BRACE = "」"
	FULLSTOP    = "。"
	COMMA       = "，"

	ADD      = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	// ADD = "加"
	// MINUS = "減"
	// MULTIPLY = "乘"
	// DIVIDE = "除"

	INCREMENT = "大D"
	DECREMENT = "細D"

	INITIALIZE = "叫佢"

	ASSIGN = "塞"
	TO     = "入"

	EQUAL_TO     = "係"
	LESS_THAN    = "細過"
	GREATER_THAN = "大過"

	AND = "同埋"
	OR  = "或者"
	NOT = "唔係"

	IF   = "如果"
	GEWA = "嘅話"

	THEN = "就"

	WHILE = "當"
	SI    = "時"

	IDENTIFIER = "IDENTIFIER"
	INVALID    = "INVALID"
	COMMENT    = "COMMENT"
	NUMBER     = "NUMBER"
	EOF        = "EOF"
)

type Token struct {
	TokenType    string
	TokenLiteral string
}
