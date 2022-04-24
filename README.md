## Equinox

This shouldn't be used in any production environment, this is just a practice tool for me to learn CI/CD using Github Actions and tests in golang. But hey, unless you really want to use this package just so you can use the single endpoint implemented, hey uh, have fun.

I was recommended [Alex Pliutau](https://www.youtube.com/watch?v=evorkFq3Y5k)'s video on youtube and got curious about other aspects like tests, I decided to make a client for the Riot Games API since I am more familiar with it.

I am avoiding using other packages like [resty](https://github.com/go-resty/resty) instead of the `net/http` package go provides to improve my golang knowledge, currently just using [testify](https://github.com/stretchr/testify).

## TODO

### Tests

I am not sure if tests are 'good enough', I am just checking if errors are `Nil` and the response is `NotNil` using `testify`.

### Logging

I don't believe the current logging and 'debugging mode' is done right or that its 'done' at all to be honest.
