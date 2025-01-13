# tzgolk-in

Goals

Things left to do

- [x] Enhance display to display player resources
- [x] add corn accumulation
- [x] Add in wealth tiles
- [x] Add age 1 buildings
- [x] add age 2 buildings
- [x] add free research to getoptions
- [x] Add first player space
- [x] add light side / dark side
- [x] add ages, food days, wealth days, point days, and such
- [x] set up currentBuildings
- [x] deal out age 2 buildings
- [x] add corn exchange
- [x] add pay corn to get less
- [x] pay corn for moves (and only generate playable moves)
- [x] workers that fall off come back
- [x] add theo 1
- [x] add second building
- [x] add "player doesn't pay for 1 worker" and "player pays 1 food for worker" effects
- [x] display research 
- [x] display temples
- [ ]  confirm all of research is there (who knows lol)
- [x] add building color
- [x] add monuments and set up currentMonuments
- [x] end game
- [ ] beg for corn
- [ ] pay corn to go lower (first attempt broke)
- [ ] uxmal exchange (first attempt broke)
- [ ] make game clone itself not calendar
- [ ] can't use same resources twice



- [ ] change age1 and age2 stuff to just be arrays
- [ ] add cdata.full to display
- [ ] display buildings & monuments

## AI 

- [ ] Be able to generate all legal moves
  - [x] Placement
  - [ ] Retrieval
    - [ ] order matters (e.g. retrieving 2 1 is different than 1 2)
    - [ ] can pay corn to access lower actions
    - [ ] decisions
      - [ ] needs to be wheel-specific?
      - [ ] how about: decisions is a number on *wheels*, not *positions*
        - [ ] palenque
          - [ ] 1: go to 1, take corn
          - [ ] 2: go to 1, take no tile but get corn (only w/ research)
          - [ ] 3: go to 1, 
      - [ ] or: decisions is a number on positions
      - [ ] when generating retrieval moves
        - [ ] each position returns a list of possibilities
        - [ ] yes this works better & it works for mirrors as well
        - [ ] since now a mirror is an expanded list of the others
      - [ ] palenque is a special wheel (?)
      - [ ] no - each palenque position contains nWheat and nWood
        - [ ] so does each position need this?
      - [ ] so change position to an interface
      - [ ] and then PalenquePosition implements it
- [ ] execute moves in a calendar
- [ ] implement food days & temples
- [ ] implement research
- [ ] implement buildings & monuments