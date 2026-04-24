// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import (
	"strings"

	"github.com/expr-lang/expr/ast"
)

type Visitor struct {
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (this *Visitor) Visit(node *ast.Node) {
	switch realNode := (*node).(type) {
	case *ast.CallNode:
		var calleeName = realNode.Callee.String()
		switch calleeName {
		case "new":
			if len(realNode.Arguments) == 1 {
				switch realNode.Arguments[0].String() {
				case "Date": // new(Date)
					idNode, ok := realNode.Callee.(*ast.IdentifierNode)
					if ok {
						idNode.Value = "NewDate"
						realNode.Arguments = nil
					}
				}
			}
		case "RegExp":
			idNode, ok := realNode.Callee.(*ast.IdentifierNode)
			if ok {
				idNode.Value = "NewRegExp"
			}
		case "URL":
			idNode, ok := realNode.Callee.(*ast.IdentifierNode)
			if ok {
				idNode.Value = "NewURL"
			}
		default:
			if strings.HasPrefix(calleeName, "ctx.req.") || strings.HasPrefix(calleeName, "ctx.resp.") {
				var lastIndex = strings.LastIndex(calleeName, ".")
				if lastIndex > 0 {
					memberNode, ok := realNode.Callee.(*ast.MemberNode)
					if ok {
						property, isStringNode := memberNode.Property.(*ast.StringNode)
						if isStringNode {
							switch property.Value {
							case "url", "uri":
								property.Value = strings.ToUpper(property.Value)
							default:
								property.Value = strings.ToUpper(property.Value[:1]) + property.Value[1:]
							}
						}
					}

				}
			}
		}
	}
}
