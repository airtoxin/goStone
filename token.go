package stone

type Token interface {
	NewToken(line int) *Token
	LineNumber(t *Token) int
	IsIdentifier(t *Token) bool
	IsNumber(t *Token) bool
	IsString(t *Token) bool
	Number(t *Token) int
	Text(t *Token) string
}
