# Blackjack exercise

[Source](https://courses.calhoun.io/lessons/les_goph_63)

## Exercise details

Using the Deck of Cards package, Build a relatively simple card game with a simple AI.

Specifically, we are going to create a blackjack game that supports the following rules:

1. Every player is dealt 2 cards. For simplicity, our first version will only support two players - the dealer and the human player. The dealing starts at the first player, and continues around the table until the dealer is dealt and then repeats, starting again with the first player, until all players have two cards. The dealer only has one visible card. The other is “face down” and isn’t visible to players. All player cards are visible.

2. The player’s turn. In our limited version of blackjack, the player will have two options: Hit or Stand. If a player chooses to hit, they are dealt a new card and will then be allowed to choose between the hit and stand options again. If a player chooses to stand their turn ends and the next player is up.

3. The dealer’s turn. If they have a score of 16 or less, or a soft 17, they will hit. This means we will need to implement scoring, and will be able to determine which player has won the game.

4. Determining the winner. The winner is the player who has the highest score without going over 21. Cards 2-10 are worth their face value in points. Aces are worth 1 or 11 points (whichever is best without busting), and face cards (J, Q, K) are all worth 10 points. A “soft 17” is a score of 17 in which 11 of the points come from an Ace card. If the player busts during their turn, the dealer automatically wins.

Blackjack occurs when a player has an Ace with a face card (J, Q, K), or a 10 card. In traditional blackjack there are special rules for this, but for our simple game we won’t be adding that.
