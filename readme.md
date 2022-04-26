### Running
```
 go run . --expr "x_(λy.x)_y_(z_z)"
```


### Grammar

Origin lambda-calculus grammar:

* `Λ ⟶ v`
* `Λ ⟶ Λ _ Λ`
* `Λ ⟶ λ v . Λ`

lambda-calculus grammar, which is converted to LL(1) 
* `Λ ⟶ v Λs | λ v . Λ Λs`
* `Λs ⟶ ε | _ Λ`


###  First and Follow
* `FIRST(T) = { λ v }`
* `FIRST(T_SUB) = { _ ε }`

Computed by https://mikedevice.github.io/first-follow/

## Automata for lexical analyzer

![Lexical Analyzer Automata](https://github.com/DamirJann/math-parser/blob/master/img/automata_for_lexical_analyzer.drawio.png)

## Supported 
* Operations:
  * MULTIPLICATION - `*`
  * DIVISION - `/`
  * PLUS - `+`
  * MINUS - `-`
* Numbers:
  * Integer

