# High-level Go Package Design

Most standard library or third-party packages are wrapped around an abstraction. e.g. `net/http` is not exposed to consumers, but a `github.com/username/repo/http` package instead. Inside the wrapper package is a struct with private fields and then public methods to interact with. The idea being to only expose the necessary functionality for the app to work, not expose the entire third-party package.

The goal is to hopefully provide benefits in:

1. Abstraction: It allows you to hide the implementation details of the underlying packages and provide a cleaner, more maintainable interface to your code.

2. Extensibility: By wrapping the packages, you can extend the functionality to meet your specific needs and make it easier to switch to alternative packages in the future if necessary.

3. Testability: Wrapping the packages makes it easier to write tests for your code and to isolate the dependencies for easier testing.

4. Consistency: You can enforce a consistent interface across different packages and maintain a consistent code style throughout your application.

However, it's important to not over-abstract as that can make your code more complex and harder to understand. Also, make sure to document your abstractions clearly so that other developers can understand the purpose and use of the abstractions you have created.

### When are you over-abstracting?

1. The abstractions add more complexity than value. If the abstraction makes the code harder to understand and does not add significant benefits, then it's over-abstracted.

2. The abstractions are not used consistently throughout the code. If the abstractions are only used in a few places, then it's probably over-abstracted.

3. The abstractions do not provide clear benefits. If the abstractions do not simplify the code or provide significant benefits, then it's over-abstracted.

4. The abstractions are not well-documented. If the abstractions are not well-documented, then it's over-abstracted.

5. The abstractions require a lot of boilerplate code. If the abstractions require a lot of boilerplate code to be written, then it's over-abstracted.

Remember, the goal of abstraction is to simplify the code and make it easier to maintain, not to add complexity. If the abstraction does not meet these goals, then it's likely over-abstracted.
