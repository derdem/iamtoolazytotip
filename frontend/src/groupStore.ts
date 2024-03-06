import { create } from "cypress/types/lodash";
import { createStore } from "solid-js/store";

export enum Strength {
  Weak = 1,
  Medium = 2,
  Strong = 3
}

export interface Country {
  name: string
  strength: Strength
}

export interface GroupInStore {
  groupName: string
  countries: Country[],
  index: Symbol | null
}

export interface MatchInStore {
  groupIndex: Symbol | null
  matchIndex: number
  country1: string
  country2: string
}

const createGroupStore = () => {
  return createStore<GroupInStore[]>([])
}

const createGroupIndexStore = () => {
  return createStore<Symbol[]>([])
}

const createMatchStore = () => {
  return createStore<MatchInStore[]>([])
}

export const createEmptyMatch: () => MatchInStore = () => (
  {groupIndex: null, matchIndex: 0, country1: "", country2: ""}
)

export const createGroupsEmptyMatches = (groupIndex: Symbol) => {
  const indices = [...Array(6).keys()]
  const matches: MatchInStore[] = []
  indices.forEach((matchIndex) => {
      matches.push({groupIndex, matchIndex, country1: "", country2: ""})
  })

  return matches
}

export const [groups, setGroups] = createGroupStore()

// changing this to a signal array rerenders the component, while user is
// typing inside the group card this signal stays untouched
export const [groupIndex, setGroupIndex] = createGroupIndexStore()

export const [matches, setMatches] = createMatchStore()

export const getGroup = (groupIndex: Symbol) => {
  const myGroups = groups.filter((group) => group.index === groupIndex);
  if (myGroups.length !== 1) {
    throw new Error("Group not found");
  }
  return myGroups[0];
}

