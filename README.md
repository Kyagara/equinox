## Equinox

This shouldn't be used in any production environment, this is just a practice tool for me to learn CI/CD using Github actions and tests in golang. But hey, unless you really want to use this package just so you can use for the single endpoint implemented, hey uh, have fun.

I was recommended [Alex Pliutau](https://www.youtube.com/watch?v=evorkFq3Y5k)'s video on youtube and got curious about other aspects like tests.

## TODO

### Improve project structure

I think the best approach for this project would be to separate LOL and others Riot games in modules, although I am not certain on how I would implement this yet. Currently, even if I supported just LOL endpoints, the project would quickly become a mess to navigate. This should be solved before I even try to implement other endpoints.

### Tests

I am not sure if tests are 'good enough', I am just checking if errors are Nil and the response is NotNil using [testify](https://github.com/stretchr/testify).
