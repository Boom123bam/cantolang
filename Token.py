# OPEN_PAREN = "OPEN_PAREN"
# CLOSE_PAREN = "CLOSE_PAREN"
# FULLSTOP = "FULLSTOP"

# INITIALIZE = "INITIALIZE"

# ASSIGN = "ASSIGN"
# TO = "TO"

# EQUAL_TO = "EQUAL_TO"
# LESS_THAN = "LESS THAN"

# AND = "AND"
# OR = "OR"
# NOT = "NOT"

# IF = "IF"
# DO = "DO"
# WHILE = "WHILE"

OPEN_PAREN = "（"
CLOSE_PAREN = "）"
OPEN_BRACE = "「"
CLOSE_BRACE = "」"
FULLSTOP = "。"
COMMA = "，"

INITIALIZE = "叫佢"

ASSIGN = "塞"
TO = "入"

EQUAL_TO = "係"
LESS_THAN = "細過"
GREATER_THAN = "多過"

AND = "同埋"
OR = "或者"
NOT = "唔係"

IF = "如果"
GEWA = "嘅話"

THEN = "就"

WHILE = "當"
SI = "時"

IDENTIFIER = "IDENTIFIER"
INVALID = "INVALID"
COMMENT = "COMMENT"

class Token:
    def __init__(self, type, literal):
        self.type = type
        self.literal = literal
