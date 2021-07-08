### Factory Method pattern 
Factory Method pattern defines an interface for creating an object. It will let its subclass instantiate the objects according to some logic they want. 

A good usecase is in game design where we have to generate enemies in space.
- Factory for generating enemies
- We can use multiple factories for multi levels too or we can pass some params to a factory too

Another usecase is when we are creating a memory intensive objects eg. db connections.
We will keep a pool of connections and using DB Factory, we can return connections that are available using round robin.

Explanation Video:
https://www.youtube.com/watch?v=EcFVTgRHJLM

Good resource: 
https://refactoring.guru/design-patterns/factory-method

- Use the Factory Method when you don’t know beforehand the exact types and dependencies of the objects your code should work with.
- Use the Factory Method when you want to provide users of your library or framework with a way to extend its internal components. Android uses this.
- Use the Factory Method when you want to save system resources by reusing existing objects instead of rebuilding them each time.
