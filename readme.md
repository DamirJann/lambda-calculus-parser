### Running
```
 go run . --expr="x_(λy.x)_y_(z_z)"
 go run . --red="beta" --expr="(λy.x)_y_(z_z)" 
 go run . --red="alpha" --expr="(λy.x)_y_(z_z)" --sub="z=t,y=q"  
 
```


### Grammar

Origin lambda-calculus grammar:

* `Λ ⟶ v`
* `Λ ⟶ Λ _ Λ`
* `Λ ⟶ λ v . Λ`

lambda-calculus grammar, which is converted to LL(1) 
* `Λ ⟶ v Λs | λ v . Λ Λs | ( Λ ) Λs`
* `Λs ⟶ ε | _ Λ`


###  First and Follow
* `FIRST(Λ) = { λ v ( }`
* `FIRST(Λs) = { _ ε }`

Computed by https://mikedevice.github.io/first-follow/

## Automata for lexical analyzer

![Lexical Analyzer Automata](https://github.com/DamirJann/lambda-calculus-parser/blob/master/img/automata_for_lexical_analyzer.drawio.png)


