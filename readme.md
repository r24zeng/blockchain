# tttgame

**tttgame** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `playerID invites sb to gameID`

A player invites another player to join the game.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| playerID  | Y        | String          | Player may exist or not                        |
| gameID | Y        | String | Game may exist or not |

### `playerID accept to join gameID`

A player invites another player to join the game.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| playerID  | Y        | String          | Player may exist or not                        |
| gameID | Y        | String | Game may exist or not |

### `playerID place a piece to (X, Y) on the game board`

A player invites another player to join the game.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| playerID  | Y        | String          | Player may exist or not                        |
| X | Y        | String | in [0, 1, 2] |
| Y | Y        | String | in [0, 1, 2] |

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
