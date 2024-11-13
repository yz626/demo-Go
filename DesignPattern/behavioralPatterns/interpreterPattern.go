package behavioralPatterns

import "strings"

// 解释器模式
// 是一种面对对象语言的编程思想，我们可以创建一种新的语言，
// 这种语言拥有自己的表达式和结构，即文法规则，这些问题的实例将对应为该语言中的句子。
// 类似于重载操作符的思想，通过定义一个表达式接口，然后实现一个终结符表达式和组合表达式，

// example:

// Expression 表达式, 包含一个解释方法
type Expression interface {
	Interpret(context string) bool
}

// terminalExpression 终结符表达式, 判断句子是否包含某字符串
type terminalExpression struct {
	matchData string
}

// Interpret 实现解释方法, 判断句子是否包含某字符串
func (t *terminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.matchData) {
		return true
	}
	return false
}

// NewTerminalExpression 创建终结符表达式
func NewTerminalExpression(matchData string) *terminalExpression {
	return &terminalExpression{matchData}
}

// AndExpression and表达式
type AndExpression struct {
	left, right Expression
}

// Interpret 实现解释方法
// 调用Expression表达式实现的Interpret方法
func (a *AndExpression) Interpret(context string) bool {
	return a.left.Interpret(context) && a.right.Interpret(context)
}

func NewAndExpression(left, right Expression) *AndExpression {
	return &AndExpression{left, right}
}

// OrExpression or表达式
type OrExpression struct {
	left, right Expression
}

// Interpret 实现解释方法
// 调用Expression表达式实现的Interpret方法
func (o *OrExpression) Interpret(context string) bool {
	return o.left.Interpret(context) || o.right.Interpret(context)
}

func NewOrExpression(left, right Expression) *OrExpression {
	return &OrExpression{left, right}
}
