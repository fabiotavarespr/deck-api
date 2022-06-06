# Deck of Cards API

Thank you for your time reviewing this code :)

## Stack

- [Golang](https://go.dev/)
- [Postgres](https://www.postgresql.org/)

## Libraries

- [Testify](https://github.com/stretchr/testify)
- [Zap](https://github.com/uber-go/zap)
- [Gorm](https://gorm.io/index.html)
- [Gin](https://github.com/gin-gonic/gin)
- [Ozzo-Validation](https://github.com/go-ozzo/ozzo-validation)

## Execution

This project contains a [make](Makefile) file to make it easy. To execute the application, just execute the
command `make start` or `docker-compose up -d`.

## Tests

The tests could be executed by your favorite IDE or executing the command `make tests` or even executing `go test ./...`
.

## API

| Method | Resource | Params |
| :---: | :---: | :---: |
| POST | /decks | shuffle: boll, optional, default false </br> cards: string, optional, default empty |
| GET  | /decks/<deck_id> | -
| PATCH | /decks/<deck_id>/drawn | count: integer, optional, default 1

Import the [postman collections](/api/deck-api.postman_collection.json) for more details.

## Assumptions

- Standard ordering
    - Spades, diamonds, clubs and hearts.


- Partial/Custom deck
    - Use **0** (zero) to represent **Ten** cards in the parameter **cards**.
        - 0S: Ten of Spades
        - 0D: Ten of Diamonds
        - 0C: Ten of Clubs
        - 0H: Ten of Hearts
    - Once a Custom deck is created and the parameter **shuffle**=false, the cards will be saved on the exact way
      received in parameter **cards**.
    - Case there is/are one or more invalid card(s) within cards parameter, the deck will be not created and, a 400
      error will be returned.


- Draw cards
    - If the **count** parameter value informed is greater than the remaining cards in the deck, a 400 error will be
      returned.

## What can be improved?

There are a couple of improvements which could be done with more time:

- Use env vars to set up the components in the application. I usually implement this with libraries
  as [godotev](https://github.com/joho/godotenv).
- Use properly migrations to create tables and indexes - could be done by [Gorm](https://gorm.io/index.html) itself or
  another migration library.
- Add more tests.
    - Most of the logic was tested in the usecase, models and dtos. But I would add tests for the handler and
      repositories.
    - For repositories depending on the situation we can use [test-containers](https://www.testcontainers.org/).
- In some projects, patterns as [golang-standards](https://github.com/golang-standards/project-layout) are widely used
  to organize the code. For this challenge I just kept the structures in the root, to make the code review easier and
  also, to avoid misunderstandings.

*At the business side, I would implement a rule to avoid insert duplicated cards on partial decks.