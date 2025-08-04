# golangRaylibEngine

This is just a setup project for creating an app or a game 

the objective is to have a unityLike setup and run

There is a Scene manager, you choose a scene in the scenes list,
In each Scene we have a list of Gameobjects
in each GameObject we have a list of component
a component is a class like scripts in unity
in each component , declare a function start, update , draw
most time draw will be empty ( only if we need to draw something, do not put logic in there)
the scene manager calls the update funcs of each component of the choosed scene

Reminder: avoid declaring variables in the update or draw functions  and do not put anything else in the main for loop in AppLoop
