# Writing Testable Go Code

Our goal is to make improve on what we did in [Writing Untestable Go Code](https://github.com/PeterAtEmmaSleep/writing-untestable-go-code).

# What we did better

- We do not mix data and behaviour (we do not write things like "_Object! Validate yourself!_"). This allows us to mock things without trying to satisfy complex business logic in tests.
- We split the logic into services which we can mock. This allows us write tests without spnning up the whole application

# What we still have to improve
- we still have tight coupling between the business logic and its representation -> our app write http responses and this makes tests hard to write
- our app is still doing to much stuff and this makes the tests very long and complicated