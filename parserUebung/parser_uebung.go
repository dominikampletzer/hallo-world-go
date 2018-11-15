/**
  (C) Copyright 2018 Armin Heller

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/**
  Start from the bottom of this file and finish at the top.
  Implement the functions below such that all the tests succeed.
  In the end, use the main function to experiment with the finished parser.
*/

package parserUebung

import (
	list2 "container/list"
	"github.com/QAhell/Parser-Gombinators/parse"
	"github.com/jweigend/concepts-of-programming-languages/oop/ast"
)

/** parseExpression parses the following grammar:
  Expression := Or Spaces*
  The syntax tree is exactly the one returned by Or.
*/
func parseExpression(input parse.ParserInput) parse.ParserResult {
	return parse.Parser(parseOr).AndThen(parse.ExpectSpaces.Optional()).First()(input)
}

/** parseOr parses the following grammar:
  Or := And ^ ("|" ^ Or)?
  If the parser for ("|" ^ Or)? produces nothing then parseOr will return the
  tree returned by And. Otherwise parseOr will return a new Or Node containing
  the sub-trees returned by the recursive calls. parseOr uses expect to parse
  the symbol "|", i. e. it actually allows for Space* ^ "|".
*/
func parseOr(input parse.ParserInput) parse.ParserResult {
	return parse.Parser(parseAnd).AndThen(expect("|").AndThen(parseOr).Second().Optional()).Convert(makeOr)(input)
}

/** parseAnd parses the following grammar:
  And := Not ^ ("&" ^ And)?
  If the parser for ("&" ^ And)? produces nothing then parseAnd will return the
  tree returned by Not. Otherwise parseAnd will return a new And Node containing
  the sub-trees returned by the recursive calls. parseAnd uses expect to parse
  the symbol "&", i. e. it actually allows for Space* ^ "&".
*/
func parseAnd(input parse.ParserInput) parse.ParserResult {
	return parse.Parser(parseNot).AndThen(expect("&").AndThen(parseAnd).Second().Optional()).Convert(makeAnd)(input)
}

/** parseNot parses the following grammar:
  Not := "!"* ^ Atom
  It delegates parsing "!"* to parseExclamationMarks and the construction of Not
  nodes to makeNots. If there's no exclamation mark then parseNot will return
  the tree parsed by parseAtom. Otherwise parseNot will wrap the atom in as many
  Not nodes as there are exclamation marks.
*/
func parseNot(input parse.ParserInput) parse.ParserResult {
	return parseExclamationMarks.AndThen(parseAtom).Convert(func(arg interface{}) interface{} {
		var pair = arg.(parse.Pair)
		return makeNot(pair.First.(int), pair.Second.(ast.Node))
	})(input)

	//return parseExclamationMarks.Repeated().Convert(func(arg interface{}) interface{} {
	//	list := arg.([]interface{})
	//	slice := []interface{}{}
	//	for _, el := range list {
	//		slice = append(slice, el.(ast.Node))
	//	}
	//	return slice
	//})(input)
}

/** parseExclamationMarks parses the following grammar:
  "!"*
  It returns the number of exclamation marks in ParserResult.Result as an int.
  parseExclamationMarks uses expect to parse the symbol "!", i. e. it actually
  allows for Space* ^ "!".
*/
var parseExclamationMarks parse.Parser = func(input parse.ParserInput) parse.ParserResult {
	return expect("!").Repeated().Convert(func(arg interface{}) interface{} {
		var list = arg.(*list2.List)
		return list.Len()
	})(input)
}

/** parseAtom parses the followiong grammar:
  Atom := Variable
        | "(" ^ Expression ^ ")"
  The parenthesis won't appear in the abstract syntax tree. parseAtom uses
  Parser.First() and Parser.Second() to extract the tree returned by
  parseExpression.
*/
func parseAtom(input parse.ParserInput) parse.ParserResult {
	return parseVariable.OrElse(expect("(").AndThen(parseExpression).AndThen(expect(")")).First().Second())(input)
}

/** parseVariable parses the following grammar:
  Variable := [a-zA-Z_][a-zA-Z_0-9]*
  It delegates parsing the variable name to ExpectIdentifier from the parser
  combinators package and uses the Convert method on parsers to create the
  ast.Val node.
*/
var parseVariable parse.Parser = func(input parse.ParserInput) parse.ParserResult {
	var foo = func(a interface{}) interface{} { return ast.Val{a.(string)} }
	return parse.ExpectIdentifier.Convert(foo)(input)
}

/** makeNot wraps the node into num ast.Not nodes.
 */
func makeNot(num int, node ast.Node) ast.Node {
	//if num == 0 {
	//	return node
	//}
	//return makeNot(num-1, ast.Not{node})
	if num <= 0 {
		return node
	} else {
		return ast.Not{makeNot(num-1, node)}
	}
}

/** makeAnd takes a parse.Pair of ast.Node as an argument and returns an
  ast.Node. If the second component of the pair is equal to Nothing{} then it
  returns the first component of the Pair. If the second component is a Node
  then makeAnd will create an ast.And node containing the first and the second
  component of the Pair as sub-nodes.
*/
func makeAnd(argument interface{}) interface{} {
	var pair = argument.(parse.Pair)
	if pair.Second == (parse.Nothing{}) {
		return pair.First
	} else {
		var firstNode = pair.First.(ast.Node)
		var secondNode = pair.Second.(ast.Node)
		return ast.And{firstNode, secondNode}
	}
	//var arg = argument.(parse.Pair)
	//switch arg.Second.(type) {
	//case parse.Nothing:
	//	return arg.First
	//case ast.Node:
	//	return ast.And{arg.First.(ast.Node), arg.Second.(ast.Node)}
	//default:
	//	return nil
	//}
	//
	//return ast.And{ast.Val{"a"}, ast.Val{"b"}}
}

/** makeOr takes a parse.Pair of ast.Node as an argument and returns an
  ast.Node. If the second component of the pair is equal to Nothing{} then it
  returns the first component of the Pair. If the second component is a Node
  then makeOr will create an ast.Or node containing the first and the second
  component of the Pair as sub-nodes.
*/
func makeOr(argument interface{}) interface{} {
	var pair = argument.(parse.Pair)
	if pair.Second == (parse.Nothing{}) {
		return pair.First
	} else {
		var firstNode = pair.First.(ast.Node)
		var secondNode = pair.Second.(ast.Node)
		return ast.Or{firstNode, secondNode}
	}
	//var arg = argument.(parse.Pair)
	//switch arg.Second.(type) {
	//case parse.Nothing:
	//	return arg.First
	//case ast.Node:
	//	return ast.Or{arg.First.(ast.Node), arg.Second.(ast.Node)}
	//default:
	//	return nil
	//}
}

/** expect expects the string s at the beginning of the input and ignores
  leading spaces. */
func expect(s string) parse.Parser {
	return parse.MaybeSpacesBefore(parse.ExpectString(s))
}
