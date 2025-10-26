The Factory Pattern is a creational design pattern that provides a way to create objects without exposing the creation logic to the client (the code that uses them).

✅ Advantages of Factory Pattern
Advantage	Explanation
🧩 Encapsulates Object Creation	Clients don’t need to know how objects are created.
🧠 Loose Coupling	Code depends on interfaces, not concrete classes.
🧱 Easier to Extend	Add new product types by extending the factory logic.
🧹 Centralized Object Management	Object creation logic is in one place, easier to change.
🧪 Improves Testing	You can mock or stub factories easily.


⚠️ Disadvantages
Disadvantage	Explanation
🧱 More Classes	Can increase the number of files/classes in large projects.
🔁 Complexity	For small apps, simple new might be easier.
⚙️ Limited Control for Client	Factory decides what object to create, not the caller.
🧩 Conditional Logic May Grow	In large systems, factory’s switch/if logic can get big (you can later replace it with Abstract Factory or Registration pattern).


🧭 When to Use Factory Pattern

Use when:

You have a family of related objects (e.g., different robot types).

You want to decouple object creation from object usage.

You expect to add new types frequently.

The creation process is complex or involves setup steps.



🏗️ Factory + Strategy Together

You can combine both patterns beautifully:

Factory → decides which strategies (behaviors) to inject.

Strategy → decides how the robot behaves at runtime.