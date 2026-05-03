// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package generators

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/tjbrains/flexlang/internal/utils"
)

var exprRegex = regexp.MustCompile(`expr:"([\\$\w]+)"`)
var prototypeDocRegex = regexp.MustCompile(`@prototype\s+(.+)`)
var exprDocRegex = regexp.MustCompile(`@expr\s+([\\$\w]+)`)
var markdownLinkRegex = regexp.MustCompile(`\[\w+]\(.+\)`)

type Doc struct {
	Name      string
	Type      string
	Prototype string
}

type DocGenerator struct {
	docWriter io.Writer
	preNode   *ast.Ident

	autoCompleteCodes []map[string]string
}

func NewDocGenerator() *DocGenerator {
	return &DocGenerator{}
}

func (this *DocGenerator) Run() error {
	var rootDir = utils.RootDir()

	var mdFile = rootDir + "/docs/references.md"
	fp, err := os.OpenFile(mdFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	this.docWriter = fp

	defer func() {
		_ = fp.Close()
	}()

	err = this.write("# 变量和函数\n")
	if err != nil {
		return err
	}

	{
		var path = rootDir + "/pkg/flexlang/context/context_request.go"
		err = this.readFile(path, "")
		if err != nil {
			return err
		}
	}

	{
		var path = rootDir + "/pkg/flexlang/context/context_basic.go"
		err = this.readFile(path, "")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## Request对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/pkg/flexlang/context/request.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## Response对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/pkg/flexlang/context/response.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## CryptoHMACHash对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/internal/functions/crypto_hmac.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## Date对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/internal/functions/dates.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## RegExp对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/internal/functions/regexp.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## HTTPHeader对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/pkg/flexlang/context/http_header.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## RequestNodeInfo对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/pkg/flexlang/context/request_node.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## RequestServerInfo对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/pkg/flexlang/context/request_server.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## URL对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/internal/functions/url.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	{
		err = this.write("## URLQuery对象\n")
		if err != nil {
			return err
		}

		var path = rootDir + "/pkg/flexlang/context/url_query.go"
		err = this.readFile(path, "[object]")
		if err != nil {
			return err
		}
	}

	// auto complete
	autoCompleteJSON, err := json.Marshal(this.autoCompleteCodes)
	if err != nil {
		return err
	}
	err = os.WriteFile(rootDir+"/docs/auto-complete.json", autoCompleteJSON, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (this *DocGenerator) readFile(path string, objName string) error {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var fileSet = token.NewFileSet()
	node, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	var lastErr error
	ast.Inspect(node, func(subNode ast.Node) bool {
		if subNode == nil {
			return true
		}

		ts, isTs := subNode.(*ast.TypeSpec)
		if isTs {
			err = this.processTypeSpec(ts, fileSet, fileData, "", objName)
			if err != nil {
				lastErr = err
			}
		}

		it, isIt := subNode.(*ast.InterfaceType)
		if isIt {
			var parentName = ""
			if this.preNode != nil {
				parentName = this.preNode.Name
			}
			err = this.processInterfaceType(it, fileSet, fileData, objName, parentName)
			if err != nil {
				lastErr = err
			}
		}

		// methods
		fd, isFd := subNode.(*ast.FuncDecl)
		if isFd {
			err = this.processFuncDecl(fd, fileSet, fileData, objName)
			if err != nil {
				lastErr = err
			}
		}

		in, isIn := subNode.(*ast.Ident)
		if isIn {
			this.preNode = in
		}

		return true
	})

	return lastErr
}

func (this *DocGenerator) processFuncDecl(fd *ast.FuncDecl, fileSet *token.FileSet, fileData []byte, objName string) error {
	var ft = fd.Type
	if ft == nil {
		return nil
	}

	var typeName string
	if fd.Recv != nil && len(fd.Recv.List) > 0 && fd.Recv.List[0].Type != nil {
		typeName = string(fileData[fileSet.Position(fd.Recv.List[0].Type.Pos()).Offset:fileSet.Position(fd.Recv.List[0].Type.End()).Offset])
	}

	// 跳过内部调用的
	if strings.Contains(fd.Doc.Text(), "@internal") {
		return nil
	}

	var docText = fd.Doc.Text()

	var funcName = this.lcFirst(fd.Name.Name)
	var fullName = strings.TrimPrefix(objName+"."+funcName, ".")

	// expr
	var exprSubmatch = exprDocRegex.FindStringSubmatch(docText)
	var exprName string
	if len(exprSubmatch) > 0 {
		exprName = exprSubmatch[1]
		docText = strings.ReplaceAll(docText, exprSubmatch[0], "") // remove expr definition
		fullName = strings.TrimPrefix(objName+"."+exprName, ".")
	}

	// prototype
	var prototypeSubmatch = prototypeDocRegex.FindStringSubmatch(docText)
	var prototype string
	if len(prototypeSubmatch) > 0 {
		prototype = prototypeSubmatch[1]
		docText = strings.ReplaceAll(docText, prototypeSubmatch[0], "") // remove prototype definition
	} else {
		var start = fileSet.Position(fd.Name.End()).Offset
		var end = fileSet.Position(ft.End()).Offset
		var def = string(fileData[start:end])
		prototype = funcName + def
	}

	// auto complete
	this.createAutoComplete(typeName, fullName, true, prototype, docText)

	prototype = this.escapeForMarkdown(prototype)

	// write name
	err := this.write(strings.Repeat("#", strings.Count(fullName, ".")+2) + " " + fullName + "\n")
	if err != nil {
		return err
	}

	// write prototype
	if len(prototype) > 0 {
		if objName != "" {
			prototype = objName + "." + prototype
		}
		err = this.write("> " + prototype + "\n\n")
		if err != nil {
			return err
		}
	}

	// write doc
	err = this.write(strings.TrimSpace(docText) + "\n\n")
	if err != nil {
		return err
	}

	return nil
}

func (this *DocGenerator) processTypeSpec(ts *ast.TypeSpec, fileSet *token.FileSet, fileData []byte, parentName string, objName string) error {
	st, ok := ts.Type.(*ast.StructType)
	if !ok {
		return nil
	}
	var processErr = this.processStructType(ts.Name.Name, fileSet, fileData, st, parentName, objName)
	if processErr != nil {
		return processErr
	}

	return nil
}

func (this *DocGenerator) processInterfaceType(it *ast.InterfaceType, fileSet *token.FileSet, fileData []byte, objName string, typeName string) error {
	for _, field := range it.Methods.List {
		var fieldName = field.Names[0].String()
		if fieldName[0] < 'A' && fieldName[0] > 'Z' {
			continue
		}

		var docText = field.Doc.Text()
		docText = strings.TrimPrefix(docText, fieldName)

		var fullName = strings.TrimPrefix(objName+"."+this.lcFirst(fieldName), ".")

		// expr
		var exprSubmatch = exprDocRegex.FindStringSubmatch(docText)
		var exprName string
		if len(exprSubmatch) > 0 {
			exprName = exprSubmatch[1]
			docText = strings.ReplaceAll(docText, exprSubmatch[0], "") // remove expr definition
			fullName = strings.TrimPrefix(objName+"."+exprName, ".")
		}

		var prototypeSubmatch = prototypeDocRegex.FindStringSubmatch(docText)
		var prototype string
		if len(prototypeSubmatch) > 0 {
			prototype = prototypeSubmatch[1]
			docText = strings.ReplaceAll(docText, prototypeSubmatch[0], "") // remove prototype definition
		} else {
			funcType, isFuncType := field.Type.(*ast.FuncType)
			if isFuncType {
				start, end := fileSet.Position(funcType.Pos()).Offset, fileSet.Position(funcType.End()).Offset
				prototype = strings.ReplaceAll(string(fileData[start:end]), "functions.", "")
				if len(exprName) > 0 {
					prototype = exprName + prototype
				} else {
					prototype = this.lcFirst(fieldName) + prototype
				}
			}
		}

		// auto complete
		if typeName == "" {
			typeName = objName
		}
		this.createAutoComplete(typeName, fullName, true, prototype, docText)

		prototype = this.escapeForMarkdown(prototype)

		// write name
		err := this.write(strings.Repeat("#", strings.Count(fullName, ".")+2) + " " + fullName + "\n")
		if err != nil {
			return err
		}

		// write prototype
		if len(prototype) > 0 {
			if objName != "" {
				prototype = objName + "." + prototype
			}
			err = this.write("> " + prototype + "\n\n")
			if err != nil {
				return err
			}
		}

		// write doc
		err = this.write(strings.TrimSpace(docText) + "\n\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *DocGenerator) processStructType(typeName string, fileSet *token.FileSet, fileData []byte, st *ast.StructType, parentName string, objName string) error {
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		var fieldName = field.Names[0].String()
		if fieldName[0] >= 'A' && fieldName[0] <= 'Z' {
			var tagName string
			if field.Tag != nil {
				tagName = field.Tag.Value

				var submatch = exprRegex.FindStringSubmatch(tagName)
				if len(submatch) > 0 {
					fieldName = submatch[1]
				}
			}

			var docText = field.Doc.Text()
			if parentName == "" {
				parentName = objName
			}
			var fullName = strings.TrimPrefix(parentName+"."+fieldName, ".")

			var prototypeSubmatch = prototypeDocRegex.FindStringSubmatch(docText)
			var prototype string
			funcType, isFuncType := field.Type.(*ast.FuncType)

			if len(prototypeSubmatch) > 0 {
				prototype = prototypeSubmatch[1]
				docText = strings.ReplaceAll(docText, prototypeSubmatch[0], "") // remove prototype definition
			} else {
				if isFuncType {
					start, end := fileSet.Position(funcType.Pos()).Offset, fileSet.Position(funcType.End()).Offset
					prototype = strings.ReplaceAll(string(fileData[start:end]), "functions.", "")
					prototype = strings.TrimPrefix(prototype, "func")
					prototype = fieldName + prototype
				} else {
					prototype = fieldName
				}
			}

			// auto complete
			this.createAutoComplete(typeName, fullName, isFuncType, prototype, docText)

			prototype = this.escapeForMarkdown(prototype)

			// write name
			err := this.write(strings.Repeat("#", strings.Count(fullName, ".")+2) + " " + fullName + "\n")
			if err != nil {
				return err
			}

			// write prototype
			if len(prototype) > 0 {
				if parentName != "" {
					prototype = parentName + "." + prototype
				}
				err = this.write("> " + prototype + "\n\n")
				if err != nil {
					return err
				}
			}

			// write doc
			err = this.write(strings.TrimSpace(docText) + "\n\n")
			if err != nil {
				return err
			}

			fieldSt, fieldOk := field.Type.(*ast.StructType)
			if fieldOk {
				err = this.processStructType(typeName, fileSet, fileData, fieldSt, fullName, objName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (this *DocGenerator) write(s string) error {
	_, err := this.docWriter.Write([]byte(s))
	return err
}

func (this *DocGenerator) lcFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func (this *DocGenerator) escapeForMarkdown(s string) string {
	// 排除 [OBJECT](LINK)
	if markdownLinkRegex.MatchString(s) {
		return s
	}

	s = strings.ReplaceAll(s, "[", "\\[")
	s = strings.ReplaceAll(s, "`", "\\`")
	return s
}

func (this *DocGenerator) createAutoComplete(typeName string, fullName string, isMethod bool, prototype string, docText string) {
	// [object]
	var apply = strings.ReplaceAll(fullName, "[object].", "")
	fullName = strings.ReplaceAll(fullName, "[object]", "["+typeName+"]")

	// doc
	docText = strings.TrimSpace(docText)
	docText = strings.ReplaceAll(docText, "~~~javascript", "~~~")

	if isMethod {
		this.autoCompleteCodes = append(this.autoCompleteCodes, map[string]string{
			"label":  fullName,
			"type":   "variable",
			"apply":  apply + "()",
			"detail": prototype,
			"info":   docText,
		})
	} else {
		this.autoCompleteCodes = append(this.autoCompleteCodes, map[string]string{
			"label":  fullName,
			"type":   "variable",
			"apply":  apply,
			"detail": prototype,
			"info":   docText,
		})
	}
}
