# tzgolk-in

Goals

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