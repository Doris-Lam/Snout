package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"monkey/ast"
	"net/http"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string {
	if b.Value {
		return "vrai" // French for true
	}
	return "faux" // French for false
}

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "nul" } // French for null

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERREUR : " + e.Message }

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string {
	return TranslateToFrench(s.Value)
}

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fonction") // French for "function"
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// You may already have these types elsewhere in your codebase:
type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Hashable interface {
	HashKey() HashKey
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// Implement HashKey for Integer, Boolean, and String

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

var manualDict = map[string]string{
	"hi":        "bonjour",
	"dog":       "chien",
	"goodbye":   "au revoir",
	"true":      "vrai",
	"false":     "faux",
	"null":      "nul",
	"function":  "fonction",
	"return":    "retour",
	"error":     "erreur",
	"builtin":   "fonction intégrée",
	"array":     "tableau",
	"hash":      "table de hachage",
	"integer":   "entier",
	"boolean":   "booléen",
	"string":    "chaîne",
	"nil":       "nul",
	"and":       "et",
	"or":        "ou",
	"not":       "pas",
	"if":        "si",
	"else":      "sinon",
	"for":       "pour",
	"let":       "laisser",
	"in":        "dans",
	"of":        "de",
	"to":        "à",
	"with":      "avec",
	"func":      "fonction",
	"undefined": "indéfini",
}

func TranslateToFrench(s string) string {
	type response struct {
		Translation string `json:"translation"`
	}
	url := fmt.Sprintf("https://lingva.ml/api/v1/en/fr/%s", s)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("API error:", err)
		if fr, ok := manualDict[s]; ok {
			return fr
		}
		return s
	}
	defer resp.Body.Close()
	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Decode error:", err)
		fmt.Println("API response body:", string(body))
		if fr, ok := manualDict[s]; ok {
			return fr
		}
		return s
	}
	// If translation is empty or unchanged, use manualDict
	if res.Translation == "" || strings.EqualFold(res.Translation, s) {
		if fr, ok := manualDict[s]; ok {
			return fr
		}
		return s
	}
	return res.Translation
}
