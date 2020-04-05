# Julepe

## Rules

For now we will assume this is always a 5 player game.

### Deck 

We play with 30 cards; that is, 10 cards of 3 suits.

### The pot

The pot is the money pool in the middle of the table. Every round the pot increases by 50 cents, unless it is split (or won entirely). We will explore this more later.

Before the game begins, each player places 50 cents in the pot - so we start with 2.50.

### Dealing the cards

The dealer changes each round, so for the first round let's assume player 1 is the dealer. The next round, player 2 is the dealer, and so on.

The dealer will deal 5 cards to each player, and a 6th to the dealer himself which is placed face up on the table - this card's suit is the trump suit _for that round_; i.e. the trump suit is selected at the beginning of each round.

### In, or not

Now each player will say, in order clockwise from the dealer, whether or not they wish to participate in this round.

This decision is based on the trump suit:

* Do I have any cards of the trump suit?
* If yes, are the cards I have of the trump suit a high value?

However, _a player must play if he has the ace of the trump suit_.

Additionally, if at least one player decides to play, _the dealer must play_.

The players who do not play, their hands are placed in the discard pile.

### Switching cards for a better hand

Each normal player has 5 cards; he may change 1 - 5 cards in exchange for the same number of cards from the 4 cards that were left over from dealing - the "leftovers". 

If there are no cards left in the "leftovers", the player must draw cards from the discard pile, i.e. the cards which the players who are not playing discarded.

_This is optional_, in the sense that if the player has a good hand he is not forced to discard any cards.

The dealer has 6 cards (5 normal cards, plus the card which was placed face up to determine the trump suit); the dealer may change 1 - 6 cards in exchange for the same number of cards from either the "leftovers" (if there are any), or discard pile. Finally, the dealer must discard a final card such that he has 5 cards.

Now we are ready to play.

### The round

The player to the left (i.e. clockwise) of the dealer will be the first to play a card. Each player _must_ play the same suit and a higher value (if possible).

If they have the same suit but not a higher card, they may play a lower card. However, if they do not have anything of that suit, they must play the trump suit. If they don't have anything of the trump suit, they can play any other card.

The first player sets the "type" for that hand, so just because someone else plays a different type does not mean the type for that hand has changed.

Of course, each round will have five hands.

Let's take an example:

Trump: Copas
1) 10 of Bastos
2) 12 of Bastos
3) 2 of Copas

Player 3 wins, because of the trump card.

Trump: Copas
1) 3 of Bastos
2) 2 of Copas
3) 5 of Copas

Player 3 wins again: both 2 and 3 played the trump card, but player 3 played the highest card.

Trump: Copas
1) 3 of Bastos
2) 2 of Bastos (doesn't have higher than the 3)
3) 1 of Bastos

Player 3 wins: he has played the ace. Notice that player 2 did not play a higher card than player 1 - he doesn't have any, but _does have cards of the same suit_.

#### Safety

In order for a player to be "safe", he needs to have won at least 2 of 5 hands. We'll look at what this means in the `What happens to the pot?` section...

#### What happens to the pot?

##### 2 Players: Split

Player 1 wins 3, player 2 wins 2. Both players are "safe" - this means they split what's in the middle.

3 euro pot: 1.50 each
3.50 euro pot: 1.50 each, leaving 0.50 in the pot

##### 2 Players: Julepe

Player 1 wins 4, player 2 wins 1. Only player 1 is safe. As such, player 2 must pay the amount equal to the value of the pot to player 1. 

_The most a player can ever have to pay for a round is what is in the pot_. Again, the loss will never be greater than the pot value. Of course, the pot can sometimes be as big as 10 - 20 euros.

##### 3 Players: Julepe only, always

Now with 3 players, let's take an example; player 1 has won 2 hands, player 2 has won 2 hands and player 3 has won 1 hand. So player 1 and 2 are "safe" - as such, player 3 will need to pay the amount that's in the pot, and split it between player 1 and player 2: e.g. 2.50 in the pot, player 3 gives 1 to player 1 and 1 to player 2, and 0.50 to the pot.

Like the title says, _with 3 or more players it will always be a "Julepe" outcome_ (i.e. the loser pays the value of the pot, split between the winning players).

##### Before the next round

Following a "split": all players add 0.50 to the pot.

Following a "Julepe": the next player to become the dealer (clockwise from the previous dealer) places 0.50 to the pot.

#### No one plays: dealer vs house

If no player decides to play, the dealer may, depending on the rules of the village (literally):

Vilalba de los Alcores - _may not look at his cards_ and decide if he wants to play "against the house"
Urones - look at his cards and decide if he wants to play "against the house"

In either case, the "house" is the remaining cards which were not dealt, plus two taken from the top of the _shuffled_ discard pile (i.e. all of the hands of the players who did not want to play the round... but, like, shuffled once they have handed them back). The "house" will then select a card to discard.

The dealer has his six cards (5 cards of the normal hand, plus the trump card - as usual), and may now choose to discard any of the cards in his hand and take cards from the discard pile. If he chooses to do this, he will again have 6 cards - he must choose 1 to discard.

Both the dealer and the "house" now have 5 cards, and the game can begin. The "house" begins, and the game is played as usual.

##### Dealer vs. house: the pot

If the dealer wins: he takes the entire pot
If the dealer and house split the pot: dealer takes half of the pot
If the dealer loses, he must pay the value of the pot to the pot: i.e. _the pot doubles_

## TO DO 

* Split up `main.go`, it's a mess!
* Find a way of selecting a dealer each round
* Pot system (table and players)

Complex:

* Probablity system for which card is best to play
* Some kind of "learning mode" where a player can be shown which card is best to play - probablity values?
