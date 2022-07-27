# Deck of Cards exercise

[Source](https://courses.calhoun.io/lessons/les_goph_54)

## Exercise details

Create a package that can be used to build decks of cards and implement the following options:

- An option to sort the cards with a user-defined comparison function. The sort package in the standard library can be used here, and expects a less(i, j int) bool function.
- A default comparison function that can be used with the sorting option.
- An option to shuffle the cards.
- An option to add an arbitrary number of jokers to the deck.
- An option to filter out specific cards. Many card games are played without 2s and 3s, while others might filter out other cards. We can provide a generic way to handle this as an option.
- An option to construct a single deck composed of multiple decks. This is used often enough in games like blackjack that having an option to build a deck of cards with say 3 standard decks can be useful.
