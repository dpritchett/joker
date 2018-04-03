// This file is generated by generate-std.joke script. Do not edit manually!

package html

import (
  "html"
  . "github.com/candid82/joker/core"
)

var htmlNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.html"))

var escape_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := html.EscapeString(s)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}

var unescape_ Proc = func(args []Object) Object {
  c := len(args)
  switch  {
  case c == 1:
    
    s := ExtractString(args, 0)
    res := html.UnescapeString(s)
    return MakeString(res)

  default:
    PanicArity(c)
  }
  return NIL
}


func init() {

htmlNamespace.ResetMeta(MakeMeta(nil, "Provides functions for escaping and unescaping HTML text.", "1.0"))

htmlNamespace.InternVar("escape", escape_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Escapes special characters like < to become &lt;. It escapes only five such characters: <, >, &, ' and ".`, "1.0"))

htmlNamespace.InternVar("unescape", unescape_,
  MakeMeta(
    NewListFrom(NewVectorFrom(MakeSymbol("s"))),
    `Unescapes entities like &lt; to become <.`, "1.0"))

}