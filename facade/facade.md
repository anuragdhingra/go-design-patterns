## Facade pattern 
Structural design patterns 
- Adapter 
- Facade
- Proxy 
- Decorator 
- Composite

Facade pattern hides the complexities of underlying classes and provides a unified interface to the client to access any underlying interfaces.

In the code example, we want to show that all complex business logic is hidden in Facade class and we just need to call methods given by facade.

Video explanation:
https://www.youtube.com/watch?v=K4FkHVO5iac

Golang example:
https://golangbyexample.com/facade-design-pattern-in-golang/

### Principle of least knowledge 
For loose coupling, the ones using a function should be the only one dependent on it and the ones using the previous ones should not be affected by change. 
   