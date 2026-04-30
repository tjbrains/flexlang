// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package visitors

import (
	"reflect"
	"slices"
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
			if strings.Contains(calleeName, ".") {
				memberNode, ok := realNode.Callee.(*ast.MemberNode)
				if ok {
					var memberNames []string

					if memberNode.Node != nil {
						var nodeType = memberNode.Node.Type()
						switch nodeType.Kind() {
						case reflect.Struct:
							memberNames = this.lookupMemberNames(nodeType)
						case reflect.Interface:
							for method := range nodeType.Methods() {
								memberNames = append(memberNames, method.Name)
							}
						default:
						}
					}

					// 转换为实际的 字段 或 方法 名
					property, isStringNode := memberNode.Property.(*ast.StringNode)
					if isStringNode {
						realName, found := this.lookupRealName(memberNames, property.Value)
						if found {
							property.Value = realName
						} else {
							property.Value = strings.ToUpper(property.Value[:1]) + property.Value[1:]
						}
					}
				}
			}
		}
	}
}

func (this *Visitor) Reset() {
}

func (this *Visitor) lookupMemberNames(nodeType reflect.Type) []string {
	var memberNames []string
	if nodeType.Kind() == reflect.Struct {
		for field := range nodeType.Fields() {
			value, lookOk := field.Tag.Lookup("expr")
			if lookOk {
				memberNames = append(memberNames, value)
			}
		}
		for method := range nodeType.Methods() {
			memberNames = append(memberNames, method.Name)
		}
	}

	return memberNames
}

func (this *Visitor) lookupRealName(names []string, currentName string) (result string, found bool) {
	// 先精准查找
	if slices.Contains(names, currentName) {
		return currentName, true
	}

	// 转换为小写查找
	var lowerName = strings.ToLower(currentName)
	for _, name := range names {
		if strings.ToLower(name) == lowerName {
			return name, true
		}
	}

	// 找不到原样返回
	return currentName, false
}
