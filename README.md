# bsr

### arch
- client-server
- pvp and pve (implement game ai)
- game notation (like chess PGN)
- parameters (number of games, etc.)
- server allows client to connect and create a new game


### game rules
- 3 rounds:
  1. 2 lives, no items, no sudden death
  2. 4 lives, 2 items, no sudden death
  3. 5 lives, 4 items, sudden death
- items:
  - magnifying glass (reveal loaded shell to player)
  - cigarettes (+1 life, does not work on round 3)
  - beer (eject loaded shell)
  - handsaw (2x dmg)
  - handcuffs (skip opponents turn)