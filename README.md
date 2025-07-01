# 🐶 Snout Interpreter

**Snout** is a programming language interpreter with a unique twist:  
**All output—including errors, booleans, null, and strings—appears in 🇫🇷 French!**

---

## ✨ Features

- 📦 **Variables**: `let x = 5;`
- ➕ **Arithmetic**: `+`, `-`, `*`, `/`
- ✅ **Booleans**: `true`, `false` (output as `vrai`, `faux`)
- 🚫 **Null**: `null` (output as `nul`)
- 🔀 **Conditionals**: `if`, `else`
- 🔙 **Return statements**: `return`
- 🧑‍💻 **Functions**: `fn(x) { ... }`
- ☎️ **Function calls**: `add(2, 3);`
- 🗃️ **Arrays**: `[1, 2, 3]`
- 🗝️ **Hashes (maps)**: `{"key": "value"}`
- 🛠️ **Built-in functions**: `len`, `first`, `last`, `rest`, `push`, `puts`
- 🐾 **Error handling**: All errors are in French (with an ASCII dog for parser errors)
- 🇫🇷 **String output**: All string values are translated to French (via API or manual dictionary)
- 💻 **REPL**: Interactive prompt

---

## 🚀 Getting Started

### 1. Clone and Build

```sh
git clone https://github.com/Doris-Lam/Snout.git
cd Snout
go build
```

### 2. Run the Interpreter

```sh
go run main.go
```

You’ll see:

```
Hello YOURNAME! This is the Snout programming language!
All output (including errors, booleans, null, and strings) will appear in French!
Feel free to type in commands
>>
```

---

## 📝 Language Examples

### 📦 Variables and Arithmetic

```monkey
let a = 10;
let b = 20;
a + b * 2
```

### 🧑‍💻 Functions

```monkey
let add = fn(x, y) { return x + y; };
add(5, 7)
```

### 🔀 Conditionals

```monkey
if (10 > 5) { "hi" } else { "bye" }
```

### 🗃️ Arrays

```monkey
let arr = ["hi", "dog"];
puts(arr)
```
**Output:**  
```
[bonjour, chien]
```

### 🗝️ Hashes

```monkey
let h = {"dog": "chien", "cat": "chat"};
puts(h)
```

### 🛠️ Built-in Functions

```monkey
len("bonjour")      // 7
len([1, 2, 3])      // 3
first([1, 2, 3])    // 1
puts("hi")          // prints: bonjour
```

---

## 🇫🇷 French Output

- All string values are translated to French (using an API or manual dictionary).
- Booleans print as `vrai`/`faux`.
- Null prints as `nul`.
- Errors print in French.

---

## 🐶 Error Handling & ASCII Dog

Whenever you make a parser error, Snout prints an ASCII dog and a French error message.

**Example:**
```monkey
dog
```
Output:
```
  / \__
 (    @\___
 /         O
/   (_____/
/_____/ U

ERREUR : identifiant introuvable: dog
```

---

## 🧪 Testing

Try these in the REPL:

```monkey
puts("hi")           // bonjour
puts([true, false])  // [vrai, faux]
5 + true             // ERREUR : décalage de type : INTEGER + BOOLEAN
unknownVar           // ERREUR : identifiant introuvable: unknownVar
```

---

## 🛠️ Implementation Notes

- **String translation** uses the [Lingva Translate API](https://lingva.ml/) and a manual French dictionary as fallback.
- **ASCII dog** is shown for all parser errors.
- **All output is in French** for a fun, immersive experience.

---

## 🤝 Contributing

Pull requests welcome!  
Feel free to add more built-in functions, improve French translations, or enhance the language.

---

## 📄 License

MIT

---

**Bon codage avec Snout! 🐾**
