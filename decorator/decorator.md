### Decorator pattern
Decorator Pattern is used to attach the additional responsibility to a class dynamically.
Without this, we might have
- Class explosion
- Contradict design principles 

Decorator pattern has a component interface (abstract class) and as many decorators as we want (abstract of component). 
Decorator is `is-a` Component and `has-a` component. 

It works something like recursion where base concrete component acts as basecase. 

Also, order doesnot matter in decorator pattern.

Explanation video -
https://www.youtube.com/watch?v=GCraGHx6gso